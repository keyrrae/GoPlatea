<?php

function addPositive($n) {
  $sum = 0;
  for ($i = 0; $i < $n; $i++) {
    $sum = $sum + $n;
  }
  return $sum;
}

$res = addPositive(100);
?>
