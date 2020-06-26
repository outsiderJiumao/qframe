package tcolumn

import "time"

func sum(values []int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

var aggregations = map[string]func([]time.Time) time.Time{
	//"sum": sum,
}
