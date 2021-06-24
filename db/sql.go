package database

import "gorm.io/gorm"

type SqlDatabase struct {
	db *gorm.DB
}

func (db *SqlDatabase) Create(record *interface{}) error {
	 return db.create(record)
}

func (db *SqlDatabase) connect(config gorm.Dialector) error {
	d,err := gorm.Open(config)
	db.db = d;
	return err
}
func (db *SqlDatabase) findOne(record *interface{}, conditions interface{}) error {
	  db.db.First(record,conditions)
	 return nil;
}

func (db *SqlDatabase) getMany(record *[]interface{}, conditions interface{}) error {
	return nil
}
func (db *SqlDatabase) update(record *interface{}, conditions interface{}) error {
	return nil
}
func (db *SqlDatabase) create(record *interface{}) error {
	return nil
}
func (db *SqlDatabase) delete(conditions interface{}) error {
	return nil
}
