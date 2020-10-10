package main

import "fmt"

var arr = []int{4, 10, 3, 5, 6}

func main() {
	insert(8)
	maxHeapifyUpBottom()
	fmt.Println("maxHeapify => ", arr)
	popedEle := maxPop()
	fmt.Println("maxPop =>", popedEle, "maxPopHeap => ", arr)
	// fmt.Println("minHeapify => ", minHeapify(arr))

}
func insert(key int) {
	arr = append(arr, key)
	maxHeapifyUpBottom()
	fmt.Println("After Insertion of =>", key, ", Heapify Heap is =>", arr)
}
func maxHeapifyBottomUp() {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		p := (i - 1) / 2 //parent index
		if p < 0 {
			break
		} else {
			if arr[i] > arr[p] {
				arr = swap(p, i, arr)
			}
		}
	}
}
func minHeapifyBottomUp() {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		p := (i - 1) / 2 //parent index
		if p < 0 {
			break
		} else {
			if arr[i] < arr[p] {
				arr = swap(p, i, arr)
			}
		}
	}
}
func maxHeapifyUpBottom() {
	n := len(arr)
	for i := 0; i < n; i++ {
		// Left child index
		l := (2 * i) + 1
		// right child index
		r := (2 * i) + 2
		if l >= n {
			break
		} else if r < n {
			//if left child value is greater than swap it with parent
			if arr[l] > arr[r] {
				if arr[l] > arr[i] {
					arr = swap(i, l, arr)
				}

			} else if arr[r] > arr[l] /*if  right child is greater than swap it with parent*/ {
				if arr[r] > arr[i] {
					arr = swap(i, r, arr)
				}
			}
		} else {
			if arr[l] > arr[i] {
				arr = swap(i, l, arr)
			}
		}
	}
}
func minHeapifyUpBottom() {
	n := len(arr)
	for i := 0; i < n; i++ {
		// Left child
		l := (2 * i) + 1
		// right child
		r := (2 * i) + 2
		if l >= n {
			break
		} else if r < n {
			//if left child value is greater than swap it with parent
			if arr[l] < arr[r] {
				if arr[l] < arr[i] {
					arr = swap(i, l, arr)
				}

			} else if arr[r] < arr[l] /*if  right child is greater than swap it with parent*/ {
				if arr[r] < arr[i] {
					arr = swap(i, r, arr)
				}
			}
		} else {
			if arr[l] < arr[i] {
				arr = swap(i, l, arr)
			}
		}
	}
}
func maxPop() int {
	n := len(arr)
	arr[n-1], arr[0] = arr[0], arr[n-1]
	popedElement := arr[0]
	fmt.Println(arr[:n-1])
	arr = arr[:n-1]
	maxHeapifyUpBottom()
	return popedElement
}
func swap(p, c int, arr []int) []int {
	//fmt.Printf("before swapping of %d = %d\n", p, arr)
	arr[c], arr[p] = arr[p], arr[c]
	//fmt.Printf("after swapping %d\n", arr)
	return arr
}
