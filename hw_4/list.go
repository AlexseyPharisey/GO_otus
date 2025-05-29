package main

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(any) *ListItem
	PushBack(any) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value any
	Next  *ListItem
	Prev  *ListItem
}

type listLru struct {
	length int
	front  *ListItem
	back   *ListItem
}

func NewList() List {
	return &listLru{}
}

func (l *listLru) Len() int {
	return l.length
}

func (l *listLru) Front() *ListItem {
	return l.front
}

func (l *listLru) Back() *ListItem {
	return l.back
}

func (l *listLru) PushFront(v any) *ListItem {
	item := &ListItem{Value: v, Next: l.front}
	if l.front != nil {
		l.front.Prev = item
	} else {
		l.back = item
	}
	l.front = item
	l.length++

	return item
}

func (l *listLru) PushBack(v any) *ListItem {
	item := &ListItem{Value: v, Prev: l.back}
	if l.back != nil {
		l.back.Next = item
	} else {
		l.front = item
	}
	l.back = item
	l.length++

	return item
}

func (l *listLru) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}
	l.length--
}

func (l *listLru) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}
	if i == l.back {
		l.back = i.Prev
		l.back.Next = nil
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	i.Prev = nil
	i.Next = l.front
	l.front.Prev = i
	l.front = i
}
