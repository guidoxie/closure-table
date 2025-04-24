package closure_table

// INode 节点
type INode interface {
	GetID() uint // 节点ID
	GetParentID() uint
}

// ClosureTable 闭包表
type ClosureTable struct {
	TreeID     uint // 树ID，根节点ID
	Ancestor   uint // 祖先节点ID
	Descendant uint // 后代节点ID
	Distance   uint // 祖先距离后代的距离
}

// tree 树
type tree struct {
	ID       uint // 节点ID
	ParentID uint // 父节点ID
	Sons     []*tree
	Father   *tree
	level    uint
}

// Generate 生成闭包表
func Generate(nodes []INode) []*ClosureTable {
	tree := generateTree(nodes)
	return generateClosureTable(tree)
}

// AttachNode 添加子节点
func (n *tree) attachNode(node *tree) bool {
	if n.ID == node.ParentID {
		n.Sons = append(n.Sons, node)
		node.Father = n
		node.level = n.level + 1
		return true
	}
	for _, son := range n.Sons {
		if son.attachNode(node) {
			return true
		}
	}
	return false
}

// generateTree 生成树
func generateTree(nodes []INode) *tree {
	// 构建树形结构,虚拟头节点
	root := &tree{}
	for _, node := range nodes {
		root.attachNode(&tree{
			ID:       node.GetID(),
			ParentID: node.GetParentID(),
		})
	}
	return root
}

// findTreeRoot 寻找节点所属的树根节点
func findTreeRoot(node *tree) uint {
	// 遍历到顶层非虚拟节点
	current := node
	for current.Father != nil && current.Father.ID != 0 {
		current = current.Father
	}
	return current.ID
}

// generateClosureTable 生成闭包表
func generateClosureTable(node *tree) []*ClosureTable {
	ct := make([]*ClosureTable, 0)
	for _, son := range node.Sons {
		// 查找该节点所属的树根ID
		treeID := findTreeRoot(son)
		// 自身到自身的距离为零
		self := &ClosureTable{
			TreeID:     treeID,
			Ancestor:   son.ID,
			Descendant: son.ID,
			Distance:   0,
		}
		ct = append(ct, self)
		// 往上找祖先，father.ID !=0 是为了不计算虚拟头节点
		fatherCt := make([]*ClosureTable, 0)
		for father := son.Father; father != nil && father.ID != 0; father = father.Father {
			fatherCt = append(fatherCt, &ClosureTable{
				TreeID:     treeID,
				Ancestor:   father.ID,
				Descendant: son.ID,
				Distance:   son.level - father.level,
			})
		}
		ct = append(ct, fatherCt...)
		ct = append(ct, generateClosureTable(son)...)
	}
	return ct
}
