package pkg

import (
	"strconv"
)

// conver string list to int list
func StringToInt(str []string) ([]int, error) {
	var res []int
	for _, n := range str {
		if n != "" {
			num, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			res = append(res, num)
		}

	}
	return res, nil
}
