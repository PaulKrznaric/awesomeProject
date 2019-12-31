package iteration

const repeatedCount = 5

func Repeat(r string) string {
	var repeated string
	for i := 0; i < repeatedCount; i++ {
		repeated += r
	}
	return repeated
}
