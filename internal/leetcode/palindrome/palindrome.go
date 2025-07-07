package palindrome

func IsPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	if x%10 == 0 && x != 0 {
		return false
	}

	reversed := 0

	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return x == reversed || x == reversed/10
}
