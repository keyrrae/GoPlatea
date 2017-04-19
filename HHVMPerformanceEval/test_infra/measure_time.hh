<?hh

class test_infra<Ta, Tr> {
  public static function measure_time((function(Ta): Tr) $fut, Ta $args): int {
    $i = microtime(true);

    for ($t = 0; $t < 1000000000; $t++)
      ;

    $j = microtime(true);
    return $j - $i;
  }
}
