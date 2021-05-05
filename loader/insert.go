package loader

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

const uri = "neo4j://localhost:7687"

func CreateNode(id, name, nodeType string) (string, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth("", "", ""))
	if err != nil {
		return "", err
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	node, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run("CREATE (n:Node { id: $id, name: $name, type: $nodeType }) RETURN n.id, n.name, n.type", map[string]interface{}{
			"id":       id,
			"name":     name,
			"nodeType": nodeType,
		})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return "", err
	}

	return node.(string), nil
}

func CreateRel(pid, id string) error {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth("", "", ""))
	if err != nil {
		return err
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	rel, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		query := fmt.Sprintf(`
		MATCH
			(a:Node),
  			(b:Node)
		WHERE a.id = '%s' AND b.id = '%s'
		CREATE (a)-[r:friends]->(b)
		RETURN r
		`, pid, id)
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
