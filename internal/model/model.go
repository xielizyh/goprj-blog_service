package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/xielizyh/goprj-blog_service/global"
	"github.com/xielizyh/goprj-blog_service/pkg/setting"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	// 打开数据库：root:123456@tcp(127.0.0.1)/blog_service?charset=utf8&parseTime=true&loc=local
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		// 日志模式
		db.LogMode(true)
	}
	// 单数形式
	db.SingularTable(true)
	// 最大空闲连接数
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	// 最大连接数
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
