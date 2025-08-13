/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	rmap := make(map[int]*Node)
	result := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}
	rmap[node.Val] = result

	queue := []*Node{node}
	var current *Node
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		for _, n := range current.Neighbors {

			if _, ok := rmap[n.Val]; !ok {
				r := &Node{
					Val:       n.Val,
					Neighbors: []*Node{},
				}
				rmap[n.Val] = r
				queue = append(queue, n)
			}
			rmap[current.Val].Neighbors = append(rmap[current.Val].Neighbors, rmap[n.Val])
		}
	}
	return result
}

// @lc code=end

