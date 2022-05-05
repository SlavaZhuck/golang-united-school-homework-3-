package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32 = 0.0
	for i := 0; i < len(input)-1; i++ {
		sum += input[i]
	}

	return sum / float32(len(input))
}
