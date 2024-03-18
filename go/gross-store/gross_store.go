package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	//first find if item exists in units
	unitQty, unitExists := units[unit]
	if !unitExists {
		return false
	}
	billQty := bill[item]
	bill[item] = billQty + unitQty
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	//first find if item exists in units
	unitQty, unitExists := units[unit]
	if !unitExists {
		return false
	}
	billQty, billExists := bill[item]
	if !billExists {
		return false
	}
	if billQty-unitQty < 0 {
		return false
	}
	if billQty-unitQty == 0 {
		delete(bill, item)
		return true
	}
	bill[item] = billQty - unitQty
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	billQty, billExists := bill[item]
	if !billExists {
		return 0, false
	}
	return billQty, true
}
