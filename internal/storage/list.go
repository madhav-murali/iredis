package storage

import "strconv"

// TODO: change from map to a linked list/ deque for o(1) LPUSH
type List struct {
	Items map[string][]string
}

func NewList() *List {
	return &List{
		Items: make(map[string][]string),
	}
}

// func (l *List) RPUSH(key, val string) int {
// 	l.Items[key] = append(l.Items[key], val)
// 	return len(l.Items[key])
// }

func (l *List) RPUSH(key string, val []string) int {
	for _, v := range val {
		l.Items[key] = append(l.Items[key], v)
	}
	return len(l.Items[key])
}

func (l *List) LRANGE(key, start, end string) []string {
	var ret []string
	length := len(l.Items[key])
	s, err := strconv.Atoi(start)
	if err != nil {
		return nil
	}
	e, err := strconv.Atoi(end)
	if err != nil {
		return nil
	}
	if s < 0 {
		s = max(length+s, 0)
	}
	if e < 0 {
		e = length + e
	}

	for i := s; i <= e && i < length; i++ {
		ret = append(ret, l.Items[key][i])
	}
	return ret
}

// uses copy to extend the slice
func (l *List) LPUSH(key string, val []string) int {
	length := len(val)
	l.Items[key] = append(l.Items[key], make([]string, length)...)
	copy(l.Items[key][length:], l.Items[key])
	i := length - 1
	for _, v := range val {
		l.Items[key][i] = v
		i--
	}
	return len(l.Items[key])
}

func (l *List) LLEN(key string) int {
	return len(l.Items[key])
}
