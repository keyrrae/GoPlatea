<?hh
function fibonacci(int $i): int {
  if ($i < 2)
    return $i;
  return fibonacci($i - 2) + fibonacci($i - 1);
}

if (count($argv) > 1) {
  echo fibonacci(intval($argv[1]));
} else {
  echo fibonacci(34);
}
