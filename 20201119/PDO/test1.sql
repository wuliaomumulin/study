/*
Navicat MySQL Data Transfer

Source Server         : 182.92.148.54
Source Server Version : 50552
Source Host           : 182.92.148.54:3306
Source Database       : plana

Target Server Type    : MYSQL
Target Server Version : 50552
File Encoding         : 65001

Date: 2017-11-29 17:05:00
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for test1
-- ----------------------------
DROP TABLE IF EXISTS `test1`;
CREATE TABLE `test1` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` char(20) DEFAULT NULL,
  `sex` tinyint(1) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of test1
-- ----------------------------
INSERT INTO `test1` VALUES ('1', '张佳佳', '1', '大兴区');
INSERT INTO `test1` VALUES ('2', '王结盟', '0', '官庄');
INSERT INTO `test1` VALUES ('3', '刘铁锁', '1', '东二旗');
INSERT INTO `test1` VALUES ('8', '张亚欣', '0', '北京市');
INSERT INTO `test1` VALUES ('9', '张薪薪', '0', '物资学院');
