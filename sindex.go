/*
Sindex is a slice indexing library. It maintains an ordered list, by mapping list positions to slice indexes.

List implementation types: List (Basic slice list), LinkedList (Todo)
*/
package sindex

//var IndexOutOfBounds = errors.New("List index out of bounds")

// An iterator iterates over a list and is able to modify the list.
// The state of an iterator is always at a position or is invalid.
// An iterator of a an empty list is always invalid.
//
// Thinking of an iterator as a cursor here are all the positions:
//                    invalid  Pos(0)  Pos(1)  Pos(2) ... Pos(len-1)  invalid
//  cursor positions:    ^       ^       ^       ^          ^            ^
//
//
// Given a list of length four and an iterator at Pos(1):
//                    invalid  Pos(0)  Pos(1)  Pos(2)  Pos(3)  invalid
//  cursor position:                     ^
// After a Remove():
//                    invalid  Pos(0)  invalid  Pos(1)  Pos(2)  invalid
//  cursor position:                      ^
// And then a call to Next():
//                    invalid  Pos(0)  Pos(1)  Pos(2)  invalid
//  cursor position:                     ^
type IteratorInterface interface {
	// Moves iterator to next position in a list. Returns true if next position
	// exists else false. This is also the validity state of the iterator. If
	// list is not empty and Next() returns false, to indicate end of the list,
	// a call to Prev() will return true and move the iterator to last item in
	// the list. (The opposite also applies when Prev() returns false to
	// indicate there are no more preceding items).
	Next() (valid bool)
	Prev() (valid bool)
	// Returns the index of the current position of the iterator. Must be called
	// on valid iterator.
	Pos() (index int)
	// Insert item before current position of the iterator. Returns the index of
	// the newly inserted item. The item to which the iterator points to does
	// not change but it's position in the list will have incremented. Must be
	// called on valid iterator.
	Insert() (index int)
	// Removes list item at position of iterator.
	// Changes the state of the iterator to invalid until a subsequent call to
	// Next()/Prev() returns true. Must be called on valid iterator.
	Remove()
	//	Err() (err error)
}

// Implementations provide this interface to access the slice.
type Interface interface {
	// Returns the index of where to put newly appended list item.
	Append() (index int)
	// Remove item at given position.
	Remove(pos int)
	// Insert item before given position. Returns index of where to put newly
	// inserted list item.
	Insert(pos int) (index int)
	// Return the index of the given position.
	Pos(pos int) (idex int)
	// Return length of list.
	Len() int
	// Return iterator. It is initially invalid.
	// The given position will be the position of the iterator after a Next() or
	// Prev() call.
	Iterator(pos int) IteratorInterface
	// Clear the list to be empty.
	Clear()
	//	Err() (err error)
}

type Option struct{}

func (o *Option) listOption() {}

type OptionInterface interface {
	listOption()
}
