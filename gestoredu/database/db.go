package database

import (
	"fmt"

	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/api/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var DB *gorm.DB

func Db() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/defense"
	//var v = "Não conseguiu conectar ao banco de dados"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Não está conectando ao banco de dados!")
	}
	// DB = db
	db.AutoMigrate(&entities.Users{})
	// db.AutoMigrate(&entities.Administrador{})
	// db.AutoMigrate(&entities.Aluno{})
	// db.AutoMigrate(&entities.Professor{})
	// db.AutoMigrate(&entities.Responsavel{})
	
	return db

}
