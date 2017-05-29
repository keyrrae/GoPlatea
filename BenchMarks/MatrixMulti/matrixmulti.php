<?php

function matrixmulti($_A,$_B) {
    // AxB outcome is C with A's rows and B'c cols
 $r = count($_A);
 $c = count($_B[0]);
 $in= count($_B); // or $_A[0]. $in is 'inner' count

 if ( $in != count($_A[0]) ) {
   print("ERROR: need to have inner size of matrices match.\n");
   print("     : trying to multiply a ".count($_A)."x".count($_A[0])." by a ".count($_B)."x".count($_B[0])." matrix.\n");
   print("\n");
   exit(1);
 }

 // allocate retval
 $retval = array();
 for($i=0;$i< $r; $i++) { $retval[$i] = array(); }
   // multiplication here
   for($ri=0;$ri<$r;$ri++) {
     for($ci=0;$ci<$c;$ci++) {
       $retval[$ri][$ci] = 0.0;
       for($j=0;$j<$in;$j++) {
         $retval[$ri][$ci] += $_A[$ri][$j] * $_B[$j][$ci];
       }
     }
   }
   return $retval;

}

function genMatrix($n) {
  $res = array($n);
  for($i = 0; $i < $n; $i++){
    $row = array($n);
    for($j = 0; $j < $n; $j++){
      $row[$j] = mt_rand(0, 100);
    }
    $res[$i] = $row;
  }
  return $res;
}

$n = 800;
$mat1 = genMatrix($n);
$mat2 = genMatrix($n);
matrixmulti($mat1,$mat2);

?>
