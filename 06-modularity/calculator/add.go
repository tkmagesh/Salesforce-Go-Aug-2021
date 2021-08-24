package calculator

var OperationCount int

func Add(x, y int) int {
	OperationCount++
	return x + y
}
