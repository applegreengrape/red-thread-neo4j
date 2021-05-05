package loader

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func ClearDB() (error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth("", "", ""))
	if err != nil {
		return err
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	rel, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		query := fmt.Sprintf(`MATCH (n)
		DETACH DELETE n`)
		result, err := transaction.Run(query, map[string]interface{}{})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return err
	}
	fmt.Printf("%v", rel)
	return nil
}
