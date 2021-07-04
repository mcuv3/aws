package auth

func inArr(arr []string, value string) bool {
	isIn := false
	for _, v := range arr {
		isIn = v == value
		if isIn { break }
	}
	return isIn
}

func inMap(m map[interface{}]interface{},key,val interface{}) {
	for k,v := range m { 
		 if k == val || v == key { 

		 }
	}
}