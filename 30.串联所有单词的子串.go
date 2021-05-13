/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
/*
	1. sliding window since word length always same
	2. check if exist, every step check hash table all false
*/
func findSubstring(s string, words []string) []int {
	wc := len(words)
	subStr := []int{}
	if s == "" || wc == 0 || wc*len(words[0]) > len(s) {
		return subStr
	}

	wl := len(words[0])
	// word count map, if 0 delete key, init fill -1
	wm := wordmap(make(map[string]int))

	// outter word length
	for i := 0; i < wl; i++ {
		if (i + wc*wl) > len(s) {
			break
		}
		wm.reset(words)
		// initial add counted word
		j, n := i, 0
		for j+wl <= len(s) && n < wc {
			wm.add(s[j : j+wl])
			j += wl
			n++
		}
		if wm.check() {
			subStr = append(subStr, i)
		}

		// sliding window
		l, r := i, j
		for r+wl <= len(s) {
			wm.add(s[r : r+wl])
			wm.minus(s[l : l+wl])
			if wm.check() {
				subStr = append(subStr, l+wl)
			}
			l = l + wl
			r = r + wl
		}
	}

	return subStr
}

type wordmap map[string]int

func (w *wordmap) reset(words []string) {
	for k := range *w {
		delete(*w, k)
	}

	for _, word := range words {
		cur, ok := (*w)[word]
		if !ok {
			(*w)[word] = -1
		} else {
			(*w)[word] = cur - 1
		}
	}
}

func (w *wordmap) add(word string) {
	cur, ok := (*w)[word]
	if !ok {
		(*w)[word] = 1
		return
	}
	if cur+1 == 0 {
		delete(*w, word)
		return
	}
	(*w)[word] = cur + 1
}

func (w *wordmap) minus(word string) {
	cur, ok := (*w)[word]
	if !ok {
		(*w)[word] = -1
		return
	}
	if cur-1 == 0 {
		delete(*w, word)
		return
	}
	(*w)[word] = cur - 1
}

func (w *wordmap) check() bool {
	if len(*w) == 0 {
		return true
	}
	return false
}

// @lc code=end

