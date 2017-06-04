<?php

function add($n) {
  $count = 0;

  for($i = 0; $i < $n; $i++){
    $count++;
  }
  return $count;
}

for ($i = 0; $i < $argv[1]; $i++) {
  add(1000);
}
