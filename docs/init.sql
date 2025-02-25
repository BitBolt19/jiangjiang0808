# ************************************************************
# Sequel Pro SQL dump
# Version 5446
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 8.0.32)
# Database: nez
# Generation Time: 2025-02-25 07:56:47 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table account_token_balance
# ------------------------------------------------------------

DROP TABLE IF EXISTS `account_token_balance`;

CREATE TABLE `account_token_balance` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账户',
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '代币',
  `balance` decimal(30,18) unsigned NOT NULL COMMENT '余额',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_account_token` (`account`,`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# Dump of table account_token_balance_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `account_token_balance_log`;

CREATE TABLE `account_token_balance_log` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int unsigned NOT NULL COMMENT '关联订单ID',
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账户',
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '代币',
  `amount` decimal(30,18) unsigned NOT NULL COMMENT '数量',
  `old_balance` decimal(30,18) unsigned NOT NULL COMMENT '旧余额',
  `new_balance` decimal(30,18) unsigned NOT NULL COMMENT '新余额',
  `type` int unsigned NOT NULL COMMENT '流水类型',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# Dump of table nft_box
# ------------------------------------------------------------

DROP TABLE IF EXISTS `nft_box`;

CREATE TABLE `nft_box` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账户',
  `contract_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'NFT合约地址',
  `token_id` int unsigned NOT NULL COMMENT 'NFT tokenId',
  `tx_hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '交易hash',
  `tx_time` datetime NOT NULL COMMENT '开盲盒时间',
  `tx_event_id` int unsigned DEFAULT NULL COMMENT '交易事件id',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# Dump of table nft_hold
# ------------------------------------------------------------

DROP TABLE IF EXISTS `nft_hold`;

CREATE TABLE `nft_hold` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账户',
  `contract_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'NFT合约地址',
  `token_id` int unsigned NOT NULL COMMENT 'NFT tokenId',
  `hold_time` datetime NOT NULL COMMENT '持有时间',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# Dump of table nft_staking
# ------------------------------------------------------------

DROP TABLE IF EXISTS `nft_staking`;

CREATE TABLE `nft_staking` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账户',
  `contract_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'NFT合约地址',
  `token_id` int unsigned NOT NULL COMMENT 'NFT tokenId',
  `tx_hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `tx_time` datetime NOT NULL,
  `tx_event_id` int unsigned NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# Dump of table nft_staking_claim
# ------------------------------------------------------------

DROP TABLE IF EXISTS `nft_staking_claim`;

CREATE TABLE `nft_staking_claim` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账户',
  `contract_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'NFT合约地址',
  `token_id` int unsigned NOT NULL COMMENT 'NFT tokenId',
  `amount` decimal(30,18) unsigned NOT NULL COMMENT '领取数量',
  `tx_hash` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `tx_time` datetime NOT NULL,
  `tx_event_id` int unsigned NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# Dump of table nft_transfer
# ------------------------------------------------------------

DROP TABLE IF EXISTS `nft_transfer`;

CREATE TABLE `nft_transfer` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账户',
  `contract_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'NFT合约地址',
  `from_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '发送方',
  `to_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '接收方',
  `token_id` int unsigned NOT NULL COMMENT 'NFT tokenId',
  `tx_hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '交易hash',
  `tx_event_id` int unsigned NOT NULL COMMENT '交替事件id',
  `tx_time` datetime NOT NULL COMMENT '交易时间',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# Dump of table reward_pool_claim
# ------------------------------------------------------------

DROP TABLE IF EXISTS `reward_pool_claim`;

CREATE TABLE `reward_pool_claim` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账户',
  `pool_id` int unsigned NOT NULL COMMENT '轮次',
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '代币合约地址',
  `amount` decimal(30,18) unsigned NOT NULL COMMENT '数量',
  `tx_hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `tx_time` datetime NOT NULL,
  `tx_event_id` int unsigned NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# Dump of table reward_pool_join
# ------------------------------------------------------------

DROP TABLE IF EXISTS `reward_pool_join`;

CREATE TABLE `reward_pool_join` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账户',
  `pool_id` int unsigned NOT NULL COMMENT '轮次',
  `tx_hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `tx_time` datetime NOT NULL,
  `tx_event_id` int unsigned NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# Dump of table reward_pool_settlement
# ------------------------------------------------------------

DROP TABLE IF EXISTS `reward_pool_settlement`;

CREATE TABLE `reward_pool_settlement` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `pool_id` int unsigned NOT NULL COMMENT '轮次',
  `owners` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '中奖人列表',
  `tx_hash` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `tx_time` datetime NOT NULL,
  `tx_event_id` int NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
