package sortarray

import (
	"sort"

	"github.com/its-rav/dwarves-go-23/ass-d2/helpers"
)

func IsSortTypeValid(sortType string) bool {
	// check if sortType in [int,string,mixed]
	return helpers.Contains([]string{"int", "string", "float"}, sortType)
}

func Sort(sortType string, arr []string) []string {
	if !IsSortTypeValid(sortType) {
		panic("invalid sort type")
	}

	switch sortType {
	case "int":
		intArr := helpers.ConvertToIntArray(arr)
		sort.Slice(intArr, func(i, j int) bool { return intArr[i] < intArr[j] })
		return helpers.ConvertIntArrToStringArr(intArr)

	case "float":
		floatArr := helpers.ConvertToFloatArray(arr)
		sort.Float64s(floatArr)
		return helpers.ConvertFloatArrToStringArr(floatArr)
	case "string":
		sort.Strings(arr)
		return arr
	default:
		panic("invalid sort type")
	}
}
