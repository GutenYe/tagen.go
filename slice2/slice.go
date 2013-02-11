package slice2

func Contains(array []interface{}, value interface{}) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func ContainsFunc(array []interface{}, callback func(value interface{}) bool) bool {
	for _, v := range array {
		if callback(v) {
			return true
		}
	}
	return false
}
