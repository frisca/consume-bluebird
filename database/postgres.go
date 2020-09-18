package database

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	Dbcon     *gorm.DB
	Errdb     error
	dbuser    string
	dbpass    string
	dbname    string
	dbaddres  string
	dbport    string
	dbdebug   bool
	dbtype    string
	sslmode   string
	dbtimeout string
)

func init() {
	dbtype = beego.AppConfig.DefaultString("db.type", "POSTGRES")
	dbuser = beego.AppConfig.DefaultString("db.postgres.user", "postgres")
	dbpass = beego.AppConfig.DefaultString("db.postgres.pass", "otto123")
	dbname = beego.AppConfig.DefaultString("db.postgres.name", "bluebird")
	dbaddres = beego.AppConfig.DefaultString("db.postgres.addres", "localhost")
	dbport = beego.AppConfig.DefaultString("db.postgres.port", "5432")
	sslmode = beego.AppConfig.DefaultString("db.postgres.sslmode", "disable")
	dbdebug = beego.AppConfig.DefaultBool("db.postgres.debug", true)
	dbtimeout = beego.AppConfig.DefaultString("db.postgres.timeout", "30")
	if DbOpen() != nil {
		fmt.Println("Can Open db Postgres")
	}
}

func DbOpen() error {
	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%s ", dbaddres, dbport, dbuser, dbpass, dbname, sslmode, dbtimeout)
	Dbcon, Errdb = gorm.Open("postgres", args)
	fmt.Println("isi postgres sett ", args)
	if Errdb != nil {
		logs.Error("open db Err ", Errdb)
		return Errdb
	}
	if errping := Dbcon.DB().Ping(); errping != nil {
		return errping
	}
	fmt.Println("Database connected [", dbaddres, "] [", dbname, "] [", dbuser, "] !")
	return nil
}

func GetDbConnect() *gorm.DB {
	if errping := Dbcon.DB().Ping(); errping != nil {
		logs.Error("Db Not Connect test Ping :", errping)
		errping = nil
		if errping = DbOpen(); errping != nil {
			logs.Error("try to connect again but error :", errping)
		}
	}

	Dbcon.LogMode(dbdebug)
	Dbcon.DB().SetMaxIdleConns(200)
	Dbcon.DB().SetMaxOpenConns(200)
	Dbcon.DB().SetConnMaxLifetime(2 * time.Hour)

	return Dbcon
}
