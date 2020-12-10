package model

import "github.com/neo4j/neo4j-go-driver/neo4j"

func Parse(record neo4j.Record) map[string]interface{} {
	return (record.GetByIndex(0)).(neo4j.Node).Props()
}
