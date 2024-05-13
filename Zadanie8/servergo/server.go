package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Do wprowadzenia: dane do linijek 28, 29, 34, 35.
// Wymagane uprawnienie w przypadku google - dostęp do informacji o adresie e-mail
var googleConfiguration = &oauth2.Config{
	RedirectURL:  "http://localhost:22222/googcallback",
	ClientID:     "POPRAWNY CLIENT ID",
	ClientSecret: "POPRAWNY CLIENT SEECRET",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}
var githubConfiguration = &oauth2.Config{
	ClientSecret: "POPRAWNY CLIENT SECRET",
	ClientID:     "POPRAWNY CLIENT ID",
	RedirectURL:  "http://localhost:22222/gitcallback",
	Scopes:       []string{"user:email"},
	Endpoint:     github.Endpoint,
}

func main() {
	backend := echo.New()
	backend.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: []string{"http://localhost:22222", "http://localhost:3000"}, AllowMethods: []string{http.MethodGet, http.MethodPost}}))
	database, err := gorm.Open(sqlite.Open("users.db"))
	if err != nil {
		panic("Nie można ustanowić połączenia z bazą danych")
	}
	database.AutoMigrate(&loginDB{})
	database.AutoMigrate(&oauthDB{})
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
	backend.GET("/google", googleOauth)
	backend.GET("/googcallback", googleCallback)
	backend.POST("/gdata", gdata)
	backend.GET("/git", githubOauth)
	backend.GET("/gitcallback", githubCallback)
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
	return c.Redirect(http.StatusTemporaryRedirect, googleConfiguration.AuthCodeURL("state", oauth2.AccessTypeOffline))
}

func googleCallback(c echo.Context) error {
	responsecode := c.QueryParam("code")
	token, err := googleConfiguration.Exchange(context.Background(), responsecode)
	if err != nil {
		return c.String(http.StatusBadRequest, "Nie udało się zalogować za pomocą Google")
	}
	client := googleConfiguration.Client(context.Background(), token)
	resp, err2 := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	//resp, err2 := client.Get("https://www.googleapis.com/plus/v1/people/me")
	if err2 != nil {
		return c.String(http.StatusInternalServerError, "Nie udało się pobrać informacji o adresie email"+" "+err2.Error())
	}
	defer resp.Body.Close()
	userInfo := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&userInfo)
	responder := new(responder)
	responder.USER = userInfo["email"].(string)
	responder.TOKEN = token.AccessToken
	db := c.Get("db").(*gorm.DB)
	usr := new(oauthDB)
	db.First(&usr, "EMAIL = ?", responder.USER)
	usr.EMAIL = responder.USER
	usr.TOKEN = responder.TOKEN
	usr.PROVIDER = "Google"
	db.Save(usr)
	value := fmt.Sprintf("%d", usr.ID)
	//return c.JSON(http.StatusOK, responder)
	nurl := "http://localhost:3000/ghidden/"
	u, _ := url.Parse(nurl)
	q := u.Query()
	q.Add("id", value)
	u.RawQuery = q.Encode()
	return c.Redirect(http.StatusMovedPermanently, u.String())
}

func gdata(c echo.Context) error {
	req := new(info)
	db := c.Get("db").(*gorm.DB)
	if err := c.Bind(req); err != nil {
		return err
	}
	oauthdb := new(oauthDB)
	err00 := db.First(&oauthdb, "ID = ?", req.ID).Error
	if err00 == gorm.ErrRecordNotFound {
		return c.String(http.StatusBadRequest, "Nie odnaleziono użytkownika o podanych danych")
	}
	responder0 := new(responder)
	responder0.TOKEN = oauthdb.TOKEN
	responder0.USER = oauthdb.EMAIL
	return c.JSON(http.StatusOK, responder0)
}

func githubOauth(c echo.Context) error {
	return c.Redirect(http.StatusTemporaryRedirect, githubConfiguration.AuthCodeURL("state"))
}

func githubCallback(c echo.Context) error {
	responsecode := c.QueryParam("code")
	token, err := githubConfiguration.Exchange(context.Background(), responsecode)
	if err != nil {
		return c.String(http.StatusBadRequest, "Nie udało się zalogować za pomocą GitHub")
	}
	client := githubConfiguration.Client(context.Background(), token)
	resp, err2 := client.Get("https://api.github.com/user/emails")
	if err2 != nil {
		return c.String(http.StatusInternalServerError, "Nie udało się pobrać informacji o adresie email"+" "+err2.Error())
	}
	defer resp.Body.Close()
	userInfo := make([]map[string]interface{}, 0)
	json.NewDecoder(resp.Body).Decode(&userInfo)
	responder := new(responder)
	responder.USER = userInfo[0]["email"].(string)
	responder.TOKEN = token.AccessToken
	db := c.Get("db").(*gorm.DB)
	usr := new(oauthDB)
	db.First(&usr, "EMAIL = ?", responder.USER)
	usr.EMAIL = responder.USER
	usr.TOKEN = responder.TOKEN
	usr.PROVIDER = "GitHub"
	db.Save(usr)
	value := fmt.Sprintf("%d", usr.ID)
	nurl := "http://localhost:3000/ghidden/"
	u, _ := url.Parse(nurl)
	q := u.Query()
	q.Add("id", value)
	u.RawQuery = q.Encode()
	return c.Redirect(http.StatusMovedPermanently, u.String())
}
