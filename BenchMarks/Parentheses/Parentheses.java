import java.util.*;

public class Parentheses {
  public static List<String> removeInvalidParentheses(String s) {
      List<String> res = new ArrayList<>();
      // sanity check
      if (s == null) return res;

      Set<String> visited = new HashSet<>();
      Queue<String> queue = new LinkedList<>();
      // initialize
      queue.add(s);
      visited.add(s);

      boolean found = false;

      while (!queue.isEmpty()) {
          s = queue.poll();
          if (isValid(s)) {
              // found an answer, add to the result
              res.add(s);
              found = true;
          }
          if (found) continue;

          // generate all possible states
          for (int i = 0; i < s.length(); i++) {
              // we only try to remove left or right paren
              if (s.charAt(i) != '(' && s.charAt(i) != ')') continue;

              String t = s.substring(0, i) + s.substring(i + 1);

              if (!visited.contains(t)) {
                  // for each state, if it's not visited, add it to the queue
                  queue.add(t);
                  visited.add(t);
              }
          }
      }

      return res;
  }

  // helper function checks if string s contains valid parantheses
  static boolean isValid(String s) {
      int count = 0;

      for (int i = 0; i < s.length(); i++) {
          char c = s.charAt(i);
          if (c == '(') count++;
          if (c == ')' && count-- == 0) return false;
      }
      return count == 0;
  }

  public static void main(String[] args){
    System.out.println(args.length);
    StringBuilder sb = new StringBuilder();
    char[] par = new char[]{'(', ')'};
    int n = 20;
    if(args.length > 0){
      n = Integer.parseInt(args[0]);
    }
    Random rand = new Random(1234);
    for(int i = 0; i < n; i++){
      sb.append(par[Math.abs(rand.nextInt())%2]);
    }
    System.out.println(sb.toString());
    removeInvalidParentheses(sb.toString());
  }
}
