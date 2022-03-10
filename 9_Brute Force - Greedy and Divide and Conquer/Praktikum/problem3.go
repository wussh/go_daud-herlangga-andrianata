package main

import (
	"fmt"
	"sort"
)

func dragonofloowater(dragonhead, knightheight []int) {
	sort.Ints(dragonhead)
	sort.Ints(knightheight)
	lastindexknight := -1
	totalheight := 0
	for _, dragon := range dragonhead {
		dragonkilled := false
		for indexknight := lastindexknight + 1; indexknight < len(knightheight); indexknight++ {
			knight := knightheight[indexknight]
			if dragon <= knight {
				lastindexknight += indexknight
				totalheight += knight
				dragonkilled = true
				break
			}
		}
		if !dragonkilled {
			fmt.Println("knight fail")
			return
		}
	}
	fmt.Println(totalheight)
}

func main() {
	dragonofloowater([]int{5, 4}, []int{7, 8, 4})
	dragonofloowater([]int{5, 10}, []int{5})
	dragonofloowater([]int{7, 2}, []int{4, 3, 1, 2})
	dragonofloowater([]int{7, 2}, []int{2, 1, 8, 5})
}
