package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go2sql/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type CategoryController struct{}

func (cc *CategoryController) Addcategory(sth echo.Context) error {
	json0 := new(models.CategoryJ)
	category := new(models.Category)
	if err := sth.Bind(json0); err != nil {
		return err
	}
	db := sth.Get("db").(*gorm.DB)
	if len(json0.NAME) > 0 {
		category.NAME = json0.NAME
		err2 := db.Create(&models.Category{NAME: category.NAME}).Error
		if err2 != nil {
			return sth.String(http.StatusBadRequest, "Istnieje już kategoria o podanej nazwie")
		} else {
			return sth.String(http.StatusOK, "Dodano kategorię o nazwie "+category.NAME)
		}
	} else {
		return sth.String(http.StatusBadRequest, "Nie podano nazwy kategorii")
	}
}

func (cc *CategoryController) Renamecategory(sth echo.Context) error {
	id := sth.Param("id")
	if len(id) == 0 {
		return sth.String(http.StatusBadRequest, "Nie podano ID")
	}
	ct := new(models.CategoryJ)
	if err := sth.Bind(ct); err != nil {
		return err
	}
	db := sth.Get("db").(*gorm.DB)
	index, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return sth.String(http.StatusBadRequest, "Podany indeks nie jest liczbą naturalną...")
	}
	var category = models.Category{ID: index}
	err2 := db.First(&category).Error
	if err2 != nil {
		return sth.String(http.StatusNotFound, "Podany indeks nie istnieje w bazie danych")
	}
	if len(ct.NAME) > 0 {
		category.NAME = ct.NAME
	}
	db.Save(&category)
	return sth.String(http.StatusOK, "Zmodyfikowano nazwę kategorii o indeksie "+id+" na "+category.NAME)
}
func (cc *CategoryController) DeleteCategory(sth echo.Context) error {
	id := sth.Param("id")
	db := sth.Get("db").(*gorm.DB)
	if len(id) == 0 {
		return sth.String(http.StatusBadRequest, "Nie podano ID")
	}
	index, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return sth.String(http.StatusBadRequest, "Podany indeks nie jest liczbą naturalną...")
	}
	if index == 1 {
		return sth.String(http.StatusBadRequest, "Nie można usunąć kategorii domyślnej!")
	}
	var category = models.Category{ID: index}
	err2 := db.First(&category).Error
	if err2 != nil {
		return sth.String(http.StatusNotFound, "Podany indeks nie istnieje w bazie danych.")
	}
	var prod = models.Product{CATEGORYID: index}
	if err3 := db.Model(&models.Product{}).Where("category_id = ?", index).First(&prod).Error; err3 != nil {
		db.Unscoped().Delete(&category)
		//Wyłączony soft delete, aby pozwalać na tworzenie nowych elementów o odpowiednich nazwach
		return sth.String(http.StatusOK, "Usunięto kategorię o indeksie "+id)
	} else {
		return sth.String(http.StatusBadRequest, "Nadal istnieją produkty powiązane z kategorią, którą próbowano usunąć...")
	}
}

/*
	func (cc *CategoryController) Getallcategories(sth echo.Context) error {
		db := sth.Get("db").(*gorm.DB)
		var categ = []models.CategoryJ2{}
		db.Model(&models.Category{}).Find(&categ)
		return sth.JSON(http.StatusOK, &categ)
	}
*/
func (cc *CategoryController) Getallcategories(sth echo.Context) error {
	db := sth.Get("db").(*gorm.DB)
	var categ = []models.Category{}
	db.Scopes(PRLD).Find(&categ)
	var categJ2 = []models.CategoryJ2{}
	for _, cat := range categ {
		var prodsJ2 = []models.ProductJ2{}
		for _, prod := range cat.PRODUCTS {
			prodsJ2 = append(prodsJ2, models.ProductJ2{
				ID:         int(prod.ID),
				NAME:       prod.NAME,
				PRICE:      prod.PRICE,
				CATEGORYID: int(prod.CATEGORYID),
			})
		}
		categJ2 = append(categJ2, models.CategoryJ2{
			ID:       int(cat.ID),
			NAME:     cat.NAME,
			PRODUCTS: prodsJ2,
		})
	}
	return sth.JSON(http.StatusOK, &categJ2)
}
func (cc *CategoryController) Getcategory(sth echo.Context) error {
	/*
		id := sth.Param("id")
		db := sth.Get("db").(*gorm.DB)
		if len(id) == 0 {
			return sth.String(http.StatusBadRequest, "Nie podano ID")
		}
		index, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return sth.String(http.StatusBadRequest, "Podany indeks nie jest liczbą naturalną...")
		}
		var categ = models.Category{ID: index}
		err2 := db.First(&categ).Error
		if err2 != nil {
			return sth.String(http.StatusNotFound, "Podany indeks nie istnieje w bazie danych.")
		}
		return sth.String(http.StatusOK, "Oto kategoria o indeksie "+id+": Nazwa - "+categ.NAME)
	*/
	id := sth.Param("id")
	db := sth.Get("db").(*gorm.DB)
	if len(id) == 0 {
		return sth.String(http.StatusBadRequest, "Nie podano ID")
	}
	index, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return sth.String(http.StatusBadRequest, "Podany indeks nie jest liczbą naturalną...")
	}
	var categ = models.Category{ID: index}
	err2 := db.Scopes(PRLD).First(&categ).Error
	if err2 != nil {
		return sth.String(http.StatusNotFound, "Podany indeks nie istnieje w bazie danych.")
	}
	productsStr := ""
	for _, prod := range categ.PRODUCTS {
		productsStr += fmt.Sprintf("Product ID: %d, Name: %s, Price: %.2f, CategoryID: %d\n", prod.ID, prod.NAME, prod.PRICE, prod.CATEGORYID)
	}
	return sth.String(http.StatusOK, fmt.Sprintf("Oto kategoria o indeksie %s: Nazwa - %s, Produkty:\n%s", id, categ.NAME, productsStr))
}

// Sprawny scope, użyty 2 razy...
func PRLD(db *gorm.DB) *gorm.DB { return db.Preload("PRODUCTS") }
