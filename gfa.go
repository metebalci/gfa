package gfa

import (
  "math/big"
)

var bigZero = big.NewInt(0)

// GF(p)
type GaloisField struct {
  p *big.Int
}

func init() {
}

func New(p *big.Int) *GaloisField {
  gf := new(GaloisField)
  gf.p = p
  return gf
}

type GaloisFieldArithmetic interface {
  Add(x, y *big.Int) *big.Int
  Sub(x, y *big.Int) *big.Int
  Mul(x, y *big.Int) *big.Int
  Inv(a *big.Int) *big.Int
  Div(x, y *big.Int) *big.Int
  Exp(x, y *big.Int) *big.Int
}

func (gf *GaloisField) Add(x, y *big.Int) *big.Int {
  return new(big.Int).Mod(new(big.Int).Add(x, y), gf.p)
}

func (gf *GaloisField) Sub(x, y *big.Int) *big.Int {
  return new(big.Int).Mod(new(big.Int).Sub(x, y), gf.p)
}

func (gf *GaloisField) Mul(x, y *big.Int) *big.Int {
  return new(big.Int).Mod(new(big.Int).Mul(x, y), gf.p)
}

func (gf *GaloisField) Inv(a *big.Int) *big.Int {
  x := new(big.Int)
  y := new(big.Int)
  new(big.Int).GCD(x, y, a, gf.p)
  return x
}

func (gf *GaloisField) Div(x, y *big.Int) *big.Int {
  return gf.Mul(x, new(big.Int).ModInverse(y, gf.p))
}

func (gf *GaloisField) Exp(x, y *big.Int) *big.Int {
  return new(big.Int).Exp(x, y, gf.p)
}
