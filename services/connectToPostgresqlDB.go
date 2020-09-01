package services

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/postgres" // IMPORTANT TO orm.RegisterDataBase
	//_ "github.com/lib/pq" // IMPORTANT TO orm.RegisterDataBase
)

func ConnToPSQL() *gorm.DB {
	port := viper.GetString("port_db")
	host := viper.GetString("host")
	username := viper.GetString("username")
	dbName := viper.GetString("db_name")
	password := viper.GetString("password")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, username, dbName, password) //Str connection
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Connected to DB!")
	return conn
}
