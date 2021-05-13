/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
/*
	dfs
*/
func generateParenthesis(n int) []string {
	re := &[]string{}
	// dfsPt(0, n, 0, "", re)
	dfsPrune(n, 0, 0, "", re)
	return *re
}

// using left and right number to check if need prune
func dfsPrune(n, l, r int, cur string, re *[]string) {
	// ending
	if len(cur) == 2*n {
		*re = append(*re, cur)
		return
	}
	// prune left bracket need smaller than n
	if l < n {
		dfsPrune(n, l+1, r, cur+"(", re)
	}

	// prune right bracket need smaller or eq to left
	if r < l {
		dfsPrune(n, l, r+1, cur+")", re)
	}
}

// using num for remind close bracket count
// re for result holding
// cur for each dfs branch data hold
func dfsPt(layer, max, num int, cur string, re *[]string) {
	// ending
	if layer == max {
		// some length not good, clean
		if len(cur) == max*2 {
			*re = append(*re, cur)
		}
		return
	}

	// add current layer open
	cur += "("
	num++
	dfsPt(layer+1, max, num, cur, re)
	// close each one and next
	for i := 0; i < num; i++ {
		cur += ")"
		dfsPt(layer+1, max, num-i-1, cur, re)
	}
}

// @lc code=end

