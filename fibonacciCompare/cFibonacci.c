#include <stdio.h>

int fibonacci(int i){
	if(i<2) return i;
	return fibonacci(i-2) + fibonacci(i-1);
}

int main(){
	printf("%d\n", fibonacci(34));
}