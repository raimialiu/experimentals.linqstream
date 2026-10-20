package linqstream

var (
	ItemModError = "Item modification is not permitted"
)

type ArrayListIterator[E any] struct {
	_elementData      []E
	_cursor           int
	_object           ArrayList[E]
	_expectedModCount int
	_lastResult       int
}

func (a *ArrayListIterator[E]) HasNext() bool {
	return a._cursor < len(a._elementData)
}

func (a *ArrayListIterator[E]) Next() E {
	a.checkForModification()
	i := a._cursor

	if !a.HasNext() {
		var zero E
		if a._lastResult == -1 {
			return zero
		}
		return a._elementData[a._lastResult]
	}

	if i > len(a._elementData) {
		panic(ItemModError)
	}

	a._cursor += 1
	a._lastResult = i
	element := a._elementData[a._lastResult]
	return element
}

func (a *ArrayListIterator[E]) Done() {
	if !a.HasNext() {
		panic("iteration already completed")
	}
	a._cursor = len(a._elementData)
}

func (a *ArrayListIterator[E]) Remove() {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayListIterator[E]) ForEachRemaining(consumer Consumer[E]) {
	for a.HasNext() {
		current := a.Next()
		consumer.Apply(current)
	}
}

func (a *ArrayListIterator[E]) checkForModification() {
	if a._expectedModCount != a._object.ModCount {
		panic(ItemModError)
	}
}
