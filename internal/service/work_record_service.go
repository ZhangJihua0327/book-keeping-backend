package service

import (
	"book-keeping-backend/internal/model"
	"book-keeping-backend/internal/repository"
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/xuri/excelize/v2"
)

type WorkRecordService struct {
	repo *repository.WorkRecordRepository
}

func NewWorkRecordService(repo *repository.WorkRecordRepository) *WorkRecordService {
	return &WorkRecordService{repo: repo}
}

func (s *WorkRecordService) AddRecord(record *model.WorkRecord) error {
	if record.RecordID == "" {
		dateStr := time.Time(record.Date).Format("20060102")
		randomCode := generateRandomString(6)
		record.RecordID = fmt.Sprintf("wr-%s-%s", dateStr, randomCode)
	}
	return s.repo.Create(record)
}

func generateRandomString(n int) string {
	const letters = "0123456789abcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		ret[i] = letters[num.Int64()]
	}
	return string(ret)
}

func (s *WorkRecordService) GetRecordsByDate(dateStr string) ([]model.WorkRecord, error) {
	return s.repo.GetByDate(dateStr)
}

func (s *WorkRecordService) UpdateRecord(recordID string, updates map[string]interface{}) error {
	return s.repo.UpdateByRecordID(recordID, updates)
}

func (s *WorkRecordService) DeleteRecord(recordID string) error {
	return s.repo.DeleteByRecordID(recordID)
}

func (s *WorkRecordService) ExportRecords(filter model.WorkRecordFilter) (*bytes.Buffer, error) {
	records, err := s.repo.GetByFilter(filter)
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	sheetName := "WorkRecords"
	// Rename default sheet
	f.SetSheetName("Sheet1", sheetName)

	headers := []string{"ID", "记录ID", "车型", "日期", "客户名称", "施工地点", "数量", "价格", "是否已收费", "备注"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, h)
	}

	for i, r := range records {
		row := i + 2
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), r.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), r.RecordID)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), r.TrunkModel)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), r.Date.String())
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), r.CustomerName)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), r.ConstructionSite)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), r.Quantity)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), r.Price)
		charged := "否"
		if r.Charged != nil && *r.Charged {
			charged = "是"
		}
		f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), charged)
		f.SetCellValue(sheetName, fmt.Sprintf("J%d", row), r.Remark)
	}

	return f.WriteToBuffer()
}
