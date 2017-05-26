public class Test{
 public static int getCount(int n){
    int count = 0;
    for(int i = 0; i < n; i++){
      count++;
    }
    return count;
  }

  public static void main(String args[]){
    //System.out.println(args.length);
    int n = Integer.parseInt(args[0]);
int c = 0;
    for (int i = 0; i < n; i++){

      c = getCount(1000);

    }
System.out.println(c);
  }
}
