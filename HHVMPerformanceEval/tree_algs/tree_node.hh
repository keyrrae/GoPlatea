<?hh

class TreeNode {
  public ?TreeNode $left;
  public ?TreeNode $right;

  public function __construct(public int $val) {
    $this->left = null;
    $this->right = null;
  }

  public static function serializeTree(?TreeNode $root): string {
    if ($root == null) {
      return "";
    }

    $res = "";
    $q = new SplQueue();
    $q->enqueue($root);
    while (!$q->isEmpty()) {
      $node = $q->dequeue();
      if ($node == null) {
        $res .= "null,";
        continue;
      }
      $res .= $node?->val;
      $res .= ",";
      $q->enqueue($node?->left);
      $q->enqueue($node?->right);
    }
    return substr($res, 0, -1);
  }

  public static function deserializeTree(string $str): ?TreeNode {
    if ($str == "") {
      return null;
    }

    $q = new SplQueue();
    $arr = explode(",", $str);
    $len = count($arr);

    $root = new TreeNode(intval($arr[0]));
    $q->enqueue($root);
    for ($i = 1; $i < $len; $i++) {
      $node = $q->dequeue();
      if ($arr[$i] != "null") {
        $left = new TreeNode(intval($arr[$i]));
        $node->left = $left;
        $q->enqueue($left);
      }
      if ($arr[++$i] != "null") {
        $right = new TreeNode(intval($arr[$i]));
        $node->right = $right;
        $q->enqueue($right);
      }
    }
    return $root;
  }
}
