package fse

type History struct {
	max     int
	data    *HistoryList
	current *HistoryNode
}

func NewHistory() *History {
	list := newHistoryList()
	return &History{
		max:     3000,
		data:    list,
		current: list.root,
	}
}

func (h *History) ResetCurrent() {
	h.current = h.data.root
}

func (h *History) Add(line string) {
	h.data.append(line)
	if h.data.size > h.max {
		h.data.removeTail()
	}
}

func (h *History) Prev() string {
	if !h.current.hasPrev() {
		return ""
	}

	h.current = h.current.prev
	return h.current.data
}

func (h *History) Current() string {
	return h.current.data
}

func (h *History) Next() string {
	if !h.current.hasNext() {
		return ""
	}

	h.current = h.current.next
	return h.current.data
}

type HistoryNode struct {
	data string
	prev *HistoryNode
	next *HistoryNode
}

func (n *HistoryNode) hasPrev() bool {
	return n.prev != nil
}

func (n *HistoryNode) hasNext() bool {
	return n.next != nil && n.next.data != ""
}

type HistoryList struct {
	root     *HistoryNode
	tailNode *HistoryNode
	set      map[string]*HistoryNode
	size     int
}

func newHistoryList() *HistoryList {
	emptyNode := new(HistoryNode)
	return &HistoryList{
		root:     emptyNode,
		tailNode: emptyNode,
		set:      make(map[string]*HistoryNode),
	}
}

func (l *HistoryList) append(data string) {
	if l.contains(data) {
		return
	}

	node := &HistoryNode{data: data}

	if l.root.prev != nil {
		node.prev = l.root.prev
		node.prev.next = node
	}

	l.root.prev = node
	node.next = l.root

	if l.tailNode == l.root {
		l.tailNode = node
	}

	l.set[data] = node
	l.size++
}

func (l *HistoryList) contains(data string) bool {
	_, ok := l.set[data]
	return ok
}

func (l *HistoryList) remove(data string) bool {
	node := l.set[data]
	if node == nil {
		return false
	}

	if l.tailNode == node {
		l.tailNode = node.next
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev = nil
	node.next = nil

	delete(l.set, data)

	l.size--
	return true
}

func (l *HistoryList) removeTail() bool {
	return l.remove(l.tailNode.data)
}

func (l *HistoryList) isEmpty() bool {
	return l.size == 0
}
