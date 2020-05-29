SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE IF NOT EXISTS `jz_cloud_game` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

-- 静态表不用status字段，只用is_deleted来表示是否有效
-- 客户端渠道信息表 
DROP TABLE IF EXISTS  `t_client_channel_info`;
CREATE TABLE `t_client_channel_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `client_type` varchar(32) DEFAULT NULL COMMENT '客户端，iOS；Android ; web; h5; wechat;',
  `channel` varchar(32) NOT NULL DEFAULT 'jiaozi' COMMENT '渠道',
  `product_name` varchar(128) NOT NULL DEFAULT 'com.jzbro.cloudgame' COMMENT '产品唯一标识，以包名识别',
  `product_type` int(11) DEFAULT 0 COMMENT '是否是合集：0-单包；1-合集',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL ,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除，1-删除; 0-未删除;',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_client_channel` (`client_type`,`channel`,`product_name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='客户端渠道信息';

-- 客户端版本
DROP TABLE IF EXISTS  `t_client_version`;
CREATE TABLE `t_client_version` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `client_channel_id` int(11) NOT NULL COMMENT '客户端渠道id',
  `version` varchar(32) NOT NULL DEFAULT '1.0.0' COMMENT '版本字符串',
  `version_name` varchar(32) NOT NULL DEFAULT 'V1.0.0' COMMENT '版本名称',
  `version_code` int(11) DEFAULT 19010100 COMMENT '版本code，命名规则 年后两位+月+日+第几次更新，如19110100',
  `desc` text COMMENT '更新简介',
  `url` varchar(255) DEFAULT NULL COMMENT '下载链接',
  `is_valid` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否启用，1-启用; 0-未启用;',
  `publish_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '启用时间',
  `update_type` int(11) NOT NULL DEFAULT 0 COMMENT '更新方式，0-正常更新; 1-静默更新',
  `force_update` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否强制更新，1-强制；0-非强制;',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除，1-删除; 0-未删除;',
   PRIMARY KEY (`id`),
  KEY `idx_client_channel` (`client_channel_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='客户端版本';


-- 游戏信息表
DROP TABLE IF EXISTS  `t_game_info`;
CREATE TABLE `t_game_info` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '游戏名称，全局唯一',
  `name_en` varchar(64) NOT NULL DEFAULT '' COMMENT '游戏英文名，比如SSFIV.',
  `name_py` varchar(32) DEFAULT NULL COMMENT '游戏拼音，便于搜索',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT '1-正常；0-下线;2-待上线；',
  `title` varchar(128) DEFAULT NULL COMMENT '游戏一句话介绍',
  `summary` varchar(255) DEFAULT NULL COMMENT '游戏简介',
  `desc` text DEFAULT NULL COMMENT '游戏详情',
  `company` varchar(64) NOT NULL DEFAULT '' COMMENT '厂商',
  `website` varchar(255) NOT NULL DEFAULT '' COMMENT '官网',
  `publish_date` date DEFAULT NULL COMMENT '游戏上线时间',
  `contact` varchar(255) NOT NULL DEFAULT '' COMMENT '联系方式，格式，QQ:231111;email:xx@xx.com',
  `video_width` int(11) NOT NULL DEFAULT 1280 COMMENT '默认宽',
  `video_height` int(11) NOT NULL DEFAULT 720 COMMENT '默认高',
  `video_bitrate` int(11) NOT NULL DEFAULT 1500 COMMENT '默认码率',
  `game_type` int(11) NOT NULL DEFAULT 1 COMMENT '游戏类型：单机、同屏联机、跨屏联机',
  `max_player` int(11) NOT NULL DEFAULT 1 COMMENT '支持玩家数',
  `cpu_load` int(11) NOT NULL DEFAULT 1000 COMMENT 'cpu负载',
  `gpu_load` int(11) NOT NULL DEFAULT 1000 COMMENT 'gpu负载',
  `memory_load` int(11) NOT NULL DEFAULT 1000 COMMENT '内存负载',
  `profile_type` int(11) NOT NULL DEFAULT 0 COMMENT '存档类型，默认0，不存档；1，普通存档。', 
  `control_type` int(11) NOT NULL DEFAULT 1 COMMENT '1，手柄；2，全键盘；4，鼠标；8, 自定义',
  `use_client_resolution` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否使用客户端分辨率，默认0，不使用',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除，1-删除; 0-未删除;',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_game_name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='游戏信息表';

-- 游戏分类，暂时支持1级分类
DROP TABLE IF EXISTS  `t_game_catagory`;
CREATE TABLE `t_game_catagory` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT 0 COMMENT '父类ID',
  `name` varchar(64) NOT NULL DEFAULT ''  COMMENT '分类名称',
  `desc` text DEFAULT NULL COMMENT '分类描述',
  `weight` int(11) NOT NULL DEFAULT 1 COMMENT '权重，越大排名越前',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除，1-删除; 0-未删除;',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='游戏分类';

-- 游戏分类关联表
DROP TABLE IF EXISTS  `t_game_category_relation`;
CREATE TABLE `t_game_category_relation` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_id` int(11) DEFAULT NULL COMMENT '游戏id',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '分类Id',
  `catagory_name` varchar(64) NOT NULL DEFAULT ''  COMMENT '分类名称',
  `weight` int(11) NOT NULL DEFAULT 1 COMMENT '权重，越大排名越前',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除，1-删除; 0-未删除;',
  PRIMARY KEY (`id`),
  KEY `game_id` (`game_id`) USING BTREE,
  KEY `category_id` (`category_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='游戏分类关联表';

-- 游戏标签关联表
DROP TABLE IF EXISTS  `t_game_tag`;
CREATE TABLE `t_game_tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_id` int(11) DEFAULT NULL COMMENT '游戏id',
  `name` varchar(64) NOT NULL DEFAULT ''  COMMENT '标签名称',
  `desc` text DEFAULT NULL COMMENT '标签描述',
  `weight` int(11) NOT NULL DEFAULT 1 COMMENT '权重，越大排名越前',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除，1-删除; 0-未删除;',
  PRIMARY KEY (`id`)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='游戏标签关联表';

-- 客户端渠道和游戏关联表
DROP TABLE IF EXISTS  `t_game_channel`;
CREATE TABLE `t_game_channel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `client_channel_id` int(11) DEFAULT NULL COMMENT '客户端渠道id',
  `game_id` int(11) DEFAULT NULL COMMENT '游戏id',
  `client_version` int(11) DEFAULT 19010100 COMMENT '最低版本code',
  `weight` int(11) NOT NULL DEFAULT 1 COMMENT '权重，越大排名越前',
  `video_width` int(11) NOT NULL DEFAULT 1280 COMMENT '默认宽',
  `video_height` int(11) NOT NULL DEFAULT 720 COMMENT '默认高',
  `video_bitrate` int(11) NOT NULL DEFAULT 1500 COMMENT '默认码率',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除，1-删除; 0-未删除;',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_client_channel` (`client_channel_id`,`game_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='客户端渠道和游戏关联';


-- 游戏图片
DROP TABLE IF EXISTS  `t_game_image`;
CREATE TABLE `t_game_image` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_id` int(11) NOT NULL COMMENT '游戏ID',
  `img_type` int(11) NOT NULL DEFAULT 0 COMMENT '1-图片；2-gif；3-视频',
  `position_type` int(11) NOT NULL DEFAULT '0' COMMENT '1-封面；2-loading；3-icon',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT 'URL',
  `desc` text DEFAULT NULL COMMENT '描述',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除，1-删除; 0-未删除;',
  PRIMARY KEY (`id`),
  KEY `idx_game` (`game_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='游戏图片';


-- 游戏按键布局
DROP TABLE IF EXISTS  `t_game_button_setting`;
CREATE TABLE `t_game_button_setting` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `game_id` int(11) NOT NULL COMMENT '游戏id',
  `status` int(11) DEFAULT 0 COMMENT '状态：0-未启用，1-已启用',
  `device_type` int(11) DEFAULT 0 COMMENT '设备类型：1-Android；2-iOS;3-Web',
  `config` text COMMENT '游戏app按键设置，json结构',
  `height` int(11) NOT NULL DEFAULT 667 COMMENT '分辨率高',
  `width` int(11) NOT NULL DEFAULT 375 COMMENT '分辨率宽',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除，1-删除; 0-未删除;',
  PRIMARY KEY (`id`),
  KEY `idx_game_user` (`user_id`,`game_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户游戏按键设置,gameid=0 通用，userid=0 该游戏默认按键';


-- 游戏存档
DROP TABLE IF EXISTS  `t_game_profile`;
CREATE TABLE `t_game_profile` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL COMMENT '用户id',
  `game_id` int(11) DEFAULT NULL COMMENT '游戏id',
  `device_id` varchar(64) DEFAULT NULL COMMENT '设备id',
  `gs_ip` varchar(64) DEFAULT NULL COMMENT 'gs_ip',
  `gp_id` varchar(64) DEFAULT NULL COMMENT 'gs_id',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT '存档状态：1-上传成功，2-上传中，3-上传失败',
  `url` varchar(255) DEFAULT NULL COMMENT '存档链接',
  `upload_time` timestamp NULL DEFAULT NULL COMMENT '存档上传时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除，1-删除; 0-未删除;',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_game_user` (`user_id`,`game_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='游戏存档';

SET FOREIGN_KEY_CHECKS = 1;

