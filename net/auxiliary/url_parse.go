package auxiliary

import (
	"fmt"
	"net/url"
)

func UrlParseSnippet() {
	rawUrl := "https://twitter.com/golang?view_as=guest"
	urlParsed, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", urlParsed)
	fmt.Printf("%s://%s:%s%s\n", urlParsed.Scheme, urlParsed.Hostname(), urlParsed.Port(), urlParsed.Path)
	fmt.Printf("RawQuery: %s\n", urlParsed.RawQuery)
	fmt.Printf("RequestURI: %s\n", urlParsed.RequestURI())
	println("=========================")
	println()

	rawUrl = "https://username:password@twitter.com:443/golang?view_as=guest"
	urlParsed, err = url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", urlParsed)
	fmt.Printf("%s://%s:%s%s\n", urlParsed.Scheme, urlParsed.Hostname(), urlParsed.Port(), urlParsed.Path)
	fmt.Printf("RawQuery: %s\n", urlParsed.RawQuery)
	fmt.Printf("RequestURI: %s\n", urlParsed.RequestURI())
	println("=========================")
	println()

	rawUrl = "//user:pass:pass:ew@twitter.com@facebook.com:443/golang?view_as=guest"
	urlParsed, err = url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", urlParsed)
	fmt.Printf("%s://%s:%s%s\n", urlParsed.Scheme, urlParsed.Hostname(), urlParsed.Port(), urlParsed.Path)
	fmt.Printf("RawQuery: %s\n", urlParsed.RawQuery)
	fmt.Printf("RequestURI: %s\n", urlParsed.RequestURI())
	println("=========================")
	println()
}
