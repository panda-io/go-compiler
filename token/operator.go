package token

var operatorRoot *operatorNode

func init() {
	operatorRoot = &operatorNode{
		children: make(map[byte]*operatorNode),
		token:    ILLEGAL,
	}
	for i := operatorBegin + 1; i < operatorEnd; i++ {
		operatorRoot.insert(tokenStrings[i])
	}
}

// ReadOperator greedy read operators by tree search
func ReadOperator(bytes []byte) (Token, int) {
	return operatorRoot.find(bytes)
}

type operatorNode struct {
	children map[byte]*operatorNode
	token    Token
}

func (node *operatorNode) insert(operator string) {
	node.insertOperator(operator, 0)
}

func (node *operatorNode) find(bytes []byte) (Token, int) {
	return node.findOperator(bytes, 0)
}

func (node *operatorNode) findOperator(bytes []byte, offset int) (Token, int) {
	if child, ok := node.children[bytes[offset]]; ok {
		offset++
		if offset < len(bytes) {
			return child.findOperator(bytes, offset)
		}
		return child.token, offset
	} else if offset > 0 {
		return ReadToken(string(bytes[:offset])), offset
	}
	return ILLEGAL, 1
}

func (node *operatorNode) insertOperator(operator string, position int) {
	if position < len(operator) {
		char := operator[position]
		if _, ok := node.children[char]; !ok {
			node.children[char] = &operatorNode{
				children: make(map[byte]*operatorNode),
				token:    ILLEGAL,
			}
		}
		position++
		node.children[char].insertOperator(operator, position)
	} else {
		node.token = ReadToken(operator)
	}
}
