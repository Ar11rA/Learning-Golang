package main

func pop(arr []int, x int) (int, []int) {
	n := arr[x]
	return n, append(arr[:x], arr[x+1:]...)[:len(arr)-1]
}

func shift(arr []int) (int, []int) {
	n := arr[0]
	return n, append(arr[:0], arr[1:]...)
}

func unshift(arr []int, x int) []int {
	return append([]int{x}, arr...)
}

func reverse(arr []int) []int {
	var narr []int
	for i := len(arr) - 1; i >= 0; i-- {
		narr = append(narr, arr[i])
	}
	return narr
}

func maps(arr []int, f func(int) int) []int {
	for i, v := range arr {
		arr[i] = f(v)
	}
	return arr
}

func reduce(arr []int, f func(int, int) int) int {
	var r int
	for _, v := range arr {
		r = f(r, v)
	}
	return r
}
