package arraysAndSlices

import (
	"fmt"
	"reflect"
	"sort"
)

func init() {
	fmt.Println("init run in slicesLesson. runs before main but after a: arraysLesson")
}

func SliceInfo() {
	sliceCreation_incFromArray_AndSorting()
	sliceCopying()
	slicesAsParamsAndComparing()
}

func sliceCreation_incFromArray_AndSorting() {
	//slices are like ArrayLists:
	// variable slices
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]

	//can be defined like arrays but without size limit:
	sli := []int{}
	sli = append(sli, 6)

	//these are also valid ways to init slices
	var sli2 [][]int //this is a 0 value slice aka nil slice
	sli3 := []int{1}

	fmt.Println(slice, sli, sli2 == nil, sli3)

	//using make function
	var s1 = make([]int, 4, 7) //current length, capacity of underlying array
	fmt.Println(s1, len(s1), cap(s1))
	s1 = make([]int, 7) //capacity is optional. here, curr. length = capacity too.
	fmt.Println(s1, len(s1), cap(s1))

	//can iterate over slice just like an array:
	for index, elem := range slice {
		fmt.Println(index, elem)
	}

	//slices ARE REFERENCES, so editing it affects the underlying array:
	slice[0] = 100
	fmt.Println(arr, slice)

	//sorting: go has separate sorting funcs for elem types:
	sort.Ints(slice)
	fmt.Println("That part of array is also sorted:", arr, slice)
}

func sliceCopying() {
	var source = []int{1, 2, 3}

	//slices CAN be copied:
	//define dest as having same length
	dest := make([]int, len(source))
	//copy source over. return value is # of copied elems (i.e. smallest of 2 lens)
	count := copy(dest, source)

	//can skip make function
	var d2 []int
	c := copy(d2, source)
	fmt.Println(dest, d2, count, c)
}

func slicesAsParamsAndComparing() {
	var src = []int{1, 2, 3}

	//slices are passed as reference, so this will double all elems
	func(x []int) {
		for s := range src {
			src[s] *= 2
		}
	}(src)

	fmt.Println(src)

	var sr2 = []string{"1", "2", "3"}
	fmt.Println(reflect.DeepEqual(src, sr2))
}
