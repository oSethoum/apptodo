package resolvers

import (
	"app/backend/db" 
	"app/backend/ent"
	"app/backend/graph/generated"

	"github.com/99designs/gqlgen/graphql"
	
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	Client *ent.Client
	
	TodoListenners        		    map[*chan *ent.Todo]TodoListenner
	TodoListennersMutext 			  sync.Mutex
	TodosListenners        map[*chan *ent.TodoConnection]TodosListenner
	TodosListennersMutext  sync.Mutex
}

var schema *graphql.ExecutableSchema

func ExecutableSchema() graphql.ExecutableSchema {
	if schema == nil {
		schema = new(graphql.ExecutableSchema)
		*schema = generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{
			Client: db.Client,
			TodoListenners:        			make(map[*chan *ent.Todo]TodoListenner),
			TodoListennersMutext:  			sync.Mutex{},
			TodosListenners:       make(map[*chan *ent.TodoConnection]TodosListenner),
			TodosListennersMutext: sync.Mutex{},
		}})
	}

	return *schema
}
