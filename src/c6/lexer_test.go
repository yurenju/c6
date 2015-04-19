package c6

import "testing"
import "github.com/stretchr/testify/assert"

func TestLexerNext(t *testing.T) {
	l := NewLexerWithString(`.test {  }`)
	assert.NotNil(t, l)

	var r rune
	r = l.next()
	assert.Equal(t, '.', r)

	r = l.next()
	assert.Equal(t, 't', r)

	r = l.next()
	assert.Equal(t, 'e', r)

	r = l.next()
	assert.Equal(t, 's', r)

	r = l.next()
	assert.Equal(t, 't', r)
}

func TestLexerMatch(t *testing.T) {
	l := NewLexerWithString(`.foo {  }`)
	assert.NotNil(t, l)
	assert.False(t, l.match(".bar"))
	assert.True(t, l.match(".foo"))
}

func TestLexerAccept(t *testing.T) {
	l := NewLexerWithString(`.foo {  }`)
	assert.NotNil(t, l)
	assert.True(t, l.accept("."))
	assert.True(t, l.accept("f"))
	assert.True(t, l.accept("o"))
	assert.True(t, l.accept("o"))
	assert.True(t, l.accept(" "))
	assert.True(t, l.accept("{"))
}

func TestLexerIgnoreSpace(t *testing.T) {
	l := NewLexerWithString(`       .test {  }`)
	assert.NotNil(t, l)

	l.ignoreSpaces()

	var r rune
	r = l.next()
	assert.Equal(t, '.', r)

	l.backup()
	assert.True(t, l.match(".test"))
}

func TestLexerString(t *testing.T) {
	l := NewLexerWithString(`   "foo"`)
	output := l.getOutput()
	assert.NotNil(t, l)
	l.til("\"")
	lexString(l)
	token := <-output
	assert.Equal(t, T_QQ_STRING, token.Type)
}

func TestLexerTil(t *testing.T) {
	l := NewLexerWithString(`"foo"`)
	assert.NotNil(t, l)
	l.til("\"")
	assert.Equal(t, 0, l.Offset)
	l.next() // skip the quote

	l.til("\"")
	assert.Equal(t, 4, l.Offset)
}

func TestLexerAtRule(t *testing.T) {
	l := NewLexerWithString(`@import "test.css";`)
	assert.NotNil(t, l)
	l.run()
	AssertTokenSequence(t, l, []TokenType{T_IMPORT, T_QQ_STRING, T_SEMICOLON})
	l.close()
}

func TestLexerClassNameSelector(t *testing.T) {
	l := NewLexerWithString(`.class { }`)
	assert.NotNil(t, l)
	l.run()
	AssertTokenSequence(t, l, []TokenType{T_CLASS_SELECTOR, T_BRACE_START, T_BRACE_END})
	l.close()
}

func TestLexerRuleWithOneProperty(t *testing.T) {
	l := NewLexerWithString(`.test { color: #fff; }`)
	assert.NotNil(t, l)
	l.run()
	AssertTokenSequence(t, l, []TokenType{
		T_CLASS_SELECTOR,
		T_BRACE_START,
		T_PROPERTY_NAME, T_COLON, T_HEX_COLOR, T_SEMICOLON,
		T_BRACE_END})
	l.close()
}

func TestLexerRuleWithTwoProperty(t *testing.T) {
	l := NewLexerWithString(`.test { color: #fff; background: #fff; }`)
	assert.NotNil(t, l)
	l.run()
	AssertTokenSequence(t, l, []TokenType{
		T_CLASS_SELECTOR,
		T_BRACE_START,
		T_PROPERTY_NAME, T_COLON, T_HEX_COLOR, T_SEMICOLON,
		T_PROPERTY_NAME, T_COLON, T_HEX_COLOR, T_SEMICOLON,
		T_BRACE_END})
	l.close()
}
