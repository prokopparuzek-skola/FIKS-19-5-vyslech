package main

import "fmt"

type Word struct {
	end  bool
	next map[string]Word
}

func main() {
	var T int
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var N, K int
		fmt.Scan(&N, &K)
		var sentences map[string]Word
		sentences = make(map[string]Word)
		for i := 0; i < N; i++ {
			var M int
			var w string
			var tmp Word
			s := sentences

			fmt.Scan(&M)
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
		fmt.Println(sentences)
	}
}
