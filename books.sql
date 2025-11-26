/*
 Navicat Premium Dump SQL

 Source Server         : godb
 Source Server Type    : MySQL
 Source Server Version : 80407 (8.4.7)
 Source Host           : localhost:3306
 Source Schema         : mysql

 Target Server Type    : MySQL
 Target Server Version : 80407 (8.4.7)
 File Encoding         : 65001

 Date: 26/11/2025 20:15:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS `books`;
CREATE TABLE `books`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `author` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `price` double(11, 2) NOT NULL,
  `sales` int NOT NULL,
  `stock` int NOT NULL,
  `img_path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 46 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of books
-- ----------------------------
INSERT INTO `books` VALUES (3, '中国哲学史', '冯友兰', 44.50, 100, 1001, '/static/images/default.jpg');
INSERT INTO `books` VALUES (5, '苏东坡传', '林语堂', 19.30, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (6, '百年孤独2', '马尔克斯', 29.50, 101, 99, '/static/images/default.jpg');
INSERT INTO `books` VALUES (8, '给孩子的诗', '北岛', 22.20, 104, 96, 'static/img/default.jpg');
INSERT INTO `books` VALUES (9, '为奴十二年', '所罗门', 16.50, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (10, '平凡的世界', '路遥', 55.00, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (11, '悟空传', '今何在', 14.00, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (12, '硬派健身', '斌卡', 31.20, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (13, '从晚清到民国', '唐德刚', 39.90, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (14, '三体', '刘慈欣', 56.50, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (15, '看见', '柴静', 19.50, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (16, '活着', '余华', 11.00, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (17, '小王子', '安托万', 19.20, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (18, '我们仨', '杨绛', 11.30, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (19, '生命不息,折腾不止', '罗永浩', 25.20, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (20, '皮囊', '蔡崇达', 23.90, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (21, '恰到好处的幸福', '毕淑敏', 16.40, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (22, '测试图书12', '测试', 88.88, 200, 100, '/static/images/default.jpg');
INSERT INTO `books` VALUES (23, '人月神话', '布鲁克斯', 55.90, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (24, 'C语言入门经典1', '霍尔顿', 45.00, 100, 100, '/static/images/default.jpg');
INSERT INTO `books` VALUES (25, '数学之美', '吴军', 29.90, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (26, 'Java编程思想', '埃史尔', 70.50, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (27, '设计模式之禅', '秦小波', 20.20, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (28, '图解机器学习', '杉山将', 33.80, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (44, '1', '1', 41.00, 2, 0, '/static/images/default.jpg');
INSERT INTO `books` VALUES (45, '1', '1', 1.00, 2, 0, '/static/images/default.jpg');

-- ----------------------------
-- Table structure for cart_items
-- ----------------------------
DROP TABLE IF EXISTS `cart_items`;
CREATE TABLE `cart_items`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `COUNT` int NOT NULL,
  `amount` double(11, 2) NOT NULL,
  `book_id` int NOT NULL,
  `cart_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `book_id`(`book_id` ASC) USING BTREE,
  INDEX `cart_id`(`cart_id` ASC) USING BTREE,
  CONSTRAINT `cart_items_ibfk_1` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `cart_items_ibfk_2` FOREIGN KEY (`cart_id`) REFERENCES `carts` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 40 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cart_items
-- ----------------------------
INSERT INTO `cart_items` VALUES (7, 1, 44.50, 3, '5cf07515-9473-4c3f-987e-022eea613bb8');
INSERT INTO `cart_items` VALUES (18, 1, 16.50, 9, '5cf07515-9473-4c3f-987e-022eea613bb8');
INSERT INTO `cart_items` VALUES (19, 2, 39.90, 13, '5cf07515-9473-4c3f-987e-022eea613bb8');
INSERT INTO `cart_items` VALUES (20, 1, 56.50, 14, '5cf07515-9473-4c3f-987e-022eea613bb8');
INSERT INTO `cart_items` VALUES (21, 3, 11.00, 16, '5cf07515-9473-4c3f-987e-022eea613bb8');
INSERT INTO `cart_items` VALUES (22, 3, 22.20, 8, '5cf07515-9473-4c3f-987e-022eea613bb8');
INSERT INTO `cart_items` VALUES (23, 1, 14.00, 11, '5cf07515-9473-4c3f-987e-022eea613bb8');
INSERT INTO `cart_items` VALUES (24, 1, 31.20, 12, '5cf07515-9473-4c3f-987e-022eea613bb8');
INSERT INTO `cart_items` VALUES (25, 2, 29.50, 6, '5cf07515-9473-4c3f-987e-022eea613bb8');
INSERT INTO `cart_items` VALUES (36, 1, 44.50, 3, '348428ef-1ddf-4ba3-81b3-f4547ac6f58b');
INSERT INTO `cart_items` VALUES (37, 1, 19.30, 5, '348428ef-1ddf-4ba3-81b3-f4547ac6f58b');

-- ----------------------------
-- Table structure for carts
-- ----------------------------
DROP TABLE IF EXISTS `carts`;
CREATE TABLE `carts`  (
  `id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `total_count` int NOT NULL,
  `total_amount` double(11, 2) NOT NULL,
  `user_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  CONSTRAINT `carts_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of carts
-- ----------------------------
INSERT INTO `carts` VALUES ('348428ef-1ddf-4ba3-81b3-f4547ac6f58b', 2, 63.80, 17);
INSERT INTO `carts` VALUES ('5cf07515-9473-4c3f-987e-022eea613bb8', 15, 401.10, 15);

-- ----------------------------
-- Table structure for order_items
-- ----------------------------
DROP TABLE IF EXISTS `order_items`;
CREATE TABLE `order_items`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `count` int NOT NULL,
  `amount` double(11, 2) NOT NULL,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `author` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `price` double(11, 2) NOT NULL,
  `img_path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `order_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `order_id`(`order_id` ASC) USING BTREE,
  CONSTRAINT `order_items_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order_items
-- ----------------------------
INSERT INTO `order_items` VALUES (1, 1, 200.00, '测试书1', '测试1', 200.00, '/static/img/default.jpg', '6688');
INSERT INTO `order_items` VALUES (2, 1, 100.00, '测试书2', '测试2', 100.00, '/static/img/default.jpg', '6688');
INSERT INTO `order_items` VALUES (3, 3, 66.60, '给孩子的诗', '北岛', 22.20, 'static/img/default.jpg', '081b94de-b6d7-4a84-835a-4e0ecf1440a9');
INSERT INTO `order_items` VALUES (4, 1, 1.00, '1', '1', 1.00, '/static/images/default.jpg', '6a4a4290-16fb-4051-acc3-0985f0cf3c5b');
INSERT INTO `order_items` VALUES (5, 1, 41.00, '1', '1', 41.00, '/static/images/default.jpg', 'c72e2a14-6a81-417e-b7d5-d864bd1a0aa9');
INSERT INTO `order_items` VALUES (6, 1, 29.50, '百年孤独2', '马尔克斯', 29.50, '/static/images/default.jpg', '2397ec05-6cfb-4cc4-969d-668c9368c9d6');
INSERT INTO `order_items` VALUES (7, 1, 22.20, '给孩子的诗', '北岛', 22.20, 'static/img/default.jpg', '2397ec05-6cfb-4cc4-969d-668c9368c9d6');

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`  (
  `id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `create_time` datetime NOT NULL,
  `total_count` int NOT NULL,
  `total_amount` double(11, 2) NOT NULL,
  `state` int NOT NULL,
  `user_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of orders
-- ----------------------------
INSERT INTO `orders` VALUES ('081b94de-b6d7-4a84-835a-4e0ecf1440a9', '2025-11-03 03:25:22', 3, 66.60, 2, 17);
INSERT INTO `orders` VALUES ('2397ec05-6cfb-4cc4-969d-668c9368c9d6', '2025-11-26 11:49:42', 2, 51.70, 1, 57);
INSERT INTO `orders` VALUES ('6688', '2025-11-03 02:03:35', 2, 300.00, 1, 25);
INSERT INTO `orders` VALUES ('6a4a4290-16fb-4051-acc3-0985f0cf3c5b', '2025-11-03 03:29:58', 1, 1.00, 2, 17);
INSERT INTO `orders` VALUES ('c72e2a14-6a81-417e-b7d5-d864bd1a0aa9', '2025-11-03 03:31:30', 1, 41.00, 1, 17);

-- ----------------------------
-- Table structure for sessions
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions`  (
  `session_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_id` int NOT NULL,
  PRIMARY KEY (`session_id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  CONSTRAINT `sessions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sessions
-- ----------------------------
INSERT INTO `sessions` VALUES ('3fbd9ec6-db3c-46eb-9afb-00ef9c29dc4b', 'qqq', 57);
INSERT INTO `sessions` VALUES ('f0c50a95-8fd4-42e8-9684-fb3e73f5ee63', 'admin66', 17);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 58 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (15, 'admin', '123', 'admin@qq.com');
INSERT INTO `users` VALUES (16, 'admin2', '666', 'admin666@qq.com');
INSERT INTO `users` VALUES (17, 'admin66', 'admin77', 'admin88');
INSERT INTO `users` VALUES (20, 'admin661', 'admin77', 'admin88');
INSERT INTO `users` VALUES (23, '123', '123456', '1@qq.com');
INSERT INTO `users` VALUES (25, 'ccc', 'cccccc', 'cc@qq.com');
INSERT INTO `users` VALUES (51, '12342', '123123', 'qq@qq.com');
INSERT INTO `users` VALUES (54, '1233', '123123', '123@qq.colm');
INSERT INTO `users` VALUES (57, 'qqq', '201206', 'qqq@qq.com');

SET FOREIGN_KEY_CHECKS = 1;
