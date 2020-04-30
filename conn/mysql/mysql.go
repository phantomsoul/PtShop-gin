package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"pt-gin/modules/cf"
	"pt-gin/modules/log"
	"strconv"
	"time"
)

type SqlTxStruct struct {
	Tx *sql.Tx
}

// 只会执行一次在执行程序启动的时候
func NewDB(c *cf.MysqlConfig) (Orm *gorm.DB) {
	// 初始化默认连接
	var (
		OrmErr	error
		//SqlErr	error
	)
	var dbConn = mysql.Config{
		User:                 c.User,
		Passwd:               c.Pwd,
		Addr:                 c.DBHost,
		DBName:               c.DBName,
		Collation:            c.Collation,
		Net:                  "tcp",
		AllowNativePasswords: true,
	}
	log.Info(dbConn.FormatDSN())
	//SqlDB, SqlErr = sql.Open("mysql", dbConn.FormatDSN())
	//if SqlErr != nil {
	//	panic(SqlErr.Error())
	//}

	//设置数据库最大连接 减少TimeWait 正式环境调大
	maxIdleConn, _ := strconv.Atoi(c.MaxIdleConn)
	maxOpenConn, _ := strconv.Atoi(c.MaxOpenConn)
	//SqlDB.SetMaxIdleConns(maxIdleConn) // 连接池连接数 = mysql最大连接数/2
	//SqlDB.SetMaxOpenConns(maxOpenConn) // 最大打开连接 = mysql最大连接数

	// 设置连接重置时间
	//SqlDB.SetConnMaxLifetime(80 * time.Second)
	//defer SqlDB.Close()

	// gorm连接方式
	Orm, OrmErr = gorm.Open("mysql", dbConn.FormatDSN())
	if OrmErr != nil {
		panic(OrmErr.Error())
	}
	Orm.SingularTable(true)
	err := Orm.DB().Ping()
	if err != nil {
		panic(err.Error())
	}
	Orm.LogMode(true)
	Orm.DB().SetMaxIdleConns(maxIdleConn)
	Orm.DB().SetMaxOpenConns(maxOpenConn)
	Orm.DB().SetConnMaxLifetime(80 * time.Second)
	defer Orm.Close()
	return Orm
}