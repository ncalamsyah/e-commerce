package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const (
	SECRET = "supersecret"
	bearer = "Bearer"
)

// func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ec echo.Context) error {
// 		jwtToken := SECRET
// 		l := len(bearer)
// 		auth := ec.Request().Header.Get("Authorization")
// 		if len(auth) <= l+1 || strings.EqualFold(auth[:1], bearer) {
// 			return ec.JSON(http.StatusUnauthorized, helper.UnauthorizedResponse)
// 		}
// 		data, err := jwtToken.ValidateToken(auth[l+1:])
// 		if err != nil {
// 			return ec.JSON(http.StatusUnauthorized, err.Error)
// 		}
// 		ctx := ec.Request().Context()
// 		ctx = context.WithValue(ctx, "user", data)
// 		ec.SetRequest(ec.Request().WithContext(ctx))
// 		return next(ec)
// 	}
// }

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET))
}

func ExtractTokenId(e echo.Context) int {
	users := e.Get("user").(*jwt.Token)
	if users.Valid {
		claims := users.Claims.(jwt.MapClaims)
		userId := claims["user_id"].(float64)
		return int(userId)
	}
	return 0
}
