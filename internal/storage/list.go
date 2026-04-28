package storage

import "strconv"

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

func (l *List) MultiRPUSH(key string, val []string) int {
	for _, v := range val {
		l.Items[key] = append(l.Items[key], v)
	}
	return len(l.Items[key])
}

func (l *List) LRANGE(key string, s, e string) []string {
	var ret []string
	sInt, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	eInt, err := strconv.Atoi(e)
	if err != nil {
		return nil
	}
	end := min(eInt, len(l.Items[key]))
	for i := sInt; i < end; i++ {
		ret = append(ret, l.Items[key][i])
	}
	return ret
}

//rpush using a key to a list
