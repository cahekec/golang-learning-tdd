package iteration

func Repeat(s string, c int) string {
	var repeated string

	for i := 0; i < c; i++ {
		repeated = repeated + s
	}

	return repeated
}
