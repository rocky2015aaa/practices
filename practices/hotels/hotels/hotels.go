package hotels

import (
	"sort"
	"strings"
)

type IdCountPair struct {
	Id    int
	Count int
}

type PairList []IdCountPair

func SortHotels(keywords, reviews []string, hotelIds []int) []int {
	idCounts := map[int]int{}
	for _, keyword := range keywords {
		for i := range reviews {
			index := strings.Index(reviews[i], keyword)
			if index > -1 {
				if count, ok := idCounts[hotelIds[i]]; ok {
					idCounts[hotelIds[i]] = count + 1
				} else {
					idCounts[hotelIds[i]] = 1
				}
			}
		}
	}
	length := len(idCounts)

	pl := make(PairList, length, length)
	i := 0
	for k, v := range idCounts {
		pl[i] = IdCountPair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))

	result := make([]int, length, length)
	for i := range pl {
		result[i] = pl[i].Id
	}

	return result
}

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].Count == p[j].Count {
		return p[i].Id < p[j].Id
	}
	return p[i].Count < p[j].Count
}
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
