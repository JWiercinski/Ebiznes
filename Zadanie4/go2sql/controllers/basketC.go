package controllers

import (
	"github.com/labstack/echo/v4"
	"go2sql/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type BasketController struct{}

func (bc *BasketController) AddBasket(sth echo.Context) error {
	basket := new(models.BasketJ)
	dbb := new(models.Basket)
	if err := sth.Bind(basket); err != nil {
		return err
	}
	db := sth.Get("db").(*gorm.DB)
	bugz := 0
	if len(basket.USER) > 0 {
		dbb.USER = basket.USER
	} else {
		bugz = bugz + 1
	}
	if basket.COUNT <= 0 {
		bugz = bugz + 1
	} else {
		dbb.COUNT = basket.COUNT
	}
	if basket.VALUE <= 0 {
		bugz = bugz + 1
	} else {
		checker := strconv.FormatFloat(basket.VALUE, 'f', 2, 64)
		theprice, thebug := strconv.ParseFloat(checker, 64)
		if thebug == nil {
			dbb.VALUE = theprice
		} else {
			bugz = bugz + 1
		}
	}
	if bugz == 0 {
		db.Create(&models.Basket{USER: dbb.USER, COUNT: dbb.COUNT, VALUE: dbb.VALUE})
		returner := string("Stworzono koszyk użytkownika " + dbb.USER + " o wartości " + strconv.FormatFloat(dbb.VALUE, 'f', 2, 64) + ", składający się z " + strconv.Itoa(dbb.COUNT) + " produktów.")
		return sth.String(http.StatusCreated, returner)
	} else {
		return sth.String(http.StatusBadRequest, "Otrzymany koszyk nie spełnia wymagań dołączenia do bazy danych")
	}

	return nil
}
func (bc *BasketController) Getbasket(sth echo.Context) error {
	id := sth.Param("id")
	db := sth.Get("db").(*gorm.DB)
	if len(id) == 0 {
		return sth.String(http.StatusBadRequest, "Nie podano ID")
	}
	index, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return sth.String(http.StatusBadRequest, "Podany indeks nie jest liczbą naturalną...")
	}
	var basket = models.Basket{ID: index}
	err2 := db.First(&basket).Error
	if err2 != nil {
		return sth.String(http.StatusNotFound, "Podany indeks nie istnieje w bazie danych.")
	}
	return sth.String(http.StatusOK, "Oto koszyk o indeksie "+id+": Nazwa użytkownika - "+basket.USER+", Wartość - "+strconv.FormatFloat(basket.VALUE, 'f', 2, 64)+", Ilość produktów - "+strconv.Itoa(basket.COUNT))
}
func (bs *BasketController) Getbaskets(sth echo.Context) error {
	db := sth.Get("db").(*gorm.DB)
	var baskets = []models.BasketJ2{}
	db.Model(&models.Basket{}).Find(&baskets)
	return sth.JSON(http.StatusOK, baskets)
}
func (bc *BasketController) Deletebasket(sth echo.Context) error {
	id := sth.Param("id")
	db := sth.Get("db").(*gorm.DB)
	if len(id) == 0 {
		return sth.String(http.StatusBadRequest, "Nie podano ID")
	}
	index, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return sth.String(http.StatusBadRequest, "Podany indeks nie jest liczbą naturalną...")
	}
	var basket = models.Basket{ID: index}
	err2 := db.First(&basket).Error
	if err2 != nil {
		return sth.String(http.StatusNotFound, "Podany indeks nie istnieje w bazie danych.")
	}
	db.Delete(&basket)
	return sth.String(http.StatusOK, "Usunięto koszyk o indeksie "+id)
}
func (bc *BasketController) Modifybasket(sth echo.Context) error {
	id := sth.Param("id")
	if len(id) == 0 {
		return sth.String(http.StatusBadRequest, "Nie podano ID")
	}
	basker := new(models.BasketJ)
	if err := sth.Bind(basker); err != nil {
		return err
	}
	db := sth.Get("db").(*gorm.DB)
	index, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return sth.String(http.StatusBadRequest, "Podany indeks nie jest liczbą naturalną...")
	}
	var basket = models.Basket{ID: index}
	err2 := db.First(&basket).Error
	if err2 != nil {
		return sth.String(http.StatusNotFound, "Podany indeks nie istnieje w bazie danych")
	}
	if len(basker.USER) > 0 {
		basket.USER = basker.USER
	}
	if basker.VALUE > 0 {
		checker := strconv.FormatFloat(basker.VALUE, 'f', 2, 64)
		theprice, thebug := strconv.ParseFloat(checker, 64)
		if thebug == nil {
			basket.VALUE = theprice
		}
	}
	if basker.COUNT > 0 {
		basket.COUNT = basker.COUNT
	}
	db.Save(&basket)
	return sth.String(http.StatusOK, "Zmodyfikowano koszyk o indeksie "+id+". Nazwa odpowiedzialnego użytkownika to "+basket.USER+", a wartość "+strconv.Itoa(basket.COUNT)+" produktów w nim zawartych to "+strconv.FormatFloat(basket.VALUE, 'f', 2, 64))
}
