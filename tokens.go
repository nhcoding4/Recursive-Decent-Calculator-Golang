package main

import "fmt"

// Token definitions.

type Token interface {
	String() string
	Eval() float64
}

// --------------------------------------------------------------------------------------------------------------------

type Add struct {
	left  Token
	right Token
}

func (a *Add) String() string {
	return fmt.Sprintf("(%v + %v)", a.left.Eval(), a.right.Eval())
}

func (a *Add) Eval() float64 {
	return a.left.Eval() + a.right.Eval()
}

// --------------------------------------------------------------------------------------------------------------------

type Subtract struct {
	left  Token
	right Token
}

func (s *Subtract) String() string {
	return fmt.Sprintf("(%v - %v)", s.left.Eval(), s.right.Eval())
}

func (s *Subtract) Eval() float64 {
	return s.left.Eval() - s.right.Eval()
}

// --------------------------------------------------------------------------------------------------------------------

type Divide struct {
	left  Token
	right Token
}

func (d *Divide) String() string {
	return fmt.Sprintf("(%v / %v)", d.left.Eval(), d.right.Eval())
}

func (d *Divide) Eval() float64 {
	return d.left.Eval() / d.right.Eval()
}

// --------------------------------------------------------------------------------------------------------------------

type Multiply struct {
	left  Token
	right Token
}

func (m *Multiply) String() string {
	return fmt.Sprintf("(%v * %v)", m.left.Eval(), m.right.Eval())
}

func (m *Multiply) Eval() float64 {
	return m.left.Eval() * m.right.Eval()
}

// --------------------------------------------------------------------------------------------------------------------

type Negative struct {
	right Token
}

func (n *Negative) String() string {
	return fmt.Sprintf("(-%v)", n.right.Eval())
}

func (n *Negative) Eval() float64 {
	return -n.right.Eval()
}

// --------------------------------------------------------------------------------------------------------------------

type Number struct {
	value float64
}

func (n *Number) String() string {
	return fmt.Sprintf("%v", n.value)
}

func (n *Number) Eval() float64 {
	return n.value
}

// --------------------------------------------------------------------------------------------------------------------

type Operator struct {
	value string
}

func (o *Operator) String() string {
	return o.value
}

func (o *Operator) Eval() float64 {
	return 0
}

// --------------------------------------------------------------------------------------------------------------------
