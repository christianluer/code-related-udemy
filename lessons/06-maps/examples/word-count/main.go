// Counting words in a book: the opening of The Great Gatsby (HOE #52).
package main

import (
	"fmt"
	"sort"
	"strings"
)

const gatsby = `In my younger and more vulnerable years my father gave me some advice
that I've been turning over in my mind ever since. Whenever you feel like criticizing
anyone, he told me, just remember that all the people in this world haven't had the
advantages that you've had. He didn't say any more, but we've always been unusually
communicative in a reserved way, and I understood that he meant a great deal more
than that. In consequence, I'm inclined to reserve all judgements, a habit that has
opened up many curious natures to me and also made me the victim of not a few veteran
bores. The abnormal mind is quick to detect and attach itself to this quality when it
appears in a normal person, and so it came about that in college I was unjustly accused
of being a politician, because I was privy to the secret griefs of wild, unknown men.`

func main() {
	counts := map[string]int{}
	for _, w := range strings.Fields(gatsby) {
		w = strings.ToLower(strings.Trim(w, ".,;:!?"))
		counts[w]++
	}

	words := make([]string, 0, len(counts))
	for w := range counts {
		words = append(words, w)
	}
	sort.Slice(words, func(i, j int) bool {
		if counts[words[i]] != counts[words[j]] {
			return counts[words[i]] > counts[words[j]]
		}
		return words[i] < words[j]
	})

	fmt.Printf("%d distinct words. Top 10:\n", len(counts))
	for _, w := range words[:10] {
		fmt.Printf("%-6s %s %d\n", w, strings.Repeat("■", counts[w]), counts[w])
	}
}
