DROP DATABASE IF EXISTS `ccnu_course_db`;

CREATE DATABASE `ccnu_course_db`;

USE `ccnu_course_db`;

CREATE TABLE `class` (
  `id`           INT unsigned NOT NULL AUTO_INCREMENT,
  `name`         VARCHAR(40)  NOT NULL,
  `academy`      VARCHAR(25)  NOT NULL,
  `course_code`  VARCHAR(8)   NOT NULL,
  `class_code`   VARCHAR(25)  NOT NULL,
  `cap`          VARCHAR(3)   NOT NULL,
  `teaching_way` VARCHAR(8)   NOT NULL,
  `teachers`     VARCHAR(150) NOT NULL,
  `duty`         VARCHAR(15)  NOT NULL,
  `time1`        VARCHAR(20)  NOT NULL,
  `place1`       VARCHAR(15)  NOT NULL,
  `time2`        VARCHAR(20)  NOT NULL,
  `place2`       VARCHAR(15)  NOT NULL,
  `time3`        VARCHAR(20)  NOT NULL,
  `place3`       VARCHAR(15)  NOT NULL,

  PRIMARY KEY (`id`),
  UNIQUE  KEY `code` (`course_code`, `class_code`),
  FULLTEXT KEY (`name`, `course_code`, `teachers`) WITH PARSER ngram
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `course` (
  `id`          INT unsigned NOT NULL AUTO_INCREMENT,
  `name`        VARCHAR(40)  NOT NULL,  
  `course_code` VARCHAR(8)   NOT NULL,
  `credit`      VARCHAR(5)   NOT NULL,

  PRIMARY  KEY (`id`),
  UNIQUE   KEY `code` (`course_code`),
  FULLTEXT KEY `text` (`name`, `course_code`) WITH PARSER ngram
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;