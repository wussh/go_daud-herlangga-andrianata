package main

import "fmt"

func cari(array []int, minindex, maxindex, x int) int {
	if array[minindex] > x || array[maxindex] < x {
		return -1
	}
	if maxindex == minindex && array[minindex] == x {
		return minindex
	}
	awal := minindex
	akhir := maxindex
	tengah := (akhir + awal) / 2
	hasilkiri := cari(array, awal, tengah, x)
	if hasilkiri != -1 {
		return hasilkiri
	}
	hasilkanan := cari(array, tengah+1, akhir, x)
	return hasilkanan
}
func binarysearch(array []int, x int) {
	hasil := cari(array, 0, len(array)-1, x)
	fmt.Println(hasil)
}

func main() {
	binarysearch([]int{1, 1, 3, 5, 6, 7}, 3)
	binarysearch([]int{1, 1, 3, 5, 6, 8, 10}, 5)
	binarysearch([]int{12, 15, 19, 24, 31, 53, 59, 60}, 100)

}
