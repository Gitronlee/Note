/*
 * @Author: ronlee
 * @Date: 2021-08-05 22:25:48
 * @LastEditors: ronlee
 * @LastEditTime: 2021-08-05 23:02:44
 * @Description: file content
 * @FilePath: \Note\Leetcode\5.最长回文子串.go
 */
/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (34.88%)
 * Likes:    3912
 * Dislikes: 0
 * Total Accepted:    667.8K
 * Total Submissions: 1.9M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出："bb"
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "a"
 * 输出："a"
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "ac"
 * 输出："a"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 仅由数字和英文字母（大写和/或小写）组成
 *
 *
 */

// @lc code=start
func longestPalindrome(s string) string {
	//此问题常规解法都是n2的复杂度，马拉车算法为n的复杂度。此处解法为中心扩散法
	//如果长度小于等于1直接返回
	if len(s) <= 1 {
		return s
	}
	//否则开始找个地方两边扩散，最初以保证maxlen = 1了
	the_start := 0
	maxlen := 1
	for i := 0; i < len(s); {
		//i即是中心点位置
		if i+maxlen/2 >= len(s) {
			//当前已经是最长了，后续无意义直接退出
			break
		}
		//从中心分别往两边扩散
		j, k := i, i
		//找到从i处往后的连续相同字符串，末尾位置为k
		for k < len(s)-1 && s[k] == s[k+1] {
			k++
		}
		//这块既然相同了。我的下次的中心i就不用i++了，应直接k+1，所以上次的i固然不是相同字符
		i = k + 1

		//同上，不是相同字符没必要向左查找了
		//往左找到第一个不等于i处的字符 位置为j
		// for j>1 && s[j] == s[j-1]{
		// 	j--
		// }
		//两边扩散
		for j > 0 && k < len(s)-1 && s[j-1] == s[k+1] {
			k++
			j--
		}
		//最终两边都停留在回文串的边界处
		if k-j+1 > maxlen {
			maxlen = k - j + 1
			//由于最终返回的是字符串，所以我们得记录下开始坐标
			the_start = j
		}

	}
	return s[the_start : the_start+maxlen]
}

// @lc code=end

