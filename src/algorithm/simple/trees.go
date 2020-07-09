package simple

import (
	"container/list"
	"fmt"
)

//二叉树专题

//队列-链表实现--------------------------------
type queue struct {
	l *list.List
}
func NewQueue() *queue {
	return &queue{
		list.New(),
	}
}
func (this *queue)InQueue(val interface{}){
	this.l.PushBack(val)
}
func (this *queue)OutQueue() interface{} {
	ele := this.l.Front()
	if ele == nil {
		return nil
	}

	this.l.Remove(ele)
	return ele.Value
}
func (this *queue)Len() int {
	return this.l.Len()
}

//栈-数组实现-----------------------------------
type stack struct {
	arr []interface{}
}
func NewStack() *stack {
	return &stack{arr:make([]interface{},0)}
}
func (this *stack)Push(val interface{})  {
	this.arr = append(this.arr,val)
}
func (this *stack)Pop() interface{} {
	if len(this.arr) > 0 {
		val     := this.arr[len(this.arr)-1]
		this.arr = this.arr[:len(this.arr)-1]
		return val
	}
	return nil
}
func (this *stack)GetLast() interface{} {
	if len(this.arr) > 0 {
		return this.arr[len(this.arr)-1]
	}
	return nil
}
func (this *stack)Len() int {
	return len(this.arr)
}


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//前序----递归
func (this *TreeNode)PreOrder(f func(curNode *TreeNode))  {
	preOrder(this,f)
}
func preOrder(current *TreeNode,f func(curNode *TreeNode))  {
	if current != nil {
		f(current)
		preOrder(current.Left,f)
		preOrder(current.Right,f)
	}
}
//前序----栈
func (this *TreeNode)PreOrderStack(f func(node *TreeNode))  {
	preOrderStack(this,f)
}
func preOrderStack(root *TreeNode,f func(node *TreeNode))  {
	sta  := NewStack()
	curr := root
	for curr != nil || sta.Len() > 0  {
		for curr != nil  {
			f(curr)
			sta.Push(curr)
			curr = curr.Left
		}

		for sta.Len() > 0 {
			curr = sta.Pop().(*TreeNode)
			curr = curr.Right
		}
	}
}

//中序----
func (this *TreeNode)InOrder(f func(curNode *TreeNode))  {
	inOrder(this,f)
}
func inOrder(current *TreeNode,f func(curNode *TreeNode))  {
	if current != nil {
		inOrder(current.Left,f)
		f(current)
		inOrder(current.Right,f)
	}
}

//后序
func (this *TreeNode)PostOrder(f func(curNode *TreeNode))  {
	postOrder(this,f)
}
func postOrder(current *TreeNode,f func(curNode *TreeNode))  {
	if current != nil {
		postOrder(current.Left,f)
		postOrder(current.Right,f)
		f(current)
	}
}

//广度优先遍历-队列实现
func (this *TreeNode)LevelOrder(f func(node *TreeNode))  {
	//根节点入队
	qu := NewQueue()
	qu.InQueue(this)

	for {
		first := qu.OutQueue()
		if first == nil {
			break
		}
		current := first.(*TreeNode)
		f(current)
		if current.Left != nil {
			qu.InQueue(current.Left)
		}
		if current.Right != nil {
			qu.InQueue(current.Right)
		}
	}

}
func (this *TreeNode)DepthLevel() int {
	return depth1(this)
}
func (this *TreeNode)DepthPost() int {
	return depth2(this)
}
//广度优先遍历-实现获取树的深度
func depth1(root *TreeNode) int {
	//根节点入队
	depth := 0
	qu := NewQueue()
	qu.InQueue(root)
	depth++
	tmp := NewQueue()
	for {
		if qu.Len() == 0 && tmp.Len() == 0 {
			break
		}
		if qu.Len() == 0 && tmp.Len() > 0 {
			depth++
			qu  = tmp
			tmp = NewQueue()
		}
		first := qu.OutQueue()


		current := first.(*TreeNode)
		//f(current)
		if current.Left != nil {
			tmp.InQueue(current.Left)
		}
		if current.Right != nil {
			tmp.InQueue(current.Right)
		}
	}

	return depth
}
//深度优先遍历-后序-递归-实现获取树的深度
func depth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth2(root.Left),depth2(root.Right)) + 1
}

//https://leetcode-cn.com/problems/minimum-height-tree-lcci/
func (*Ref)STBST()  {
	nums     := []int{-10,-3,0,5,9}

	if len(nums) == 0 {
		return
	}
	node := &TreeNode{}
	c(node,nums)

	node.PreOrder(func(curNode *TreeNode) {
		fmt.Println(curNode.Val)
	})
}
func c(current *TreeNode,arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	rootNode := len(arr)/2

	current.Val   = arr[rootNode]
	if current.Left == nil {
		current.Left = &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		}
	}
	if current.Right == nil {
		current.Right = &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		}
	}
	current.Left  = c(current.Left,arr[:rootNode])
	current.Right = c(current.Right,arr[rootNode+1:])

	return current
}

//https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
func (*Ref)Mt()  {
	rootA     := []int{2,3,-1,1}
	root      := ArrToNode(rootA)
	toMirror(root)
	root.LevelOrder(func(node *TreeNode) {
		fmt.Println(node.Val)
	})
}
func toMirror(current *TreeNode) *TreeNode {
	if current == nil {
		return nil
	}
	tmp          := current.Left
	current.Left  = toMirror(current.Right)
	current.Right = toMirror(tmp)
	return current
}

//层级遍历输出的数组 转为 树
func ArrToNode(arr []int) *TreeNode {
	rootNode := &TreeNode{
		Val:   arr[0],
		Left:  nil,
		Right: nil,
	}
	currentNode := rootNode

	for i := 0; i < len(arr); i++ {
		left  := 2*i+1  //1,2|3,4
		right := left+1


		if left < len(arr) && arr[left] != -1 {
			currentNode.Left = &TreeNode{
				Val:   arr[left],
				Left:  nil,
				Right: nil,
			}
		}
		if right < len(arr) && arr[right] != -1 {
			currentNode.Right = &TreeNode{
				Val:   arr[right],
				Left:  nil,
				Right: nil,
			}
		}
		if left == len(arr) {
			break
		}
		if right == len(arr) {
			break
		}
		if i+1 > len(arr) {
			break
		}
		if  arr[i+1] != -1{
			preOrder(rootNode,func(curNode *TreeNode) {
				if curNode.Val == arr[i+1] {
					currentNode = curNode
				}
			})
		}
	}
	return rootNode
}

//https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
func (*Ref)Md()  {
	rootA := []int{3,9,20,-1,-1,15,7}
	root  := ArrToNode(rootA)
	fmt.Println(root.DepthLevel())
	fmt.Println(root.DepthPost())
}

func (*Ref)RsBst()  {
	rootA := []int{15,9,21,7,13,19,23,5,-1,11,-1,17}
	root  := ArrToNode(rootA)

	fmt.Println(sum(root,19,21,0))
}

func sum(root *TreeNode,L,R,ans int) int {
	if root != nil {
		if L <= root.Val && root.Val <= R {
			ans += root.Val
		}
		if L < root.Val {
			ans = sum(root.Left,L,R,ans)
		}
		if root.Val < R {
			ans = sum(root.Right,L,R,ans)
		}
	}
	return ans
}

// https://leetcode-cn.com/problems/merge-two-binary-trees/
func (*Ref)MergeTrees() {
	t1  := ArrToNode([]int{1,3,2,5})
	t2  := ArrToNode([]int{2,1,3,-1,4,-1,7})

	t1 = preOrder1(t1,t2)
	t1.LevelOrder(func(node *TreeNode) {
		fmt.Println(node.Val)
	})
}
func preOrder1(n1 *TreeNode,n2 *TreeNode) *TreeNode {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	n1.Val  += n2.Val
	n1.Left  = preOrder1(n1.Left,n2.Left)
	n1.Right = preOrder1(n1.Right,n2.Right)
	return n1
}

//https://leetcode-cn.com/problems/invert-binary-tree/
func (*Ref)InvertTree()  {
	root := ArrToNode([]int{4,2,7,1,3,6,9})
	mirror(root)
	root.LevelOrder(func(node *TreeNode) {
		fmt.Println(node.Val)
	})
}
func mirror(node *TreeNode)  {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		return
	}

	tmp       := node.Left
	node.Left  = node.Right
	node.Right = tmp
	mirror(node.Left)
	mirror(node.Right)
}

//https://leetcode-cn.com/problems/search-in-a-binary-search-tree/
func (*Ref)SearchBST()  {
	root := ArrToNode([]int{4,2,7,1,3})
	val  := 2

	var valNode *TreeNode
	valNode = bstFindTree(root,val)

	if valNode != nil {
		valNode.LevelOrder(func(node *TreeNode) {
			fmt.Println(node.Val)
		})
	}
}
func bstFindTree(node *TreeNode,val int) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val == val {
		return node
	}

	if val > node.Val {
		return bstFindTree(node.Right,val)
	}
	return bstFindTree(node.Left,val)
}

type Node struct {
    Val int
    Children []*Node
}
//https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
func (*Ref)PO()  {
	root := &Node{
		Val:      1,
		Children: []*Node{
			{3,[]*Node{
				{5,nil},
				{6,nil},
			}},
			{2,nil},
			{4,nil},
		},
	}

	arr := []int{}
	Npo1(root, func(cn *Node) {
		arr = append(arr,cn.Val)
	})
	sta2 := NewStack()
	for _,v := range arr  {
		sta2.Push(v)
	}
	i := 0
	for sta2.Len() > 0 {
		arr[i] = sta2.Pop().(int)
		i++
	}
	fmt.Println(arr)
}
//递归实现
func Npo(n *Node,f func(cn *Node))  {
	if n == nil {
		return
	}

	for len(n.Children) > 0 {
		Npo(n.Children[0],f)
		n.Children = n.Children[1:]
	}
	f(n)
}
//栈实现
func Npo1(root *Node,f func(cn *Node))  {
	sta := NewStack()
	sta.Push(root)
	for sta.Len() > 0  {
		n := sta.Pop().(*Node)
		f(n)

		if n.Children != nil {
			for _,v := range n.Children {
				sta.Push(v)
			}
		}
	}
}

func (*Ref)KthLargest()  {
	root := ArrToNode([]int{3,1,4,-1,2})
	k    := 4

	val  := 0
	root.InOrder(func(curNode *TreeNode) {
		if k < 1 {
			return
		}
		val = curNode.Val
		k--
	})
	fmt.Println(val)
}

//https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
func (*Ref)Pre()  {
	root := &Node{
		Val:      1,
		Children: []*Node{
			{3,[]*Node{
				{5,nil},
				{6,nil},
			}},
			{2,nil},
			{4,nil},
		},
	}

	//递归
	arr := []int{}
	pre1(root, func(node *Node) {
		arr = append(arr,node.Val)
	})

	arr2 := []int{}
	pre2(root, func(node *Node) {
		arr2 = append(arr2,node.Val)
	})

	fmt.Println(arr)
	fmt.Println(arr2)
}
//递归实现
func pre1(n *Node,f func(node *Node))  {
	if n == nil {
		return
	}

	f(n)
	if n.Children != nil {
		for i := 0; i < len(n.Children); i++ {
			pre1(n.Children[i],f)
		}
	}
}

//迭代实现
func pre2(root *Node,f func(node *Node))  {
	sta := NewStack()
	sta.Push(root)
	for sta.Len() > 0 {
		cur := sta.Pop().(*Node)
		f(cur)
		if cur.Children != nil {
			for i := len(cur.Children)-1; i >= 0; i--  {
				sta.Push(cur.Children[i])
			}
		}
	}
}

//https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
func (*Ref)SAToBST()  {
	nums := []int{-10,-3,0,5,9}
	tree := SortArrToBst(nums,0,len(nums)-1)
	tree.LevelOrder(func(node *TreeNode) {
		fmt.Println(node.Val)
	})
}

//把小->大的数组转换成BST(二叉搜索树)
func SortArrToBst(arr []int,left int,right int) *TreeNode {
	if left > right {
		return nil
	}
	rootIndex := (left + right + 1)/2
	root      := &TreeNode{
		Val:   arr[rootIndex],
		Left:  nil,
		Right: nil,
	}
	root.Left = SortArrToBst(arr,left,rootIndex -1)
	root.Right = SortArrToBst(arr,rootIndex+1,right)
	return root
}


//https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/
func (*Ref)MaxNDepth()  {
	root := &Node{
		Val:      1,
		Children: []*Node{
			{3,[]*Node{
				{5,nil},
				{6,nil},
			}},
			{2,nil},
			{4,nil},
		},
	}
	fmt.Println(Ndepth(root))
}
func Ndepth(root *Node) int {
	if root == nil {
		return 0
	}
	
	max1 := 0
	if root.Children != nil {
		for _,c := range root.Children {
			cDepth := Ndepth(c)
			if max1 < cDepth {
				max1 = cDepth
			}
		}
	}
	return max1 + 1
}

//https://leetcode-cn.com/problems/increasing-order-search-tree/
func (*Ref)ICBST()  {
	root := ArrToNode([]int{5,3,6,2,4,-1,8,1,-1,-1,-1,7,9})
	ToRight(root).InOrder(func(curNode *TreeNode) {
		fmt.Println(curNode.Val)
	})
}
func ToRight(root *TreeNode) *TreeNode {
	var rootNew  *TreeNode
	var tmp  *TreeNode

	root.InOrder(func(curNode *TreeNode) {
		if rootNew == nil {
			rootNew = curNode
			tmp     = rootNew
		} else {
			tmp.Right = &TreeNode{
				Val:   curNode.Val,
				Left:  nil,
				Right: nil,
			}
			tmp       = tmp.Right
		}
	})
	return rootNew
}

//https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
func (*Ref)LO()  {
	root := ArrToNode([]int{3,9,20,-1,-1,15,7})
	fmt.Println(lo(root))
}
func lo(root *TreeNode) [][]int {
	arr := make([][]int,0)

	que := NewQueue()
	que.InQueue(root)
	tmp := NewQueue()
	arr1 := []int{}
	for que.Len() > 0 || tmp.Len() > 0  {
		if que.Len() == 0 {
			que  = tmp
			tmp  = NewQueue()
			arr  = append(arr,arr1)
			arr1 = []int{}
		}

		node := que.OutQueue().(*TreeNode)
		arr1  = append(arr1,node.Val)
		if node.Left != nil {
			tmp.InQueue(node.Left)
		}
		if node.Right != nil {
			tmp.InQueue(node.Right)
		}
	}

	arr = append(arr,arr1)
	return arr
}
