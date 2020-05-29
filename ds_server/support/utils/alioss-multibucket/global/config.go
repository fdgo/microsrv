package global

import (
	"os"
	"log"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"strings"
	"io/ioutil"
)
var GloConfig *GlobalConfig

type GlobalConfig struct {
	Sqlconfig Sqlconfig `yaml:"sqlconfig"`
}

type Sqlconfig struct {
	RedisConfig Redisrv  `yaml:"redisrv"`
	MysqlConfig Mysqlsrv `yaml:"mysqlsrv"`
	AliossConfig Aliosssrv `yaml:"aliosssrv"`
	QiniuossConfig Qiniuosssrv `yaml:"qiniuosssrv"`
}

type Redisrv struct {
	Server    string `yaml:"server"`
}

type Mysqlsrv struct {
	Server    string `yaml:"server"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Db        string `yaml:"db"`

}
type Aliosssrv struct{
	Endpoint   string       `yaml:"endpoint"`
	AccessKeyID   string     `yaml:"accessKeyID"`
	AccessKeySecret string    `yaml:"accessKeySecret"`
	BucketName  string        `yaml:"bucketName"`
}
type Qiniuosssrv struct{
	Endpoint   string       `yaml:"endpoint"`
	AccessKey string  `yaml:"accessKey"`
	SecretKey string   `yaml:"secretKey"`
	Bucket string      `yaml:"bucket"`
}

func GetConfig() ( err error) {
	configFile, err := ioutil.ReadFile(GetCurrentDirectory()+"/global/config.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
	}
	err = yaml.Unmarshal(configFile, &GloConfig)
	return err
}
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
