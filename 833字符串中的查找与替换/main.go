package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findReplaceString("abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "ffff"}))
}
func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	replaceMap := make(map[int][]string)
	for i, index := range indices {
		if strings.HasPrefix(s[index:], sources[i]) {
			replaceMap[index] = []string{sources[i], targets[i]}
		}
	}
	var res strings.Builder
	i := 0
	for i < len(s) {
		if replace, ok := replaceMap[i]; ok {
			res.WriteString(replace[1])
			i += len(replace[0])
		} else {
			res.WriteByte(s[i])
			i++
		}
	}
	return res.String()
}

// if indices的排序时从小到大的 可以使用以下方法
// func findReplaceString(s string, indices []int, sources []string, targets []string) (res string) {
// 	i, j := 0, 0
// 	for i < len(s) {
// 		if i == indices[j] {
// 			flag := false
// 			for m, n := i, 0; m < len(s) && n < len(sources[j]); m, n = m+1, n+1 {
// 				if s[m] != sources[j][n] {
// 					flag = true
// 					break
// 				}
// 			}
// 			if !flag {
// 				res += targets[j]
// 				i += len(sources[j])
// 				j++
// 			} else {
// 				res += string(s[i : i+len(sources[j])])
// 				i += len(sources[j])
// 				j++
// 			}
// 		} else {
// 			res += string(s[i])
// 			i++
// 		}
// 	}
// 	return
// }
