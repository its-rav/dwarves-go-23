package helpers

import "strconv"

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ConvertToIntArray(arr []string) []int64 {
	var intArr []int64

	for _, v := range arr {
		intVal, _ := strconv.ParseInt(v, 10, 64)
		intArr = append(intArr, intVal)
	}

	return intArr
}

func ConvertToFloatArray(arr []string) []float64 {
	var floatArr []float64

	for _, v := range arr {
		floatVal, _ := strconv.ParseFloat(v, 64)
		floatArr = append(floatArr, floatVal)
	}

	return floatArr
}

func ConvertToStringArray(arr []interface{}) []string {
	var strArr []string

	for _, v := range arr {
		strArr = append(strArr, v.(string))
	}

	return strArr
}

func ConvertIntArrToStringArr(arr []int64) []string {
	var strArr []string

	for _, v := range arr {
		strArr = append(strArr, strconv.FormatInt(v, 10))
	}

	return strArr
}

func ConvertFloatArrToStringArr(arr []float64) []string {
	var strArr []string

	for _, v := range arr {
		strArr = append(strArr, strconv.FormatFloat(v, 'f', 2, 64))
	}

	return strArr
}
