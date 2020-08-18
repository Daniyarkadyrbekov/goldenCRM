package handlers

import (
	"strconv"

	"github.com/goldenCRM.git/lib/models"
)

func getComplexCondition(flat *models.Flat) string {
	condition := ""

	if flat.PriceMax != 0 {
		condition += "price_max <= " + strconv.Itoa(flat.PriceMax) + " AND price_max <> 0"
		flat.PriceMax = 0
	}

	if flat.Square != 0 {
		condition += " AND square >= " + strconv.Itoa(flat.Square)
		flat.Square = 0
	}

	return condition
}
