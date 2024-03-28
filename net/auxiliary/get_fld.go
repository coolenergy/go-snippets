package auxiliary

import (
	"fmt"
	"golang.org/x/net/publicsuffix"
)

func GetFLD() {
	googleSubdomain := "subdomain.google.com"
	seekSubdomain := "subdomain.seek.com.au"
	suffix, err := publicsuffix.EffectiveTLDPlusOne(googleSubdomain)
	if err != nil {
		panic(err)
	}
	fmt.Println(suffix)

	suffix, err = publicsuffix.EffectiveTLDPlusOne(seekSubdomain)
	if err != nil {
		panic(err)
	}

	fmt.Println(suffix)
}
