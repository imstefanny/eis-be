package database

import (
	"eis-be/config"
	"eis-be/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
	InitialMigration()
}

func InitDB() *gorm.DB {
	dbconfig := config.ReadEnv()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbconfig.DB_USERNAME,
		dbconfig.DB_PASSWORD,
		dbconfig.DB_HOSTNAME,
		dbconfig.DB_PORT,
		dbconfig.DB_NAME,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func InitialMigration() {
	DB.AutoMigrate(&models.Blogs{})
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Applicants{})
	DB.AutoMigrate(&models.Students{})
	DB.AutoMigrate(&models.Guardians{})
	DB.AutoMigrate(&models.DocTypes{})
	DB.AutoMigrate(&models.Documents{})
	DB.AutoMigrate(&models.WorkScheds{})
	DB.AutoMigrate(&models.WorkSchedDetails{})
	DB.AutoMigrate(&models.Subjects{})
	DB.AutoMigrate(&models.Levels{})
	DB.AutoMigrate(&models.LevelHistories{})
	DB.AutoMigrate(&models.Classrooms{})
}
