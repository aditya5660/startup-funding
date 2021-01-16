# ************************************************************
# Sequel Ace SQL dump
# Version 2077
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# Host: 127.0.0.1 (MySQL 5.7.23)
# Database: startup_funding_v1_db
# Generation Time: 2021-01-16 14:42:40 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table campaign_images
# ------------------------------------------------------------

DROP TABLE IF EXISTS `campaign_images`;

CREATE TABLE `campaign_images` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `campaign_id` int(11) DEFAULT NULL,
  `file_name` varchar(100) DEFAULT NULL,
  `is_primary` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `campaign_images` WRITE;
/*!40000 ALTER TABLE `campaign_images` DISABLE KEYS */;

INSERT INTO `campaign_images` (`id`, `campaign_id`, `file_name`, `is_primary`, `created_at`, `updated_at`)
VALUES
	(1,1,'storage/images/campaign-1.jpg',1,'2020-11-29 00:40:39','2020-11-29 00:40:39'),
	(2,2,'storage/images/campaign-2.jpg',1,'2020-11-29 00:40:39','2020-11-29 00:40:39'),
	(3,3,'storage/images/campaign-3.jpg',1,'2020-11-29 00:40:39','2020-11-29 00:40:39');

/*!40000 ALTER TABLE `campaign_images` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table campaigns
# ------------------------------------------------------------

DROP TABLE IF EXISTS `campaigns`;

CREATE TABLE `campaigns` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL DEFAULT '',
  `short_description` varchar(100) DEFAULT NULL,
  `description` text,
  `perks` varchar(255) DEFAULT NULL,
  `backer_count` int(20) DEFAULT NULL,
  `goal_amount` int(20) DEFAULT NULL,
  `current_amount` int(20) DEFAULT NULL,
  `slug` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `campaigns` WRITE;
/*!40000 ALTER TABLE `campaigns` DISABLE KEYS */;

INSERT INTO `campaigns` (`id`, `user_id`, `name`, `short_description`, `description`, `perks`, `backer_count`, `goal_amount`, `current_amount`, `slug`, `created_at`, `updated_at`)
VALUES
	(1,1,'Campain 1','asdasd','asdasdasdasdasdasdasd','asdasd,asdasd,asdasd,asdasd,',0,100000000,0,'campign-1','2020-11-29 00:29:55','2020-11-29 00:30:01'),
	(2,2,'Campain 2','asdasd','asdasdasdasdasdasdasd','asdasd,asdasd,asdasd,asdasd,',0,100000000,0,'campign-1','2020-11-29 00:29:55','2020-11-29 00:30:01'),
	(3,6,'Campain 3','asdasd','asdasdasdasdasdasdasd','asdasd,asdasd,asdasd,asdasd,',0,100000000,0,'campign-1','2020-11-29 00:29:55','2020-11-29 00:30:01');

/*!40000 ALTER TABLE `campaigns` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `occupation` varchar(100) DEFAULT NULL,
  `email` varchar(155) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `role` varchar(100) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password`, `token`, `role`, `avatar`, `created_at`, `updated_at`)
VALUES
	(1,'Aditya Putra','Back End Engineer','aditya5660@gmail.com','$2a$04$fEz3B.YotClXfWEj5Z39HedI..rlM8RfpdegPfBMy/Gdrknj0byOS','semuabisa','user','adit.png','2020-11-10 00:53:27',NULL),
	(2,'Alvin Ika','Mobile Engineer','alvinikaa@gmail.com','$2a$04$fEz3B.YotClXfWEj5Z39HedI..rlM8RfpdegPfBMy/Gdrknj0byOS','semuabisa','user','alvin.png','2020-11-10 00:53:27',NULL),
	(3,'Codeinaja','CEO','codeinaja.dev@gmail.com','$2a$04$fEz3B.YotClXfWEj5Z39HedI..rlM8RfpdegPfBMy/Gdrknj0byOS','semuabisa','admin','public/images/3-Screenshot at Oct 08 13-40-20.png','2020-11-10 00:53:27','2020-11-13 21:04:58'),
	(5,'Test Dari Servis','Project Manager','admin@example.com','$2a$04$ys4HLtUfpwu6rPKk7krtneWa98wDm5vNgiHICPiGeoW5YlVVBCleO',NULL,'user','','2020-11-10 00:53:27','2020-11-10 00:53:27'),
	(6,'Andika','CTO','andika@gmail.com','$2a$04$fEz3B.YotClXfWEj5Z39HedI..rlM8RfpdegPfBMy/Gdrknj0byOS',NULL,'user','public/images/6-20201030_165551.jpg','2020-11-10 01:11:23','2020-11-11 01:09:22');

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
