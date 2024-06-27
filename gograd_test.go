package gograd

import "testing"

func assert(t *testing.T, f1, f2 float64) {

	if f1 != f2 {

		t.Error(f1, "is not equal to", f2)

	}

}

func TestGradient(t *testing.T) {

	n1 := NewNeuron(5.0)
	n2 := NewNeuron(2.1)
	n3 := n1.Add(&n2)
	n4 := NewNeuron(4.4)
	n5 := n3.Mul(&n4)
	n6 := NewNeuron(1.3)

	n7 := n5.Pow(&n6)
	n7.Gradient()

	assert(t, n7.Grad, 1.0)
	assert(t, n6.Grad, 301.9237378146226)
	assert(t, n5.Grad, 3.650536203812254)
	assert(t, n4.Grad, 25.918807047067002)
	assert(t, n3.Grad, 16.06235929677392)
	assert(t, n2.Grad, 16.06235929677392)
	assert(t, n1.Grad, 16.06235929677392)

}
