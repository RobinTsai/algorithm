package main

import "fmt"

// NumberDivide
// 题目：不适用乘法、除法实现两个 整数 的除法操作，输入 15, 2 返回 7
// 注意：1. 可以输入负数 2. 注意溢出
type NumberDivide [3]int // with testcases: num1, num2, exp

// 看起来简单，但细节很多
// 先说整体方法：
//   法1：按减法计数来做，效率会差，如 2^32 - 2
//   法2：用移位的方法对减数按指数增长，计数
// 细节：
//   一是负数的问题
//   然后是最小的负值转换为正数时会发生溢出，如 int32 型范围在 -2^31 ~ 2^31-1，由 -2^31 转换为正数是会发生溢出
// 所以思路是，由正数转换为负数，用移位+减法来计数
func (NumberDivide) Exec(a, b int) int {
	toNeg := func(in int) int {
		if in > 0 {
			return -in
		}
		return in
	}

	divideWithShift := func(num1, num2 int) (int, int) {
		base := 1
		for num1 < num2<<1 {
			num2 = num2 << 1
			base = base << 1
		}
		return base, num1 - num2
	}

	negA := toNeg(a)
	negB := toNeg(b)
	isFlagDiff := (a > 0 && b < 0) || (a < 0 && b > 0)

	result := 0
	curBase := 0
	rest := negA
	for rest <= negB {
		curBase, rest = divideWithShift(rest, negB)
		result += curBase
	}

	if isFlagDiff {
		return 0 - result
	}
	return result
}

/*----------------- test ----------------------*/
var numberDivideCases = []NumberDivide{
	{15, 2, 15 / 2},
	{-15, 2, -15 / 2},
	{4, 2, 4 / 2},
	{-4, 2, -4 / 2},
	{(-1 << 63), 2, (-1 << 63) / 2},
}

func runDivideCases() {
	fmt.Println("runDivideCases:")
	for _, testcase := range numberDivideCases {
		res := testcase.Exec(testcase[0], testcase[1])
		if res != testcase[2] {
			fmt.Println("error: ", testcase, "rel:", res)
		} else {
			fmt.Println("succeed: ", testcase)
		}
	}
	fmt.Println()
}
