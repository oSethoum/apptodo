package db

import (
	"app/backend/ent"
	"context"

	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/mattn/go-sqlite3"
)

var Client *ent.Client

func Init() {
	client, err := ent.Open("sqlite3", "file:ent.sqlite?_fk=1")

	if err != nil {
		panic(err)
	}

	err = client.Schema.Create(
		context.Background(),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
		schema.WithGlobalUniqueID(true),
	)

	if err != nil {
		panic(err)
	}

	Client = client
}
