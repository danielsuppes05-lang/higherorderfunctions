package filter

type Person struct {
	Name string
	Age  int
}

// FilterMinAge erwartet eine Liste von Personen und ein Mindestalter.
// FilterMinAge liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// die mindestens das angegebene Alter haben.
func FilterMinAge(list []Person, minAge int) []Person {
	e := []Person{}
	for _, v := range list {
		if v.Age >= minAge {
			e = append(e, v)
		}
	}
	return e
}

// FilterLongNames erwartet eine Liste von Personen und eine Mindestlänge.
// FilterLongNames liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// deren Name mindestens die angegebene Länge hat.
func FilterLongNames(list []Person, minLength int) []Person {
	e := []Person{}
	for _, v := range list {
		if len(v.Name) >= minLength {
			e = append(e, v)
		}
	}
	return e
}

// FilterNamePrefix erwartet eine Liste von Personen und einen Namenspräfix.
// FilterNamePrefix liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// deren Name mit dem angegebenen Präfix beginnt.
func FilterNamePrefix(list []Person, prefix string) []Person {
	e := []Person{}
	for _, v := range list {
		if v.Name[0] == prefix[0] {
			e = append(e, v)
		}
	}
	return e
}

// FilterChildren erwartet eine Liste von Personen.
// FilterChildren liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// die höchstens 13 Jahre alt sind.
func FilterChildren(list []Person) []Person {
	e := []Person{}
	for _, v := range list {
		if v.Age <= 13 {
			e = append(e, v)
		}
	}
	return e
}

// FilterChildrenWithLongNames erwartet eine Liste von Personen und eine Mindestlänge.
// FilterChildrenWithLongNames liefert eine neue Liste, die Personen aus der ursprünglichen Liste enthält,
// die höchstens 13 Jahre alt sind und deren Name mindestens die angegebene Länge hat.
func FilterChildrenWithLongNames(list []Person, minLength int) []Person {
	e := []Person{}
	for _, v := range list {
		if v.Age <= 13 && len(v.Name) >= minLength {
			e = append(e, v)
		}
	}
	return e
}
