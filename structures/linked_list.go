package structures

// import "fmt"

type List struct {
	Item interface{}
	Next *List
}

func (l *List) Search(item interface{}) *List {
	if l == nil {
		return nil
	}
	if l.Item == item {
		return l
	}
	return l.Next.Search(item)
}

func (l *List) Insert(item interface{}) {
	lo := *l
	list := List{item, &lo}
	*l = list
}

func (l *List) Delete(item interface{}) {
	if toDelete := l.Search(item); toDelete != nil {
		pred := l.predecessor(item)
		if pred == nil {
			*l = *toDelete.Next
		} else {
			pred.Next = toDelete.Next
		}
	}
}

func (l *List) predecessor(item interface{}) *List {
	if l == nil || l.Next == nil {
		return nil
	}

	if l.Next.Item == item {
		return l
	}
	return l.Next.predecessor(item)
}
