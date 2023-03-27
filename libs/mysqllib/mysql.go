package mysqllib

import (
	"app/utils"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	db       *gorm.DB
	dbSystem *gorm.DB
)

//返回新的连接
func getDbConnection(dbName string) *gorm.DB {
	var dsn string
	host := viper.GetString("mysql.host")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	port := viper.GetString("mysql.port")
	//database := viper.GetString("mysql.database")
	//连接数据库的时候加入参数parseTime=true 和loc=Local ，解决时间格式化问题
	dsn = username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	newDb, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql连接异常：", err.Error())
		panic(utils.StringToInterface(err.Error()))
	}
	// 获取通用数据库对象sql.DB以使用其函数
	sqlDB, err := newDb.DB()
	//SetMaxIdleConns设置空闲连接池的最大连接数。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns设置数据库打开的最大连接数。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime设置连接可重用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return newDb
}

//初始化各个数据库的连接
func InitMysqlDb() {
	db = getDbConnection("yibai_account_manage")
	dbSystem = getDbConnection("yibai_account_system")
}

//获取公共连接
func GetMysqlDb() *gorm.DB {
	return db
}

//多个数据库连接
func GetMysqlDbSystem() *gorm.DB {
	return dbSystem
}

//获取新的mysql连接（?）
//func GetNewMysqlDb() *gorm.DB {
//	return getDbConnection()
//}
