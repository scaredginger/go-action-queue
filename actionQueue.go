package main

type Action interface {
	Execute()
}

type listNode struct {
	start int
	end   int
	arr   []Action
	next  *listNode
}

type ActionQueue struct {
	firstNode *listNode
	lastNode  *listNode
}

func (l *listNode) append(a Action) {
	l.arr[l.end] = a
	l.end++
}

func (l *ActionQueue) Append(a Action) {
	if l.lastNode.end == cap(l.lastNode.arr) {
		var n = &listNode{
			start: 0,
			end:   0,
			arr:   make([]Action, 1000, 1000),
			next:  nil,
		}
		l.lastNode.next = n
		l.lastNode = n
		l.lastNode.append(a)
	} else {
		l.lastNode.append(a)
	}
}

func (l *listNode) serve() Action {
	if l.start < l.end {
		var ret = l.arr[l.start]
		l.start++
		return ret
	}
	return nil
}

func (q *ActionQueue) Serve() bool {
	var s = q.firstNode.serve()
	if s == nil {
		if q.firstNode.next == nil {
			return false
		}
		q.firstNode = q.firstNode.next
		var a = q.firstNode.serve()
		a.Execute()
		return true
	}
	s.Execute()
	return true
}

func (q *ActionQueue) EmptyQueue() {
	for q.Serve() {
	}
}

func MakeActionQueue() ActionQueue {
	var n = &listNode{
		start: 0,
		end:   0,
		arr:   make([]Action, 1000, 1000),
		next:  nil,
	}
	return ActionQueue{
		firstNode: n,
		lastNode:  n,
	}
}
