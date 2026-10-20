package linqstream

type Consumer[E any] interface {
	Accept(valueFunc func(E))
	Apply(value E)
}
