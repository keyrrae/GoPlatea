<?php

function drawChessboard($cols) {
  $chessboard = array();
  for ($i = 0; $i < count($cols); $i++) {
    $chessboard[$i] = "";
    for ($j = 0; $j < count($cols); $j++) {
        if ($j == $cols[$i]) {
            $chessboard[$i] = $chessboard[$i]."Q";
        } else {
            $chessboard[$i] = $chessboard[$i].".";
        }
    }
  }
  return $chessboard;
}

function isValid($cols, $col) {
  $row = count($cols);
  for ( $i = 0; $i < $row; $i++) {
    // same column
    if ($cols[$i] == $col) {
      return false;
    }
    // left-top to right-bottom
    if ($i - $cols[$i] == $row - $col) {
        return false;
    }
    // right-top to left-bottom
    if ($i + $cols[$i] == $row + $col) {
        return false;
    }
  }
  return true;
}

function search($n, $cols, $res) {
  if (count($cols) == $n) {
    array_push($res, drawChessboard($cols));
    return;
  }

  for ($col = 0; $col < $n; $col++) {
    if (!isValid($cols, $col)) {
      continue;
    }
    array_push($cols, $col);
    search($n, $cols, $res);
    array_pop($cols);
  }
}

function solveNQueens($n) {
  $result = array();
  if ($n <= 0) {
    return $result;
  }
  $path = array();
  search($n, $path, $result);
  return $result;
}

function main(){
  solveNQueens(13);
}

main();
?>
