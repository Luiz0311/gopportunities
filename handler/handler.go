package handler

import (
	"github.com/Luiz0311/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.NewLogger("handler")
	db = config.GetSQLite()
}
