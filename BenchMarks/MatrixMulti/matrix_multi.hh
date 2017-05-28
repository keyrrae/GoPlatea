<?hh

class MatrixMulti {

  public static function multiplicar(
    Vector<Vector<float>> $A,
    Vector<Vector<float>> $B,
  ): Vector<Vector<float>> {

    $aRows = count($A);
    $aColumns = count($A[0]);
    $bRows = count($B);
    $bColumns = count($B[0]);
    if ($aColumns != $bRows) {
      throw new Exception('A:Rows: '.$aColumns.' did not match B:Columns '.$bRows.".");
    }

    $C = new Vector(null);

    for ($i = 0; $i < $aRows; $i++) {
      $row = new Vector(null);
      for ($j = 0; $j < $bColumns; $j++) {
          $row->add(0.0);
      }
      $C->add($row);
    }

    for ($i = 0; $i < $aRows; $i++) { // aRow
      for ($j = 0; $j < $bColumns; $j++) { // bColumn
        for ($k = 0; $k < $aColumns; $k++) { // aColumn
          $C[$i][$j] += $A[$i][$k] * $B[$k][$j];
        }
      }
    }

    return $C;
  }
}

class MatrixBuilder {
  private int $x;
  private int $y;
  private bool $genRand;

  public function __construct() {
    $this->x = 0;
    $this->y = 0;
    $this->genRand = False;
  }

  public function set_x(int $xx): MatrixBuilder {
    $this->x = $xx;
    return $this;
  }

  public function set_y(int $yy): MatrixBuilder {
    $this->y = $yy;
    return $this;
  }

  public function enable_gen_rand(): MatrixBuilder {
    $this->genRand = True;
    return $this;
  }

  public static function frand($min, $max, $decimals = 0): float {
    $scale = pow(10, $decimals);
    return (float) mt_rand($min * $scale, $max * $scale) / $scale;
  }

  public function build(): ?Vector<Vector<float>> {
    if ($this->x <= 0 || $this->y <= 0) {
      return null;
    }

    $res = new Vector(null);

    for ($i = 0; $i < $this->x; $i++) {
      $row = new Vector(null);
      for ($j = 0; $j < $this->y; $j++) {
        if ($this->genRand) {
          $f = MatrixBuilder::frand(0, 9, 4);
          $row->add($f);
        } else {
          $row->add(0.0);
        }
      }
      $res->add($row);
    }

    return $res;
  }
}

$dim = 100;
$A = (new MatrixBuilder())->set_x($dim)->set_y($dim)->enable_gen_rand()->build();
$B = (new MatrixBuilder())->set_x($dim)->set_y($dim)->enable_gen_rand()->build();
$C = MatrixMulti::multiplicar($A, $A);
