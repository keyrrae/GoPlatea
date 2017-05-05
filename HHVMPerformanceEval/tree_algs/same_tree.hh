<?hh
require "tree_node.hh";

function isSameTree(?TreeNode $p, ?TreeNode $q): bool {
  if ($p == null && $q == null) {
    return True;
  }
  if ($p == null || $q == null) {
    return False;
  }
  if ($q->val != $p->val) {
    return False;
  }
  return isSameTree($q->left, $p->left) && isSameTree($q->right, $p->right);
}

$p = TreeNode::deserializeTree("1,2,3");
$q = TreeNode::deserializeTree("1,2,3,4,5,6,7");

var_dump(isSameTree($p, $q));
