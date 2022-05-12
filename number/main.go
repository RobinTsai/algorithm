package main

// 数值的右移位操作是带符号的，注意不是和二进制中符号位相关，而是和数字的类型相关
// 即：0b1000... >> 1 结果有可能是 0b0100... 也有可能是 0b1100...，结果是跟 0b1000... 这个数值的定义类型相关的

func main() {
	runDivideCases()
	runBinaryStringAddCases()
}
