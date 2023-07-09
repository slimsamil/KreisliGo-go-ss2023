package db

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	err := Connect(os.Getenv("DB_CONNECT"))
	if err != nil {
		panic(err)
	}
}

func Connect(connect string) error {
	// init DB
	var err error
	dsn := fmt.Sprintf("root:root@tcp(%s)/kreisligo?charset=utf8&parseTime=True&loc=Local", connect)
	log.Info("Using DSN for DB:", dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("failed to connect to database")
	}
	log.Info("Starting automatic migrations")
	if err := DB.Debug().AutoMigrate(&model.Association{}); err != nil {
		return err
	}
	if err := DB.Debug().AutoMigrate(&model.League{}); err != nil {
		return err
	}
	if err := DB.Debug().AutoMigrate(&model.Team{}); err != nil {
		return err
	}
	if err := DB.Debug().AutoMigrate(&model.Player{}); err != nil {
		return err
	}
	if err := DB.Debug().AutoMigrate(&model.Event{}); err != nil {
		return err
	}
	log.Info("Automatic migrations finished")
	return nil
}
