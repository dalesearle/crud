package structs

type Library []*Book

func (l Library) Len() int           { return len(l) }
func (l Library) Less(i, j int) bool { return l[i].Id < l[j].Id }
func (l Library) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
