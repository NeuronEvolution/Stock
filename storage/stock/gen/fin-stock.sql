-- MySQL dump 10.13  Distrib 5.7.17, for macos10.12 (x86_64)
--
-- Host: 127.0.0.1    Database: fin-stock
-- ------------------------------------------------------
-- Server version	5.7.19

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `exchange`
--

DROP TABLE IF EXISTS `exchange`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `exchange` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `exchange_id` varchar(32) NOT NULL COMMENT '交易所ID',
  `exchange_name_cn` varchar(128) NOT NULL COMMENT '交易所中文名称',
  `exchange_name_en` varchar(128) NOT NULL COMMENT '交易所英文名称',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `exchange_id_UNIQUE` (`exchange_id`),
  KEY `idx_update_time` (`update_time`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `stock`
--

DROP TABLE IF EXISTS `stock`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `stock` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `stock_id` varchar(32) NOT NULL COMMENT '股票ID',
  `exchange_id` varchar(32) NOT NULL COMMENT '交易所ID',
  `stock_code` varchar(32) NOT NULL COMMENT '股票代码',
  `stock_name_cn` varchar(32) NOT NULL COMMENT '股票中文名称',
  `stock_name_en` varchar(32) NOT NULL COMMENT '股票英文名称',
  `launch_date` datetime NOT NULL COMMENT '上市日期',
  `company_name_cn` varchar(128) NOT NULL COMMENT '公司中文名称',
  `company_name_en` varchar(128) NOT NULL COMMENT '公司英文名称',
  `website_url` varchar(128) NOT NULL COMMENT '网站地址',
  `industry_name` varchar(32) NOT NULL COMMENT '行业名称',
  `city_name_cn` varchar(128) NOT NULL COMMENT '城市中文名称',
  `city_name_en` varchar(128) NOT NULL COMMENT '城市英文名称',
  `province_name_cn` varchar(128) NOT NULL COMMENT '省份中文名称',
  `province_name_en` varchar(128) NOT NULL COMMENT '省份英文名称',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `stock_id_UNIQUE` (`stock_id`),
  UNIQUE KEY `idx_exchang_id_stock_code` (`exchange_id`,`stock_code`),
  KEY `idx_update_time` (`update_time`),
  KEY `idx_industry_name` (`industry_name`)
) ENGINE=InnoDB AUTO_INCREMENT=933 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-09-27  9:53:20
