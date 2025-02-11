package model

type Step int

func (b *Step) RemainStep() int { return int(MaxEdge) - int(*b) }

func (b *Step) Gaming() bool { return int(*b) < int(MaxEdge) }
