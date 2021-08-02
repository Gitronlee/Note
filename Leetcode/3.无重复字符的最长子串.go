/*
 * @Author: ronlee
 * @Date: 2021-08-02 21:53:39
 * @LastEditors: ronlee
 * @LastEditTime: 2021-08-02 22:11:58
 * @Description: file content
 * @FilePath: \Note\Leetcode\3.无重复字符的最长子串.go
 */
/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (37.64%)
 * Likes:    5846
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 3M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 * 示例 4:
 *
 *
 * 输入: s = ""
 * 输出: 0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * s 由英文字母、数字、符号和空格组成
 *
 *
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	//滑动窗口：r-l = cur。什么时候更新l，当无重复时不更新，有重复则l++，直到不重复
	//什么时候更新r， 不重复的时候r++, 以求cur更大。重复的时候r不变。所以r是导致重复的第一人
	//怎么记录重复， 用map记录之前字母是否已经出现。待l++前置为false
	lenth := len(s)
	left, right, res := 0, 0, 0
	var bitSet [256]bool
	for left < lenth {
		if bitSet[s[right]] == true {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		//此处不能加1 因为之前right++后还未判定是否有重复
		if right-left > res {
			res = right - left
		}
		//此处注意res + left判断,是否 此次执行完已确定不会有再大了
		if res+left >= lenth || right >= lenth {
			break
		}
	}
	return res

}

// @lc code=end

