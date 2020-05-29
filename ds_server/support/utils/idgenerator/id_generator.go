package idgenerator

import (
	"fmt"
	"strconv"
	"sync"

	db2 "ds_server/support/lib/mysqlex"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	l   sync.Mutex
	gen *SQLIdGenerator
	db  *gorm.DB
)

type IDSpace struct {
	SpaceName string `gorm:"primary_key"`
	Prefix    string
	Suffix    string
	Seed      int64
	BatchSize int64
	start     int64 `gorm:"-"`
	end       int64 `gorm:"-"`
}

func (ids IDSpace) TableName() string {
	return "id_space"
}

type IDGenerator interface {
	Reload(spaceName string)
	GenerateStringID(spaceName string) string
	GenerateLongID(spaceName string) int64
	GenerateGuardID(spaceName string) string
	ChaosID(id int64, key string) string
	RestoreID(messed string, key string) int64
	Prefix(spaceName string) string
	Suffix(spaceName string) string
}

func Instance() IDGenerator {
	factory := ServiceInstance(ID_GENERATOR)
	return factory().(IDGenerator)
}

type SQLIdGenerator map[string]*IDSpace

func newSQLIdGenerator() interface{} {
	if gen == nil {
		l.Lock()
		if gen == nil {
			gen = &SQLIdGenerator{}
			Init()
		}
		l.Unlock()
	}
	return gen
}

//Init 在系统启动时加载id space
func Init() {
	if gen == nil {
		gen = &SQLIdGenerator{}
	}
	if db == nil {
		db = db2.MysqlInstanceg()
	}
	idSpaces := []*IDSpace{}
	db.Find(&idSpaces)
	for _, idSpace := range idSpaces {
		gen.Reload(idSpace.SpaceName)
	}
}

func SetDB(pDB *gorm.DB) {
	db = pDB
}

//Reload 重新加载id space
func (gen *SQLIdGenerator) Reload(spaceName string) {
	var idspace *IDSpace
	var ok bool
	idspace, ok = (*gen)[spaceName]
	if !ok {
		idspace = &IDSpace{}
	}

	db.Where("space_name = ?", spaceName).Find(idspace)
	rows := db.Model(&IDSpace{}).Where("space_name = ?", spaceName).Where("seed = ?", idspace.Seed).Update("seed", idspace.Seed+idspace.BatchSize).RowsAffected
	if rows > 0 {
		if idspace.start < idspace.Seed {
			idspace.start = idspace.Seed + 1 // 避免边界值的时候服务器关闭,下次启动有可能产生重复单号
		} else {
			idspace.start++
		}
		idspace.end = idspace.Seed + idspace.BatchSize
		(*gen)[idspace.SpaceName] = idspace
	} else {
		gen.Reload(spaceName)
	}
}

//GenerateStringID 生成字符串ID， prefix + id + suffix，
//返回：order_34354_000
func (gen *SQLIdGenerator) GenerateStringID(spaceName string) string {
	l.Lock()
	defer l.Unlock()

	idspace := (*gen)[spaceName]
	if idspace == nil {
		Log.WithFields(logrus.Fields{
			"spaceName": spaceName,
			"idspace":   gen,
		}).Panic("space name not found")
	}
	if idspace.start >= idspace.end {
		gen.Reload(idspace.SpaceName)
	} else {
		idspace.start++
	}

	if idspace.Suffix != "" {
		return fmt.Sprintf("%s_%d_%s", idspace.Prefix, idspace.start, idspace.Suffix)
	}
	return fmt.Sprintf("%s_%d", idspace.Prefix, idspace.start)
}

//GenerateLongID 生成纯数字ID
//返回: 34354000
func (gen *SQLIdGenerator) GenerateLongID(spaceName string) int64 {
	l.Lock()
	defer l.Unlock()

	idspace := (*gen)[spaceName]
	if idspace == nil {
		Log.WithFields(logrus.Fields{
			"spaceName": spaceName,
			"idspace":   gen,
		}).Panic("space name not found")
	}
	if idspace.start >= idspace.end {
		gen.Reload(idspace.SpaceName)
	} else {
		idspace.start++
	}

	return idspace.start
}

//GenerateGuardID 生成混淆ID, suffix为混淆字符串, seed需要设置较大, 先将seed转换成36进制, 再按照混淆字符串进行转换
//example: spaceName:order	prefix:FO	suffix:3Q0SCUBKNMERTYIOPAWDFG24HJL579ZV68X1	seed:784367845356546	batch_size:1000
//返回: FOKLQM2STDTR
func (gen *SQLIdGenerator) GenerateGuardID(spaceName string) string {
	l.Lock()
	defer l.Unlock()

	idspace := (*gen)[spaceName]
	if idspace == nil {
		Log.WithFields(logrus.Fields{
			"spaceName": spaceName,
			"idspace":   gen,
		}).Panic("space name not found")
	}
	if idspace.start >= idspace.end {
		gen.Reload(idspace.SpaceName)
	} else {
		idspace.start++
	}

	num := gen.ChaosID(idspace.start, idspace.Suffix)
	return idspace.Prefix + num
}

func (gen *SQLIdGenerator) Prefix(spaceName string) string {
	l.Lock()
	defer l.Unlock()

	idspace := (*gen)[spaceName]
	if idspace == nil {
		Log.WithFields(logrus.Fields{
			"spaceName": spaceName,
			"idspace":   gen,
		}).Panic("space name not found")
	}

	return idspace.Prefix
}

func (gen *SQLIdGenerator) Suffix(spaceName string) string {
	l.Lock()
	defer l.Unlock()

	idspace := (*gen)[spaceName]
	if idspace == nil {
		Log.WithFields(logrus.Fields{
			"spaceName": spaceName,
			"idspace":   gen,
		}).Panic("space name not found")
	}
	return idspace.Suffix
}

//ChaosID 混淆一个int64类型的ID
func (gen *SQLIdGenerator) ChaosID(id int64, key string) string {
	base := len(key)
	//将seed转换为36进制字符串
	baseInt := strconv.FormatInt(id, base)

	var num string
	for _, char := range baseInt {
		//按照混淆字符串进行映射
		i, _ := strconv.ParseInt(string(char), base, 64)
		num = num + string(key[i])
	}
	return num
}

//RestoreID 还原被混淆过的int64类型ID
func (gen *SQLIdGenerator) RestoreID(messed string, key string) int64 {
	var baseInt string
	base := len(key)
	//按照混淆字符串进行逆映射
	for _, char := range messed {
		for i, c := range key {
			if char == c {
				s := strconv.FormatInt(int64(i), base)
				baseInt = baseInt + s
			}
		}
	}

	//将36进制字符串转化为10进制
	num, err := strconv.ParseInt(baseInt, base, 64)
	if err != nil {
		fmt.Print(err)
		return 0
	}
	return num
}

func (gen *SQLIdGenerator) CreateTable() {
	db := db2.MysqlInstanceg()
	if db.HasTable(&IDSpace{}) {
		return
	}
	db.CreateTable(&IDSpace{})
}
