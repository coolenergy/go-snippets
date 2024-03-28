package text

import (
	"fmt"
	"regexp"
)

func regexGroupSnippet() {
	regexStatus := "status: (?P<Status>\\w+),"
	regexAnswers := "ANSWER: (?P<AnswerCount>\\d*), AUTHORITY: (?P<AuthorityCount>\\d*),"

	r := regexp.MustCompile(regexStatus)
	r2 := regexp.MustCompile(regexAnswers)

	fmt.Printf("%#v\n", r.FindStringSubmatch(`;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 18233`))
	fmt.Printf("%#v\n", r.SubexpNames())

	fmt.Printf("%#v\n", r2.FindStringSubmatch(`;; flags: qr rd ra ; QUERY: 1, ANSWER: 2, AUTHORITY: 4, ADDITIONAL: 5`))
	fmt.Printf("%#v\n", r2.SubexpNames())

	resCount := r2.FindStringSubmatch(`;; flags: qr rd ra ; QUERY: 1, ANSWER: 2, AUTHORITY: 4, ADDITIONAL: 5`)[1:]
	answerCount := resCount[0]
	authorityCount := resCount[1]

	println(answerCount)
	println(authorityCount)
}
