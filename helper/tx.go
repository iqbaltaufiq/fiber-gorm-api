package helper

import (
	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	errRecover := recover()
	if errRecover != nil {
		if errRollback := tx.Rollback().Error; errRollback != nil {
			panic(errRollback)
		}
		panic(errRecover)
	} else {
		if errCommit := tx.Commit().Error; errCommit != nil {
			panic(errCommit)
		}
	}
}
