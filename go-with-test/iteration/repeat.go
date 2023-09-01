package iteration

func Repeat(s string, num int) string {

	var repeated string
	for i := 0; i < num; i++ {
		repeated += s
	}
	return repeated
}
