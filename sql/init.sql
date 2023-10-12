CREATE DATABASE IF NOT EXISTS dev_db;
-- CREATE DATABASE IF NOT EXISTS test;
USE dev_db;

CREATE TABLE IF NOT EXISTS `products` (
  `id` int(11) NOT NULL auto_increment,
  `nome` varchar(250)  NOT NULL default '',
  `gtin`  varchar(13) NOT NULL default '',
  `inclusao` datetime not null default current_timestamp,
   PRIMARY KEY  (`id`)
);