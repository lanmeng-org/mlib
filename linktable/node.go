package linktable


type node struct {
	PreviousNode *node
	NextNode *node
	Content interface{}
}

func NewLinkedNode() *node {
	return &node{}
}

