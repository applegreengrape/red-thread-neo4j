package main

import (
	"encoding/json"
	"fmt"

	"github.com/applegreengrape/red-thread-neo4j/loader"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"

	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
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
		q := fmt.Sprintf(`MATCH (a:Node {id:"vKY8Uw"} ), 
			(x:Node {id: "EXIhQM"}),
			p = shortestPath((a)-[*]-(x))
  			RETURN p`)
		result, err := transaction.Run(q, map[string]interface{}{})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return nil, err
	}

	return query.(interface{}), nil
}

type Users []struct {
	Id string 
	Name string 
}

type Rels []struct {
	Id string 
	Pid string 
}

func loadNode(name string) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./svc.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		// todo
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		// todo
	}

	users, err := client.Collection("nodes").Doc(name).Get(ctx)
	if err != nil {
		// todo
	}
	user := users.Data()

	rawUser, err := json.Marshal(user["data"])
	if err != nil {
	}

	var userData Users
	err = json.Unmarshal(rawUser, &userData)
	if err != nil {
	}

	for  _, user := range userData {
		res, err := loader.CreateNode(user.Id, user.Name, name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}

	defer client.Close()
}

func loadRel(){
	ctx := context.Background()
	sa := option.WithCredentialsFile("./svc.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		// todo
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		// todo
	}
	rel, err := client.Collection("rel").Doc("data").Get(ctx)
	if err != nil {
		//
	}
	rels := rel.Data()

	rawRels, err := json.Marshal(rels["data"])
	if err != nil {
	}

	var r Rels
	err = json.Unmarshal(rawRels, &r)
	if err != nil {
	}

	for _, rel := range r {
		fmt.Println("create rels: ",rel.Id, rel.Pid)
		err := loader.CreateRel(rel.Id, rel.Pid)
		if err != nil {
			//
		}
	}

	defer client.Close()
}

func reload(){
	loader.ClearDB()
	loadNode("users")
	loadNode("friends")
	loadRel()
}

func main() {
	reload()

	/*
	res, err := query()
	if err != nil {
		//
	}
	p :=res.(neo4j.Path)
	fmt.Println(p)
	fmt.Println(p.Nodes[1].Props["name"])
	*/

}
