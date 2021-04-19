/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
/*
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
*/
func intToRoman(num int) string {
	reps := []byte("IVXLCDM")
	re := []byte{}
	n, i := 0, 0
	for num > 0 {
		n = num % 10
		num = num / 10
		re = append(compose(reps, i*2, n), re...)
		i++
	}

	return string(re)
}

func compose(rep []byte, i, n int) []byte {
	if n == 0 {
		return nil
	}

	b := []byte{}
	if n < 4 {
		for n > 0 {
			b = append(b, rep[i])
			n--
		}
	}

	if n == 4 {
		b = append(b, rep[i], rep[i+1])
	}

	if n < 9 && n >= 5 {
		b = append(b, rep[i+1])
		for n > 5 {
			b = append(b, rep[i])
			n--
		}
	}

	if n == 9 {
		b = append(b, rep[i], rep[i+2])
	}

	return b
}

// @lc code=end
