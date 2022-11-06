//go:build ignore
// +build ignore

package main

import (
	"log"

	"app/backend/entgqlplus"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereInputs(true),
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("../graph/schemas/schema.graphqls"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	eex := entgqlplus.NewExtension(
		entgqlplus.WithConfigPath("../gqlgen.yml"),
		entgqlplus.WithDatabase(entgqlplus.SQLite, "db.sqlite?_fk=1"),
		entgqlplus.WithMutation(true),
		entgqlplus.WithEchoServer(true),
		entgqlplus.WithSubscription(true),
		entgqlplus.WithFileUpload(true),
		entgqlplus.WithJWTAuth(true),
		entgqlplus.WithPrivacy(true),
	)

	opts := []entc.Option{
		entc.Extensions(ex, eex),
		entc.FeatureNames("privacy"),
	}

	if err := entc.Generate("./schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
