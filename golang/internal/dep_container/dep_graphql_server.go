package dep_container

import (
	"context"
	"strings"

	graphql2 "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/sarulabs/di"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
	"silverspase/blockchain-api/golang/internal/service/transport/graphql"
	gen_server "silverspase/blockchain-api/golang/internal/service/transport/graphql/generated"
)

const graphqlServerDefName = "graphql-server"

// RegisterGraphqlServer registers Graphql Server dependency.
func RegisterGraphqlServer(builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: graphqlServerDefName,
		Build: func(ctn di.Container) (interface{}, error) {
			resolver := ctn.Get(resolverServiceDefName).(*graphql.Resolver)
			cfg := gen_server.Config{
				Resolvers: resolver,
			}
			server := handler.NewDefaultServer(gen_server.NewExecutableSchema(cfg))

			server.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
				if strings.Contains(err.Error(), "access denied") {
					zap.L().Info("muted access denied err", zap.String("err", err.Error()))
					return graphql2.DefaultErrorPresenter(ctx, err)

				} else if strings.Contains(err.Error(), "must not be null") {
					zap.L().Info("muted must not be null err", zap.String("err", err.Error()))
					return graphql2.DefaultErrorPresenter(ctx, err)
				}

				zap.L().Error("gql error", zap.Any("err", err))
				return graphql2.DefaultErrorPresenter(ctx, err)
			})

			return server, nil
		},
	})
}
