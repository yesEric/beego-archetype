package init

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)
// 数据库链接配置
func InitDb() {
	username:=beego.AppConfig.String("db_username")
	passwd:=beego.AppConfig.String("db_password")
	host:=beego.AppConfig.String("db_host")
	port,err:=beego.AppConfig.Int("db_port")
	dbname:=beego.AppConfig.String("db_name")
	if nil!=err{
		port=3306
	}
	orm.Debug=true
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",username, passwd, host, port, dbname))
	orm.RunSyncdb("default",false,false)
}