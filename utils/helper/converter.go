package helper

import "strconv"

func S2I(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func I2S(i int) string {
	s := strconv.Itoa(i)
	return s
}

func S2I64(s string) int64 {
	i64, _ := strconv.ParseInt(s, 10, 64)
	return i64
}

func I642S(i int64) string {
	return strconv.FormatInt(i, 10)
}

func S2Bool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}

func StrArr2Int64Arr(strArr []string) []int64 {
	int64Arr := make([]int64, 0, len(strArr))
	for _, s := range strArr {
		int64Arr = append(int64Arr, S2I64(s))
	}
	return int64Arr
}

func Int64Arr2StrArr(int64Arr []int64) []string {
	strArr := make([]string, 0, len(int64Arr))
	for _, i := range int64Arr {
		strArr = append(strArr, I642S(i))
	}
	return strArr
}
