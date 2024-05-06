package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/thangsuperman/bee-happy/config"
	"github.com/thangsuperman/bee-happy/db"
	"github.com/thangsuperman/bee-happy/utils"
)

func main() {
	db, err := db.NewMySQLStorage(mysqlCfg.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.PublicHost,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	utils.HaltOn(err)

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	utils.HaltOn(err)

	m, err := migrate.NewWithDatabaseInstance("file://cmd/migrate/migrations", "mysql", driver)
	utils.HaltOn(err)

	cmd := os.Args[(len(os.Args) - 1)]

	// TODO: should test no error no change
	if cmd == "up" {
		err := m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		fmt.Println("Migration up completed")
		return
	}

	if cmd == "down" {
		err := m.Steps(-1)
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		fmt.Println("Migration down completed")
		return
	}

	if len(os.Args) == 3 && os.Args[(len(os.Args)-2)] == "force" {
		targetVersion, err := strconv.Atoi(os.Args[len(os.Args)-1])
		utils.HaltOn(err)

		_, isDirty, err := m.Version()
		if !isDirty {
			fmt.Println("Migration is already up to date")
			return
		}
		utils.HaltOn(err)

		currentVersion, isDirty, err := m.Version()
		if targetVersion >= int(currentVersion) || !isDirty {
			log.Fatal("Can not force through the migration error at version ", currentVersion)
		}

		err = m.Force(int(targetVersion))
		utils.HaltOn(err)

		fmt.Printf("Migration version %d forced successfully\n", targetVersion)

		// TODO: should investigate this one to see the whold flow
		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		} else if err == migrate.ErrNoChange {
			fmt.Println("No pending migrations")
		} else {
			fmt.Println("Migration up completed")
		}

	}
}
