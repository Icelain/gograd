package gograd

import (
	"fmt"
	"math"
)

type Op uint

const (
	NoneOp Op = iota
	AddOp
	MulOp
	PowOp
)

type Neuron struct {
	Value        float64
	Grad         float64
	Partner      *Neuron
	Constituents [2]*Neuron // index 0 is always (first) and index 1 is always (to)
	Operation    Op // Operation used to create this Neuron
}

func (n Neuron) String() string {

	return fmt.Sprintf("Neuron{%.5f}", n.Value)

}

func (n *Neuron) Add(to *Neuron) Neuron {

	n.Partner = to
	to.Partner = n

	res := Neuron{Value: n.Value + to.Value}
	res.Constituents = [2]*Neuron{n, to}

	res.Operation = AddOp

	return res
}

func (n *Neuron) Mul(to *Neuron) Neuron {

	n.Partner = to
	to.Partner = n
	res := Neuron{Value: n.Value * to.Value}
	res.Constituents = [2]*Neuron{n, to}
	res.Operation = MulOp

	return res

}

// Here, the constituents have index 0 as the base and index 1 as the power
func (n *Neuron) Pow(to *Neuron) Neuron {

	n.Partner = to
	to.Partner = n
	res := Neuron{Value: math.Pow(n.Value, to.Value)}
	res.Constituents = [2]*Neuron{n, to}

	res.Operation = PowOp

	return res

}

func NewNeuron(value float64) Neuron {

	return Neuron{Value: value, Operation: NoneOp}

}

// first^to
func traverse(n *Neuron, prevgrad float64, operation Op, first bool) {

	head := n
		
	if head==nil {

		return

	}

	switch operation {

	case AddOp:
		head.Grad = prevgrad

	case MulOp:
		head.Grad = prevgrad * (head.Partner.Value)

	case PowOp:
		if first {
			head.Grad = prevgrad * (math.Pow(head.Value, head.Partner.Value) * (head.Partner.Value / head.Value))
		} else {

			head.Grad = prevgrad * (math.Pow(head.Partner.Value, head.Value) * (math.Log(head.Partner.Value)))

		}

	}
	
	traverse(head.Constituents[0], head.Grad, head.Operation, true)
	traverse(head.Constituents[1], head.Grad, head.Operation, false)


}

func (n *Neuron) Gradient() {

	n.Grad = 1.0

	traverse(n.Constituents[0], n.Grad, n.Operation, true)
	traverse(n.Constituents[1], n.Grad, n.Operation, false)

}
