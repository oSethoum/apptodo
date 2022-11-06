package schema

import (
	"app/backend/ent/privacy"
	"app/backend/entgqlplus"

	"app/backend/auth"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().NotEmpty().
			Annotations(entgql.OrderField("USERNAME")),

		field.String("email").Unique().NotEmpty().
			Annotations(entgql.OrderField("EMAIL")),

		field.String("password").Sensitive().
			Annotations(entgql.Skip(entgql.SkipType, entgql.SkipWhereInput)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todos", Todo.Type).Annotations(entgql.RelayConnection()),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgqlplus.SchemaAnnotation(entgqlplus.AuthSchema),
	}
}

// Policy defines the privacy policy of the User.
func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			auth.MutationPrivacy("User"),
		},
		Query: privacy.QueryPolicy{
			auth.QueryPrivacy("User"),
		},
	}
}
