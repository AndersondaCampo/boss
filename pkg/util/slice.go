package util

func ContainsStr(arr []string, value string) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}

func Contains(arr []interface{}, value interface{}) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}
