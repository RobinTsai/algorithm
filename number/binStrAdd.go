package main

import (
	"fmt"
)

// BinaryStringAdd
// 以字符串表示的二进制，进行加法运算
// 如： "1101" + "1111" = "11100"
// 误区：
//   不能用转换称数值，因为字符串可以很长，数值可能覆盖不了字符串的表达值
// 方法：
//   从右向左累加
// 细节：1. 注意进位
type BinaryStringAdd [3]string

func (*BinaryStringAdd) Exec(str1, str2 string) string {
	buf := make([]byte, 0, 20)

	getVal := func(str string, idx int) int {
		if idx < 0 {
			return 0
		}
		return int(str[idx] - '0')
	}

	idx1, idx2 := len(str1)-1, len(str2)-1
	overflow := 0
	for idx1 >= 0 || idx2 >= 0 || overflow == 1 {
		val1 := getVal(str1, idx1)
		val2 := getVal(str2, idx2)
		switch val1 + val2 + overflow {
		case 1:
			buf = append(buf, '1')
			overflow = 0
		case 2:
			buf = append(buf, '0')
			overflow = 1
		case 3:
			buf = append(buf, '1')
			overflow = 1
		default:
			panic(fmt.Sprint("no way:", idx1, idx2, overflow))
		}
		idx1--
		idx2--
	}
	reverse := func(in []byte) []byte {
		for i := 0; i < len(in)/2; i++ {
			temp := in[i]
			in[i] = in[len(in)-1-i]
			in[len(in)-1-i] = temp
		}
		return in
	}

	return string(reverse(buf))
}

/*--------------------- test ------------------*/
var binaryStringAddCases = []*BinaryStringAdd{
	{"1101", "1111", "11100"},
	{"1111111111111111111111111111", "1", "10000000000000000000000000000"},
	{"1111111111111111111111111111", "1111111111111111111111111111", "11111111111111111111111111110"},
}

func runBinaryStringAddCases() {
	fmt.Println("runBinaryStringAddCases:")
	for _, v := range binaryStringAddCases {
		res := v.Exec(v[0], v[1])
		if res != v[2] {
			fmt.Println("error:", v, res)
		} else {
			fmt.Println("succeed:", v)
		}
	}
	fmt.Println()
}
