package v1

func (l *Links) Len() int {
	return len(l.GetLink())
}

func (l *Links) Less(i, j int) bool {
	return l.GetLink()[i].GetCreatedAt().GetTime().Before(l.GetLink()[j].GetCreatedAt().GetTime())
}

func (l *Links) Swap(i, j int) {
	l.link[i], l.link[j] = l.GetLink()[j], l.GetLink()[i]
}
