package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/thangsuperman/bee-happy/cmd/api"
	AppSetting "github.com/thangsuperman/bee-happy/config"
	"github.com/thangsuperman/bee-happy/db"
	_ "github.com/thangsuperman/bee-happy/docs"
)

// @title						            Bee happy API
// @version					            1.0
// @description		              Manage feeds, chat with fiends. It also provides endpoints for searching feed by keyword
// @contact.name		            Thang Phan
// @contact.url				          http://thangphan.com
// @contact.email				        thanglearndevops@gmail.com
// @host						            localhost:3000
// @SecurityDefinitions.apiKey	Bearer
// @in							            header
// @name						            Authorization
// @BasePath					          /
func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 AppSetting.Envs.DBUser,
		Passwd:               AppSetting.Envs.DBPassword,
		Addr:                 AppSetting.Envs.DBAddress,
		DBName:               AppSetting.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":"+AppSetting.Envs.Port, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to Mysql")
}
