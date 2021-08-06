/*
 * @Author: ronlee
 * @Date: 2021-08-05 23:04:51
 * @LastEditors: ronlee
 * @LastEditTime: 2021-08-06 12:42:52
 * @Description: file content
 * @FilePath: \Note\Leetcode\6.z-字形变换.go
 */
/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 *
 * https://leetcode-cn.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (50.50%)
 * Likes:    1228
 * Dislikes: 0
 * Total Accepted:    281.9K
 * Total Submissions: 558K
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
 *
 * 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
 *
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 * 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
 *
 * 请你实现这个将字符串进行指定行数变换的函数：
 *
 *
 * string convert(string s, int numRows);
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "PAYPALISHIRING", numRows = 3
 * 输出："PAHNAPLSIIGYIR"
 *
 * 示例 2：
 *
 *
 * 输入：s = "PAYPALISHIRING", numRows = 4
 * 输出："PINALSIGYAHRPI"
 * 解释：
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "A", numRows = 1
 * 输出："A"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 由英文字母（小写和大写）、',' 和 '.' 组成
 * 1
 *
 *
 */

// @lc code=start
func convert(s string, numRows int) string {
	//以v型队列方式
	//每个v队的个数为 n= 2*（ numRows - 1）
	strList := make([]string, numRows)
	if len(s) <= 1 {
		return s
	}
	if numRows <= 1 {
		return s
	}
	n := 2 * (numRows - 1)
	//对应的i位应该被放在v队列的位置 i%n
	for i, v := range s {
		vindex := i % n
		//如果vindex 超过了最大坐标 numRows-1
		if vindex > numRows-1 {
			//vindex = numRows - 1 - (vindex - numRows - 1 )
			vindex = n - vindex
		}
		strList[vindex] += string(v)
	}
	return strings.Join(strList, "")
}

// @lc code=end

