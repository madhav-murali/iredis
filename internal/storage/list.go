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

func (l *List) LRANGE(key, start, end string) []string {
	var ret []string
	s, err := strconv.Atoi(start)
	if err != nil {
		return nil
	}
	e, err := strconv.Atoi(end)
	if err != nil {
		return nil
	}
	if s < 0 {
		s = len(l.Items[key]) + s
	}
	if e < 0 {
		e = len(l.Items[key]) + e
	}

	for i := s; i <= e && i < len(l.Items[key]); i++ {
		ret = append(ret, l.Items[key][i])
	}
	return ret
}

//rpush using a key to a list
