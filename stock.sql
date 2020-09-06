/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : localhost:3306
 Source Schema         : stock

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 06/09/2020 14:34:44
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for stock_history
-- ----------------------------
DROP TABLE IF EXISTS `stock_history`;
CREATE TABLE `stock_history` (
  `stock_code` varchar(12) NOT NULL COMMENT '股票代码',
  `stock_name` varchar(12) DEFAULT NULL COMMENT '股票名称',
  `stock_date` datetime DEFAULT NULL COMMENT '股票开盘日期',
  `stock_closing` float(10,2) DEFAULT NULL COMMENT '股票收盘价',
  `stock_high` float(10,2) DEFAULT NULL COMMENT '股票最高价',
  `stock_low` float(10,2) DEFAULT NULL COMMENT '股票最低价',
  `stock_opening` float(10,2) DEFAULT NULL COMMENT '股票开盘价',
  `stock_prior` varchar(12) DEFAULT NULL COMMENT '股票前收盘价',
  `stock_change_amount` varchar(12) DEFAULT NULL COMMENT '涨跌额',
  `stock_amplitude` float(10,2) DEFAULT NULL COMMENT '涨跌幅'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
