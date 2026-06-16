package generateparentheses

func GenerateParentheses(n int) []string {
	var res []string
	var generate func(*[]string, string, int, int)

	generate = func(ss *[]string, s string, l, r int) {
		if l == 0 && r == 0 {
			*ss = append(*ss, s)
			return
		}

		if l > 0 {
			generate(ss, s+"(", l-1, r+1)
		}

		if r > 0 {
			generate(ss, s+")", l, r-1)
		}
	}
	return res
}
