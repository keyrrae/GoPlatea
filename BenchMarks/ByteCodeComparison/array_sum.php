<?php

function addPositive($n) {
  $sum = 0;
  for ($i = 0; $i < $n; $i++) {
    $sum = $sum + $n;
  }
  return $sum;
}

for($i = 0; $i < 100000000; $i++){
$res = addPositive(100);
}
?>
