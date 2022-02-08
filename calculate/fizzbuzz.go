package calculate

func FizzBuzz(num int) bool {
	if Fizz(num) && Buzz(num) {
		return true
	}
	return false
}
