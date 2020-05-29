package viper

import (
	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}
func LoadConfigFromYaml(path string) (*Config, error) {
	c := &Config{}
	c.v = viper.New()
	//设置配置文件的名字
	c.v.SetConfigName("conf")
	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
	c.v.AddConfigPath(path)
	//设置配置文件类型
	c.v.SetConfigType("yaml")
	if err := c.v.ReadInConfig(); err != nil {
		return nil, err
	}
	return c, nil
}

//c, err := LoadConfigFromYaml("./config")
//if err != nil {
//	return
//}
//name := c.v.GetString("information.name")
//alise := c.v.GetStringSlice("information.Alise")
//age := c.v.GetString("information.age")
//fmt.Println(name, alise, age)