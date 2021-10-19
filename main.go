package main

import (
	"fmt"

	"github.com/yanfrimmel/generic-crud-service/api"
	"github.com/yanfrimmel/generic-crud-service/common"
)

func main() {
	fmt.Println("Hello World")
	field := common.Field{Name: "id", Value: 0}
	record := common.Record([]common.Field{field})
	table := common.Table{Name: "test", Records: []common.Record{record}}
	api.Create(table)
	api.Read("test")
	api.Update()
	api.Delete()
}
