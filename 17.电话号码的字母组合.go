/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
/*
	dfs
	end condition: layer end, append str to result
	every layer, append to str, going to next layer, restore str
*/
func letterCombinations(digits string) []string {
	re := []string{}
	if digits == "" {
		return re
	}
	rl := []rune(digits)
	m := prepare()
	// dfs every layer
	dfsStr(rl, 0, m, "", &re)
	return re
}

func dfsStr(rl []rune, layer int, m map[rune]string, cur string, re *[]string) {
	// end statement
	if layer == len(rl) {
		*re = append(*re, cur)
		return
	}

	// recursion
	curStr := m[rl[layer]]
	for i := 0; i < len(curStr); i++ {
		// append to str
		cur = cur + string(curStr[i])
		// dfs
		dfsStr(rl, layer+1, m, cur, re)
		// restore
		cur = cur[:len(cur)-1]
	}
}

func prepare() map[rune]string {
	m := make(map[rune]string)
	m['2'] = "abc"
	m['3'] = "def"
	m['4'] = "ghi"
	m['5'] = "jkl"
	m['6'] = "mno"
	m['7'] = "pqrs"
	m['8'] = "tuv"
	m['9'] = "wxyz"
	return m
}

// @lc code=end

