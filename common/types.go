package common

type Record []Field

//Table in the database
type Table struct {
	Name   string
	Record Record
}

// Field has a name a value
type Field struct {
	Name  string
	Value interface{}
}
