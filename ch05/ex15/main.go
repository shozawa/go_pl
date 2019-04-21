package vargs

func Max(vals ...int) (max int) {
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return
}

func Min(vals ...int) (min int) {
	if len(vals) < 1 {
		return 0
	}
	min = vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return
}

func Max2(vals ...int) (int, bool) {
	if len(vals) < 1 {
		return 0, false
	}
	return Max(vals...), true
}

func Min2(vals ...int) (int, bool) {
	if len(vals) < 1 {
		return 0, false
	}
	return Min(vals...), true
}
