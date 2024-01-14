package pkg

// Cov(X, Y) = Σ(Xi-mean_x)(Yi-mean_y) / n
func Covariance(mean_x, mean_y float64, array_x, array_y []int) float64 {
	length := len(array_x)
	i := 0
	var result float64
	for i < length {
		result += ((float64(array_x[i]) - mean_x) * (float64(array_y[i]) - mean_y))
		i++
	}
	return result / float64(length)
}
