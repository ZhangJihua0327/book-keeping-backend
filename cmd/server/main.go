package main

import (
	_ "book-keeping-backend/docs"
	"book-keeping-backend/internal/handler"
	"book-keeping-backend/internal/repository"
	"book-keeping-backend/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Book Keeping API
// @version 1.0
// @description This is a sample server for a book keeping application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	// 1. Initialize DB
	repository.InitDB()

	// 2. Initialize Repositories
	customerRepo := repository.NewCustomerRepository()
	trunkRepo := repository.NewTrunkModelRepository()
	recordRepo := repository.NewWorkRecordRepository()

	// 3. Initialize Services
	customerService := service.NewCustomerService(customerRepo)
	trunkService := service.NewTrunkModelService(trunkRepo)
	recordService := service.NewWorkRecordService(recordRepo)

	// 4. Initialize Handlers
	customerHandler := handler.NewCustomerHandler(customerService)
	trunkHandler := handler.NewTrunkModelHandler(trunkService)
	recordHandler := handler.NewWorkRecordHandler(recordService)

	// 5. Setup Router
	r := gin.Default()

	// Root URL
	r.GET("/", func(c *gin.Context) {
		c.String(200, "welcome to bookeeper")
	})

	// Optional: Add CORS middleware if needed
	// r.Use(cors.Default())

	// 6. Define Routes
	api := r.Group("/api")
	{
		// 1. 添加客户
		api.POST("/customers", customerHandler.AddCustomer)
		// 2. 返回所有客户列表
		api.GET("/customers", customerHandler.GetAllCustomers)

		// 3. 添加车型
		api.POST("/models", trunkHandler.AddTrunkModel)
		// 4. 返回所有车型
		api.GET("/models", trunkHandler.GetAllTrunkModels)

		// 5. record: 用于填写work_record
		api.POST("/records", recordHandler.AddRecord)
		// 6. 查询指定日期创建的所有记录 按照主键id 倒序
		// usage: GET /api/records?date=2023-10-27
		api.GET("/records", recordHandler.GetRecords)
		// 7. 对record_id 记录进行更正
		api.PUT("/records/:id", recordHandler.UpdateRecord)
		// 8. 导出 excel
		api.GET("/records/export", recordHandler.ExportRecords)
	}

	// 8. Swagger URL
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 9. Start Server
	r.Run(":8080")
}
