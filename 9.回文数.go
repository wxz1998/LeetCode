
/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	val := 0
	temp := x
	for x != 0 {
		val = val*10 + x%10
		x = x / 10
	}
	if val == temp {
		return true
	} else {
		return false
	}
}

// @lc code=end
// wxz
