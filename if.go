package gox

var (
	_ = If[int]
	_ = Ifx[int]
)

// If 模块条件表达式，主要用法是减少大括号
func If[T any](condition bool, result T) (t T) {
	if condition {
		t = result
	}

	return
}

// Ifx 模拟三元表达式，主要用法是减少大括号
func Ifx[T any](condition bool, first T, second T) (t T) {
	return Ternary[T](condition, first, second)
}

// Ternary 模拟三元表达式，主要用法是减少大括号
func Ternary[T any](condition bool, first T, second T) (t T) {
	if condition {
		t = first
	} else {
		t = second
	}

	return
}
