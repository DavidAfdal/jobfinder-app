package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/DavidAfdal/workfinder/pkg/response"
	"github.com/DavidAfdal/workfinder/pkg/route"
	"github.com/DavidAfdal/workfinder/pkg/token"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


type Server struct {
	*echo.Echo
}

func NewServer(serverName string, publicRoutes, privateRoutes []*route.Route, secretKey string) *Server {
	e := echo.New()


	e.Use(
		middleware.Logger(),
		middleware.CORS(),
	)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Welcome to WorkFinder API", nil))
	})

	v1 := e.Group(fmt.Sprintf("/%s/v1", serverName))

	if len(publicRoutes) > 0 {
		for _, route := range publicRoutes {
			v1.Add(route.Methode, route.Path, route.Handler)
		}
	}
	if len(privateRoutes) > 0 {
		for _, route := range privateRoutes {
			v1.Add(route.Methode, route.Path, route.Handler, JWTProtection(secretKey))
		}
	}

	fmt.Println("server running on port 8080")

	return &Server{e}
}


func (srv *Server) Run()  {
	runServer(srv)
	gracefulShutdown(srv)
}

func runServer(srv *Server) {
	go func() {
		err := srv.Start(":8080")
		log.Fatal(err)
	}()
}

func gracefulShutdown(srv *Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal("Server Shutdown:", err)
		}
	}()
}

func JWTProtection(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(token.JwtCustomClaims)
		},
		SigningKey: []byte(secretKey),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "anda harus login untuk mengakses resource ini"))
		},
	})
}



