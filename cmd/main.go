package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/thangsuperman/bee-happy/cmd/api"
	"github.com/thangsuperman/bee-happy/config"
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
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.PublicHost,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":3000", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
