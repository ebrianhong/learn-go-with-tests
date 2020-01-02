package iteration

func Repeat(s string, r int) (output string) {
	output = ""
	for i := 0; i < r; i++ {
		output += s
	}
	return
}
