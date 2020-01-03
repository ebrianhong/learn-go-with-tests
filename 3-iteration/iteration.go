package iteration

func Repeat(s string, r int) (output string) {
	for i := 0; i < r; i++ {
		output += s
	}
	return
}
