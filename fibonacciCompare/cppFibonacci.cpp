constexpr int fibonacci(const int i){
	if(i<2) return i;
	return fibonacci(i-2) + fibonacci(i-1);
}

int main(){
	fibonacci(34);
	return 0;
}