package dep_container

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/sarulabs/di"
	"go.uber.org/zap"
	"silverspase/blockchain-api/golang/internal/config"
	"silverspase/blockchain-api/golang/internal/service/transport/http_handlers"
)

const (
	httpServerDefName = "http-server"
)

// RegisterHTTPServer registers HTTP Server dependency.
func RegisterHTTPServer(builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: httpServerDefName,
		Build: func(ctn di.Container) (interface{}, error) {
			router := initRouter()

			cfg := ctn.Get(configDefName).(*config.Config)
			graphqlServer := ctn.Get(graphqlServerDefName).(*handler.Server)
			httpHandlers := ctn.Get(httpHandlersDefName).(http_handlers.HttpProvider)

			router.Handle("/", playground.Handler("GraphQL playground", "/query"))
			router.Handle("/query", graphqlServer)
			router.HandleFunc("/liveness", httpHandlers.LivenessCheck)

			zap.L().Info(
				"connected to GraphQL playground",
				zap.String("graphql address", fmt.Sprintf("http://localhost:%s/", cfg.HttpPort)),
			)
			zap.L().Fatal("", zap.Error(http.ListenAndServe(":"+cfg.HttpPort, router)))

			return router, nil
		},
	})
}

// RunHTTPServer runs HTTP Server dependency.
func (c Container) RunHTTPServer() {
	c.container.Get(httpServerDefName)
}

func initRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	}).Handler)

	return router
}
