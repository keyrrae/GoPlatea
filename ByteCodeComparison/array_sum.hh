<?hh

function addPositive(int $n): int {
  $sum = 0;
  for ($i = 0; $i < $n; $i++) {
    $sum = $sum + $n;
  }
  return $sum;
}

$res = addPositive(100);
