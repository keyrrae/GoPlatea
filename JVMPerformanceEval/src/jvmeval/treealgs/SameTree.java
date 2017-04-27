package jvmeval.treealgs;

/**
 * Created by xuanwang on 4/24/17.
 */
public class SameTree {
  public static boolean isSameTree(TreeNode p, TreeNode q) {
    if(p == null && q == null){
      return true;
    }
    if(p == null || q == null){
      return false;
    }
    if(q.val != p.val){
      return false;
    }
    return isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
  }
}
