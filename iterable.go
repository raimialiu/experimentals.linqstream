package linqstream

type Iterable[E any] interface {
	Iterate() Iterator[E]
	ForEach(consumer Consumer[E])
}
