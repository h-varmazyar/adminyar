package model

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func Parse(record neo4j.Record) interface{} {
	//if reflect.TypeOf(record.GetByIndex(0)) == reflect.(neo4j.Node) {
	//
	//}
	if node, ok := (record.GetByIndex(0)).(neo4j.Node); ok {
		return node.Props()
	}
	if relation, ok := (record.GetByIndex(0)).(neo4j.Relationship); ok {
		return relation.Props()
	}
	if num, ok := (record.GetByIndex(0)).(int64); ok {
		return num
	}
	return nil
}
