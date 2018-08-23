package bracketPermutation

func Permu(n int) []string {
	sb := []rune{}
	re := &[]string{}
	if n <= 0 {
		return *re
	}
	dfs(n, 0, 0, 0, sb, re)
	return *re
}

func dfs(n, lf, rt, le int, sb []rune, re *[]string) {
	//basecase
	if n*2 == le {
		*re = append(*re, string(sb))
		return
	}
	//recur
	if lf < n {
		sb = append(sb, '(')
		dfs(n, lf+1, rt, le+1, sb, re)
		sb = sb[:len(sb)-1]
	}
	if rt < lf {
		sb = append(sb, ')')
		dfs(n, lf, rt+1, le+1, sb, re)
		sb = sb[:len(sb)-1]
	}
}
