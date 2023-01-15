package iteration

const defaultRepeatCount = 5

// Repeat takes a single character as string and repeats it 5 times
func Repeat(character string, count int) (repeated string) {
	if count == 0 {
		count = defaultRepeatCount
	}
	for i := 0; i < count; i++ {
		repeated += character
	}
	return
}
