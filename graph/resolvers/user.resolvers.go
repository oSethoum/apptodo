package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/backend/ent"
	"context"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.Create().SetInput(input).Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*ent.User, error) {
	entity, _ := r.Client.User.Get(ctx, id)
	return entity, ent.FromContext(ctx).User.DeleteOneID(id).Exec(ctx)
}

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context, id int) (*ent.User, error) {
	return r.Client.User.Get(ctx, id)
}
