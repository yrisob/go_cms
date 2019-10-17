package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func GetDBEngine() (*xorm.Engine, error) {
	return xorm.NewEngine("mysql", "root:gfhjkm1986@/cms_project?parseTime=true")
}
