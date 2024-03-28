package mem

import (
	"fmt"
)

func MapCopy() {
	roles := map[string]string{
		"admin":     "ROLE_ADMIN",
		"anonymous": "ROLE_ANONYMOUS",
	}

	println("Dumping original map")

	for key, value := range roles {
		fmt.Printf("%s %s\n", key, value)
	}
	println()
	println()

	rolesCopy := roles
	rolesCopy["user"] = "ROLE_USER"

	println("Dumping copy map")
	for key, value := range rolesCopy {
		fmt.Printf("%s %s\n", key, value)
	}
	println()
	println()

	println("Dumping original map")
	for key, value := range roles {
		fmt.Printf("%s %s\n", key, value)
	}
	println()
}
