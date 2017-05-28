package main

func drawChessboard(cols []int) []string {
  chessboard := make([]string, len(cols))
  for i := 0; i < len(cols); i++ {
    chessboard[i] = ""
    for j := 0; j < len(cols); j++ {
        if j == cols[i] {
            chessboard[i] += "Q";
        } else {
            chessboard[i] += ".";
        }
    }
  }
  return chessboard
}

func isValid(cols []int, col int) bool {
  row := len(cols)
  for i := 0; i < row; i++ {
    // same column
    if cols[i] == col {
      return false
    }
    // left-top to right-bottom
    if (i - cols[i] == row - col) {
        return false;
    }
    // right-top to left-bottom
    if (i + cols[i] == row + col) {
        return false;
    }
  }
  return true;
}

func search(n int, cols []int, res [][]string) {
  if len(cols) == n {
    res = append(res, drawChessboard(cols))
    return
  }

  for col := 0; col < n; col++ {
    if !isValid(cols, col) {
      continue
    }
    cols = append(cols, col)
    search(n, cols, res)
    cols = cols[:len(cols)-1]
  }
}

func solveNQueens(n int) [][]string {
  var result [][]string
  if n <= 0 {
    return result
  }
  var path []int
  search(n, path, result)
  return result
}

func main(){
  solveNQueens(13)
}
