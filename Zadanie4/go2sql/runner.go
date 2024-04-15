package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go2sql/controllers"
	"net/http"
)

func main() {
	echoapp := echo.New()
	echoapp.Use(middleware.CORS())
	db := dbconn()
	pc := controllers.ProductController{}
	bc := controllers.BasketController{}
	echoapp.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(cont echo.Context) error {
			cont.Set("db", db)
			return next(cont)
		}
	})
	echoapp.GET("/", func(mainer echo.Context) error {
		return mainer.String(http.StatusOK, "Serwer gotowy do działania")
	})
	echoapp.GET("/kill", kill)
	echoapp.POST("/product", pc.Addproduct)
	echoapp.PUT("/product/:id", pc.Updateproduct)
	echoapp.DELETE("/product/:id", pc.Deleteproduct)
	echoapp.GET("/product/:id", pc.Getoneproduct)
	echoapp.GET("/product", pc.Getallproducts)
	echoapp.POST("/basket", bc.AddBasket)
	echoapp.GET("/basket/:id", bc.Getbasket)
	echoapp.GET("/basket", bc.Getbaskets)
	echoapp.DELETE("/basket/:id", bc.Deletebasket)
	echoapp.PUT("/basket/:id", bc.Modifybasket)
	echoapp.Start(":22222")
}

//Przykładowe komendy do testowania endpointów znajdują się w pliku testcommands.txt
