package maze

// A robot is located at the top-left corner of a m x n grid
// The robot can only move either down or right at any point
// in time. The robot is trying to reach the bottom-right corner of the grid
// return total numbers of all possible unique paths
func UniquePaths(m int, n int) int {
	//if m > 1 && n > 1 {
	//	return UniquePaths(m-1, n) + UniquePaths(m, n-1)
	//} else if m > 1 {
	//	return UniquePaths(m-1, n)
	//} else if n > 1 {
	//	return UniquePaths(m, n-1)
	//} else if m == 1 && n == 1 {
	//	return 1
	//} else {
	//	return 0
	//}
	if m > n {
		return UniquePaths(n, m)
	}
	vec := make([]int, m)
	for i := 0; i < m; i++ {
		vec[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			vec[j] += vec[j-1]
		}
	}
	return vec[m-1]
}

// some obstacles are added to the grids
// An obstacle and empty space is marked as 1 and 0 respectively in the grid.
// 1 means obstacle
// return total numbers of all possible unique paths
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	width := len(obstacleGrid[0])
	vec := make([]int, width)
	vec[0] = 1
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < width; j++ {
			if obstacleGrid[i][j] == 1 {
				vec[j] = 0
			} else if j > 0 {
				vec[j] += vec[j-1]
			}
		}
	}
	return vec[width-1]
}
