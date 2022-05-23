package set

type set[T comparable] map[T]struct{}

// 创建 new
func New[T comparable]() set[T] {
	return make(set[T])
}

// 增 add
func (s set[T]) Add(elems ...T) set[T] {
	for _, elem := range elems {
		s[elem] = struct{}{}
	}
	return s
}

// 删 remove
func (s set[T]) Remove(elems ...T) set[T] {
	for _, elem := range elems {
		delete(s, elem)
	}
	return s
}

// 改 update
func (s set[T]) Update(old, new T) set[T] {
	s.Remove(old)
	s.Add(new)
	return s
}

// 查 has
func (s set[T]) Has(elem T) bool {
	_, ok := s[elem]
	return ok
}

// len
func (s set[T]) Len() int {
	return len(s)
}

// 返回所有元素 members
func (s set[T]) Members() []T {
	result := make([]T, 0, s.Len())
	for elem := range s {
		result = append(result, elem)
	}
	return result
}

// 随机返回一个并删除 pop
func (s set[T]) Pop() T {
	var elem T
	for k := range s {
		elem = k
		s.Remove(k)
	}
	return elem
}

// 随机返回一个不删除 one
func (s set[T]) One() T {
	var elem T
	for k := range s {
		elem = k
	}
	return elem
}

// 并集 union
func (s set[T]) Union(other set[T]) set[T] {
	for elem := range other {
		s[elem] = struct{}{}
	}
	return s
}

func Union[T comparable](set1, set2 set[T]) set[T] {
	for elem := range set2 {
		set1[elem] = struct{}{}
	}
	return set1
}

// 差集 diff
func (s set[T]) Diff(other set[T]) set[T] {
	for k1 := range other {
		if _, ok := s[k1]; ok {
			s.Remove(k1)
		}
	}
	return s
}

func Diff[T comparable](s1, s2 set[T]) set[T] {
	for k := range s1 {
		if _, ok := s2[k]; ok {
			delete(s1, k)
		}
	}
	return s1
}

// 交集 intersection
func (s set[T]) Intersection(other set[T]) set[T] {
	for k := range s {
		if _, ok := other[k]; !ok {
			s.Remove(k)
		}
	}
	return s
}

func Intersection[T comparable](s1, s2 set[T]) set[T] {
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			delete(s1, k)
		}
	}
	return s1
}
