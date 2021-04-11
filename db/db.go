package db

import "github.com/yanfrimmel/generic-crud-service/common"

//Read a record from table
// gets a table name and fields values to search for, in no fields received returns whole table
func Read(tableName string, fields ...common.Field) ([]common.Record, error) {
	record := common.Record(fields)
	records := []common.Record{record}
	return records, nil
}
