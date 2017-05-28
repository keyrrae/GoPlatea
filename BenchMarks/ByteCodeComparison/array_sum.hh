<?hh

function addPositive(int $n): int {
  $sum = 0;
  for ($i = 0; $i < $n; $i++) {
    $sum = $sum + $n;
  }
  return $sum;
}

for ($i = 0; $i < 100; $i++) {
  $res = addPositive(100);
}
