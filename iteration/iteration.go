package iteration

func Repeat(r string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += r
	}
	return repeated
}
