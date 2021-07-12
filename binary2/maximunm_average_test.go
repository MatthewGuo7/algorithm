package binary2

type SumRet struct {
	Sum int
	num int
}

var sum = SumRet{
	Sum: 0,
	num: 0,
}

func maximumAverageSubtree(root *TreeNode) float64 {
	return float64(maximumAverageSubtreeHelper(root).Sum)
}

func maximumAverageSubtreeHelper(root *TreeNode) SumRet {
	if root == nil {
		return SumRet{
			Sum: 0,
			num: 0,
		}
	}

	leftSumRet := maximumAverageSubtreeHelper(root.Left)
	rightSumRet := maximumAverageSubtreeHelper(root.Right)

	ret := SumRet{
		Sum: leftSumRet.Sum + rightSumRet.Sum + root.Val,
		num: leftSumRet.num + rightSumRet.num + 1,
	}

	if ret.Sum*sum.num > sum.Sum*ret.num {
		sum = ret
	}

	return ret
}
