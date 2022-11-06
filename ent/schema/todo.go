package schema

import (
	"app/backend/auth"
	"app/backend/ent/privacy"
	"app/backend/entgqlplus"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").Unique().Annotations(entgql.OrderField("TEXT")),
		field.Bool("done").Default(false).Annotations(entgql.OrderField("DONE")),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("todos").Unique(),
	}
}

// Annotations of the Todo.
func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgqlplus.SchemaAnnotation(entgqlplus.Subscription),
	}
}

// Policy defines the privacy policy of the Todo.
func (Todo) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			auth.MutationPrivacy("Todo"),
		},
		Query: privacy.QueryPolicy{
			auth.QueryPrivacy("Todo"),
		},
	}
}
