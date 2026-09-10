package client

import "context"

type contextKey string

const AuthorizationContextKey contextKey = "authorization"

func SetAuthorizationToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, AuthorizationContextKey, token)
}

func GetAuthorizationToken(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(AuthorizationContextKey).(string)
	return token, ok
}
