package mem

import (
	"fmt"
	"time"
)

func getMap() map[string]string {

	roles := map[string]string{
		"admin":     "ROLE_ADMIN",
		"anonymous": "ROLE_ANONYMOUS",
	}
	go func() {
		time.Sleep(2 * time.Second)
		println("Dumping original map")
		for key, value := range roles {
			fmt.Printf("%s %s\n", key, value)
		}
		println()

	}()
	return roles
}

func MapAssignReturn() {

	println("Dumping original map")
	roles := getMap()
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

	time.Sleep(time.Second * 5)
}
