package nsgaii

//ByDistance ::
type ByDistance []Individual

func (s ByDistance) Len() int {
	return len(s)
}

func (s ByDistance) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByDistance) Less(i, j int) bool {
	return s[i].Distance < s[j].Distance
}
