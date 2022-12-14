// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var ts, ss, eq int
	//ar := []int{1, 5, 11, 5}
	//ar := []int{1, 3, 3, 3, 3, 1}
	//ar := []int{4, 5, 1, 1, 1}
	//ar := []int{3, 34, 4, 12, 5, 2}
	ar := []int{34, 3, 4, 12, 5, 2, 8}
	var s1, s2 = []int{}, []int{}
	for i := 0; i < len(ar); i++ {
		ts += ar[i]
	}
	fmt.Println("total sum", ts)
	eq = ts / 2
	if ts%2 != 0 {
		fmt.Println("Cannot divide equally")
		return
	}
	candivide := false
	for i := 0; i < len(ar); i++ {
		ss += ar[i]

		if ss < eq {
			s1 = append(s1, ar[i])
		}

		if ss == eq {
			s1 = append(s1, ar[i])
			candivide = true
		}

		if ss > eq {
			ss -= ar[i]
			s2 = append(s2, ar[i])
		}
	}
	
	if !candivide {
		fmt.Println("Unable to divide in equally subarray")
		return
	}
	fmt.Println("equal sub arrays are",s1, s2)
}
