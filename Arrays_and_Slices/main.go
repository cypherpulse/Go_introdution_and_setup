// This program demonstrates the use of arrays in Go. It initializes an array of ages and prints it to the console.

package main

import "fmt"

func main() {
	//arrays in Go are fixed-size collections of elements of the same type. They are defined using the syntax [size]type, where size is the number of elements in the array and type is the data type of the elements.
	// var ages [3]int = [3]int{25, 30, 35}
	var ages = [3]int{25, 30, 35}
	names := [4]string{"Alice", "Bob", "Charlie", "Bato"}
	names[3] = "Brighton"

	fmt.Println(names, len(names))
	fmt.Println(ages, len(ages))

	// slices are more flexible than arrays and can grow and shrink in size. They are built on top of arrays and provide a more convenient way to work with collections of data.
	var scores = []int{90, 85, 80}
	fmt.Println(scores, len(scores))
	scores[2] = 95
	fmt.Println(scores, len(scores))
	scores = append(scores, 100)
	fmt.Println(scores, len(scores))

	//slice ranges are a powerful feature in Go that allows you to create new slices from existing ones. You can specify a range of indices to create a new slice that includes only the elements within that range. The syntax for slicing is slice[low:high], where low is the starting index (inclusive) and high is the ending index (exclusive).
	subScores := scores[1:3]
	fmt.Println(subScores, len(subScores))
	nameRange := names[1:3]
	fmt.Println(nameRange, len(nameRange))
	//This gives us a new slice that includes the elements at indices 1 and 2 of the scores slice. The resulting subScores slice will contain the values [85, 95]. Similarly, we can create a new slice from the names array using slicing. The nameRange slice will contain the values ["Bob", "Charlie"].
	nameRangeBeggining := names[:2]
	fmt.Println(nameRangeBeggining, len(nameRangeBeggining))
	//nameRangeEnd slice give us a new slice that includes all the elements from index 2 to the end of the names array. The resulting nameRangeEnd slice will contain the values ["Charlie", "Bato"]. This is a convenient way to create a new slice without having to specify the ending index when you want to include all remaining elements.
	nameRangeEnd := names[2:]
	fmt.Println(nameRangeEnd, len(nameRangeEnd))

	//appending on slice
	nameRange = append(nameRange, "David")
	fmt.Println(nameRange, len(nameRange))

}
