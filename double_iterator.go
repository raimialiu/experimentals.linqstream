package linqstream

type DoubleIterator interface {
	Iterator[float64]
	NextDouble() float64
}
