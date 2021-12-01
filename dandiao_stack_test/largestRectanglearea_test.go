package dandiao_stack

import "math"

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	maxArea := 0
	for start:= 0; start< len(heights); start++{
		for end := start; end < len(heights); end++{
			maxHeight := math.MaxInt32
			for i := start; i <= end; i++ {
				maxHeight = int(math.Min(float64(maxHeight), float64(heights[i])))
			}
			maxArea = int(math.Max(float64(maxArea), float64((start-end+1)*maxHeight)))
		}
	}

	return maxArea
}

type Rec struct {
	Width int
	Height int
}

func largestRectangleArea2(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	heights = append(heights, 0)

	maxArea := 0
	stack := make([]Rec, 0, len(heights))

	for i := 0; i < len(heights);i++{
		accumulatedWidth := 0

		for len(stack) > 0 && stack[len(stack)-1].Height > heights[i] {
			top := stack[len(stack)-1]

			accumulatedWidth+=top.Width
			maxArea = int(math.Max(float64(top.Height)*float64(accumulatedWidth), float64(maxArea)))

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, Rec{
			Width:  accumulatedWidth+1,
			Height: heights[i],
		})
	}

	return maxArea
}
