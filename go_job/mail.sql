-- MySQL dump 10.13  Distrib 5.7.19, for Linux (x86_64)
--
-- Host: localhost    Database: mail
-- ------------------------------------------------------
-- Server version	5.7.19-0ubuntu0.16.04.1

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
-- Table structure for table `mail_0`
--

DROP TABLE IF EXISTS `mail_0`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail_0` (
  `id` int(11) NOT NULL,
  `company_id` int(11) NOT NULL,
  `task_id` int(11) NOT NULL,
  `contact_info` varchar(1000) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '打开 点击 软弹 硬弹 退',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `created_user_task` (`created_at`,`company_id`,`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mail_1`
--

DROP TABLE IF EXISTS `mail_1`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail_1` (
  `id` int(11) NOT NULL,
  `company_id` int(11) NOT NULL,
  `task_id` int(11) NOT NULL,
  `contact_info` varchar(1000) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '打开 点击 软弹 硬弹 退',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `created_user_task` (`created_at`,`company_id`,`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mail_2`
--

DROP TABLE IF EXISTS `mail_2`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail_2` (
  `id` int(11) NOT NULL,
  `company_id` int(11) NOT NULL,
  `task_id` int(11) NOT NULL,
  `contact_info` varchar(1000) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '打开 点击 软弹 硬弹 退',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `created_user_task` (`created_at`,`company_id`,`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mail_3`
--

DROP TABLE IF EXISTS `mail_3`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail_3` (
  `id` int(11) NOT NULL,
  `company_id` int(11) NOT NULL,
  `task_id` int(11) NOT NULL,
  `contact_info` varchar(1000) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '打开 点击 软弹 硬弹 退',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `created_user_task` (`created_at`,`company_id`,`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mail_4`
--

DROP TABLE IF EXISTS `mail_4`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail_4` (
  `id` int(11) NOT NULL,
  `company_id` int(11) NOT NULL,
  `task_id` int(11) NOT NULL,
  `contact_info` varchar(1000) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '打开 点击 软弹 硬弹 退',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `created_user_task` (`created_at`,`company_id`,`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mail_5`
--

DROP TABLE IF EXISTS `mail_5`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail_5` (
  `id` int(11) NOT NULL,
  `company_id` int(11) NOT NULL,
  `task_id` int(11) NOT NULL,
  `contact_info` varchar(1000) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '打开 点击 软弹 硬弹 退',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `created_user_task` (`created_at`,`company_id`,`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mail_6`
--

DROP TABLE IF EXISTS `mail_6`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail_6` (
  `id` int(11) NOT NULL,
  `company_id` int(11) NOT NULL,
  `task_id` int(11) NOT NULL,
  `contact_info` varchar(1000) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '打开 点击 软弹 硬弹 退',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `created_user_task` (`created_at`,`company_id`,`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mail_7`
--

DROP TABLE IF EXISTS `mail_7`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail_7` (
  `id` int(11) NOT NULL,
  `company_id` int(11) NOT NULL,
  `task_id` int(11) NOT NULL,
  `contact_info` varchar(1000) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '打开 点击 软弹 硬弹 退',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `created_user_task` (`created_at`,`company_id`,`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mail_allot`
--

DROP TABLE IF EXISTS `mail_allot`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail_allot` (
  `id` int(11) NOT NULL,
  `version` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-10-23 18:54:37
