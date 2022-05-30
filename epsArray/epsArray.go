package epsArray

// func indexOfInt(a []int, n int) int {
// 	for i, v := range a {
// 		if v == n {
// 			return i
// 		}
// 	}
// 	return -1
// }
//
// func indexOfInt64(a []int64, n int64) int {
// 	for i, v := range a {
// 		if v == n {
// 			return i
// 		}
// 	}
// 	return -1
// }

func IndexOfString(a []string, n string) int {
	for i, v := range a {
		if v == n {
			return i
		}
	}
	return -1
}
