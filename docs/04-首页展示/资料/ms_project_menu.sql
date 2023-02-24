/*
 Navicat Premium Data Transfer

 Source Server         : msproject
 Source Server Type    : MySQL
 Source Server Version : 80020
 Source Host           : localhost:3309
 Source Schema         : msproject

 Target Server Type    : MySQL
 Target Server Version : 80020
 File Encoding         : 65001

 Date: 10/12/2022 12:53:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ms_project_menu
-- ----------------------------
DROP TABLE IF EXISTS `ms_project_menu`;
CREATE TABLE `ms_project_menu`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父id',
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `icon` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `url` varchar(400) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '链接',
  `file_path` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件路径',
  `params` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '链接参数',
  `node` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '#' COMMENT '权限节点',
  `sort` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '菜单排序',
  `status` tinyint(0) UNSIGNED NULL DEFAULT 1 COMMENT '状态(0:禁用,1:启用)',
  `create_by` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `is_inner` tinyint(1) NULL DEFAULT 0 COMMENT '是否内页',
  `values` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '参数默认值',
  `show_slider` tinyint(1) NULL DEFAULT 1 COMMENT '是否显示侧栏',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 176 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '项目菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ms_project_menu
-- ----------------------------
INSERT INTO `ms_project_menu` VALUES (120, 0, '工作台', 'appstore-o', 'home', 'home', ':org', '#', 0, 1, 0, 0, '', 0);
INSERT INTO `ms_project_menu` VALUES (121, 0, '项目管理', 'project', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (122, 121, '项目列表', 'branches', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (124, 0, '系统设置', 'setting', '#', '#', '', '#', 100, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (125, 124, '成员管理', 'unlock', '#', '#', '', '#', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (126, 125, '账号列表', '', 'system/account', 'system/account', '', 'project/account/index', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (127, 122, '我的组织', '', 'organization', 'organization', '', 'project/organization/index', 30, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (130, 125, '访问授权', '', 'system/account/auth', 'system/account/auth', '', 'project/auth/index', 20, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (131, 125, '授权页面', '', 'system/account/apply', 'system/account/apply', ':id', 'project/auth/apply', 30, 1, 0, 1, '', 1);
INSERT INTO `ms_project_menu` VALUES (138, 121, '消息提醒', 'info-circle-o', '#', '#', '', '#', 30, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (139, 138, '站内消息', '', 'notify/notice', 'notify/notice', '', 'project/notify/index', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (140, 138, '系统公告', '', 'notify/system', 'notify/system', '', 'project/notify/index', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (143, 124, '系统管理', 'appstore', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (144, 143, '菜单路由', '', 'system/config/menu', 'system/config/menu', '', 'project/menu/menuadd', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (145, 143, '访问节点', '', 'system/config/node', 'system/config/node', '', 'project/node/save', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (148, 124, '个人管理', 'user', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (149, 148, '个人设置', '', 'account/setting/base', 'account/setting/base', '', 'project/index/editpersonal', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (150, 148, '安全设置', '', 'account/setting/security', 'account/setting/security', '', 'project/index/editpersonal', 0, 1, 0, 1, '', 1);
INSERT INTO `ms_project_menu` VALUES (151, 122, '我的项目', '', 'project/list', 'project/list', ':type', 'project/project/index', 0, 1, 0, 0, 'my', 1);
INSERT INTO `ms_project_menu` VALUES (152, 122, '回收站', '', 'project/recycle', 'project/recycle', '', 'project/project/index', 20, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (153, 121, '项目空间', 'heat-map', 'project/space/task', 'project/space/task', ':code', '#', 20, 1, 0, 1, '', 1);
INSERT INTO `ms_project_menu` VALUES (154, 153, '任务详情', '', 'project/space/task/:code/detail', 'project/space/taskdetail', ':code', 'project/task/read', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (155, 122, '我的收藏', '', 'project/list', 'project/list', ':type', 'project/project/index', 10, 1, 0, 0, 'collect', 1);
INSERT INTO `ms_project_menu` VALUES (156, 121, '基础设置', 'experiment', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (157, 156, '项目模板', '', 'project/template', 'project/template', '', 'project/project_template/index', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (158, 156, '项目列表模板', '', 'project/template/taskStages', 'project/template/taskStages', ':code', 'project/task_stages_template/index', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (159, 122, '已归档项目', '', 'project/archive', 'project/archive', '', 'project/project/index', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (160, 0, '团队成员', 'team', '#', '#', '', '#', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (161, 153, '项目概况', '', 'project/space/overview', 'project/space/overview', ':code', 'project/index/info', 20, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (162, 153, '项目文件', '', 'project/space/files', 'project/space/files', ':code', 'project/index/info', 10, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (163, 122, '项目分析', '', 'project/analysis', 'project/analysis', '', 'project/index/info', 5, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (164, 160, '团队成员', '', '#', '#', '', '#', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (166, 164, '团队成员', '', 'members', 'members', '', 'project/department/index', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (167, 164, '成员信息', '', 'members/profile', 'members/profile', ':code', 'project/department/read', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (168, 153, '版本管理', '', 'project/space/features', 'project/space/features', ':code', 'project/index/info', 20, 1, 0, 1, '', 0);

SET FOREIGN_KEY_CHECKS = 1;
