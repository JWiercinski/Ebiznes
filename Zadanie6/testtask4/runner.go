package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"testtask4/controllers"
)

// Dwa scope'y, po jednym w produktach i kategorii. Scope zdefiniowany na spodzie kontrolera
func main() {
	echoapp := echo.New()
	echoapp.Use(middleware.CORS())
	db := dbconn()
	pc := controllers.ProductController{}
	bc := controllers.BasketController{}
	cc := controllers.CategoryController{}
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
	echoapp.POST("/category", cc.Addcategory)
	echoapp.PUT("/category/:id", cc.Renamecategory)
	echoapp.DELETE("/category/:id", cc.DeleteCategory)
	echoapp.GET("/category", cc.Getallcategories)
	echoapp.GET("/category/:id", cc.Getcategory)
	echoapp.Start(":22222")
}

//Przykładowe komendy do testowania endpointów znajdują się w pliku testcommands.txt
