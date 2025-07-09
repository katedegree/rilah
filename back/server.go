package main

import (
	"back/domain/entity"
	"back/infrastructure"
	"back/infrastructure/graphql"
	"back/infrastructure/graphql/directive"
	"back/infrastructure/graphql/resolver"
	"back/pkg"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	pkg.LoadEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	orm := infrastructure.NewGorm()
	storage := infrastructure.NewS3()
	validator := infrastructure.NewValidate()
	authUserContext := infrastructure.NewContext[*entity.UserEntity]("authUser")
	httpRequestContext := infrastructure.NewContext[*http.Request]("httpRequest")

	authDirective := directive.NewAuthDirective(orm, authUserContext, httpRequestContext)

	srv := handler.New(graphql.NewExecutableSchema(graphql.Config{
		Resolvers: &resolver.Resolver{
			Orm:                orm,
			Storage:            storage,
			Validator:          validator,
			AuthUserContext:    authUserContext,
			HttpRequestContext: httpRequestContext,
		},
		Directives: graphql.DirectiveRoot{
			Auth: authDirective.Execute,
		},
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := httpRequestContext.Set(r.Context(), r)
		srv.ServeHTTP(w, r.WithContext(ctx))
	}))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
