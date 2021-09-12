package lambda

import (
	"github.com/MauricioAntonioMartinez/aws/model"
	"gorm.io/gorm"
)

type LambdaRepo struct {
	db *gorm.DB
}

func (l *LambdaRepo) findFunction(fn *model.Function) error {
	tx := l.db.Where(fn).First(&fn)
	return tx.Error
}
