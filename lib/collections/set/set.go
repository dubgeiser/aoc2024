package set

type Set[T comparable] struct {
	// Empty struct uses no memory
	items map[T]struct{}
}

func New[T comparable]() *Set[T] {
	s := &Set[T]{make(map[T]struct{})}
	return s
}

func (s *Set[T]) Add(v T) {
	s.items[v] = struct{}{}
}

func (s *Set[T]) Delete(v T) {
	delete(s.items, v)
}

func (s *Set[T]) Len() int {
	return len(s.items)
}

func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

func (s *Set[T]) Contains(v T) bool {
	_, ok := s.items[v]
	return ok
}

func (s *Set[T]) Values() map[T]struct{} {
	return s.items
}

func (s *Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	intersection := New[T]()
	for e := range s.items {
		if s2.Contains(e) {
			intersection.Add(e)
		}
	}
	return intersection
}
