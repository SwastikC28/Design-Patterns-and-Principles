package product

import "fmt"

type Pencil interface {
	Draw()
}

type Nataraj struct {
}

func (r *Nataraj) Draw() {
	fmt.Println("Drawn by Nataraj Pencil")
}

type Apsara struct {
}

func (p *Apsara) Draw() {
	fmt.Println("Drawn by Apsara Pencil")
}
