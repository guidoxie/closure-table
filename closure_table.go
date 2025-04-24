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

// Generate 生成闭包表
func Generate(nodes []INode) []*ClosureTable {
	// 映射节点 ID 到其父节点 ID
	parentMap := make(map[uint]uint)
	for _, node := range nodes {
		parentMap[node.GetID()] = node.GetParentID()
	}
	// 缓存每个节点的 TreeID（即其根节点 ID）
	treeIDMap := make(map[uint]uint)

	var closures []*ClosureTable
	for _, node := range nodes {
		treeID := findRootID(node.GetID(), parentMap, treeIDMap)
		// 添加自身闭包
		closures = append(closures, &ClosureTable{
			TreeID:     treeID,
			Ancestor:   node.GetID(),
			Descendant: node.GetID(),
			Distance:   0,
		})
		// 向上遍历祖先，逐级生成闭包
		depth := 1
		parent := parentMap[node.GetID()]
		for parent != 0 {
			closures = append(closures, &ClosureTable{
				TreeID:     treeID,
				Ancestor:   parent,
				Descendant: node.GetID(),
				Distance:   uint(depth),
			})
			parent = parentMap[parent]
			depth++
		}
	}
	return closures
}

// 向上查找根节点 ID 作为 TreeID，并缓存
func findRootID(id uint, parentMap map[uint]uint, cache map[uint]uint) uint {
	if treeID, ok := cache[id]; ok {
		return treeID
	}
	node, exists := parentMap[id]
	if !exists || node == 0 {
		cache[id] = id
		return id
	}

	treeID := findRootID(node, parentMap, cache)
	cache[id] = treeID
	return treeID
}
