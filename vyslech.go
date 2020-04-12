package main

import (
	"fmt"
)

const (
	ODD  = false
	EVEN = true
)

type Word struct {
	end  bool
	next map[string]Word
}

type Branch struct {
	odd    bool
	even   bool
	firstO bool
	firstE bool
}

func even(test int) bool {
	return test%2 == 0
}

func parse(word Word, depth int) (out Branch) {
	//fmt.Println(word, depth)
	var back Branch

	if word.end {
		if even(depth) {
			out.even = true
			if len(word.next) == 0 {
				out.firstE = true
			}
		} else {
			out.odd = true
			if len(word.next) == 0 {
				out.firstO = true
			}
		}
	}
	depth++
	for _, w := range word.next {
		back = parse(w, depth)
		out.even = out.even || back.even
		out.odd = out.odd || back.odd
		if len(word.next) > 1 || word.end {
			if !even(depth) {
				out.firstE = out.firstE || back.firstE
				out.firstO = out.firstO || back.firstO
			} else {
				if out.even && !out.odd || !out.even && out.odd {
					out.firstE = back.even
					out.firstO = back.odd
				} else {
					out.firstE = false
					out.firstO = false
				}
			}
		} else {
			out.firstE = out.firstE || back.firstE
			out.firstO = out.firstO || back.firstO
		}
	}
	//fmt.Println(depth-1, out)
	return
}

func main() {
	var T int
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var N, K int
		var SEven, AOdd, AEven bool
		AEven = true
		AOdd = true
		var sentences map[string]Word
		sentences = make(map[string]Word)
		var test Branch

		fmt.Scan(&N, &K)
		SEven = even(K)
		for i := 0; i < N; i++ {
			var M int
			var w string
			var tmp Word
			s := sentences

			fmt.Scan(&M)
			if even(M) && AEven {
				AEven = false
			} else if !even(M) && AOdd {
				AOdd = false
			}
			for j := 1; j <= M; j++ {
				fmt.Scan(&w)
				_, err := s[w]
				if !err {
					tmp.next = make(map[string]Word)
					tmp.end = false
					s[w] = tmp
				}
				if j == M {
					tmp = s[w]
					tmp.end = true
					s[w] = tmp
				}
				s = s[w].next
			}
		}
		//fmt.Println(sentences)
		if SEven && (AEven || AOdd) {
			fmt.Println("Rassmo je vychytraly")
			continue
		} else if !SEven && AOdd {
			fmt.Println("Rassmo je vychytraly")
			continue
		} else if !SEven && AEven {
			fmt.Println("Rassmo se priznal")
			continue
		}
		for _, s := range sentences {
			p := parse(s, 1)
			test.even = test.even || p.even
			test.odd = test.odd || p.odd
			test.firstE = test.firstE || p.firstE
			test.firstO = test.firstO || p.firstO
		}
		if test.even == test.firstE || test.odd == test.firstO {
			fmt.Println("Rassmo se priznal")
		} else {
			fmt.Println("Rassmo je vychytraly")
		}
	}
}
