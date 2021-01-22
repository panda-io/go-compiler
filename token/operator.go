package token

//TO-DO use array instead map

var operatorRoot *operatorNode

func ReadOperator(bytes []byte) (Token, int) {
	return operatorRoot.find(bytes)
}

type operatorNode struct {
	children map[byte]*operatorNode
	token    Token
}

func (n *operatorNode) insert(operator string) {
	n.insertOperator(operator, 0)
}

func (n *operatorNode) find(bytes []byte) (Token, int) {
	return n.findOperator(bytes, 0)
}

func (n *operatorNode) findOperator(bytes []byte, offset int) (Token, int) {
	if child, ok := n.children[bytes[offset]]; ok {
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

func (n *operatorNode) insertOperator(operator string, position int) {
	if position < len(operator) {
		char := operator[position]
		if _, ok := n.children[char]; !ok {
			n.children[char] = &operatorNode{
				children: make(map[byte]*operatorNode),
				token:    ILLEGAL,
			}
		}
		position++
		n.children[char].insertOperator(operator, position)
	} else {
		n.token = ReadToken(operator)
	}
}
