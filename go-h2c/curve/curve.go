package curve

import (
	"math/big"

	GF "github.com/armfazh/hash-to-curve-ref/go-h2c/field"
)

// Point is an elliptic curve point
type Point interface {
	IsIdentity() bool
	IsEqual(Point) bool
	Copy() Point
	X() GF.Elt
	Y() GF.Elt
}

// EllCurve is an elliptic curve
type EllCurve interface {
	Field() GF.Field
	Order() *big.Int
	Cofactor() *big.Int
	NewPoint(x, y GF.Elt) Point
	Identity() Point
	Neg(Point) Point
	Add(Point, Point) Point
	Double(Point) Point
	IsOnCurve(Point) bool
	ClearCofactor(Point) Point
}

// RationalMap is
type RationalMap interface {
	Domain() EllCurve
	Codomain() EllCurve
	Push(Point) Point
	Pull(Point) Point
}

//
// type Model int
//
// const (
// 	Weierstrass Model = iota
// 	Montgomery
// 	Edwards
// )
