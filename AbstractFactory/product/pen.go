package product

import "fmt"

type Pen interface {
	Sign()
}

type Reynolds struct {
}

func (r *Reynolds) Sign() {
	fmt.Println("Signed by Reynolds pen")
}

type Pilot struct {
}

func (p *Pilot) Sign() {
	fmt.Println("Signed by Pilot pen")
}
