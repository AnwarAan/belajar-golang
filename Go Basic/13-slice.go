package main

import "fmt"

func main() {

	month := [12]string{
		"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December",
	}
	slice1 := month[4:7]

	fmt.Println(month)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(append(slice1, ""))

	score := [...]int{
		1, 2, 3, 4, 5,
	}
	fmt.Println(cap(score))

	days := [...]string{
		"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu",
	}
	sliceDay := days[:5]
	sliceDay2 := append(sliceDay, "sabtu", "minggu")
	sliceDay[2] = "pasar rebo"
	sliceDay[3] = "pasar kemis"
	fmt.Println(sliceDay)
	fmt.Println(sliceDay2)
	fmt.Println(days)

	makeSlice := make([]string, 3, 10)
	makeSlice[0] = "mca"
	makeSlice[1] = "mcnwr"
	fmt.Println(makeSlice)
	fmt.Println(len(makeSlice))
	fmt.Println(cap(makeSlice))

	copySlice := make([]string, len(makeSlice), cap(makeSlice))
	copySlice2 := make([]string, 1, cap(makeSlice))
	copy(copySlice, makeSlice)
	fmt.Println(copySlice)
	fmt.Println(copySlice2)

	array := [...]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(array)
	fmt.Println(slice)

	// sliceProb := days[0:2]
	sliceProb := days[4:5]
	sliceProb2 := append(sliceProb, "ahad")
	sliceProb2[0] = "libur"
	fmt.Println(sliceProb)
	fmt.Println(sliceProb2)
	fmt.Println(days)

}

// go bulid hello-world.go
// go run hello-world.go
