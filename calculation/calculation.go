package calculation

func Cal(a, b int) int {

	var c int
	if a > 0 && b > 0 {
		c = 1
	} else if a > 0 && b < 0 {
		c = 2
	} else if a < 0 && b > 0 {
		c = 3
	} else {
		c = 4
	}

	switch c {
	case 1:
		c = a * b
	case 2:
		c = b - a
	case 3:
		c = a / b
	case 4:
		c = a + b
	}

	return c

}
