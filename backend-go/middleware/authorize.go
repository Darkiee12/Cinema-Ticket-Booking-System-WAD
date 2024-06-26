package middleware

import (
	"cinema/common"
	"cinema/component/appctx"
	"cinema/module/user/usermodel"
	"cinema/plugin/tokenprovider/jwt"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//Authorization : Bearn{token}
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

// RequireAuth 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user_id to find from DB
func RequireAuth(appCtx appctx.AppContext, authStore AuthenStore) func(ctx *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())
	return func(c *gin.Context) {
		tracer := appCtx.GetTracer()
		ctx1, span := (*tracer).Start(c.Request.Context(), "RequireAuth")

		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		//db := appCtx.GetMyDBConnection()
		//store := userstore.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		//user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})
		//if err != nil {
		//	//c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		//	panic(err)
		//}

		user, err := authStore.FindUser(ctx1, map[string]interface{}{"id": payload.UserId})
		if err != nil {
			panic(err)
		}
		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(false)
		span.End()

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}

func NotRequireAuth(appCtx appctx.AppContext, authStore AuthenStore) func(ctx *gin.Context) {

	tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			c.Next()
			return
		}

		//db := appCtx.GetMyDBConnection()
		//store := userstore.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			c.Next()
			return
		}

		//user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})
		//if err != nil {
		//	//c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		//	panic(err)
		//}

		user, err := authStore.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if user.Status == 0 {
			c.Next()
			return
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
