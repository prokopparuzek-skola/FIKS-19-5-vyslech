package main

import "fmt"

const (
	ODD  = false
	EVEN = true
)

type Word struct {
	end  bool
	next map[string]Word
}

func even(test int) bool {
	return test%2 == 0
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
		}
		if !SEven && AOdd {
			fmt.Println("Rassmo je vychytraly")
		}
		if !SEven && AEven {
			fmt.Println("Rassmo se priznal")
		}
	}
}
