package linqstream

type ArrayList[E any] struct {
	elementData [100]E
	cursor      int
	Size        int
	ModCount    int
}
