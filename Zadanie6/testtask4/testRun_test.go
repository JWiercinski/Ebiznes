package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"testtask4/controllers"
	"testtask4/models"
)

func TestRun(t *testing.T) {
	t.Run("Sprawdzenie poprawności uruchomienia lub stworzenia bazy danych", func(t *testing.T) {
		//8(8) asercje
		echoapp := echo.New()
		db := dbconn()
		assert.FileExists(t, "./godb.db", "go.db nie istnieje")
		echoapp.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(cont echo.Context) error {
				cont.Set("db", db)
				assert.Equal(t, db, cont.Get("db"), "go.db nie zostaje właściwie zapisana w kontekście")
				return next(cont)
			}
		})
		assert.Equal(t, false, db.Migrator().HasTable("koszyki"), "Baza danych posiada tabele których posiadać nie miała")
		assert.Equal(t, true, db.Migrator().HasTable("products"), "Baza danych nie zawiera tabeli products")
		assert.Equal(t, true, db.Migrator().HasTable("baskets"), "Baza danych nie zawiera tabeli baskets")
		if assert.Equal(t, true, db.Migrator().HasTable("categories"), "Baza danych nie zawiera tabeli categories") {
			c1 := new(models.Category)
			db.Model(&models.Category{}).First(&c1)
			assert.Equal(t, uint64(1), c1.ID)
			assert.Equal(t, "INNE", c1.NAME)
		}

	})
	//
	t.Run("Sprawdzenie działania terminacji programu", func(t *testing.T) {
		//3(11) asercje
		echoapp := echo.New()
		echoapp.Start("22222")
		killer := kill
		resp, err := http.Get("http://localhost:22222")
		assert.NotEmpty(t, killer, "Funkcja terminacyjna nie została wykonana")
		assert.Nil(t, resp, "Aplikacja echo zwróciła odpowiedź")
		assert.NotEmpty(t, err, "Nie otrzymano błędu, wbrew oczekiwaniom")
	})
	t.Run("Sprawdzenie poprawności struktur danych", func(t *testing.T) {
		//8(19) asercje
		product := new(models.ProductJ)
		assert.Empty(t, product.NAME, "Nazwa produktu została zainicjalizowana przed dokonaniem jakiegokolwiek działania")
		assert.Empty(t, product.CATEGORYID, "Kategoria produktu została zainicjalizowana przed dokonaniem jakiegokolwiek działania")
		assert.Empty(t, product.PRICE, "Wartość produktu została zainicjalizowana przed dokonaniem jakiegokolwiek działania")
		product.NAME = "Computer"
		wrongid := uint64(3000)
		product.CATEGORYID = wrongid
		product.PRICE = 12
		productDB := new(models.Product)
		assert.Empty(t, productDB.ID, "Id zostało nadane bez kontaktu z bazą danych")
		assert.Empty(t, productDB.PRICE, "Cena została nadana mimo braku ingerencji użytkownika")
		assert.Empty(t, productDB.NAME, "Nazwa została nadana wbrew użytkownikowi")
		assert.Empty(t, productDB.CATEGORYID, "Kategoria została nadana, mimo że defaultowa wartość ma się pojawić dopiero w bazie danych")
		productDB.CATEGORYID = product.CATEGORYID
		assert.Equal(t, product.CATEGORYID, productDB.CATEGORYID, "Kategorie nie zostały właściwie przekazane")
	})
	t.Run("Test zwrócenia wszystkich kategorii", func(t *testing.T) {
		// 4(23) asercje
		echoapp := echo.New()
		db := dbconn()
		cc := new(controllers.CategoryController)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := echoapp.NewContext(req, rec)
		c.Set("db", db)
		if assert.NoError(t, cc.Getallcategories(c), "Wystąpił błąd w wywołaniu funkcji Getallcategories") {
			assert.Equal(t, http.StatusOK, rec.Code, "Otrzymano kod inny niż HTTP status OK")
			if assert.NotEmpty(t, rec.Body, "Puste body odpowiedzi serwera") {
				assert.Contains(t, fmt.Sprintf("%v", rec.Body), "{\"ID\":1,\"NAME\":\"INNE\"", "Nie zwrócono defaultowej kategorii")
			}
		}
	})
	t.Run("Test zwrócenia pojedynczej kategorii", func(t *testing.T) {
		//14(37) asercji
		echoapp := echo.New()
		cc := new(controllers.CategoryController)
		db := dbconn()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := echoapp.NewContext(req, rec)
		c.Set("db", db)
		c.SetParamNames("id")
		c.SetParamValues("1")
		if assert.NoError(t, cc.Getcategory(c), "Wystąpił błąd w wywołaniu funkcji Getcategory") {
			assert.Equal(t, http.StatusOK, rec.Code, "Otrzymano kod inny niż HTTP status OK")
			if assert.NotEmpty(t, rec.Body, "Puste body odpowiedzi serwera") {
				if assert.Contains(t, fmt.Sprintf("%v", rec.Body), "Oto kategoria o indeksie", "Nie zwrócono żadnej kategorii") {
					assert.Contains(t, fmt.Sprintf("%v", rec.Body), ("1"), "Zwrócona kategoria nie jest kategorią o indeksie 1")
					assert.Contains(t, fmt.Sprintf("%v", rec.Body), ("INNE"), "Zwrócona kategoria nie jest kategorią domyślną")

				}
			}
		}
		req = httptest.NewRequest(http.MethodGet, "/", nil)
		rec = httptest.NewRecorder()
		n := echoapp.NewContext(req, rec)
		n.Set("db", db)
		n.SetParamNames("id")
		n.SetParamValues("as-2")
		if assert.NoError(t, cc.Getcategory(n), "Wystąpił błąd w wywołaniu funkcji Getcategory") {
			assert.Equal(t, http.StatusBadRequest, rec.Code, "Otrzymano kod inny niż HTTP status BadRequest, przy niewłaściwym parametrze id")
			if assert.NotEmpty(t, rec.Body, "Serwer nie zwrócił body odpowiedzi") {
				assert.Equal(t, "Podany indeks nie jest liczbą naturalną...", fmt.Sprintf("%v", rec.Body), "Nie otrzymano oczekiwanej odpowiedzi")
			}
		}
		req = httptest.NewRequest(http.MethodGet, "/", nil)
		rec = httptest.NewRecorder()
		g := echoapp.NewContext(req, rec)
		g.Set("db", db)
		g.SetParamNames("id")
		g.SetParamValues("200000000000009")
		if assert.NoError(t, cc.Getcategory(g), "Wystąpił błąd w wywołaniu funkcji Getcategory") {
			assert.Equal(t, http.StatusNotFound, rec.Code, "Otrzymano kod inny niż HTTP status StatusNotFound, przy niewłaściwym indeksie zapytania")
			if assert.NotEmpty(t, rec.Body, "Serwer nie zwrócił body odpowiedzi") {
				assert.Equal(t, "Podany indeks nie istnieje w bazie danych.", fmt.Sprintf("%v", rec.Body), "Nie otrzymano oczekiwanej odpowiedzi")
			}
		}
	})
	t.Run("Próba dodania niewłaściwego produktu do bazy danych", func(t *testing.T) {
		// 10(47) asercji
		echoapp := echo.New()
		db := dbconn()
		wrongproduct := new(models.ProductJ)
		wrongproduct.PRICE = -1
		wrongproduct.NAME = "ASDF"
		prodj, bug := json.Marshal(wrongproduct)
		pc := new(controllers.ProductController)
		if assert.Empty(t, bug, "Impossible to serialize the JSON representation of product") {
			req := httptest.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(prodj))
			rec := httptest.NewRecorder()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := echoapp.NewContext(req, rec)
			c.Set("db", db)
			if assert.NoError(t, pc.Addproduct(c), "Wystąpił błąd w wywołaniu funkcji Getcategory") {
				if assert.NotEmpty(t, rec.Body, "Serwer nie zwrócił body odpowiedzi") {
					assert.Equal(t, http.StatusBadRequest, rec.Code, "Otrzymano inny kod http niż BadRequest")
					assert.Equal(t, "Otrzymano albo pustą nazwę, albo cenę która nie jest ceną...", fmt.Sprintf("%v", rec.Body), "Nie otrzymano oczekiwanej odpowiedzi")
				}
			}
		}
		wrongproduct.PRICE = 12
		wrongproduct.NAME = ""
		prodj2, err2 := json.Marshal(wrongproduct)
		if assert.Empty(t, err2, "Impossible to serialize the JSON representation of product") {
			req := httptest.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(prodj2))
			rec := httptest.NewRecorder()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			n := echoapp.NewContext(req, rec)
			n.Set("db", db)
			if assert.NoError(t, pc.Addproduct(n), "Wystąpił błąd w wywołaniu funkcji Getcategory") {
				if assert.NotEmpty(t, rec.Body, "Serwer nie zwrócił body odpowiedzi") {
					assert.Equal(t, http.StatusBadRequest, rec.Code, "Otrzymano inny kod http niż BadRequest")
					assert.Equal(t, "Otrzymano albo pustą nazwę, albo cenę która nie jest ceną...", fmt.Sprintf("%v", rec.Body), "Nie otrzymano oczekiwanej odpowiedzi")
				}
			}
		}
	})
	t.Run("Próba modyfikacji koszyka bez indeksu", func(t *testing.T) {
		// 4(51) asercji
		echoapp := echo.New()
		db := dbconn()
		req := httptest.NewRequest(http.MethodPut, "/product", bytes.NewBuffer(nil))
		rec := httptest.NewRecorder()
		n := echoapp.NewContext(req, rec)
		n.Set("db", db)
		bc := new(controllers.BasketController)
		if assert.NoError(t, bc.Modifybasket(n), "Wystąpił błąd w wywołaniu funkcji Modifybasket") {
			if assert.NotEmpty(t, rec.Body, "Serwer nie zwrócił body odpowiedzi") {
				assert.Equal(t, http.StatusBadRequest, rec.Code, "Otrzymano inny kod http niż BadRequest")
				assert.Equal(t, "Nie podano ID", fmt.Sprintf("%v", rec.Body), "Nie otrzymano oczekiwanej odpowiedzi")
			}
		}
	})
}
