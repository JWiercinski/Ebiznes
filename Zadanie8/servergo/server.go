package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	backend := echo.New()
	backend.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: []string{"http://localhost:22222", "http://localhost:3000"}, AllowMethods: []string{http.MethodGet, http.MethodPost}}))
	database, err := gorm.Open(sqlite.Open("users.db"))
	if err != nil {
		panic("Nie można ustanowić połączenia z bazą danych")
	}
	database.AutoMigrate(&loginDB{})
	backend.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(cont echo.Context) error {
			cont.Set("db", database)
			return next(cont)
		}
	})
	backend.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "Serwer gotowy do działania") })
	backend.GET("/kill", killInstance)
	backend.POST("/register", register)
	backend.POST("/classic", classicLogin)
	backend.POST("/google", googleOauth)
	err2 := backend.Start(":22222")
	if err2 != nil {
		fmt.Println(err)
	}

}

func killInstance(c echo.Context) error {
	err := c.Echo().Shutdown(context.Background())
	if err != nil {
		if err != http.ErrServerClosed {
			c.Echo().Logger.Fatal("Serwer wyłączony")
		}
	}
	return nil
}

func register(c echo.Context) error {
	user := new(registerJ)
	if err := c.Bind(user); err != nil {
		return err
	}
	db := c.Get("db").(*gorm.DB)
	if len(user.USER) < 2 || len(user.USER) > 30 || len(user.NAME) < 3 || len(user.NAME) > 30 || len(user.EMAIL) < 7 || len(user.EMAIL) > 50 || len(user.PASSWORD) < 10 || len(user.PASSWORD) > 50 {
		return c.String(http.StatusBadRequest, "Podane elementy nie spełniają wymogów długości")
	}
	if strings.Contains(user.EMAIL, "@") != true || strings.Contains(user.EMAIL, ".") != true {
		return c.String(http.StatusBadRequest, "Email nie posiada małpy bądź kropki...")
	}
	err00 := db.First(&loginDB{USER: user.USER}).Error
	if !errors.Is(err00, gorm.ErrRecordNotFound) {
		fmt.Println(err00)
		return c.String(http.StatusBadRequest, "Istnieje już użytkownik o podanej nazwie")
	}
	newpassword := hash(user.PASSWORD)
	token := genToken()
	db.Create(&loginDB{USER: user.USER, NAME: user.NAME, EMAIL: user.EMAIL, PASSWORD: newpassword, TOKEN: token})
	responderSON := new(responder)
	responderSON.TOKEN = token
	responderSON.USER = user.USER
	return c.JSON(http.StatusCreated, responderSON)
}

func hash(password string) string {
	hasz := sha256.New()
	hasz.Write([]byte(password))
	return hex.EncodeToString(hasz.Sum(nil))
}

func genToken() string {
	rand.Seed(time.Now().UnixNano())
	const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZðÐ1234567890ąĄŻżŹźĆćśŚΩłŁəƏóÓ"
	b := make([]byte, 30)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}
	return fmt.Sprintf("%d", b)
}

func classicLogin(c echo.Context) error {
	user := new(loginJ)
	if err := c.Bind(user); err != nil {
		return err
	}
	db := c.Get("db").(*gorm.DB)
	pass := user.PASSWORD
	hashed := hash(pass)
	userdb := new(loginDB)
	err00 := db.First(&userdb, "USER = ? AND PASSWORD = ?", user.USER, hashed).Error
	if err00 == gorm.ErrRecordNotFound {
		return c.String(http.StatusBadRequest, "Nie odnaleziono użytkownika o podanych danych")
	}
	userdb.USER = user.USER
	token := genToken()
	db.First(&userdb)
	userdb.TOKEN = token
	db.Save(&userdb)
	responderSON := new(responder)
	responderSON.TOKEN = token
	responderSON.USER = user.USER
	return c.JSON(http.StatusCreated, responderSON)
}

func googleOauth(c echo.Context) error {
	return nil
}
