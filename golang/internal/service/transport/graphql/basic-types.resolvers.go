package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	gen_server "silverspase/blockchain-api/golang/internal/service/transport/graphql/generated"
	"time"
)

// Ping is the resolver for the ping field.
func (r *mutationResolver) Ping(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// ServerTime is the resolver for the serverTime field.
func (r *mutationResolver) ServerTime(ctx context.Context) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

// Ping is the resolver for the ping field.
func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// ServerTime is the resolver for the serverTime field.
func (r *queryResolver) ServerTime(ctx context.Context) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns gen_server.MutationResolver implementation.
func (r *Resolver) Mutation() gen_server.MutationResolver { return &mutationResolver{r} }

// Query returns gen_server.QueryResolver implementation.
func (r *Resolver) Query() gen_server.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
