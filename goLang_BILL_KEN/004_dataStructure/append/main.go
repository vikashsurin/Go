package main

import "fmt"

/*
  while appending ,slice doubles the memory allocated space ,when the capacity is reached.
*/
func main() {
	// u := make([]string, 1, 1)
	// fmt.Println(u, len(u), cap(u))
	// u = append(u, "dfdfdfdf", "fdfdf", "PPPP", "OOO", "dfdfdfdf", "fdfdf", "PPPP", "OOO", "dfdfdfdf", "fdfdf", "PPPP", "OOO")
	// fmt.Println(u, &u[0], len(u), cap(u))

	/*
				creating slice of a slice doesnot make a copy of a slice.
				it refers to the original slice.as a result any modification  made to the
				new slice will also affect the original or underlying slice.

		        TO MAKE A NEW COPY OF THE SLICE ,  one must use copy() functions.

	*/

	s := []string{"a", "b", "c", "d", "e"}
	fmt.Println(&s[0])
	fmt.Println(s)

	ss := s[0:4]
	fmt.Println(&ss[0])
	fmt.Println(ss)

	ss[0] = "ok"
	fmt.Println(s)
	fmt.Println(ss)

	sss := make([]string, len(s))
	// making a copy of the slice.
	copy(sss, s)
	fmt.Println(&sss[0])
}
