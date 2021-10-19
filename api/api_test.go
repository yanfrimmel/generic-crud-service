package api

import (
	"testing"

	"github.com/yanfrimmel/generic-crud-service/common"
)

func TestCreate(t *testing.T) {
	field := common.Field{Name: "id", Value: 0}
	record := common.Record([]common.Field{field})
	table := common.Table{Name: "test", Records: []common.Record{record}}
	error := Create(table)
	if error != nil {
		t.Error(error)
	}
}
