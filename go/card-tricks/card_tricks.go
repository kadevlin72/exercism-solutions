package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < 0 {
		return -1
	}
	if len(slice) < index+1 {
		return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 {
		return append(slice, value)
	}
	if len(slice) < index+1 {
		return append(slice, value)
	}
	slice[index] = value
	return slice

}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	if len(values) == 0 {
		return slice
	}
	var newSlice []int
	for _, value := range values {
		newSlice = append(newSlice, value)
	}
	return append(newSlice, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	for i, _ := range slice {
		if i == index {
			return append(slice[:index], slice[index+1:]...)
		}
	}
	return slice
}
