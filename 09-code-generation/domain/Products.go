
	package domain
	type Products []Product

	func (items *Products) IndexOf(item Product) int {
		for idx, p := range *items {
			if p == item {
				return idx
			}
		}
		return -1
	}

	func (items *Products) Includes(item Product) bool {
		return items.IndexOf(item) != -1
	}

	func (items *Products) Any(criteria func(Product) bool) bool {
		for _, item := range *items {
			if criteria(item) {
				return true
			}
		}
		return false
	}
