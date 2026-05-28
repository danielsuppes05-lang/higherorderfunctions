package sorting

import (
	"slices"
)

// SortAscending sortiert die Liste aufsteigend.
// Die Funktion verwendet dafür die Funktion slices.Sort aus dem Paket slices.
// Diese erwartet eine Liste und eine Funktion, die zwei Elemente der Liste vergleicht.
// Die Vergleichsfunktion (ein *Komparator*) erwartet zwei Position und liefert true,
// falls das Element an der ersten Position kleiner ist als das Element an der zweiten Position.
func SortAscending(list []int) {
	slices.SortFunc(list, func(a, b int) int {
		return a - b
	})
}

// SortStringsByLengthAscending erwartet eine Liste von Strings.
// SortStringsByLengthAscending sortiert die Liste nach der Länge der Strings aufsteigend.
// Die Funktion verwendet dafür die Funktion slices.Sort aus dem Paket slices.
func SortStringsByLengthAscending(list []string) {
	slices.SortFunc(list, func(a, b string) int {
		return len(a) - len(b)
	})
}

// Person ist (mal wieder) ein Datentyp für Personen.
type Person struct {
	Name string
	Age  int
}

// SortByAgeDescending sortiert die Liste nach dem Alter der Personen absteigend.
// Die Funktion verwendet dafür die Funktion slices.Sort aus dem Paket slices.
func SortByAgeDescending(list []Person) {
	// TODO
}

// SortByNameAscending sortiert die Liste nach dem Namen der Personen aufsteigend.
// Die Funktion verwendet dafür die Funktion slices.SortFunc aus dem Paket slices.
func SortByNameAscending(list []Person) {
	// TODO
}
