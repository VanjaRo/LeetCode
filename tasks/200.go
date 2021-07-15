package main

func main() {

}

func numIslands(grid [][]byte) int {
	ret := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				dfs(i, j, &grid)
				ret++
			}
		}
	}
	return ret
}

func dfs(i, j int, grid *[][]byte) {
	if i < 0 || j < 0 || i >= len((*grid)) || j >= len((*grid)[0]) || (*grid)[i][j] == byte('0') {
		return
	}

	(*grid)[i][j] = '0'
	dfs(i-1, j, grid)
	dfs(i+1, j, grid)
	dfs(i, j-1, grid)
	dfs(i, j+1, grid)
}
