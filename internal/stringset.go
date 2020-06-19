package internal

type StringSet map[string]struct{}

func NewStringSet() StringSet {
	return make(StringSet)
}

func (s StringSet) Add(i string) {
	s[i] = struct{}{}
}

func (s StringSet) Remove(i string) {
	delete(s, i)
}

func (s StringSet) Contains(i string) bool {
	_, ok := s[i]
	return ok
}
