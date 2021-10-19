package api

import (
	"fmt"

	"github.com/yanfrimmel/generic-crud-service/common"
	"github.com/yanfrimmel/generic-crud-service/db"
)

// Create a table with fields
func Create(table common.Table) error {
	fmt.Println(table.Name + " Created")
	return nil
}

// Read a record from table
// gets a table name and fields values to search for, in no fields received returns whole table
func Read(tableName string, fields ...common.Field) ([]common.Record, error) {
	return db.Read(tableName, fields...)
}

// Update a record from table
func Update() {
	fmt.Println("Update")
}

// Delete a record from table
func Delete() {
	fmt.Println("Delete")
}
