package config

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type AppConfig struct {
	DB_DRIVER   string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     int
	DB_NAME     string
	SERVER_PORT int16
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock() //buat kunci  dan tidak bisa di rubah
	defer lock.Unlock()

	if appConfig == nil { //jika tidak ada isinya
		appConfig = InitConfig() // untuk iniatialisasi
	}
	return appConfig
}

func InitConfig() *AppConfig {
	var defaultconfig AppConfig

	// if _, exist := os.LookupEnv("SECRET"); !exist {
	// 	if err := godotenv.Load(".env"); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// SECRET = os.Getenv("SECRET")
	cnv, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal("Cannot parse DB Port variable")
		return nil
	}
	defaultconfig.SERVER_PORT = int16(cnv)
	defaultconfig.DB_NAME = os.Getenv("DB_NAME")
	defaultconfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	defaultconfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	defaultconfig.DB_HOST = os.Getenv("DB_HOST")
	cnv, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("cannot parse db port variable")
		return nil
	}
	defaultconfig.DB_PORT = cnv
	return &defaultconfig
}

// db_username := os.getenv("DB_USERNAME")
// db_password := os.getenv("DB_PASSWORD")
// db_port := os.getenv("DB_PORT")
// db_host:= os.getenv("DB_Host")
// db_name := os.getenv("DB_NAME")
