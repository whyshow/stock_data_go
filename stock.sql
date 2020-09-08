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

 Date: 08/09/2020 21:12:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for stock_history
-- ----------------------------
DROP TABLE IF EXISTS `stock_history`;
CREATE TABLE `stock_history` (
  `stock_code` varchar(15) NOT NULL COMMENT '股票代码',
  `stock_wy_code` varchar(15) DEFAULT NULL COMMENT '网易股票代码',
  `stock_hy_code` varchar(15) DEFAULT NULL COMMENT '行业代码',
  `stock_name` varchar(15) DEFAULT NULL COMMENT '名称',
  `stock_price` decimal(10,2) DEFAULT NULL COMMENT '价格',
  `stock_open` decimal(15,2) DEFAULT NULL COMMENT '今开',
  `stock_five_minute` decimal(15,6) DEFAULT NULL COMMENT '5分钟涨跌额',
  `stock_high` decimal(15,2) DEFAULT NULL COMMENT '最高价',
  `stock_hs` decimal(15,4) DEFAULT '0.0000' COMMENT '换手率',
  `stock_lb` decimal(15,2) DEFAULT NULL COMMENT '量比',
  `stock_low` decimal(15,2) DEFAULT NULL COMMENT '最低',
  `stock_mcap` decimal(15,2) DEFAULT NULL COMMENT '流通市值',
  `stock_mfratio2` decimal(15,2) DEFAULT NULL COMMENT '净利润',
  `stock_mfratio10` decimal(15,2) DEFAULT NULL COMMENT '主营收',
  `stock_mfsum` decimal(15,6) DEFAULT NULL COMMENT '每股收益',
  `stock_pe` decimal(15,6) DEFAULT NULL COMMENT '市盈率',
  `stock_percent` decimal(15,6) DEFAULT NULL COMMENT '涨跌幅',
  `stock_tcap` decimal(15,2) DEFAULT NULL COMMENT '总市值',
  `stock_turnover` decimal(15,2) DEFAULT NULL COMMENT '成交额',
  `stock_updown` decimal(15,2) DEFAULT NULL COMMENT '涨跌额',
  `stock_volume` int(15) DEFAULT NULL COMMENT '成交量',
  `stock_wb` decimal(15,2) DEFAULT NULL COMMENT '委比',
  `stock_yestclose` decimal(15,2) DEFAULT NULL COMMENT '昨收',
  `stock_zf` decimal(15,6) DEFAULT NULL COMMENT '振幅',
  `stock_time` datetime DEFAULT NULL COMMENT '日期'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for stock_hy
-- ----------------------------
DROP TABLE IF EXISTS `stock_hy`;
CREATE TABLE `stock_hy` (
  `hy_code` varchar(10) NOT NULL,
  `hy_name` varchar(25) DEFAULT NULL,
  `hy_company_amount` int(5) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for stock_listing
-- ----------------------------
DROP TABLE IF EXISTS `stock_listing`;
CREATE TABLE `stock_listing` (
  `stock_code` varchar(15) NOT NULL COMMENT '股票代码',
  `stock_wy_code` varchar(15) DEFAULT NULL COMMENT '网易股票代码',
  `stock_hy_code` varchar(15) DEFAULT NULL COMMENT '行业代码',
  `stock_name` varchar(15) DEFAULT NULL COMMENT '名称',
  `stock_price` decimal(10,2) DEFAULT NULL COMMENT '价格',
  `stock_open` decimal(15,2) DEFAULT NULL COMMENT '今开',
  `stock_five_minute` decimal(15,6) DEFAULT NULL COMMENT '5分钟涨跌额',
  `stock_high` decimal(15,2) DEFAULT NULL COMMENT '最高价',
  `stock_hs` decimal(15,4) DEFAULT '0.0000' COMMENT '换手率',
  `stock_lb` decimal(15,2) DEFAULT NULL COMMENT '量比',
  `stock_low` decimal(15,2) DEFAULT NULL COMMENT '最低',
  `stock_mcap` decimal(15,2) DEFAULT NULL COMMENT '流通市值',
  `stock_mfratio2` decimal(15,2) DEFAULT NULL COMMENT '净利润',
  `stock_mfratio10` decimal(15,2) DEFAULT NULL COMMENT '主营收',
  `stock_mfsum` decimal(15,6) DEFAULT NULL COMMENT '每股收益',
  `stock_pe` decimal(15,6) DEFAULT NULL COMMENT '市盈率',
  `stock_percent` decimal(15,6) DEFAULT NULL COMMENT '涨跌幅',
  `stock_tcap` decimal(15,2) DEFAULT NULL COMMENT '总市值',
  `stock_turnover` decimal(15,2) DEFAULT NULL COMMENT '成交额',
  `stock_updown` decimal(15,2) DEFAULT NULL COMMENT '涨跌额',
  `stock_volume` int(15) DEFAULT NULL COMMENT '成交量',
  `stock_wb` decimal(15,2) DEFAULT NULL COMMENT '委比',
  `stock_yestclose` decimal(15,2) DEFAULT NULL COMMENT '昨收',
  `stock_zf` decimal(15,6) DEFAULT NULL COMMENT '振幅',
  `stock_time` datetime DEFAULT NULL COMMENT '日期',
  PRIMARY KEY (`stock_code`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

SET FOREIGN_KEY_CHECKS = 1;
