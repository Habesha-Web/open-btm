package users

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
)

// User struct to hold user information from context
type User struct {
	ID   int
	Role string
}

// Context key for storing user information
type contextKey string

const userContextKey = contextKey("user")

// Directive function to check user role against a list of roles
func HasProjectRoleDirective(ctx context.Context, obj interface{}, next graphql.Resolver, roles []string) (interface{}, error) {
	user, ok := ctx.Value(userContextKey).(*User)
	if !ok {
		return nil, errors.New("unauthorized")
	}

	// Check if user's role is in the list of allowed roles
	roleAllowed := false
	for _, role := range roles {
		if user.Role == role {
			roleAllowed = true
			break
		}
	}

	if !roleAllowed {
		return nil, errors.New("forbidden")
	}

	return next(ctx)
}
