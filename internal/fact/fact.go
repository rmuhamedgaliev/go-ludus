package fact

func Fact(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * Fact(number-1)
	}
}
