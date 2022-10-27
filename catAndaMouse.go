func catAndMouse(x int32, y int32, z int32) string {
	tempA := x - z
	tempB := y - z

	if tempA < 0 {
		tempA *= -1
	}

	if tempB < 0 {
		tempB *= -1
	}

	if tempA > tempB {
		return "Cat B"
	} else if tempB > tempA {
		return "Cat A"
	}

	return "Mouse C"

}
