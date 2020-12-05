package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

const (
	cookieName = "golden_crm_test_7"
	ttl        = 365 * 24 * time.Hour
)

func SignIn() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(200, "signin.html", gin.H{})
	}
}

func Authorize(l *zap.Logger, _ *gorm.DB) func(c *gin.Context) {
	l = l.With(zap.String("method", "auth.IsAuthorized"))
	return func(c *gin.Context) {

		user, ok := c.GetPostForm("name")
		if !ok {
			l.Error("no name in form")
			c.Redirect(http.StatusTemporaryRedirect, "/")
			c.AbortWithStatus(http.StatusTemporaryRedirect)
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			l.Error("no password in form")
			c.Redirect(http.StatusTemporaryRedirect, "/")
			c.AbortWithStatus(http.StatusTemporaryRedirect)
		}
		if user != "user" || password != "password" {
			l.Error("middlewareErr",
				zap.String("err", "Unauthorized"),
				zap.String("path", c.Request.URL.Path),
				zap.String("user", user),
				zap.String("password", password))
			c.Redirect(http.StatusTemporaryRedirect, "/")
			c.AbortWithStatus(http.StatusTemporaryRedirect)
		}
		c.SetCookie(cookieName, "someCookie", int(ttl.Seconds()), "/", "localhost", true, true)
		c.Redirect(http.StatusFound, "/auth")
	}
}

func IsAuthorized(l *zap.Logger, database *gorm.DB) func(c *gin.Context) {

	l = l.With(zap.String("method", "auth.IsAuthorized"))

	return func(c *gin.Context) {
		cookie, err := c.Cookie(cookieName)
		if err != nil {
			l.Error("get cookie err", zap.Error(err))
			c.Redirect(http.StatusTemporaryRedirect, "/")
			c.AbortWithStatus(http.StatusTemporaryRedirect)
		}

		l.Info("successful authorize", zap.String("cookie", cookie))
		c.Next()
	}
}