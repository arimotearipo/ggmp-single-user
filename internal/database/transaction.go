package database

import "errors"

func (db *Database) BeginTx() error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	db.TX = tx
	return nil
}

func (db *Database) CommitTx() error {
	if db.TX == nil {
		return errors.New("no transaction in progress")
	}

	err := db.TX.Commit()
	if err != nil {
		return err
	}

	// when done with transaction, reset value to nil
	db.TX = nil

	return nil
}

func (db *Database) RollbackTx(f func()) error {
	if db.TX == nil {
		return errors.New("no transaction in progress")
	}

	err := db.TX.Rollback()
	if err != nil {
		return err
	}

	// when done with transaction, reset value to nil
	db.TX = nil

	f() // execute callback

	return nil
}
