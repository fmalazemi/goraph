package testgraph

// Graph01 represents test_graph_01.
var Graph01 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"S", "A"}, 100.0},
	{[]string{"S", "B"}, 14.0},
	{[]string{"S", "C"}, 200.0},

	{[]string{"A", "S"}, 15.0},
	{[]string{"A", "B"}, 5.0},
	{[]string{"A", "D"}, 20.0},
	{[]string{"A", "T"}, 44.0},

	{[]string{"B", "S"}, 14.0},
	{[]string{"B", "A"}, 5.0},
	{[]string{"B", "D"}, 30.0},
	{[]string{"B", "E"}, 18.0},

	{[]string{"C", "S"}, 9.0},
	{[]string{"C", "E"}, 24.0},

	{[]string{"D", "A"}, 20.0},
	{[]string{"D", "B"}, 30.0},
	{[]string{"D", "E"}, 2.0},
	{[]string{"D", "F"}, 11.0},
	{[]string{"D", "T"}, 16.0},

	{[]string{"E", "B"}, 18.0},
	{[]string{"E", "C"}, 24.0},
	{[]string{"E", "D"}, 2.0},
	{[]string{"E", "F"}, 6.0},
	{[]string{"E", "T"}, 19.0},

	{[]string{"F", "D"}, 11.0},
	{[]string{"F", "E"}, 6.0},
	{[]string{"F", "T"}, 6.0},

	{[]string{"T", "A"}, 44.0},
	{[]string{"T", "D"}, 16.0},
	{[]string{"T", "F"}, 6.0},
	{[]string{"T", "E"}, 19.0},
}

// Graph02 represents test_graph_02.
var Graph02 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"S", "A"}, 100.0},
	{[]string{"S", "B"}, 14.0},
	{[]string{"S", "C"}, 200.0},

	{[]string{"A", "S"}, 15.0},
	{[]string{"A", "B"}, 5.0},
	{[]string{"A", "D"}, 20.0},
	{[]string{"A", "T"}, 44.0},

	{[]string{"B", "S"}, 14.0},
	{[]string{"B", "A"}, 5.0},
	{[]string{"B", "D"}, 30.0},
	{[]string{"B", "E"}, 18.0},

	{[]string{"C", "S"}, 9.0},
	{[]string{"C", "E"}, 24.0},

	{[]string{"E", "B"}, 18.0},
	{[]string{"E", "D"}, 2.0},
	{[]string{"E", "F"}, 6.0},
	{[]string{"E", "T"}, 19.0},

	{[]string{"F", "D"}, 11.0},
	{[]string{"F", "E"}, 6.0},
	{[]string{"F", "T"}, 6.0},

	{[]string{"T", "A"}, 44.0},
	{[]string{"T", "D"}, 16.0},
	{[]string{"T", "F"}, 6.0},
	{[]string{"T", "E"}, 19.0},
}

// Graph03 represents test_graph_03.
var Graph03 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"S", "A"}, 100.0},
	{[]string{"S", "B"}, 20.0},
	{[]string{"S", "C"}, 200.0},

	{[]string{"A", "S"}, 15.0},
	{[]string{"A", "B"}, 5.0},
	{[]string{"A", "D"}, 20.0},
	{[]string{"A", "T"}, 44.0},

	{[]string{"B", "S"}, 14.0},
	{[]string{"B", "A"}, 5.0},
	{[]string{"B", "D"}, 30.0},
	{[]string{"B", "E"}, 18.0},

	{[]string{"C", "S"}, 9.0},
	{[]string{"C", "E"}, 24.0},

	{[]string{"E", "B"}, 18.0},
	{[]string{"E", "D"}, 2.0},
	{[]string{"E", "F"}, 6.0},
	{[]string{"E", "T"}, 19.0},

	{[]string{"F", "D"}, 11.0},
	{[]string{"F", "E"}, 6.0},
	{[]string{"F", "T"}, 6.0},

	{[]string{"T", "A"}, 44.0},
	{[]string{"T", "D"}, 16.0},
	{[]string{"T", "F"}, 6.0},
	{[]string{"T", "E"}, 19.0},
}

// Graph04 represents test_graph_04.
var Graph04 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"S", "B"}, 14.0},

	{[]string{"A", "S"}, 15.0},
	{[]string{"A", "B"}, 5.0},
	{[]string{"A", "D"}, 20.0},
	{[]string{"A", "T"}, 44.0},

	{[]string{"B", "S"}, 14.0},
	{[]string{"B", "A"}, 5.0},
	{[]string{"B", "D"}, 30.0},
	{[]string{"B", "E"}, 18.0},

	{[]string{"C", "S"}, 9.0},
	{[]string{"C", "E"}, 24.0},

	{[]string{"E", "B"}, 18.0},
	{[]string{"E", "C"}, 24.0},
	{[]string{"E", "D"}, 2.0},
	{[]string{"E", "F"}, 6.0},
	{[]string{"E", "T"}, 19.0},

	{[]string{"F", "D"}, 11.0},
	{[]string{"F", "E"}, 6.0},
	{[]string{"F", "T"}, 6.0},

	{[]string{"T", "A"}, 44.0},
	{[]string{"T", "D"}, 16.0},
	{[]string{"T", "F"}, 6.0},
	{[]string{"T", "E"}, 19.0},
}

// Graph05 represents test_graph_05.
var Graph05 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"A", "B"}, 7.0},
	{[]string{"A", "C"}, 9.0},
	{[]string{"A", "F"}, 20.0},

	{[]string{"B", "A"}, 7.0},
	{[]string{"B", "C"}, 10.0},
	{[]string{"B", "D"}, 15.0},

	{[]string{"C", "A"}, 9.0},
	{[]string{"C", "B"}, 10.0},
	{[]string{"C", "D"}, 11.0},
	{[]string{"C", "E"}, 30.0},
	{[]string{"C", "F"}, 2.0},

	{[]string{"D", "B"}, 15.0},
	{[]string{"D", "C"}, 11.0},
	{[]string{"D", "E"}, 2.0},

	{[]string{"E", "F"}, 9.0},
	{[]string{"E", "C"}, 30.0},
	{[]string{"E", "D"}, 2.0},

	{[]string{"F", "A"}, 20.0},
	{[]string{"F", "C"}, 2.0},
	{[]string{"F", "E"}, 9.0},
}

// Graph06 represents test_graph_06.
var Graph06 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"A", "F"}, 1.0},

	{[]string{"B", "A"}, 1.0},

	{[]string{"D", "B"}, 1.0},
	{[]string{"D", "C"}, 1.0},

	{[]string{"E", "C"}, 1.0},
	{[]string{"E", "F"}, 1.0},
}

// Graph07 represents test_graph_07.
var Graph07 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"A", "E"}, 1.0},
	{[]string{"A", "H"}, 1.0},

	{[]string{"B", "D"}, 1.0},

	{[]string{"C", "D"}, 1.0},
	{[]string{"C", "E"}, 1.0},

	{[]string{"D", "F"}, 1.0},
	{[]string{"D", "G"}, 1.0},
	{[]string{"D", "H"}, 1.0},

	{[]string{"E", "G"}, 1.0},
}

// Graph08 represents test_graph_08.
var Graph08 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"A", "E"}, 1.0},
	{[]string{"A", "F"}, 1.0},

	{[]string{"B", "A"}, 1.0},

	{[]string{"D", "B"}, 1.0},
	{[]string{"D", "C"}, 1.0},

	{[]string{"E", "D"}, 1.0},
	{[]string{"E", "C"}, 1.0},
	{[]string{"E", "F"}, 1.0},
}

// Graph09 represents test_graph_09.
var Graph09 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"A", "B"}, 1.0},
	{[]string{"A", "E"}, 1.0},
	{[]string{"A", "H"}, 1.0},

	{[]string{"B", "C"}, 1.0},
	{[]string{"B", "D"}, 1.0},

	{[]string{"C", "D"}, 1.0},
	{[]string{"C", "E"}, 1.0},

	{[]string{"D", "F"}, 1.0},
	{[]string{"D", "G"}, 1.0},
	{[]string{"D", "H"}, 1.0},

	{[]string{"E", "A"}, 1.0},
	{[]string{"E", "G"}, 1.0},

	{[]string{"F", "E"}, 1.0},

	{[]string{"G", "H"}, 1.0},

	{[]string{"H", "F"}, 1.0},
}

// Graph10 represents test_graph_10.
var Graph10 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"A", "C"}, 9.0},
	{[]string{"A", "F"}, 20.0},

	{[]string{"B", "A"}, 1.0},
	{[]string{"B", "D"}, 15.0},

	{[]string{"C", "B"}, 10.0},
	{[]string{"C", "E"}, 30.0},

	{[]string{"D", "C"}, 11.0},
	{[]string{"D", "E"}, 2.0},

	{[]string{"E", "C"}, 30.0},
	{[]string{"E", "F"}, 9.0},

	{[]string{"F", "A"}, 20.0},
	{[]string{"F", "C"}, 2.0},
}

// Graph11 represents test_graph_11.
var Graph11 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"S", "A"}, 11.0},
	{[]string{"S", "B"}, 17.0},
	{[]string{"S", "C"}, 9.0},

	{[]string{"A", "S"}, 11.0},
	{[]string{"A", "B"}, 5.0},
	{[]string{"A", "D"}, 50.0},
	{[]string{"A", "T"}, 500.0},

	{[]string{"B", "S"}, 17.0},
	{[]string{"B", "D"}, 30.0},

	{[]string{"C", "S"}, 9.0},

	{[]string{"D", "A"}, 50.0},
	{[]string{"D", "B"}, 30.0},
	{[]string{"D", "E"}, 3.0},
	{[]string{"D", "F"}, 11.0},

	{[]string{"E", "B"}, 18.0},
	{[]string{"E", "D"}, 2.0},
	{[]string{"E", "F"}, 6.0},
	{[]string{"E", "T"}, 19.0},

	{[]string{"F", "D"}, 11.0},
	{[]string{"F", "E"}, 6.0},
	{[]string{"F", "T"}, 77.0},

	{[]string{"T", "A"}, 500.0},
	{[]string{"T", "D"}, 10.0},
	{[]string{"T", "F"}, 77.0},
	{[]string{"T", "E"}, 19.0},
}

// Graph12 represents test_graph_12.
var Graph12 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"S", "A"}, 7.0},
	{[]string{"S", "B"}, 6.0},

	{[]string{"A", "C"}, -3.0},
	{[]string{"A", "T"}, 9.0},

	{[]string{"B", "A"}, 8.0},
	{[]string{"B", "C"}, 5.0},
	{[]string{"B", "T"}, -4.0},

	{[]string{"C", "B"}, -2.0},

	{[]string{"T", "C"}, 7.0},
	{[]string{"T", "S"}, 2.0},
}

// Graph13 represents test_graph_13.
var Graph13 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"S", "A"}, 7.0},
	{[]string{"S", "B"}, 6.0},

	{[]string{"A", "C"}, -3.0},
	{[]string{"A", "T"}, 9.0},

	{[]string{"B", "A"}, -8.0},
	{[]string{"B", "C"}, 5.0},
	{[]string{"B", "T"}, -4.0},

	{[]string{"C", "B"}, -2.0},

	{[]string{"T", "S"}, 2.0},
	{[]string{"T", "C"}, 7.0},
}

// Graph14 represents test_graph_14.
var Graph14 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"A", "B"}, 4.0},
	{[]string{"A", "H"}, 8.0},

	{[]string{"B", "A"}, 4.0},
	{[]string{"B", "H"}, 11.0},
	{[]string{"B", "C"}, 8.0},

	{[]string{"C", "B"}, 8.0},
	{[]string{"C", "I"}, 2.0},
	{[]string{"C", "F"}, 4.0},
	{[]string{"C", "D"}, 7.0},

	{[]string{"D", "C"}, 7.0},
	{[]string{"D", "F"}, 14.0},
	{[]string{"D", "E"}, 9.0},

	{[]string{"E", "D"}, 9.0},
	{[]string{"E", "F"}, 10.0},

	{[]string{"F", "G"}, 2.0},
	{[]string{"F", "C"}, 4.0},
	{[]string{"F", "D"}, 14.0},
	{[]string{"F", "E"}, 10.0},

	{[]string{"G", "H"}, 1.0},
	{[]string{"G", "I"}, 6.0},
	{[]string{"G", "F"}, 2.0},

	{[]string{"H", "A"}, 8.0},
	{[]string{"H", "B"}, 11.0},
	{[]string{"H", "I"}, 7.0},
	{[]string{"H", "G"}, 1.0},

	{[]string{"I", "H"}, 7.0},
	{[]string{"I", "G"}, 6.0},
	{[]string{"I", "C"}, 2.0},
}

// Graph15 represents test_graph_15.
var Graph15 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"A", "B"}, 1.0},

	{[]string{"B", "E"}, 1.0},
	{[]string{"B", "F"}, 1.0},
	{[]string{"B", "C"}, 1.0},

	{[]string{"C", "D"}, 1.0},
	{[]string{"C", "G"}, 1.0},

	{[]string{"D", "C"}, 1.0},
	{[]string{"D", "H"}, 1.0},

	{[]string{"E", "A"}, 1.0},
	{[]string{"E", "F"}, 1.0},

	{[]string{"F", "G"}, 1.0},

	{[]string{"G", "F"}, 1.0},
	{[]string{"G", "H"}, 1.0},

	{[]string{"H", "H"}, 1.0},
}

// Graph16 represents test_graph_16.
var Graph16 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"A", "B"}, 1.0},
	{[]string{"A", "F"}, 1.0},

	{[]string{"B", "C"}, 1.0},
	{[]string{"B", "G"}, 1.0},

	{[]string{"C", "D"}, 1.0},
	{[]string{"C", "H"}, 1.0},

	{[]string{"D", "H"}, 1.0},
	{[]string{"D", "I"}, 1.0},
	{[]string{"D", "J"}, 1.0},
	{[]string{"D", "E"}, 1.0},

	{[]string{"E", "J"}, 1.0},

	{[]string{"F", "G"}, 1.0},

	{[]string{"G", "A"}, 1.0},
	{[]string{"G", "H"}, 1.0},

	{[]string{"H", "C"}, 1.0},

	{[]string{"I", "J"}, 1.0},

	{[]string{"J", "E"}, 1.0},
}

// Graph17 represents test_graph_17.
var Graph17 = []struct {
	Vertices []string
	Weight   float32
}{
	{[]string{"S", "A"}, 10.0},
	{[]string{"S", "B"}, 5.0},
	{[]string{"S", "C"}, 15.0},

	{[]string{"A", "B"}, 4.0},
	{[]string{"A", "D"}, 9.0},
	{[]string{"A", "E"}, 15.0},

	{[]string{"B", "C"}, 4.0},
	{[]string{"B", "E"}, 8.0},

	{[]string{"C", "F"}, 16.0},

	{[]string{"D", "E"}, 15.0},
	{[]string{"D", "T"}, 10.0},

	{[]string{"E", "T"}, 10.0},
	{[]string{"E", "F"}, 15.0},

	{[]string{"F", "T"}, 10.0},
	{[]string{"F", "B"}, 6.0},
}
