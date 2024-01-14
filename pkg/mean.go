package pkg

// calculate the avarge sum over the len
func Mean(data []int) float64 {
	datalen := float64(len(data))
	var datasum float64
	for _, n := range data {
		datasum += float64(n)
	}
	mean := datasum / datalen
	return mean
}