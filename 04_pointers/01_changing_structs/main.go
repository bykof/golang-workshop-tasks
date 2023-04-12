package main

type Params struct {
	A *float64
	B *float64
}

func sum(params Params) float64 {
	panic("Implement me")
}

func product(params Params) float64 {
	panic("Implement me")
}

func divide(params Params) float64 {
	panic("Implement me")
}

func main() {
	if sum(Params{}) != 0 {
		panic("Wrong")
	}

	//TODO:  Create variable a and b
	if sum(Params{A: a, B: b}) != 6 {
		panic("Wrong")
	}

	if sum(Params{A: a}) != 0 {
		panic("Wrong")
	}

	if product(Params{}) != 0 {
		panic("Wrong")
	}

	if product(Params{A: a, B: b}) != 9 {
		panic("Wrong")
	}

	if divide(Params{A: a, B: b}) != 1 {
		panic("Wrong")
	}
}
