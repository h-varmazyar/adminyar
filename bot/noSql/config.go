package noSql

import (
	"github.com/gobuffalo/envy"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var driver neo4j.Driver
var err error

func Run() error {
	username := envy.Get("NEO_USERNAME", "")
	password := envy.Get("NEO_PASSWORD", "")
	configForNeo4j40 := func(conf *neo4j.Config) { conf.Encrypted = false }

	driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth(username, password, ""), configForNeo4j40)
	if err != nil {
		return err
	}

	return addConstraints()
}

func Close() {
	driver.Close()
}

func GetSession() (neo4j.Session, error) {
	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	session, err := driver.NewSession(sessionConfig)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	return session, nil
}
