package stringsArrCircular

func IsCircular(set []string) bool {
	re := false
	dfs(set, 0, &re)
	return re
}

func dfs(set []string, l int, re *bool) {
	//basecase
	if l == len(set) {
		if l > 0 && isChained(set[l-1], set[0]) {
			*re = true
		}
		return
	}
	for i := l; i < len(set); i++ {
		if l == 0 || isChained(set[l-1], set[i]) {
			set[l], set[i] = set[i], set[l]
			dfs(set, l+1, re)
			set[l], set[i] = set[i], set[l]
		}
	}
}

func isChained(s1 string, s2 string) bool {
	if len(s2) > 0 && s1[len(s1)-1] == s2[0] {
		return true
	}
	return false
}
