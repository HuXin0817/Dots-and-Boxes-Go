package internal

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/HuXin0817/dots-and-boxes/src/model/board"
	"gopkg.in/Knetic/govaluate.v2"
)

type Interface interface {
	BestCandidateEdges(*board.BoardV2) (edges []model.Edge)
}

var ErrModelFormat = errors.New("model format error")

func NewInterface(s string) (Interface, error) {
	for len(s) >= 2 && s[0] == '(' && s[len(s)-1] == ')' {
		s = s[1 : len(s)-1]
	}
	s = strings.ToUpper(s)
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, "")
	switch s {
	case "L0", "L0()":
		return NewL0Model(), nil
	case "L1", "L1()":
		return NewL1Model(), nil
	case "L2", "L2()":
		return NewL2Model(), nil
	case "L3", "L3()":
		return DefaultL3Model(), nil
	case "L4", "L4()":
		return DefaultL4Model(), nil
	}
	if len(s) <= 4 {
		return nil, ErrModelFormat
	}
	if s[2] != '(' || s[len(s)-1] != ')' {
		return nil, ErrModelFormat
	}
	switch s[:2] {
	case "L3":
		p1 := 3
		lpCount := 0
		rpCount := 0
		for p1 < len(s) {
			if s[p1] == '(' {
				lpCount++
			}
			if s[p1] == ')' {
				rpCount++
			}
			if s[p1] == ',' && lpCount == rpCount {
				break
			}
			p1++
		}
		arg1 := s[3:p1]
		if p1 == len(s) {
			arg1 = s[3 : p1-1]
		}
		exp, err := govaluate.NewEvaluableExpression(arg1)
		if err != nil {
			return nil, err
		}
		result, err := exp.Evaluate(nil)
		if err != nil {
			return nil, err
		}
		f, err := strconv.ParseFloat(fmt.Sprint(result), 64)
		if err != nil {
			return nil, err
		}
		if p1 == len(s) {
			return NewL3Model(int(f), NewL2Model()), nil
		}
		if p1+1 >= len(s) {
			return nil, ErrModelFormat
		}
		arg2 := s[p1+1 : len(s)-1]
		M, err := NewInterface(arg2)
		if err != nil {
			return nil, err
		}
		return NewL3Model(int(f), M), nil
	case "L4":
		M, err := NewInterface(s[3 : len(s)-1])
		if err != nil {
			return nil, err
		}
		if _, ok := M.(*L3Model); !ok {
			return nil, ErrModelFormat
		}
		return NewL4Model(func() *L3Model {
			M, _ = NewInterface(s[3 : len(s)-1])
			return M.(*L3Model)
		}), nil
	}
	return nil, ErrModelFormat
}
