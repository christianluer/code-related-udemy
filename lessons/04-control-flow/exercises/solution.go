//go:build solution

package exercises

import "fmt"

func Bucket(x int) string {
	if x >= 0 && x <= 100 {
		return "between 0 and 100"
	} else if x >= 101 && x <= 200 {
		return "between 101 and 200"
	} else if x >= 201 && x <= 250 {
		return "between 201 and 250"
	}
	return "out of range"
}

func DescribePair(x, y int) string {
	switch {
	case x < 4 && y < 4:
		return "x and y are both less than 4"
	case x > 6 && y > 6:
		return "x and y are both greater than 6"
	case x >= 4 && x <= 6:
		return "x is greater than or equal to 4 and less than or equal to 6"
	case y != 5:
		return "y is not 5"
	default:
		return "none of the previous cases were met"
	}
}

func Odds(n int) []int {
	var out []int
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			continue
		}
		out = append(out, i)
	}
	return out
}

func FirstMultiple(start, k int) int {
	i := start
	for {
		if i%k == 0 {
			break
		}
		i++
	}
	return i
}

func Grid(n int) []string {
	var out []string
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			out = append(out, fmt.Sprintf("%d,%d", i, j))
		}
	}
	return out
}

func IndexValues(xi []int) []string {
	var out []string
	for i, v := range xi {
		out = append(out, fmt.Sprintf("%d:%d", i, v))
	}
	return out
}

func AgeOf(m map[string]int, name string) string {
	if age, ok := m[name]; ok {
		return fmt.Sprintf("%s is %d", name, age)
	}
	return name + " is not in the map"
}
