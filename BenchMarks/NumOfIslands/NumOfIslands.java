public class NumOfIslands {
  public static int numIslands(char[][] grid) {
    int count = 0;
    for(int i=0; i<grid.length; i++) {
        for(int j=0; j<grid[0].length; j++) {
            if(grid[i][j]=='1') {
                search(grid, i, j);
                ++count;
            }
        }
    }
    return count;
  }

  private static void search(char[][] grid, int x, int y) {
    if(x<0 || x>=grid.length || y<0 || y>=grid[0].length || grid[x][y]!='1') return;
    grid[x][y] = '0';
    search(grid, x-1, y);
    search(grid, x+1, y);
    search(grid, x, y-1);
    search(grid, x, y+1);
  }

  public static void main(String[] args){
    int n = 10000;
    char[] charSet = new char[]{'0', '1'};
    char[][] grid = new char[n][n];
    for(int i = 0; i < n; i++) {
      for(int j = 0; j < n; j++) {
        grid[i][j] = charSet[(i*17+j*13) % 2];
      }
    }
    numIslands(grid);
  }
}
