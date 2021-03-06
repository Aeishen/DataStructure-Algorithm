### 1.二叉查找树（Binary Search Tree）基础

二叉查找树也可叫做二分查找树。它不仅可以查找数据，还可以高效地插入、删除数据。

特点：

1. 若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
2. 若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值；
3. 任意节点的左、右子树也分别为二叉查找树；
4. 没有键值相等的节点。

注意：它不一定是完全的二叉树。

总结：二叉搜索树（BST）是一种特殊的二叉树，其任何节点中的值都会大于其左子树中存储的值并且小于其右子树中存储的值。



#### ①查找元素操作：

步骤：

1：从根节点开始，先比较当前节点，如果当前节点为 nil 返回查找失败

2：如果当前节点不是nil，那么和当前节点比较，如果小于节点就往左子树递归找，如果大于节点就往右子树递归找，如果等于当前节点则返回寻找成功

##### 查找时间依赖于树的拓扑结构。最佳情况是 O(log­2n)，而最坏情况是 O(n)。



#### ②添加元素操作：

步骤：

1：从根节点开始，先比较当前节点，如果当前节点为 nil 那么就插入到这个节点

2：如果当前节点不是nil，那么和当前节点比较，如果小于节点就往左子树放，如果大于节点就往右子树放

3：然后分别对左子树或者右子树递归的递归进行如上 1 、 2 步骤的操作

##### **加元素算法的复杂度与查找算法的复杂度是一样的：最佳情况是 O(log­2n)，而最坏情况是 O(n)。**



#### ③删除元素操作：

步骤：

1：若删除的是叶子节点p，则不破坏整棵树的结构，只需置空该节点即可

2：若删除的节点p只有左子树pL或右子树pR，此时只要左子树pL或右子树pR赋值到删除节点的位置即可

2：若删除的节点p存在左子树pL和右子树pR，找到pL中最大的节点或pR中最小的节点，将这个节点赋值到删除节点的位置即可

##### **删除算法的运行时间也与 BST 的拓扑结构有关，最佳情况是 O(log­2n)，而最坏情况是 O(n)。**



### 2.二叉查找树（Binary Search Tree）遍历

遍历(Traversal)是指沿着某条搜索路线，依次对树中每个结点均做一次且仅做一次访问。二叉树的遍历有三种：

- 前序遍历(Preorder Traversal)：先访问当前节点，再依次递归访问左右子树
- 中序遍历(Inorder Traversal)：先递归访问左子树，再访问自身，再递归访问右子树
- 后序遍历(Postorder Traversal)：先递归访问左右子树，最后再访问当前节点。













当二叉查找树接近形成直线，那么搜索效率将极其差，变成了线性搜索。

因此二叉查找树就需要进行改进为平衡二叉树，比较常见的 Balanced Binary Tree有：

- 红黑树
- tree
- AVL tree
- Splay tree
- Treap

