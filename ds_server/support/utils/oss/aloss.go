package oss

import (
	"ds_server/support/utils/constex"
	"fmt"
	"os"
	"sync"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var ossbLock sync.Mutex
var ossInstarnce *oss.Client

// 得到唯一的主库实例
func MyOssInstance() *oss.Client {
	if ossInstarnce != nil {
		return ossInstarnce
	}
	ossbLock.Lock()
	defer ossbLock.Unlock()
	fmt.Println("================", constex.AiossCfg.EndPoint)
	// oss.Timeout(10, 120)表示设置HTTP连接超时时间为10秒（默认值为30秒），HTTP读写超时时间为120秒（默认值为60秒）。0表示永不超时（不推荐使用）。
	ossInstarnce, err := oss.New(constex.AiossCfg.EndPoint, constex.AiossCfg.AccessKeyId, constex.AiossCfg.AccessKeySecret, oss.Timeout(30, 120))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	if ossInstarnce != nil {
		return ossInstarnce
	}
	return nil
}
