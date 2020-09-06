/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50527
 Source Host           : localhost:3306
 Source Schema         : stock

 Target Server Type    : MySQL
 Target Server Version : 50527
 File Encoding         : 65001

 Date: 06/09/2020 23:53:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for stock_history
-- ----------------------------
DROP TABLE IF EXISTS `stock_history`;
CREATE TABLE `stock_history`  (
  `stock_code` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '股票代码',
  `stock_name` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '股票名称',
  `stock_date` datetime DEFAULT NULL COMMENT '股票开盘日期',
  `stock_closing` float(15, 2) DEFAULT NULL COMMENT '股票收盘价',
  `stock_high` float(15, 2) DEFAULT NULL COMMENT '股票最高价',
  `stock_low` float(15, 2) DEFAULT NULL COMMENT '股票最低价',
  `stock_opening` float(15, 2) DEFAULT NULL COMMENT '股票开盘价',
  `stock_prior` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '股票前收盘价',
  `stock_change_amount` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '涨跌额',
  `stock_amplitude` float(15, 2) DEFAULT NULL COMMENT '涨跌幅'
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for stock_hy
-- ----------------------------
DROP TABLE IF EXISTS `stock_hy`;
CREATE TABLE `stock_hy`  (
  `hy_code` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `hy_name` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `hy_company_amount` int(5) DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for stock_now_data
-- ----------------------------
DROP TABLE IF EXISTS `stock_now_data`;
CREATE TABLE `stock_now_data`  (
  `stock_code` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `stock_wy_code` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `stock_hy_code` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `stock_name` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `stock_price` decimal(10, 2) DEFAULT NULL,
  `stock_open` decimal(15, 2) DEFAULT NULL,
  `stock_five_minute` decimal(15, 2) DEFAULT NULL,
  `stock_high` decimal(15, 2) DEFAULT NULL,
  `stock_hs` decimal(15, 2) DEFAULT NULL,
  `stock_lb` decimal(15, 2) DEFAULT NULL,
  `stock_low` decimal(15, 2) DEFAULT NULL,
  `stock_mcap` decimal(15, 2) DEFAULT NULL,
  `stock_mfratio2` decimal(15, 2) DEFAULT NULL,
  `stock_mfratio10` decimal(15, 2) DEFAULT NULL,
  `stock_mfsum` decimal(15, 2) DEFAULT NULL,
  `stock_pe` decimal(15, 2) DEFAULT NULL,
  `stock_percent` decimal(15, 2) DEFAULT NULL,
  `stock_tcap` decimal(15, 2) DEFAULT NULL,
  `stock_turnover` decimal(15, 2) DEFAULT NULL,
  `stock_updown` decimal(15, 2) DEFAULT NULL,
  `stock_volume` decimal(15, 2) DEFAULT NULL,
  `stock_wb` decimal(15, 2) DEFAULT NULL,
  `stock_yestclose` decimal(15, 2) DEFAULT NULL,
  `stock_zf` decimal(15, 2) DEFAULT NULL,
  PRIMARY KEY (`stock_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
