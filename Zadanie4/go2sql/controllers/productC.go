package controllers

import (
	"github.com/labstack/echo/v4"
	"go2sql/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ProductController struct{}

func (pc *ProductController) Addproduct(sth echo.Context) error {
	categoryname := "INNE"
	product := new(models.ProductJ)
	dbprod := new(models.Product)
	if err := sth.Bind(product); err != nil {
		return err
	}
	db := sth.Get("db").(*gorm.DB)
	errors := 0
	if product.CATEGORYID > 1 {
		otherc := new(models.Category)
		otherc.ID = product.CATEGORYID
		if err := db.Model(&models.Category{}).First(&otherc).Error; err == nil {
			dbprod.CATEGORYID = product.CATEGORYID
			categoryname = otherc.NAME
		}
	}
	if len(product.NAME) > 0 {
		dbprod.NAME = product.NAME
	} else {
		errors = errors + 1
	}
	if product.PRICE <= 0 {
		errors = errors + 1
	} else {
		checker := strconv.FormatFloat(product.PRICE, 'f', 2, 64)
		theprice, thebug := strconv.ParseFloat(checker, 64)
		if thebug == nil {
			dbprod.PRICE = theprice
		} else {
			errors = errors + 1
		}
	}
	if errors == 0 {
		db.Create(&models.Product{NAME: dbprod.NAME, PRICE: dbprod.PRICE, CATEGORYID: dbprod.CATEGORYID})
		returner := string("Stworzono produkt " + dbprod.NAME + " o cenie " + strconv.FormatFloat(product.PRICE, 'f', 2, 64) + ", należący do kategorii " + categoryname)
		return sth.String(http.StatusCreated, returner)
	} else {
		return sth.String(http.StatusBadRequest, "Otrzymano albo pustą nazwę, albo cenę która nie jest ceną...")
	}
}

func (pc *ProductController) Updateproduct(sth echo.Context) error {
	categoryname := "INNE"
	id := sth.Param("id")
	if len(id) == 0 {
		return sth.String(http.StatusBadRequest, "Nie podano ID")
	} else {
		product := new(models.ProductJ)
		if err := sth.Bind(product); err != nil {
			return err
		}
		db := sth.Get("db").(*gorm.DB)
		index, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return sth.String(http.StatusBadRequest, "Podany indeks nie jest liczbą naturalną...")
		} else {
			var prod = models.Product{ID: index}
			err2 := db.First(&prod).Error
			if err2 != nil {
				return sth.String(http.StatusNotFound, "Podany indeks nie istnieje w bazie danych")
			}
			if len(product.NAME) > 0 {
				prod.NAME = product.NAME
			}
			if product.CATEGORYID > 1 {
				otherc := new(models.Category)
				otherc.ID = product.CATEGORYID
				if err := db.Model(&models.Category{}).First(&otherc).Error; err == nil {
					prod.CATEGORYID = product.CATEGORYID
					categoryname = otherc.NAME
				}
			}
			if product.PRICE > 0 {
				checker := strconv.FormatFloat(product.PRICE, 'f', 2, 64)
				theprice, thebug := strconv.ParseFloat(checker, 64)
				if thebug == nil {
					prod.PRICE = theprice
				}
			}
			db.Save(&prod)
			returner := string("Zmodyfikowano produkt o indeksie " + id + ". Jego obecna nazwa to " + prod.NAME + ", cena to " + strconv.FormatFloat(prod.PRICE, 'f', 2, 64) + ", a kategoria to " + categoryname)
			return sth.String(http.StatusOK, returner)
		}
	}
}
func (pc *ProductController) Deleteproduct(sth echo.Context) error {
	id := sth.Param("id")
	db := sth.Get("db").(*gorm.DB)
	if len(id) == 0 {
		return sth.String(http.StatusBadRequest, "Nie podano ID")
	}
	index, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return sth.String(http.StatusBadRequest, "Podany indeks nie jest liczbą naturalną...")
	}
	var prod = models.Product{ID: index}
	err2 := db.First(&prod).Error
	if err2 != nil {
		return sth.String(http.StatusNotFound, "Podany indeks nie istnieje w bazie danych.")
	}
	db.Delete(&prod)
	return sth.String(http.StatusOK, "Usunięto element o indeksie "+id)
}
func (pc *ProductController) Getoneproduct(sth echo.Context) error {
	id := sth.Param("id")
	db := sth.Get("db").(*gorm.DB)
	if len(id) == 0 {
		return sth.String(http.StatusBadRequest, "Nie podano ID")
	}
	index, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return sth.String(http.StatusBadRequest, "Podany indeks nie jest liczbą naturalną...")
	}
	var prod = models.Product{ID: index}
	err2 := db.First(&prod).Error
	if err2 != nil {
		return sth.String(http.StatusNotFound, "Podany indeks nie istnieje w bazie danych.")
	}
	otherc := new(models.Category)
	otherc.ID = prod.CATEGORYID
	db.Model(&models.Category{}).First(&otherc)
	return sth.String(http.StatusOK, "Oto element o indeksie "+id+": Nazwa - "+prod.NAME+", Cena - "+strconv.FormatFloat(prod.PRICE, 'f', 2, 64)+", Kategoria - "+otherc.NAME)
}
func (pc *ProductController) Getallproducts(sth echo.Context) error {
	db := sth.Get("db").(*gorm.DB)
	var prods = []models.ProductJ2{}
	db.Model(&models.Product{}).Find(&prods)
	return sth.JSON(http.StatusOK, &prods)
}
