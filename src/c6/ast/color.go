package ast

type Color interface {
	CanBeColor()
}

type HexColor struct {
	Hex   string
	Token *Token
}

func (self HexColor) CanBeExpression() {}
func (self HexColor) CanBeNode()       {}
func (self HexColor) CanBeColor()      {}
func (self HexColor) CanBeValue()      {}
func (self HexColor) String() string {
	return "#" + self.Hex
}

// Factor functions
func NewHexColor(hex string, token *Token) *HexColor {
	return &HexColor{hex, token}
}

type RGBAColor struct {
	R     float64
	G     float64
	B     float64
	A     float64
	Token *Token
}

func (self RGBAColor) CanBeExpression() {}
func (self RGBAColor) CanBeNode()       {}
func (self RGBAColor) CanBeColor()      {}
func (self RGBAColor) CanBeValue()      {}

func NewRGBAColor(r, g, b, a float64, token *Token) *RGBAColor {
	return &RGBAColor{r, g, b, a, token}
}

type RGBColor struct {
	R     float64
	G     float64
	B     float64
	Token *Token
}

func (self RGBColor) CanBeExpression() {}
func (self RGBColor) CanBeNode()       {}
func (self RGBColor) CanBeColor()      {}
func (self RGBColor) CanBeValue()      {}

func NewRGBColor(r, g, b float64, token *Token) *RGBColor {
	return &RGBColor{r, g, b, token}
}
