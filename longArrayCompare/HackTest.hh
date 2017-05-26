<?hh
namespace HHVM\UserDocumentation\BasicUsage\Examples\CommandLine;

function main(array<string> $argv) {
  $element = 0;
  $iteration = 0;
  $iterations = 0;
  $innerloop = 0;
  $sum = 0.0;
  $array_length = 10000;
  $array = new Vector(null);

  if (count($argv) > 0){
    $iterations = intval($argv[1]);
  }
  for ($element = 0; $element < $array_length; $element++)
    $array->add($element);
  for ($iteration = 0; $iteration < $iterations; $iteration++) {
    for ($innerloop = 0; $innerloop < 1000; $innerloop++) {
      $sum += $array[($iteration + $innerloop) % $array_length];
    }
  }
}

main($argv);
