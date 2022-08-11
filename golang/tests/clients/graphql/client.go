package graphql

import (
	"context"

	"github.com/shurcooL/graphql"
)

// Client represents GraphQL client.
type Client struct {
	client *graphql.Client
}

// NewClient creates new GraphQL client.
func NewClient(url string) *Client {
	return &Client{
		client: graphql.NewClient(
			url, nil,
		),
	}
}

// Query runs GraphQL query.
func (c Client) Query(query interface{}, variables map[string]interface{}) error {
	return c.client.Query(context.Background(), query, variables)
}

// Mutation runs GraphQL mutation.
func (c Client) Mutation(mutation interface{}, variables map[string]interface{}) error {
	return c.client.Mutate(context.Background(), mutation, variables)
}
