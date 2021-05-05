module github.com/applegreengrape/red-thread-neo4j

go 1.15

require (
	cloud.google.com/go/firestore v1.5.0 // indirect
	firebase.google.com/go v3.13.0+incompatible
	github.com/neo4j/neo4j-go-driver/v4 v4.2.4
	google.golang.org/api v0.44.0
)

replace github.com/applegreengrape/red-thread-neo4j/loader => ./loader
