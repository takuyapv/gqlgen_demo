package common

import (
	"context"
	"demo/graph/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&model.User{},
		&model.Todo{},
		&model.Schedule{})

	return db, nil
}

func GetAll(ctx context.Context, items any) (any, error) {

	context := GetContext(ctx)
	err := context.Database.Find(&items).Error

	if err != nil {
		return items, err
	}
	return items, nil

}
