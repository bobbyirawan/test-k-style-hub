-- MySQL dump 10.13  Distrib 8.0.32, for Linux (aarch64)
--
-- Host: localhost    Database: k_style_hub_test
-- ------------------------------------------------------
-- Server version	8.0.32

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `like_reviews`
--

DROP TABLE IF EXISTS `like_reviews`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `like_reviews` (
  `review_id` bigint DEFAULT NULL,
  `member_id` bigint DEFAULT NULL,
  `like_review_id` bigint NOT NULL AUTO_INCREMENT,
  UNIQUE KEY `like_review_id` (`like_review_id`),
  UNIQUE KEY `idx_unique` (`like_review_id`),
  KEY `fk_review_products_list_like_review` (`review_id`),
  KEY `fk_members_list_like_review` (`member_id`),
  CONSTRAINT `fk_members_list_like_review` FOREIGN KEY (`member_id`) REFERENCES `members` (`member_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_review_products_list_like_review` FOREIGN KEY (`review_id`) REFERENCES `review_products` (`review_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `like_reviews`
--

LOCK TABLES `like_reviews` WRITE;
/*!40000 ALTER TABLE `like_reviews` DISABLE KEYS */;
INSERT INTO `like_reviews` VALUES (1,1,1),(2,1,2),(3,1,3),(4,1,4),(1,2,5),(2,2,6),(3,2,7),(4,2,8);
/*!40000 ALTER TABLE `like_reviews` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `members`
--

DROP TABLE IF EXISTS `members`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `members` (
  `member_id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(100) DEFAULT NULL,
  `gender` varchar(255) DEFAULT NULL,
  `skin_type` varchar(50) DEFAULT NULL,
  `skin_color` varchar(50) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`member_id`),
  UNIQUE KEY `idx_unique` (`member_id`,`username`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `members`
--

LOCK TABLES `members` WRITE;
/*!40000 ALTER TABLE `members` DISABLE KEYS */;
INSERT INTO `members` VALUES (1,'kuma','female','skinny','brown','2023-03-27 08:38:05.821','2023-03-27 08:38:05.821',NULL),(2,'jess','female','skinny','brown','2023-03-27 08:41:42.028','2023-03-27 08:42:00.265',NULL),(3,'kuma town','female','skinny','brown','2023-03-27 08:59:41.481','2023-03-27 08:59:41.481',NULL),(4,'kuma lay','female','skinny','brown','2023-03-27 09:03:06.600','2023-03-28 06:37:32.839',NULL);
/*!40000 ALTER TABLE `members` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `product_id` bigint NOT NULL AUTO_INCREMENT,
  `product_name` varchar(100) DEFAULT NULL,
  `price` double DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`product_id`),
  UNIQUE KEY `idx_unique` (`product_id`,`product_name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'pemutih',30000,'2023-03-27 08:38:05.821','2023-03-27 08:38:05.821',NULL),(2,'skin care',15000,'2023-03-27 08:59:05.821','2023-03-27 08:59:05.821',NULL);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `review_products`
--

DROP TABLE IF EXISTS `review_products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `review_products` (
  `review_id` bigint DEFAULT NULL,
  `product_id` bigint DEFAULT NULL,
  `member_id` bigint DEFAULT NULL,
  `review_desc` varchar(500) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  UNIQUE KEY `idx_unique` (`review_id`),
  KEY `fk_members_list_review` (`member_id`),
  KEY `fk_products_list_review` (`product_id`),
  CONSTRAINT `fk_members_list_review` FOREIGN KEY (`member_id`) REFERENCES `members` (`member_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_products_list_review` FOREIGN KEY (`product_id`) REFERENCES `products` (`product_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `review_products`
--

LOCK TABLES `review_products` WRITE;
/*!40000 ALTER TABLE `review_products` DISABLE KEYS */;
INSERT INTO `review_products` VALUES (1,1,1,'good','2023-03-27 08:40:05.821','2023-03-27 08:40:05.821',NULL),(2,2,2,'good','2023-03-27 08:44:05.821','2023-03-27 08:44:05.821',NULL),(3,1,3,'good','2023-03-27 08:59:05.821','2023-03-27 08:59:05.821',NULL),(4,2,4,'good','2023-03-27 08:59:05.821','2023-03-27 08:59:05.821',NULL),(5,2,1,NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `review_products` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-03-28  0:03:14
