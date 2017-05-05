package jvmeval.treealgs;

import java.util.LinkedList;
import java.util.Queue;

/**
 * Created by xuanwang on 4/24/17.
 */
public class TreeNode {
  int val;
  TreeNode left;
  TreeNode right;

  public TreeNode(int v){
    val = v;
  }

  // Encodes a tree to a single string.
  public static String serializeTree(TreeNode root) {
    if (root == null) return "";
    Queue<TreeNode> q = new LinkedList<>();
    StringBuilder res = new StringBuilder();
    q.add(root);
    while (!q.isEmpty()) {
      TreeNode node = q.poll();
      if (node == null) {
        res.append("null,");
        continue;
      }
      res.append(node.val).append(',');
      q.add(node.left);
      q.add(node.right);
    }

    String r = res.toString();
    return r.substring(0, r.length()-1);
  }

  public static TreeNode deserializeTree(String data) {
    if (data.equals("")) return null;
    Queue<TreeNode> q = new LinkedList<>();
    String[] values = data.split(",");
    TreeNode root = new TreeNode(Integer.parseInt(values[0]));
    q.add(root);
    for (int i = 1; i < values.length; i++) {
      TreeNode parent = q.poll();
      if (!values[i].equals("null")) {
        TreeNode left = new TreeNode(Integer.parseInt(values[i]));
        parent.left = left;
        q.add(left);
      }
      if (!values[++i].equals("null")) {
        TreeNode right = new TreeNode(Integer.parseInt(values[i]));
        parent.right = right;
        q.add(right);
      }
    }
    return root;
  }
}
