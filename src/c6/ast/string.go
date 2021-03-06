package ast

type String struct {
	Value string
	Token *Token
}

func (self String) CanBeValue() {}

func (self String) CanBeExpression() {}
func (self String) String() string {
	return self.Value
}

func NewString(token *Token) *String {
	return &String{token.Str, token}
}
