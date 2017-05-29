<?php
function removeInvalidParentheses($str) {
  $res = array();
  if ($str == "") {
    return $res;
  }

  $visited = array();
  $q = new SplQueue();

  $q->enqueue($str);
  $visited[$str] = true;

  $found = false;

  while (!$q->isEmpty()){
    $s = $q->dequeue();
    if (isValid($s)) {
      array_push($res, $s);
      $found = true;
    }
    if ($found) {
      continue;
    }

    for ($i = 0; $i < count($s); $i++) {
      if ($s[$i] != '(' && $s[$i] != ')'){
        continue;
      }
      $left = substr($s, 0, $i - 1);
      $right = substr($s, $i + 1);
      $t = ".".$left;
      $t = $t.$right;
      if (!array_key_exists($t, $visited)) {
        $q->enqueue($t);
        $visited[$t] = true;
      }
    }
  }
  return $res;
}

function isValid($str) {
  $count = 0;
  for ($i = 0; $i < strlen($str); $i++) {
    $c = $str[$i];
    if ($c == '(') {
      $count++;
    }
    if ($c == ')') {
      $count--;
      if ($count == -1) {
        return false;
      }
    }
  }
  return $count == 0;
}

function main(){
  $s = "";
  $n = 40;
  $par = "(()()(()(()((()))(()((())(()(()((()())((())))()";

  for ($i = 0; $i < $n; $i++) {
    $c = $par[ $i % strlen($par) ];
    $s = $s.$c;
  }
  $t = removeInvalidParentheses($s);
  var_dump($t);
}

main();
?>
