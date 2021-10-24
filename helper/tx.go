package helper

import "github.com/jmoiron/sqlx"

func CommitOrRollback(tx *sqlx.Tx) {
	err := recover()
	if err != nil {
		tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
	}
}
