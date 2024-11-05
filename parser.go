package main

// Turns tokens in a expression which can be evaluated.

type Parser struct {
	tokens                     []Token
	currentToken               Token
	totalTokens, current, peek int
}

// --------------------------------------------------------------------------------------------------------------------

func NewParser(tokens []Token) *Parser {
	parser := &Parser{
		tokens:      tokens,
		totalTokens: len(tokens),
		current:     0,
		peek:        0,
	}
	parser.readToken()

	return parser
}

// --------------------------------------------------------------------------------------------------------------------
// Parse Expressions: Addition and Subtraction.
// --------------------------------------------------------------------------------------------------------------------

func (p *Parser) ParseExpr() Token {
	term := p.parseTerm()

	for p.currentToken != nil {
		switch p.currentToken.(type) {
		case *Operator:
			if p.currentToken.String() == "+" {
				term = p.parseAddition(term)
			} else if p.currentToken.String() == "-" {
				term = p.parseSubtraction(term)
			} else {
				return term
			}
		default:
			return term
		}
	}
	return term
}

// --------------------------------------------------------

func (p *Parser) parseAddition(term Token) Token {
	p.readToken()
	term2 := p.parseTerm()
	return &Add{left: term, right: term2}
}

// --------------------------------------------------------

func (p *Parser) parseSubtraction(term Token) Token {
	p.readToken()
	term2 := p.parseTerm()
	return &Subtract{left: term, right: term2}
}

// --------------------------------------------------------------------------------------------------------------------
// Parse Terms: Multiplication and Division.
// --------------------------------------------------------------------------------------------------------------------

func (p *Parser) parseTerm() Token {
	factor := p.parseFactor()

	for p.currentToken != nil {
		p.readToken()

		switch p.currentToken.(type) {
		case *Operator:
			if p.currentToken.String() == "*" {
				factor = p.parseMultiplication(factor)
			} else if p.currentToken.String() == "/" {
				factor = p.parseDivision(factor)
			} else {
				return factor
			}
		default:
			return factor
		}
	}
	return factor
}

// --------------------------------------------------------

func (p *Parser) parseDivision(factor Token) Token {
	p.readToken()
	factor2 := p.parseFactor()
	return &Divide{left: factor, right: factor2}
}

// --------------------------------------------------------

func (p *Parser) parseMultiplication(factor Token) Token {
	p.readToken()
	factor2 := p.parseFactor()
	return &Multiply{left: factor, right: factor2}
}

// --------------------------------------------------------------------------------------------------------------------
// Parse Factors: Single digits and expressions inside parenthesis.
// --------------------------------------------------------------------------------------------------------------------

func (p *Parser) parseFactor() Token {
	switch p.currentToken.(type) {
	case *Number:
		return p.currentToken
	case *Operator:
		if p.currentToken.String() == "(" {
			return p.parseParen()
		}
		if p.currentToken.String() == "-" {
			p.readToken()
			return &Negative{right: p.parseFactor()}
		}
	}

	return p.currentToken
}

// --------------------------------------------------------

func (p *Parser) parseParen() Token {
	tokens := p.getSubTokens()
	subParser := NewParser(tokens)

	return &Number{subParser.ParseExpr().Eval()}
}

// --------------------------------------------------------

func (p *Parser) getSubTokens() []Token {
	subTokens := make([]Token, 0)

	for p.currentToken != nil {
		p.readToken()

		switch p.currentToken.(type) {
		case *Operator:
			if p.currentToken.String() == ")" {
				return subTokens
			} else if p.currentToken.String() == "(" {
				subTokens = append(subTokens, p.parseParen())
			} else {
				subTokens = append(subTokens, p.currentToken)
			}
		default:
			subTokens = append(subTokens, p.currentToken)
		}
	}

	return subTokens
}

// --------------------------------------------------------------------------------------------------------------------
// Helper functions.
// --------------------------------------------------------------------------------------------------------------------

func (p *Parser) readToken() {
	if p.peek >= p.totalTokens {
		p.currentToken = nil
	} else {
		p.current = p.peek
		p.peek++
		p.currentToken = p.tokens[p.current]
	}
}

// --------------------------------------------------------------------------------------------------------------------
