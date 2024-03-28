package auxiliary

import (
	"fmt"
	"golang.org/x/net/publicsuffix"
)

func GetTLD() {
	googleSubdomain := "subdomain.google.com"
	seekSubdomain := "subdomain.seek.com.au"

	tld, isManagedByIcann := publicsuffix.PublicSuffix(googleSubdomain)
	fmt.Printf("%s %t\n", tld, isManagedByIcann)

	tld, isManagedByIcann = publicsuffix.PublicSuffix(seekSubdomain)
	fmt.Printf("%s %t\n", tld, isManagedByIcann)

}
