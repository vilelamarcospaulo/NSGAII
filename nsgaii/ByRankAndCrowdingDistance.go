package nsgaii

//ByRankAndCrowndingDistance ::
type ByRankAndCrowdingDistance []Individual

func (s ByRankAndCrowdingDistance) Len() int {
	return len(s)
}

func (s ByRankAndCrowdingDistance) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByRankAndCrowdingDistance) Less(i, j int) bool {
	return s[i].Better(s[j])
}
