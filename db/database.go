package database



type DiskStorage interface { 
	connect(interface{}) bool
	get(record *interface{},conditions interface{}) error
	getMany(record *[]interface{},conditions interface{})  error 
	update(record *interface{},conditions interface{}) error
	create(record *interface{}) error
	delete(conditions interface{}) error
}