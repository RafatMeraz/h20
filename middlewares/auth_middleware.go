package middlewares

import (
	"github.com/RafatMeraz/h20/controllers"
	"github.com/labstack/echo/v4"
	"strings"
)

type AuthMiddleware struct{}

func (AuthMiddleware) AuthVerification(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// getting token from header
		token, err := getAuthorizationTokenFromHeader(c)
		if err != nil {
			return err
		}
		// getting claims from token
		claims, err := controllers.JWTTokenController{}.GetClaimsFromToken(token)
		if err != nil {
			return echo.ErrUnauthorized
		}
		// checking token validation
		validate := controllers.JWTTokenController{}.CheckTokenValidation(&claims)
		if !validate {
			return echo.ErrUnauthorized
		}
		c.Set("user", claims.UserId)
		return next(c)
	}
}

func getAuthorizationTokenFromHeader(c echo.Context) (string, error) {
	authorizationHeader := c.Request().Header.Get("Authorization")

	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return "", echo.ErrUnauthorized
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	return token, nil
}
