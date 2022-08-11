package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	gen_models "silverspase/blockchain-api/golang/internal/service/transport/graphql/models"
)

// Hello is the resolver for the hello field.
func (r *queryResolver) Hello(ctx context.Context, name string) (*gen_models.HelloResponse, error) {
	return &gen_models.HelloResponse{
		Response: fmt.Sprintf("Hello, %s", name),
	}, nil
}

// Hello2 is the resolver for the hello2 field.
func (r *queryResolver) Hello2(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s", name), nil
}
