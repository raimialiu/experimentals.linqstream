package linqstream

type Collection[E any] interface {
	Iterable[E]
	Add(element E)
	Remove(element E)
	AddAll(elements ...E)
	Contains(element E) bool
	Insert(index int, element E)
}
