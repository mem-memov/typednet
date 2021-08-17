package typednet

type Class uint

func newClass(number uint) Class {
	return Class(number)
}

func (c Class) toInteger() uint {
	return uint(c)
}
