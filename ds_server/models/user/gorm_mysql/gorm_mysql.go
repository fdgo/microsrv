package model

import (
	"time"
)

/******sql******
CREATE TABLE `ds_sys_info` (
  `connect_us` varchar(255) NOT NULL DEFAULT '' COMMENT '联系我们'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC
******sql******/
// DsSysInfo [...]
type DsSysInfo struct {
	ConnectUs string `gorm:"column:connect_us;type:varchar(255);not null" json:"connect_us"` // 联系我们
}

/******sql******
CREATE TABLE `ds_user_basicinfo` (
  `uuid` varchar(16) NOT NULL DEFAULT '' COMMENT '用户ID号',
  `mobile` varchar(16) NOT NULL DEFAULT '' COMMENT '手机号',
  `salt` varchar(16) NOT NULL DEFAULT '' COMMENT '密码盐',
  `hash` varchar(64) NOT NULL DEFAULT '' COMMENT '密码hash',
  `last_login_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次登录时间',
  `last_login_ip` varchar(32) NOT NULL DEFAULT '' COMMENT '最后一次登录ip',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` tinyint(4) NOT NULL COMMENT '是否删除0:未删除 1:删除',
  `real_name` varchar(16) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `age` tinyint(3) unsigned zerofill NOT NULL DEFAULT '000' COMMENT '年龄',
  `gender` tinyint(3) unsigned zerofill NOT NULL DEFAULT '000' COMMENT '性别：0:男，1:女',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像地址',
  `nick_name` varchar(16) NOT NULL DEFAULT '' COMMENT '昵称',
  `birthday` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生日',
  `status` tinyint(3) unsigned zerofill NOT NULL DEFAULT '000' COMMENT '账户状态0:正常 1:禁用  2:注销',
  PRIMARY KEY (`uuid`) USING BTREE,
  UNIQUE KEY `mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC
******sql******/
// DsUserBasicinfo [...]
type DsUserBasicinfo struct {
	UUID          string    `gorm:"primary_key;column:uuid;type:varchar(16);not null" json:"uuid"`          // 用户ID号
	Mobile        string    `gorm:"unique;column:mobile;type:varchar(16);not null" json:"mobile"`           // 手机号
	Salt          string    `gorm:"column:salt;type:varchar(16);not null" json:"salt"`                      // 密码盐
	Hash          string    `gorm:"column:hash;type:varchar(64);not null" json:"hash"`                      // 密码hash
	LastLoginTime time.Time `gorm:"column:last_login_time;type:datetime;not null" json:"last_login_time"`   // 最后一次登录时间
	LastLoginIP   string    `gorm:"column:last_login_ip;type:varchar(32);not null" json:"last_login_ip"`    // 最后一次登录ip
	CreateTime    time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`           // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime;not null" json:"update_time"`           // 更新时间
	Deleted       int8      `gorm:"column:deleted;type:tinyint(4);not null" json:"deleted"`                 // 是否删除0:未删除 1:删除
	RealName      string    `gorm:"column:real_name;type:varchar(16);not null" json:"real_name"`            // 真实姓名
	Age           int8      `gorm:"column:age;type:tinyint(3) unsigned zerofill;not null" json:"age"`       // 年龄
	Gender        int8      `gorm:"column:gender;type:tinyint(3) unsigned zerofill;not null" json:"gender"` // 性别：0:男，1:女
	Avatar        string    `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`                 // 头像地址
	NickName      string    `gorm:"column:nick_name;type:varchar(16);not null" json:"nick_name"`            // 昵称
	Birthday      time.Time `gorm:"column:birthday;type:datetime;not null" json:"birthday"`                 // 生日
	Status        int8      `gorm:"column:status;type:tinyint(3) unsigned zerofill;not null" json:"status"` // 账户状态0:正常 1:禁用  2:注销
}

/******sql******
CREATE TABLE `ds_user_member_account` (
  `uuid` varchar(16) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '用户ID',
  `mobile` varchar(16) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '用户手机号',
  `balance` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '总金额',
  `private_key` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
  `salt` varchar(16) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '支付盐',
  `hash` varchar(64) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '密码hash',
  `address_in` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '收款地址',
  `address_out` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '付款地址',
  `status` tinyint(3) unsigned zerofill NOT NULL DEFAULT '000' COMMENT '账户状态0:正常 1:禁用  2:注销',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`uuid`) USING BTREE,
  UNIQUE KEY `mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC
******sql******/
// DsUserMemberAccount [...]
type DsUserMemberAccount struct {
	UUID       string    `gorm:"primary_key;column:uuid;type:varchar(16);not null" json:"uuid"` // 用户ID
	Mobile     string    `gorm:"unique;column:mobile;type:varchar(16);not null" json:"mobile"`  // 用户手机号
	Balance    float64   `gorm:"column:balance;type:decimal(18,2);not null" json:"balance"`     // 总金额
	PrivateKey string    `gorm:"column:private_key;type:varchar(255);not null" json:"private_key"`
	Salt       string    `gorm:"column:salt;type:varchar(16);not null" json:"salt"`                      // 支付盐
	Hash       string    `gorm:"column:hash;type:varchar(64);not null" json:"hash"`                      // 密码hash
	AddressIn  string    `gorm:"column:address_in;type:varchar(255);not null" json:"address_in"`         // 收款地址
	AddressOut string    `gorm:"column:address_out;type:varchar(255);not null" json:"address_out"`       // 付款地址
	Status     uint8     `gorm:"column:status;type:tinyint(3) unsigned zerofill;not null" json:"status"` // 账户状态0:正常 1:禁用  2:注销
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`           // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null" json:"update_time"`           // 更新时间
	Ispwd      int8      `gorm:"column:ispwd;type:tinyint(3) unsigned zerofill;not null" json:"ispwd"`   // 密码是否为空
}
/******sql******
CREATE TABLE `ds_user_member_agent` (
  `uuid_self` varchar(16) NOT NULL DEFAULT '' COMMENT '新注册的用户ID号',
  `mobile_self` varchar(16) NOT NULL DEFAULT '' COMMENT '新注册的用户手机号',
  `invcode_self` varchar(16) NOT NULL DEFAULT '' COMMENT '新注册的用户自身邀请码',
  `uuid_agent` varchar(16) NOT NULL DEFAULT '' COMMENT '代理的ID号',
  `mobile_agent` varchar(16) NOT NULL DEFAULT '' COMMENT '代理的手机号',
  `invcode_agent` varchar(16) NOT NULL DEFAULT '' COMMENT '代理的邀请码',
  `memclass_self` tinyint(3) unsigned zerofill NOT NULL DEFAULT '000' COMMENT '会员等级',
  `member_tag` varchar(16) NOT NULL DEFAULT '' COMMENT '会员标识',
  `member_name` varchar(16) NOT NULL DEFAULT '' COMMENT '会员名称',
  `agent_class` tinyint(3) unsigned zerofill NOT NULL DEFAULT '000' COMMENT '代理等级',
  `agent_tag` varchar(16) NOT NULL DEFAULT '' COMMENT '代理标识',
  `agent_name` varchar(16) NOT NULL DEFAULT '' COMMENT '代理名称',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`uuid_self`) USING BTREE,
  UNIQUE KEY `mobile_self` (`mobile_self`) USING BTREE,
  UNIQUE KEY `invcode_self` (`invcode_self`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC
******sql******/
// DsUserMemberAgent [...]
type DsUserMemberAgent struct {
	UUIDSelf     string    `gorm:"primary_key;column:uuid_self;type:varchar(16);not null" json:"uuid_self"`              // 新注册的用户ID号
	MobileSelf   string    `gorm:"unique;column:mobile_self;type:varchar(16);not null" json:"mobile_self"`               // 新注册的用户手机号
	InvcodeSelf  string    `gorm:"unique;column:invcode_self;type:varchar(16);not null" json:"invcode_self"`             // 新注册的用户自身邀请码
	UUIDAgent    string    `gorm:"column:uuid_agent;type:varchar(16);not null" json:"uuid_agent"`                        // 代理的ID号
	MobileAgent  string    `gorm:"column:mobile_agent;type:varchar(16);not null" json:"mobile_agent"`                    // 代理的手机号
	InvcodeAgent string    `gorm:"column:invcode_agent;type:varchar(16);not null" json:"invcode_agent"`                  // 代理的邀请码
	MemclassSelf int8      `gorm:"column:memclass_self;type:tinyint(3) unsigned zerofill;not null" json:"memclass_self"` // 会员等级
	MemberTag    string    `gorm:"column:member_tag;type:varchar(16);not null" json:"member_tag"`                        // 会员标识
	MemberName   string    `gorm:"column:member_name;type:varchar(16);not null" json:"member_name"`                      // 会员名称
	AgentClass   int8      `gorm:"column:agent_class;type:tinyint(3) unsigned zerofill;not null" json:"agent_class"`     // 代理等级
	AgentTag     string    `gorm:"column:agent_tag;type:varchar(16);not null" json:"agent_tag"`                          // 代理标识
	AgentName    string    `gorm:"column:agent_name;type:varchar(16);not null" json:"agent_name"`                        // 代理名称
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`                         // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime;not null" json:"update_time"`                         // 更新时间
}

/******sql******
CREATE TABLE `ds_user_member_class` (
  `mem_money` decimal(18,2) NOT NULL COMMENT '会员等级金额界限',
  `mem_tag` tinyint(4) NOT NULL COMMENT '会员等级1,2,3,4,5',
  `mem_tagex` varchar(16) NOT NULL COMMENT '会员等级标记  M1，M2，M3，M4，M5',
  `mem_name` varchar(16) NOT NULL COMMENT '会员等级名称  普通卡，铜卡，银卡，金卡，钻石卡',
  KEY `mem_money` (`mem_money`) USING BTREE,
  KEY `mem_tag` (`mem_tag`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC
******sql******/
// DsUserMemberClass [...]
type DsUserMemberClass struct {
	MemMoney float64 `gorm:"index;column:mem_money;type:decimal(18,2);not null" json:"mem_money"` // 会员等级金额界限
	MemTag   int8    `gorm:"index;column:mem_tag;type:tinyint(4);not null" json:"mem_tag"`        // 会员等级1,2,3,4,5
	MemTagex string  `gorm:"column:mem_tagex;type:varchar(16);not null" json:"mem_tagex"`         // 会员等级标记  M1，M2，M3，M4，M5
	MemName  string  `gorm:"column:mem_name;type:varchar(16);not null" json:"mem_name"`           // 会员等级名称  普通卡，铜卡，银卡，金卡，钻石卡
}

/******sql******
CREATE TABLE `ds_user_agent_class` (
  `agent_money` decimal(18,2) NOT NULL COMMENT '合伙人等级金额界限',
  `agent_tag` tinyint(4) NOT NULL COMMENT '合伙人等级1,2,3,4,5',
  `agent_tagex` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '合伙人等级标记  D1，D2，D3，D4，D5',
  `agent_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '合伙人等级名称  代理合伙人，高级合伙人，城市合伙人，区域合伙人，全球合伙人',
  KEY `agent_money` (`agent_money`) USING BTREE,
  KEY `agent_tag` (`agent_tag`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC
******sql******/
// DsUserAgentClass [...]
type DsUserAgentClass struct {
	AgentMoney float64 `gorm:"index;column:agent_money;type:decimal(18,2);not null" json:"agent_money"` // 合伙人等级金额界限
	AgentTag   int8    `gorm:"index;column:agent_tag;type:tinyint(4);not null" json:"agent_tag"`        // 合伙人等级1,2,3,4,5
	AgentTagex string  `gorm:"column:agent_tagex;type:varchar(16);not null" json:"agent_tagex"`         // 合伙人等级标记  D1，D2，D3，D4，D5
	AgentName  string  `gorm:"column:agent_name;type:varchar(16);not null" json:"agent_name"`           // 合伙人等级名称  代理合伙人，高级合伙人，城市合伙人，区域合伙人，全球合伙人
}
/******sql******
CREATE TABLE `ds_user_member_deposit_history` (
  `uuid` varchar(16) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '用户ID',
  `mobile` varchar(16) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '用户手机号',
  `source_id` varchar(64) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '业务订单Id',
  `balance` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '金额',
  `rate` decimal(18,6) unsigned DEFAULT '0.000000' COMMENT '汇率',
  `balance_src` decimal(18,2) unsigned DEFAULT '0.00' COMMENT '原始金额',
  `deposit_type` tinyint(3) unsigned zerofill NOT NULL DEFAULT '000' COMMENT '充值类型 0:扣款,1:充值',
  `deposit_name` varchar(16) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '充值名字：购买商品，商品退款',
  `address_in` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '收款地址',
  `address_out` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '支付地址',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` tinyint(4) unsigned zerofill NOT NULL DEFAULT '0000' COMMENT '会员账户状态 0:正常，1:禁止 2:销户',
  `invcode_self` varchar(16) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '自身邀请码',
  `invcode_agent` varchar(16) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '代理邀请码',
  KEY `source_id` (`source_id`) USING BTREE,
  KEY `invcode_agent` (`invcode_agent`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC
******sql******/
// DsUserMemberDepositHistory [...]
type DsUserMemberDepositHistory struct {
	UUID         string    `gorm:"column:uuid;type:varchar(16);not null" json:"uuid"`                                  // 用户ID
	Mobile       string    `gorm:"column:mobile;type:varchar(16);not null" json:"mobile"`                              // 用户手机号
	SourceID     string    `gorm:"index;column:source_id;type:varchar(64);not null" json:"source_id"`                  // 业务订单Id
	Balance      float64   `gorm:"column:balance;type:decimal(18,2);not null" json:"balance"`                          // 金额
	Rate         float64   `gorm:"column:rate;type:decimal(18,6) unsigned" json:"rate"`                                // 汇率
	BalanceSrc   float64   `gorm:"column:balance_src;type:decimal(18,2) unsigned" json:"balance_src"`                  // 原始金额
	DepositType  int8      `gorm:"column:deposit_type;type:tinyint(3) unsigned zerofill;not null" json:"deposit_type"` // 充值类型 0:扣款,1:充值
	DepositName  string    `gorm:"column:deposit_name;type:varchar(16);not null" json:"deposit_name"`                  // 充值名字：购买商品，商品退款
	AddressIn    string    `gorm:"column:address_in;type:varchar(255);not null" json:"address_in"`                     // 收款地址
	AddressOut   string    `gorm:"column:address_out;type:varchar(255);not null" json:"address_out"`                   // 支付地址
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`                       // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime;not null" json:"update_time"`                       // 更新时间
	Status       uint8     `gorm:"column:status;type:tinyint(4) unsigned zerofill;not null" json:"status"`             // 会员账户状态 0:正常，1:禁止 2:销户
	InvcodeSelf  string    `gorm:"column:invcode_self;type:varchar(16);not null" json:"invcode_self"`                  // 自身邀请码
	InvcodeAgent string    `gorm:"index;column:invcode_agent;type:varchar(16);not null" json:"invcode_agent"`          // 代理邀请码
}

type CommonTotal struct {
	Total   float64   `gorm:"column:total;type:decimal(18,2) unsigned" json:"total"`
}