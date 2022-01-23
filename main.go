package main

type IntSlice []int
type JimmyTrees = IntSlice

func (src IntSlice) Splice(start, deleteCount int) IntSlice {
	dest := make(IntSlice, len(src)-deleteCount)
	copy(dest, src[:start])
	copy(dest[start:], src[start+deleteCount:])
	return dest
}

func (trees JimmyTrees) IsAesthetic() bool {
	for i := 1; i < len(trees)-1; i++ {
		if (trees[i] >= trees[i-1] && trees[i] <= trees[i+1]) ||
			(trees[i] <= trees[i-1] && trees[i] >= trees[i+1]) {
			return false
		}
	}
	return true
}

func Solution(A []int) int {
	trees := JimmyTrees(A)
	if trees.IsAesthetic() {
		return 0
	}

	solutions := 0
	for i := 0; i < len(A); i++ {
		if trees.Splice(i, 1).IsAesthetic() {
			solutions++
		}
	}

	if solutions > 0 {
		return solutions
	}

	return -1
}
