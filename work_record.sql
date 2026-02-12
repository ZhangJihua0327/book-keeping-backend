CREATE DATABASE IF NOT EXISTS bookeeper;
USE bookeeper;

CREATE TABLE IF NOT EXISTS `work_record` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    `record_id` VARCHAR(256) COMMENT '记录ID',
    `trunk_model` VARCHAR(256) COMMENT '车型',
    `date` DATE COMMENT '日期',
    `customer_name` VARCHAR(256) COMMENT '客户名称',
    `construction_site` VARCHAR(1024) COMMENT '施工地点',
    `quantity` INT UNSIGNED COMMENT '数量',
    `price` INT UNSIGNED COMMENT '价格',
    `charged` BOOLEAN COMMENT '是否已收费',
    `remark` VARCHAR(4096) COMMENT '备注',
    INDEX `idx_trunk_model` (`trunk_model`),
    INDEX `idx_date` (`date`),
    INDEX `idx_customer_name` (`customer_name`),
    INDEX `idx_record_id` (`record_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='工作记录表';

CREATE TABLE IF NOT EXISTS `customer_name` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    `customer_name` VARCHAR(256) COMMENT '客户名称'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='客户名称表';

CREATE TABLE IF NOT EXISTS `trunk_model` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    `trunk_model` VARCHAR(256) COMMENT '车型'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='车型表';

ALTER TABLE `customer_name` ADD UNIQUE INDEX `uk_customer_name` (`customer_name`);
ALTER TABLE `trunk_model` ADD UNIQUE INDEX `uk_trunk_model` (`trunk_model`);
