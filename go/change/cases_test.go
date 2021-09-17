package change

// Source: exercism/problem-specifications
// Commit: 258c807 change: Use error object instead of sentinel -1 (#1340)
// Problem Specifications Version: 1.3.0

var testCases = []struct {
	description    string
	A              []int
	B              []int
	N              int  // true => no error, false => error expected
	expectedOutput bool // true => an directed graph consists of a set of vertices and a set of edges , false => otherwise
}{
	{
		"1 -> 2 -> 3 -> 4",
		[]int{1, 2, 3},
		[]int{2, 3, 4},
		4,
		true,
	},
	{
		"1 -> 2 -> 3 -> 4",
		[]int{1, 2, 4, 3},
		[]int{2, 3, 6, 4},
		4,
		true,
	},
	{
		"1 -> 2 -> 3 -> 4",
		[]int{1, 2, 4},
		[]int{2, 3, 3},
		4,
		true,
	},
	{
		"cannot directed to 1",
		[]int{2, 4, 3},
		[]int{3, 6, 4},
		4,
		true,
	},
	{
		"cannot directed 2 to 3",
		[]int{1, 2, 3},
		[]int{2, 4, 4},
		4,
		true,
	},
}
