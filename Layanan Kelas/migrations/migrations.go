package main

import (
	config "github.com/salamanderman234/project-intro-2023/layanan-kelas/config"
	model "github.com/salamanderman234/project-intro-2023/layanan-kelas/models"
	"gorm.io/gorm"
)

func init() {
	config.SetConfig("../.env")
}

func main() {
	client, err := config.ConnectDatabase()
	if err != nil {
		panic("Failed to connect database")
	}
	client.Transaction(func(tx *gorm.DB) error {
		models := model.GetAllModel()
		for _, model := range models {
			client.Migrator().DropTable(model)
			err := client.AutoMigrate(model)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
		return nil
	})
}