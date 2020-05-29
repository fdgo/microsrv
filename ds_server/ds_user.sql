/*
 Navicat Premium Data Transfer

 Source Server         : rm-t4nwa526sm5x7s1kveo.mysqlex.singapore.rds.aliyuncs.com
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : rm-t4nwa526sm5x7s1kveo.mysqlex.singapore.rds.aliyuncs.com:3306
 Source Schema         : ds

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 08/05/2020 18:02:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ds_sys_info
-- ----------------------------
DROP TABLE IF EXISTS `ds_sys_info`;
CREATE TABLE `ds_sys_info`  (
  `connect_us` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系我们'
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ds_sys_info
-- ----------------------------
INSERT INTO `ds_sys_info` VALUES ('13666666666');

-- ----------------------------
-- Table structure for ds_user_agent_class
-- ----------------------------
DROP TABLE IF EXISTS `ds_user_agent_class`;
CREATE TABLE `ds_user_agent_class`  (
  `agent_money` decimal(18, 2) NOT NULL COMMENT '合伙人等级金额界限',
  `agent_tag` tinyint(4) NOT NULL COMMENT '合伙人等级1,2,3,4,5',
  `agent_tagex` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '合伙人等级标记  D1，D2，D3，D4，D5',
  `agent_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '合伙人等级名称  代理合伙人，高级合伙人，城市合伙人，区域合伙人，全球合伙人',
  INDEX `agent_money`(`agent_money`) USING BTREE,
  INDEX `agent_tag`(`agent_tag`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ds_user_agent_class
-- ----------------------------
INSERT INTO `ds_user_agent_class` VALUES (30000.00, 1, 'D1', '代理合伙人');
INSERT INTO `ds_user_agent_class` VALUES (90000.00, 2, 'D2', '高级合伙人');
INSERT INTO `ds_user_agent_class` VALUES (270000.00, 3, 'D3', '城市合伙人');
INSERT INTO `ds_user_agent_class` VALUES (810000.00, 4, 'D4', '区域合伙人');
INSERT INTO `ds_user_agent_class` VALUES (2430000.00, 5, 'D5', '全球合伙人');

-- ----------------------------
-- Table structure for ds_user_basicinfo
-- ----------------------------
DROP TABLE IF EXISTS `ds_user_basicinfo`;
CREATE TABLE `ds_user_basicinfo`  (
  `uuid` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户ID号',
  `mobile` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `salt` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码盐',
  `hash` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码hash',
  `last_login_time` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '最后一次登录时间',
  `last_login_ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后一次登录ip',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `deleted` tinyint(4) NOT NULL COMMENT '是否删除0:未删除 1:删除',
  `real_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '真实姓名',
  `age` tinyint(3) UNSIGNED ZEROFILL NOT NULL DEFAULT 000 COMMENT '年龄',
  `gender` tinyint(3) UNSIGNED ZEROFILL NOT NULL DEFAULT 000 COMMENT '性别：0:男，1:女',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像地址',
  `nick_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `birthday` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生日',
  `status` tinyint(3) UNSIGNED ZEROFILL NOT NULL DEFAULT 000 COMMENT '账户状态0:正常 1:禁用  2:注销',
  PRIMARY KEY (`uuid`) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ds_user_basicinfo
-- ----------------------------
INSERT INTO `ds_user_basicinfo` VALUES ('5gKJW6gn6J', '13999999999', 'DwlruRWV8WRLE3Gv', 'lGXKZ79qvXJFQFNFWFRFAoRB2djz4p', '2020-05-05 02:16:23', '127.0.0.1:40842', '2020-05-05 10:14:12', '2020-05-05 02:16:23', 0, '', 000, 000, '', 'DS974757', '2020-05-05 10:14:12', 000);
INSERT INTO `ds_user_basicinfo` VALUES ('7DQZOEubNK', '13121391353', 'T9iB7wbTab1cBpFq', 'N386lK9xWGjsxUOcPH2FJ0mQV1AjY5', '2020-05-08 13:53:53', '183.14.133.72:39117', '2020-05-08 13:53:53', '2020-05-08 13:53:53', 0, '', 000, 000, '', 'DS953382', '2020-05-08 13:53:53', 000);
INSERT INTO `ds_user_basicinfo` VALUES ('HbfH7AMsYa', '13510385413', 'ad3PAf9j6yookTma', 'kryboVl65mVIbTRF6H0uJXpGZKjEWM', '2020-05-08 17:09:18', '183.14.133.72:38625', '2020-05-07 19:21:31', '2020-05-08 17:09:18', 0, '', 000, 000, '', 'DS236918', '2020-05-07 19:21:31', 000);
INSERT INTO `ds_user_basicinfo` VALUES ('ynbGOz4QVT', '17688566605', 'L8ykgBn7dfy14cVq', '', '2020-05-08 14:12:07', '183.14.133.72:39055', '2020-05-08 14:12:07', '2020-05-08 14:12:07', 0, '', 000, 000, '', 'DS882368', '2020-05-08 14:12:07', 000);

-- ----------------------------
-- Table structure for ds_user_member_account
-- ----------------------------
DROP TABLE IF EXISTS `ds_user_member_account`;
CREATE TABLE `ds_user_member_account`  (
  `uuid` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户ID',
  `mobile` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户手机号',
  `balance` decimal(18, 2) NOT NULL DEFAULT 0.00 COMMENT '总金额',
  `private_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `salt` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '支付盐',
  `hash` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码hash',
  `address_in` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收款地址',
  `address_out` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '付款地址',
  `status` tinyint(3) UNSIGNED ZEROFILL NOT NULL DEFAULT 000 COMMENT '账户状态0:正常 1:禁用  2:注销',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `ispwd` tinyint(3) UNSIGNED ZEROFILL NOT NULL COMMENT '密码是否为空',
  PRIMARY KEY (`uuid`) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ds_user_member_account
-- ----------------------------
INSERT INTO `ds_user_member_account` VALUES ('5gKJW6gn6J', '13999999999', 0.00, '', 'DwlruRWV8WRLE3Gv', 'lGXKZ79qvXJFQFNFWFRFAoRB2djz4p', '', '', 000, '2020-05-05 10:14:12', '2020-05-05 02:16:14', 000);
INSERT INTO `ds_user_member_account` VALUES ('7DQZOEubNK', '13121391353', 1133.21, '', 'aJdQ1JPXRBBxxdIy', 'AKlorpmDRwzsPs8sgsZs2W8J5O3QVq', '', '', 000, '2020-05-08 13:53:53', '2020-05-08 16:04:59', 000);
INSERT INTO `ds_user_member_account` VALUES ('HbfH7AMsYa', '13510385413', 14140038.80, '', 'EYK3kl3VgUc6ZR2U', 'vj01xgD8e8JHRSEcniVtK2kq3JVORd', '', '', 000, '2020-05-07 19:21:31', '2020-05-08 17:11:07', 001);
INSERT INTO `ds_user_member_account` VALUES ('ynbGOz4QVT', '17688566605', 0.00, '', '', '', '', '', 000, '2020-05-08 14:12:07', '2020-05-08 14:12:07', 000);

-- ----------------------------
-- Table structure for ds_user_member_agent
-- ----------------------------
DROP TABLE IF EXISTS `ds_user_member_agent`;
CREATE TABLE `ds_user_member_agent`  (
  `uuid_self` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '新注册的用户ID号',
  `mobile_self` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '新注册的用户手机号',
  `invcode_self` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '新注册的用户自身邀请码',
  `uuid_agent` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '代理的ID号',
  `mobile_agent` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '代理的手机号',
  `invcode_agent` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '代理的邀请码',
  `memclass_self` tinyint(3) UNSIGNED ZEROFILL NOT NULL DEFAULT 000 COMMENT '会员等级',
  `member_tag` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '会员标识',
  `member_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '会员名称',
  `agent_class` tinyint(3) UNSIGNED ZEROFILL NOT NULL DEFAULT 000 COMMENT '代理等级',
  `agent_tag` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '代理标识',
  `agent_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '代理名称',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  PRIMARY KEY (`uuid_self`) USING BTREE,
  UNIQUE INDEX `mobile_self`(`mobile_self`) USING BTREE,
  UNIQUE INDEX `invcode_self`(`invcode_self`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ds_user_member_agent
-- ----------------------------
INSERT INTO `ds_user_member_agent` VALUES ('5gKJW6gn6J', '13999999999', '123456', '5gKJW6gn6J', '13999999999', '123456', 000, '', '', 005, 'D5', '全球合伙人', '2020-05-08 13:56:54', '2020-05-08 13:56:54');
INSERT INTO `ds_user_member_agent` VALUES ('7DQZOEubNK', '13121391353', '530686', '5gKJW6gn6J', '13999999999', '123456', 001, 'V1', '普通卡', 000, '', '', '2020-05-08 13:56:54', '2020-05-08 13:56:54');
INSERT INTO `ds_user_member_agent` VALUES ('HbfH7AMsYa', '13510385413', '710650', '5gKJW6gn6J', '13999999999', '123456', 005, 'V5', '钻石卡', 000, '', '', '2020-05-08 10:55:49', '2020-05-08 10:55:49');
INSERT INTO `ds_user_member_agent` VALUES ('ynbGOz4QVT', '17688566605', '517368', '5gKJW6gn6J', '13999999999', '123456', 000, '', '', 000, '', '', '2020-05-08 14:12:07', '2020-05-08 14:12:07');

-- ----------------------------
-- Table structure for ds_user_member_class
-- ----------------------------
DROP TABLE IF EXISTS `ds_user_member_class`;
CREATE TABLE `ds_user_member_class`  (
  `mem_money` decimal(18, 2) NOT NULL COMMENT '会员等级金额界限',
  `mem_tag` tinyint(4) NOT NULL COMMENT '会员等级1,2,3,4,5',
  `mem_tagex` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '会员等级标记  M1，M2，M3，M4，M5',
  `mem_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '会员等级名称  普通卡，铜卡，银卡，金卡，钻石卡',
  INDEX `mem_money`(`mem_money`) USING BTREE,
  INDEX `mem_tag`(`mem_tag`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ds_user_member_class
-- ----------------------------
INSERT INTO `ds_user_member_class` VALUES (1000.00, 1, 'V1', '普通卡');
INSERT INTO `ds_user_member_class` VALUES (5000.00, 2, 'V2', '铜卡');
INSERT INTO `ds_user_member_class` VALUES (10000.00, 3, 'V3', '银卡');
INSERT INTO `ds_user_member_class` VALUES (30000.00, 4, 'V4', '金卡');
INSERT INTO `ds_user_member_class` VALUES (50000.00, 5, 'V5', '钻石卡');

-- ----------------------------
-- Table structure for ds_user_member_deposit_history
-- ----------------------------
DROP TABLE IF EXISTS `ds_user_member_deposit_history`;
CREATE TABLE `ds_user_member_deposit_history`  (
  `uuid` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户ID',
  `mobile` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户手机号',
  `source_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '业务订单Id',
  `balance` decimal(18, 2) NOT NULL DEFAULT 0.00 COMMENT '金额',
  `rate` decimal(18, 6) UNSIGNED DEFAULT 0.000000 COMMENT '汇率',
  `balance_src` decimal(18, 2) UNSIGNED DEFAULT 0.00 COMMENT '原始金额',
  `deposit_type` tinyint(3) UNSIGNED ZEROFILL NOT NULL DEFAULT 000 COMMENT '充值类型 0:扣款,1:充值',
  `deposit_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '充值名字：购买商品，商品退款',
  `address_in` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收款地址',
  `address_out` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '支付地址',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `status` tinyint(4) UNSIGNED ZEROFILL NOT NULL DEFAULT 0000 COMMENT '会员账户状态 0:正常，1:禁止 2:销户',
  `invcode_self` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '自身邀请码',
  `invcode_agent` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '代理邀请码',
  INDEX `source_id`(`source_id`) USING BTREE,
  INDEX `invcode_agent`(`invcode_agent`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ds_user_member_deposit_history
-- ----------------------------
INSERT INTO `ds_user_member_deposit_history` VALUES ('HbfH7AMsYa', '13510385413', 'ali_pay', 141.38, 7.044550, 1000.00, 001, 'ali_pay', '', '', '2020-05-07 19:23:55', '2020-05-07 19:23:55', 0000, '710650', '123456');
INSERT INTO `ds_user_member_deposit_history` VALUES ('HbfH7AMsYa', '13510385413', 'ali_pay', 141.38, 7.044550, 1000.00, 001, 'ali_pay', '', '', '2020-05-07 19:26:11', '2020-05-07 19:26:11', 0000, '710650', '123456');
INSERT INTO `ds_user_member_deposit_history` VALUES ('HbfH7AMsYa', '13510385413', 'wechat_pay', 14138472.64, 7.044550, 100000000.00, 001, 'wechat_pay', '', '', '2020-05-07 19:29:16', '2020-05-07 19:29:16', 0000, '710650', '123456');
INSERT INTO `ds_user_member_deposit_history` VALUES ('HbfH7AMsYa', '13510385413', '78436784996606', 0.00, 1.000000, 0.00, 003, 'USDT', '', '', '2020-05-07 21:18:56', '2020-05-07 21:18:56', 0001, '', '');
INSERT INTO `ds_user_member_deposit_history` VALUES ('HbfH7AMsYa', '13510385413', '78436784996608', 0.00, 1.000000, 0.00, 003, 'USDT', '', '', '2020-05-07 21:20:55', '2020-05-07 21:20:55', 0001, '', '');
INSERT INTO `ds_user_member_deposit_history` VALUES ('HbfH7AMsYa', '13510385413', '78436784996609', 0.00, 1.000000, 0.00, 003, 'USDT', '', '', '2020-05-07 21:21:31', '2020-05-07 21:21:31', 0001, '', '');
INSERT INTO `ds_user_member_deposit_history` VALUES ('HbfH7AMsYa', '13510385413', '78436784996610', 0.00, 1.000000, 0.00, 003, 'USDT', '', '', '2020-05-07 21:22:41', '2020-05-07 21:22:41', 0001, '', '');
INSERT INTO `ds_user_member_deposit_history` VALUES ('HbfH7AMsYa', '13510385413', 'wechat_pay', 141.70, 7.028900, 1000.00, 001, 'wechat_pay', '', '', '2020-05-08 10:55:03', '2020-05-08 10:55:03', 0000, '710650', '123456');
INSERT INTO `ds_user_member_deposit_history` VALUES ('HbfH7AMsYa', '13510385413', 'ali_pay', 141.70, 7.028900, 1000.00, 001, 'ali_pay', '', '', '2020-05-08 10:55:39', '2020-05-08 10:55:39', 0000, '710650', '123456');
INSERT INTO `ds_user_member_deposit_history` VALUES ('HbfH7AMsYa', '13510385413', 'compose_pay', 1000.00, 7.028900, 1000.00, 001, 'compose_pay', '', '', '2020-05-08 10:55:49', '2020-05-08 10:55:49', 0000, '710650', '123456');
INSERT INTO `ds_user_member_deposit_history` VALUES ('7DQZOEubNK', '13121391353', 'ali_pay', 141.65, 7.031300, 1000.00, 001, 'ali_pay', '', '', '2020-05-08 13:55:31', '2020-05-08 13:55:31', 0000, '530686', '123456');
INSERT INTO `ds_user_member_deposit_history` VALUES ('7DQZOEubNK', '13121391353', 'ali_pay', 849.91, 7.031300, 6000.00, 001, 'ali_pay', '', '', '2020-05-08 13:55:53', '2020-05-08 13:55:53', 0000, '530686', '123456');
INSERT INTO `ds_user_member_deposit_history` VALUES ('7DQZOEubNK', '13121391353', 'ali_pay', 141.65, 7.031300, 1000.00, 001, 'ali_pay', '', '', '2020-05-08 13:56:54', '2020-05-08 13:56:54', 0000, '530686', '123456');

SET FOREIGN_KEY_CHECKS = 1;
