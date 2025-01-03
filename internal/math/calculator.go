package math

func Calculator(method string, a, b float64) float64 {
	switch method {
	case "sum":
		return Add(a, b)
	case "sub":
		return Subtract(a, b)
	case "div":
		return Divide(a, b)
	case "mult":
		return Multiply(a, b)
	default:
		panic("Método inválido")
	}
}
