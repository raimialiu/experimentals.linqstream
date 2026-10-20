package linqstream

type ListIterator[E any] interface {
	Iterator[E]
	HasPrevious() bool
	Previous() E
	NextIndex() int
	PreviousIndex() int
	Set(index int, element E)
	Add(element E)
}
