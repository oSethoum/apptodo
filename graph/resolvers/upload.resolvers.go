package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"io"
	"os"

	"github.com/99designs/gqlgen/graphql"
)

// Upload is the resolver for the upload field.
func (r *mutationResolver) Upload(ctx context.Context, input graphql.Upload) (bool, error) {
	out, err := io.ReadAll(input.File)
	if err != nil {
		return false, err
	}
	err = os.WriteFile(input.Filename, out, 0666)
	return true, err
}
