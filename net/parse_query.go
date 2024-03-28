package net

import (
	"net/url"
)

func parseQuery() {
	a := url.Values{
		"username":                {"go-student"},
		"accepted_terms":          {"true"},
		"programming_languages[]": {"java", "cpp", "python"},
	}.Encode()

	println(a)

}
