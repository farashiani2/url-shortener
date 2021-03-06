package drivers

import (
	"fmt"

	"github.com/khalil-farashiani/url-shortener/internal/models/url"
	"github.com/khalil-farashiani/url-shortener/internal/models/user"
	"github.com/khalil-farashiani/url-shortener/internal/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	username = utils.GetEnv("my_sql_username", "root")
	password = utils.GetEnv("my_sql_password", "root")
	host     = utils.GetEnv("my_sql_host", "127.0.0.1:3306")
	schema   = utils.GetEnv("my_sql_schema", "sholink")

	DB = &gorm.DB{}
)

func ConnectSQL() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, schema)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB.AutoMigrate(&url.Url{})
	DB.AutoMigrate(&user.User{})

	return nil
}
