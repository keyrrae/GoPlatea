public class P {
	public static void main(String[] args){
		int total = 0;
		for(int i=0; i<999; i++){
			for(int j=1; j<999; j++){
				total += (i / j);
			}
		}
		System.out.println(total);
	}
}
