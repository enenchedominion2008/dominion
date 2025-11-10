package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	menu := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}

	if item, exists := menu[order]; exists {
		return item.preptime
	}

	// Return 404 for items not found as specified in the requirements
	return 404
}
