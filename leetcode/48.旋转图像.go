/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	// 层循环，i 为当前层的起始索引
	for i := 0; 2*i < n; i++ {
		ne := n - 2*i // 当前层边长
		// 只需处理每行/列的前 ne-1 个元素（避免重复旋转角落）
		for j := 0; j < ne-1; j++ {
			// 四个角的坐标（对于当前层的偏移 j）
			top := i
			left := i
			bottom := i + ne - 1
			right := i + ne - 1

			// 暂存上边元素
			tmp := matrix[top][left+j]
			// 左边元素 → 上边
			matrix[top][left+j] = matrix[bottom-j][left]
			// 下边元素 → 左边
			matrix[bottom-j][left] = matrix[bottom][right-j]
			// 右边元素 → 下边
			matrix[bottom][right-j] = matrix[top+j][right]
			// 暂存值 → 右边
			matrix[top+j][right] = tmp
		}
	}
}

// @lc code=end

