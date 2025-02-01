package model

type Ladder struct {
	start int
	end   int
}

func NewLadder(start, end int) Entity {
	return &Ladder{
		start: start,
		end:   end,
	}
}

func (l *Ladder) GetStartPosition() int {
	return l.start
}

func (l *Ladder) GetEndPosition() int {
	return l.end
}
