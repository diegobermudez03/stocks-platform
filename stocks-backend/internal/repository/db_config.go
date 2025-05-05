package repository

import (
	"fmt"

	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


type DBConfig struct{
	Host string
	Port string
	User string
	Password string
	DbName string
	SSLMode string
}

/*
	This function is the one that opens the DB (using the Postgresql driver, compatible with cockroachDB)
	Also migrates (create stocks table)
*/
func InitDb(config DBConfig) (*gorm.DB,error){
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DbName, config.SSLMode)
	//open db
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})	//silent mode, no logs 
	if err != nil{
		return nil, err 
	}
	//clean db just in case, each server execution creates a new db
	//db.Exec(`DROP SCHEMA public CASCADE;`)
	//db.Exec(`CREATE SCHEMA public;`)

	//migrate database 
	err = db.AutoMigrate(&domain.StockModel{})
	if err != nil{
		return nil, err
	}
	return db, nil
}