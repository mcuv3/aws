package database

type SqlDatabase struct{}

func (db *SqlDatabase) Create(record *interface{}) error {

	return nil
}

func (db *SqlDatabase) connect(interface{}) bool {
	return true
}
func (db *SqlDatabase) get(record *interface{}, conditions interface{}) error {
	return nil
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
