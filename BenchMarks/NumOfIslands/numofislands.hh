<?php
function numIslands($grid) {
  $count = 0;
  for ($i = 0; $i < count($grid); $i++) {
      for ($j = 0; $j < count($grid[0]); $j++) {
          if ($grid[$i][$j] == '1') {
              search($grid, $i, $j);
              $count++;
          }
      }
  }
  return $count;
}

function search($grid, $x, $y) {
  if ($x<0 || $x >= count($grid) || $y<0 || $y >= count($grid[0]) || $grid[$x][$y] != '1') {
    return;
  }
  $grid[$x][$y] = '0';
  search($grid, $x-1, $y);
  search($grid, $x+1, $y);
  search($grid, $x, $y-1);
  search($grid, $x, $y+1);
}

function main(){
  $n = 10000;
  $charSet = "01";
  $grid = array();
  for ($i = 0; $i < $n; $i++) {
    $row = array();
    for ($j = 0; $j < $n; $j++) {
      array_push($row, $charSet[($i*17 + $j*13) % 2]);
    }
    array_push($grid, $row);
  }
  numIslands($grid);
}

main();
?>
