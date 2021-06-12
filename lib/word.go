package lib

import (
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ProcessWord(entity *string) string {
	re := regexp.MustCompile(REGEXP)
	trimmed := strings.TrimSpace(*entity)

	return re.ReplaceAllString(strings.ToLower(trimmed), "")
}

func GetWordCounters(doc *goquery.Document) map[string]int {
	wordCounters := make(map[string]int)
	doc.Find(SELECTOR).Each(func(ix int, link *goquery.Selection) {
		words := strings.Split(link.Text(), " ")
		for _, entity := range words {
			word := ProcessWord(&entity)
			if len(word) == 0 {
				continue
			}
			_, ok := wordCounters[word]
			if !ok {
				wordCounters[word] = 0
			}
			wordCounters[word] += 1
		}
	})

	return wordCounters
}

func RankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}
