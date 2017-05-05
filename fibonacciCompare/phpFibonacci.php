<?php
function fibonacci(int $i): int{
	if($i < 2) return $i;
	return fibonacci($i - 2) + fibonacci($i - 1);
}

echo fibonacci(34);
