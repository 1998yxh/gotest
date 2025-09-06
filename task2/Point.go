package main

func sum(int2 *int) {
	*int2 = *int2 + 10
}

func slice(arr *[]int) {
	for i := range *arr {
		(*arr)[i] = (*arr)[i] * 2
	}
}

func main() {
	a := 10
	sum(&a)
	println(a)

	arr := []int{1, 2, 3}
	slice(&arr)
	for _, v := range arr {
		println(v)
	}
}
