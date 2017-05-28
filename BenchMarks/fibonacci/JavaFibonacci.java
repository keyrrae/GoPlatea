public class JavaFibonacci{
	public static void main(String[] args){
		System.out.println(fibonacci(45));
	}

	static long fibonacci(long i){
		if(i<2) return i;
		return fibonacci(i-2) + fibonacci(i-1);
	}
}
