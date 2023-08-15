package domain

import "errors"

type Stack struct {
	Name string `json:"name"`
}

type StackList struct {
	Stacks []Stack `json:"stack"`
}

func NewStack(stack string) (*Stack, error) {
	newStack := &Stack{stack}
	err := newStack.Validate()
	if err != nil {
		return nil, err
	}
	return newStack, nil
}

func (s *Stack) Validate() error {
	if s.Name == "" {
		return errors.New("invalid stack")
	}
	if len(s.Name) > 32 {
		return errors.New("invalid stack length")
	}
	return nil
}

func (l *StackList) AddStack(stack Stack) {
	l.Stacks = append(l.Stacks, stack)
}

func (l *StackList) GetStacks() []Stack {
	return l.Stacks
}
