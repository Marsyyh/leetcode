package stringSpacePermutation

func SpacePerm(s string) []string {
	sb := []rune{}
	re := &[]string{}
	r := []rune(s)
	dfs(r, 0, re, sb)
	return *re
}

func dfs(r []rune, level int, re *[]string, sb []rune) {
	// basecase
	if level == len(r)-1 {
		sb = append(sb, r[level])
		*re = append(*re, string(sb))
		sb = sb[:len(sb)-1]
		return
	}

	// recursion
	sb = append(sb, r[level])
	sb = append(sb, '_')
	dfs(r, level+1, re, sb)
	sb = sb[:len(sb)-1]
	dfs(r, level+1, re, sb)
}
