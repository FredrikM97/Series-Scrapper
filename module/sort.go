package module

import "sort"

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func sort_on_keyValue(data_map map[string]int) PairList {
	p := make(PairList, len(data_map))

	i := 0
	for k, v := range data_map {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}
