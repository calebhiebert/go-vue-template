package util

import "github.com/volatiletech/sqlboiler/v4/types"

func StringSliceToInterfaceSlice(s []string) (i []interface{}) {
	for _, str := range s {
		i = append(i, str)
	}

	return
}

func IntSliceToInterfaceSlice(in []int) (i []interface{}) {
	for _, integer := range in {
		i = append(i, integer)
	}

	return
}

func StringSliceContains(s []string, c ...string) bool {
	for _, str := range s {
		for _, chk := range c {
			if chk == str {
				return true
			}
		}
	}

	return false
}

func StringPtrSliceToStringArray(s []*string) types.StringArray {
	var arr types.StringArray

	for _, str := range s {
		if str != nil {
			arr = append(arr, *str)
		}
	}

	return arr
}