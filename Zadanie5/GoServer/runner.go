package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math"
	"net/http"
)

func main() {
	products := initproducts()
	backend := echo.New()
	backend.Use(middleware.CORS())
	database := connectDB()
	backend.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			context.Set("database", database)
			context.Set("inventory", products)
			return next(context)
		}
	})
	backend.GET("/", home)
	backend.GET("/kill", kill)
	//backend.GET("/test", teststuff)
	backend.GET("/products", sendProducts)
	backend.POST("/payment", logPayment)
	backend.Start(":22222")
}

func home(context echo.Context) error {
	return context.String(http.StatusOK, "Serwer uruchomiony.")
}

func sendProducts(context echo.Context) error {
	products := context.Get("inventory")
	return context.JSON(http.StatusOK, products)
}

func logPayment(context echo.Context) error {
	paydata := new(payment)
	if err := context.Bind(&paydata); err != nil {
		return context.String(http.StatusBadRequest, "Problem przy przyjmowaniu danych. Skontaktuj się z obsługą")
	}
	database := context.Get("database").(*gorm.DB)
	error := string("")
	if len(paydata.USER) < 1 {
		error = errorParser(error, "Nazwa użytkownika jest pusta. ")
	}
	if paydata.METHOD != "CARD" && paydata.METHOD != "BANKTRANSFER" && paydata.METHOD != "PAYPAL" {
		error = errorParser(error, "Nieobsługiwana metoda płatności. ")
	}
	if paydata.AMOUNT <= 0 {
		error = errorParser(error, "Opłata mniejsza niż zero")
	}
	if checker := math.Round(paydata.AMOUNT*100) / 100; checker != paydata.AMOUNT {
		error = errorParser(error, "Opłata nie jest w odpowiednim formacie - zawiera więcej niż dwie istotne cyfry w formacie dziesiętnym")
	}
	if len(error) == 0 {
		database.Create(&paymentdb{USER: paydata.USER, METHOD: paydata.METHOD, AMOUNT: paydata.AMOUNT})
		return context.String(http.StatusOK, "OK")
	} else {
		return context.String(http.StatusBadRequest, error)
	}
}

func errorParser(error string, message string) string {
	error = error + message
	return error
}

type product struct {
	ID    int     `json:"ID"`
	NAME  string  `json:"NAME"`
	DEV   string  `json:"DEV"`
	DESC  string  `json:"DESC"`
	PRICE float64 `json:"PRICE"`
}

type payment struct {
	USER   string  `json:"USER"`
	METHOD string  `json:"METHOD"`
	AMOUNT float64 `json:"AMOUNT"`
}

type paymentdb struct {
	gorm.Model
	ID     uint64 `gorm:"primaryKey"`
	USER   string
	METHOD string
	AMOUNT float64
}

func kill(c echo.Context) error {
	err := c.Echo().Shutdown(context.Background())
	if err != nil {
		if err != http.ErrServerClosed {
			c.Echo().Logger.Fatal("Serwer wyłączony")
		}
	}
	return nil
}

func connectDB() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("payment.db"))
	if err != nil {
		panic("Nie można ustanowić połączenia z bazą danych")
	}
	database.AutoMigrate(&paymentdb{})
	return database
}

func initproducts() [10]product {
	var product1 = product{
		ID:    1,
		NAME:  "A Collection of Bad Moments",
		DEV:   "Sky Trail",
		DESC:  "Odnajdź się w samym centrum trudnych sytuacji - i wyjdź z nich cało",
		PRICE: 14.49,
	}
	var product2 = product{
		ID:    2,
		NAME:  "Miasmata",
		DEV:   "Ion FX",
		DESC:  "Eksploruj zapomnianą wyspę, znajdź lek na tajemniczą chorobę, a przede wszystkim przetrwaj",
		PRICE: 53.99,
	}
	var product3 = product{
		ID:    3,
		NAME:  "Dead Secret",
		DEV:   "Robot Invader",
		DESC:  "Rozwiąż zagadkę zabójcy zanim staniesz się następną ofiarą",
		PRICE: 53.99,
	}
	var product4 = product{
		ID:    4,
		NAME:  "Unearthed: Trail of Ibn Battuta",
		DEV:   "Semaphore",
		DESC:  "Poczuj się jak ubogi kuzyn Nathana Drake'a",
		PRICE: 17.99,
	}
	var product5 = product{
		ID:    5,
		NAME:  "Kholat",
		DEV:   "IMGN.PRO",
		DESC:  "Odkryj przyczyny tragedii na Przełęczy Diatłowa - i wyjdź z tego cało",
		PRICE: 49.99,
	}
	var product6 = product{
		ID:    6,
		NAME:  "Flatout 3",
		DEV:   "Team 6",
		DESC:  "Kultowa seria powraca w budżetowej odsłonie",
		PRICE: 8.99,
	}
	var product7 = product{
		ID:    7,
		NAME:  "Pineview Drive - Homeless",
		DEV:   "VIS Games",
		DESC:  "Sequel niszowego horroru, w pełnoprawnej odsłonie",
		PRICE: 64.99,
	}
	var product8 = product{
		ID:    8,
		NAME:  "Night Book",
		DEV:   "Wales Interactive",
		DESC:  "Thriller FMV o okultystycznym zabarwieniu. Występuje między innymi rewelacyjny Colin Salmon",
		PRICE: 67.99,
	}
	var product9 = product{
		ID:    9,
		NAME:  "雪女",
		DEV:   "Chilla's Art",
		DESC:  "Masz jedną szansę, by uwolnić dzieci porwane przez Yuki Onnę, w nowym retro horrorze stylizowanym na lata 90.",
		PRICE: 22.99,
	}
	var product10 = product{
		ID:    10,
		NAME:  "Balan Wonderworld",
		DEV:   "Square Enix",
		DESC:  "Piękne światy i nieprzemyślany gameplay - to wszystko znajdziesz w tej platformówce",
		PRICE: 165.00,
	}
	products := [10]product{product1, product2, product3, product4, product5, product6, product7, product8, product9, product10}
	return products
}

/*
func teststuff(context echo.Context) error {
	error := ""
	value1 := 3.2222222
	if checker := math.Round(value1*100) / 100; checker != value1 {
		boo := strconv.FormatFloat(checker, 'f', 2, 64)
		error = errorParser(error, "Pierwsza wartość nie jest równa wartości zmienionej. "+boo)
	}
	value2 := 3.22
	if checker := math.Round(value2*100) / 100; checker != value2 {
		error = errorParser(error, "Druga wartość nie jest równa wartości zmienionej.")
	} else {
		x := fmt.Sprintf("%d", math.Round(value2*100)/100)
		error = errorParser(error, x)
	}
	return context.String(http.StatusOK, error)
}
*/
