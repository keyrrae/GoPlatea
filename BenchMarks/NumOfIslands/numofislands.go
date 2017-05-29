package main

func numIslands(grid [][]byte) int{
  count := 0
  for i := 0; i < len(grid); i++ {
      for j := 0; j < len(grid[0]); j++ {
          if grid[i][j] == '1' {
              search(grid, i, j)
              count++
          }
      }
  }
  return count
}

func search(grid [][]byte, x, y int) {
  if x<0 || x >= len(grid) || y<0 || y >= len(grid[0]) || grid[x][y] != '1' {
    return
  }
  grid[x][y] = '0'
  search(grid, x-1, y)
  search(grid, x+1, y)
  search(grid, x, y-1)
  search(grid, x, y+1)
}

func main(){
  n := 10000
  charSet := []byte{'0', '1'}
  var grid [][]byte
  for i := 0; i < n; i++ {
    var row []byte
    for j := 0; j < n; j++ {
      row = append(row, charSet[(i*17+j*13) % 2])
    }
    grid = append(grid, row)
  }
  numIslands(grid)
}
