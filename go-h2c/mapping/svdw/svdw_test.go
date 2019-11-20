package svdw_test

import (
	"testing"

	GF "github.com/armfazh/hash-to-curve-ref/go-h2c/field"
	"github.com/armfazh/hash-to-curve-ref/go-h2c/mapping"
	"github.com/armfazh/hash-to-curve-ref/go-h2c/mapping/svdw"
	"github.com/armfazh/hash-to-curve-ref/go-h2c/toy"
)

func TestMap(t *testing.T) {
	for _, EC := range toy.WeCurves {
		E := EC.E
		F := E.Field()
		n := F.Order().Int64()
		Z := F.Elt(50)
		for _, m := range []mapping.Map{
			svdw.New(E, Z, GF.SignLE),
			svdw.New(E, Z, GF.SignBE),
		} {
			for i := int64(0); i < n; i++ {
				u := F.Elt(i)
				P := m.MapToCurve(u)
				if !E.IsOnCurve(P) {
					t.Fatalf("%vu: %v\nP: %v not on curve.", m, u, P)
				}
			}
		}
	}
}
