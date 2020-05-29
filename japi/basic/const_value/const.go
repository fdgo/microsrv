package const_value

import "time"

const CONFIG_ADDRESS = "127.0.0.1:9600"
const EncryptKey = "jzxdjzxd" //必须8位
const TokenKey = "1234567890"
const TokenExpiredDate = 3600 * 24 * 10 * time.Second
const TokenIDKeyPrefix = "token:auth:id:"