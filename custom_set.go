package stringset

var seen struct{}

type Set map[string]struct{}

func New() Set {
	s := make(Set)
	return s
}

func NewFromSlice(input []string) Set {
	s := New()
	for _, v := range input {
		s[v] = seen
	}
	return s
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Add(str string) {
	s[str] = seen
}

func (s Set) Has(str string) bool {
	_, ok := s[str]
	return ok
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	return Subset(s1, s2)
}

func Subset(s1, s2 Set) bool {
	for k1 := range s1 {
		if _, ok := s2[k1]; !ok {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for k1 := range s1 {
		if _, ok := s2[k1]; ok {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	result := New()
	for k1 := range s1 {
		if _, ok := s2[k1]; ok {
			result.Add(k1)
		}
	}
	return result
}

func Difference(s1, s2 Set) Set {
	result := New()
	for k1 := range s1 {
		if _, ok := s2[k1]; !ok {
			result.Add(k1)
		}
	}
	return result
}

func Union(s1, s2 Set) Set {
	result := New()
	for k1 := range s1 {
		result.Add(k1)
	}
	for k2 := range s2 {
		result.Add(k2)
	}
	return result
}

func (s Set) String() string {
	str := "{"
	for k := range s {
		str += `"` + k + `", `
	}
	if len(str) > 1 {
		str = str[:len(str)-2]
	}
	return str + "}"
}
