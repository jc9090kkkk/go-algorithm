// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。
// 输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
// ]
// 输出：1
//
// 解题思路: 将二维网格看成一个无向图，竖直或水平相邻的 1 之间有边相连。
// 为了求出岛屿的数量，我们可以扫描整个二维网格。如果一个位置为 1，则以其为起始节点开始进行深度优先搜索。在深度优先搜索的过程中，每个搜索到的 1 都会被重新标记为 0。
// 最终岛屿的数量就是我们进行深度优先搜索的次数。
//

package main

func NumIslands(grid [][]byte) int {
	num := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < len(grid) && i >= 0 && j < len(grid[0]) && j >= 0 && grid[i][j] == '1' {
			grid[i][j] = 0
			dfs(i, j+1)
			dfs(i, j-1)
			dfs(i-1, j)
			dfs(i+1, j)
		}
	}
	for i := 0;i < len(grid[0]);i++ {
		for j := 0;j < len(grid);j++ {
			if grid[j][i] == '1' {
				num++
				dfs(j, i)
			}
		}
	}
	return num
}
