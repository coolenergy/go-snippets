package mem

// This snippet shows if a map when used in an if statement
// we get a copy or a direct reference
type Person struct {
	FirstName string
}

func IfFoundPtr() {
	melarDev := &Person{FirstName: "MelarDev"}

	personMap := map[string]*Person{
		"admin": melarDev,
	}

	var temp *Person
	if t, found := personMap["admin"]; found {
		temp = t
		t.FirstName = "Changed"
	} else {
		temp = &Person{
			FirstName: "NotFound",
		}
	}
	println(temp.FirstName)
}
