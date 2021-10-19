package db

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yanfrimmel/generic-crud-service/common"
)

var dbPath string = "." + string(filepath.Separator) + "db"

func init() {
	error := os.Mkdir(dbPath, 0777)
	if !os.IsExist(error) {
		fmt.Println(error)
		panic(error)
	}
}

// Create a table with fields in a new file
// if a table with that name already exists return error
func Create(table common.Table) error {
	fmt.Println(table.Name + " Created")
	return nil
}

// Read a record from table
// gets a table name and fields values to search for, in no fields received returns whole table
func Read(tableName string, fields ...common.Field) ([]common.Record, error) {
	record := common.Record(fields)
	records := []common.Record{record}
	return records, nil
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
