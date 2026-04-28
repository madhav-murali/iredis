package storage

type List struct {
	Items map[string][]string
}

func NewList() *List {
	return &List{
		Items: make(map[string][]string),
	}
}

func (l *List) RPUSH(key, val string) int {
	l.Items[key] = append(l.Items[key], val)
	return len(l.Items[key])
}

//rpush using a key to a list
