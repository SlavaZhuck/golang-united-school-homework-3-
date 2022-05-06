package homework

// package main

func reverse(input []int64) (result []int64) {
	//Place your code here
	s := make([]int64, len(input))
	for i := 0; i < len(input); i++ {
		s[i] = input[len(input)-i-1]
	}
	return s
}

// func main() {

// 	primes := [6]int64{2, 3, 5, 7, 11, 13}

// 	var s []int64 = primes[0:6]

// 	fmt.Println(reverse(s))
// }
