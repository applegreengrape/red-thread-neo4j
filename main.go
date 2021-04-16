package main

import (
	"encoding/json"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	//"context"
	//firebase "firebase.google.com/go"
	// "firebase.google.com/go/auth"
	//"google.golang.org/api/option"
)


func query() (interface{}, error) {
	driver, err := neo4j.NewDriver("neo4j://localhost:7687", neo4j.BasicAuth("", "", ""))
	if err != nil {
		return nil, err
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	query, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		q:=fmt.Sprintf(`MATCH p=(a:Node { id:"vKY8Uw" })-[r:friend*1..2]-(x:Node { type:"user" })
		RETURN p`)
		result, err := transaction.Run(q, map[string]interface{}{})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			fmt.Println(result.Record().Values[0])
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return nil, err
	}

	return query.(interface{}), nil
}

func main() {
	res, err:= query()
	if err != nil {}

	resraw, err := json.Marshal(res)
    if err != nil {
    }

	var resData Res
	err = json.Unmarshal(resraw, &resData)
	if err != nil {
	}

	//fmt.Println(resData)

}

type Res struct {
	Records []Record
}

type Record struct{
	Item []Data
}

type Data struct{
	Id string
}