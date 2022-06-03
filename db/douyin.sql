/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : douyin_utf8

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 01/06/2022 23:52:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment_praises
-- ----------------------------
DROP TABLE IF EXISTS `comment_praises`;
CREATE TABLE `comment_praises`  (
  `praise_user_id` int(0) UNSIGNED NOT NULL COMMENT '用户id',
  `praise_comment_id` int(0) UNSIGNED NOT NULL COMMENT '点赞评论id',
  `praise_ispraised` tinyint(1) NOT NULL,
  PRIMARY KEY (`praise_user_id`, `praise_comment_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `comment_id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `comment_user_id` int(0) UNSIGNED NOT NULL COMMENT '用户id',
  `comment_video_id` int(0) UNSIGNED NOT NULL COMMENT '视频id',
  `comment_content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '评论信息',
  `comment_latest_time` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '评论最新修改时间',
  `comment_returnid` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '所回复的评论id（可为null）',
  PRIMARY KEY (`comment_id`) USING BTREE,
  INDEX `comment_video_id`(`comment_video_id`) USING BTREE,
  INDEX `comment_returnid`(`comment_returnid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for contributions
-- ----------------------------
DROP TABLE IF EXISTS `contributions`;
CREATE TABLE `contributions`  (
  `user_id` int(0) UNSIGNED NOT NULL COMMENT '用户id',
  `video_id` int(0) UNSIGNED NOT NULL COMMENT '视频id',
  PRIMARY KEY (`user_id`, `video_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for favorites
-- ----------------------------
DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites`  (
  `favorite_user_id` int(0) UNSIGNED NOT NULL COMMENT '所关注用户id',
  `favorite_fan_id` int(0) UNSIGNED NOT NULL COMMENT '粉丝id',
  PRIMARY KEY (`favorite_user_id`, `favorite_fan_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `user_id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
  `user_password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户密码',
  PRIMARY KEY (`user_id`) USING BTREE,
  INDEX `user_name`(`user_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `video_id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '视频id',
  `video_location` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频文件路径',
  `video_picture_location` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频图片路径',
  `video_latest_time` datetime(0) NOT NULL COMMENT '视频修改时间',
  `video_state` int(0) NOT NULL COMMENT '视频状态',
  `video_category` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频分类id集合（使用”;“隔开）',
  `video_introduction` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '视频简介',
  `video_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频标题',
  `video_authorid` int(0) NOT NULL COMMENT '作者id',
  PRIMARY KEY (`video_id`) USING BTREE,
  FULLTEXT INDEX `video_title`(`video_title`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for view_counts
-- ----------------------------
DROP TABLE IF EXISTS `view_counts`;
CREATE TABLE `view_counts`  (
  `video_id` int(0) UNSIGNED NOT NULL COMMENT '视频id',
  `video_counts` int(0) NOT NULL COMMENT '视频播放量',
  PRIMARY KEY (`video_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for watch_praises
-- ----------------------------
DROP TABLE IF EXISTS `watch_praises`;
CREATE TABLE `watch_praises`  (
  `praise_user_id` int(0) UNSIGNED NOT NULL COMMENT '用户id',
  `praise_video_id` int(0) UNSIGNED NOT NULL COMMENT '点赞视频id',
  `praise_ispraised` tinyint(1) NOT NULL,
  PRIMARY KEY (`praise_user_id`, `praise_video_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
