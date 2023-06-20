package directives

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
)

func HasAuth(ctx context.Context, obj interface{}, next graphql.Resolver, scopes []string) (res interface{}, err error) {
	fmt.Println("HasAuth")
	return next(ctx)
}

func HasPermission(ctx context.Context, obj interface{}, next graphql.Resolver, action string, resource string) (res interface{}, err error) {
	fmt.Println("HasPermission")
	return next(ctx)
}
