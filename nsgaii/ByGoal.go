package nsgaii

//ByGoal ::
type ByGoal []*Individual

func (s ByGoal) Len() int {
	return len(s)
}

func (s ByGoal) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByGoal) Less(i, j int) bool {
	goal := s[i].CurrentGoal
	return s[i].Goals[goal] < s[j].Goals[goal]
}
