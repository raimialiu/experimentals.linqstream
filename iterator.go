package linqstream

type Iterator[E any] interface {
	HasNext() bool
	Next() E
	Done()
	Remove()
	ForEachRemaining(consumer Consumer[E])
}
