package gfa

import (
  "math/big"
  "testing"
)

var tables = []struct {
  p int64
  x int64
  y int64
  a int64
  s int64
  m int64
  d int64
}{
  {2, 1, 1, 0, 0, 1, 1},
  {3, 1, 1, 2, 0, 1, 1},
  {23, 7, 10, 17, 20, 1, 3},
}

func TestGF(t *testing.T) {
  for _, table := range tables {
    var gf = New(big.NewInt(table.p))
    z := gf.Add(big.NewInt(table.x), big.NewInt(table.y))
    if z.Cmp(big.NewInt(table.a)) != 0 {
      t.Errorf("GF(%v) %v+%v should be %v but %v", table.p, table.x, table.y, table.a, z)
    }
    z = gf.Sub(big.NewInt(table.x), big.NewInt(table.y))
    if z.Cmp(big.NewInt(table.s)) != 0 {
      t.Errorf("GF(%v) %v-%v should be %v but %v", table.p, table.x, table.y, table.s, z)
    }
    z = gf.Mul(big.NewInt(table.x), big.NewInt(table.y))
    if z.Cmp(big.NewInt(table.m)) != 0 {
      t.Errorf("GF(%v) %v*%v should be %v but %v", table.p, table.x, table.y, table.m, z)
    }
    z = gf.Div(big.NewInt(table.x), big.NewInt(table.y))
    if z.Cmp(big.NewInt(table.d)) != 0 {
      t.Errorf("GF(%v) %v/%v should be %v but %v", table.p, table.x, table.y, table.d, z)
    }

  }
}
