package leetcode

func MySqrt(x int) int {
	i := 0
	for {
		if i*i == x {
			return i
		} else if i*i > x {
			return i - 1
		}
		i++
	}
}
