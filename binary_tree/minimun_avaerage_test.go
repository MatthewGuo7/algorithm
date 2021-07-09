package binary_tree

type MaxAverageRet struct {
	Sum int
	num int
}

var subRet = MaxAverageRet{
	Sum: 0,
	num: 0,
}

var subTree *TreeNode

func maximumAverageSubtree(root *TreeNode) float64 {
	_ = maximumAverageSubtreeHelper(root)
	return float64(subRet.Sum)
}

func maximumAverageSubtreeHelper(root *TreeNode) MaxAverageRet {
	if root == nil {
		return MaxAverageRet{
			Sum: 0,
			num: 0,
		}
	}

	leftRet := maximumAverageSubtreeHelper(root.Left)
	rightRet := maximumAverageSubtreeHelper(root.Right)

	ret := MaxAverageRet{
		Sum: leftRet.Sum + rightRet.Sum + root.Val,
		num: leftRet.num + rightRet.num + 1,
	}

	if (subTree == nil) || (ret.Sum *subRet.num > ret.num * subRet.Sum) {
		subTree = root
		subRet = ret
	}

	return ret
}
