package database

type NoSqlDatabase struct{}

func (db *NoSqlDatabase) Create(record *interface{}) error {

	return nil
}

func (db *NoSqlDatabase) connect(interface{}) bool {
	return true
}
func (db *NoSqlDatabase) get(record *interface{}, conditions interface{}) error {
	return nil
}
func (db *NoSqlDatabase) getMany(record *[]interface{}, conditions interface{}) error {
	return nil
}
func (db *NoSqlDatabase) update(record *interface{}, conditions interface{}) error {
	return nil
}
func (db *NoSqlDatabase) create(record *interface{}) error {
	return nil
}
func (db *NoSqlDatabase) delete(conditions interface{}) error {
	return nil
}
