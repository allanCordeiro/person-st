package domain

import (
	"errors"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Given a valid stack when create an object should be ok", func(t *testing.T) {
		expectedStack := "Go Lang"
		stack, err := NewStack(expectedStack)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if stack.Name != expectedStack {
			t.Errorf("value error. Expected %v, but received %v.", expectedStack, stack.Name)
		}
	})

	t.Run("Given a empty stack when create an object should return error", func(t *testing.T) {
		expectedError := errors.New("invalid stack")
		_, err := NewStack("")
		if err == nil {
			t.Errorf("an error was expected, but returned nil")
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("value error. Expected %v, but received %v.", expectedError.Error(), err.Error())
		}
	})

	t.Run("Given a stack with more than 32 chars when create an object should return error", func(t *testing.T) {
		expectedError := errors.New("invalid stack length")
		_, err := NewStack("This should be a nice one stack or a list of then, but thats so not right")
		if err == nil {
			t.Errorf("an error was expected, but returned nil")
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("value error. Expected %v, but received %v.", expectedError.Error(), err.Error())
		}
	})
}

func TestStackList(t *testing.T) {
	expectedLength := 4
	goLang, _ := NewStack("Go")
	python, _ := NewStack("Python")
	php, _ := NewStack("PHP")
	java, _ := NewStack("Java")
	list := StackList{}
	list.AddStack(*goLang)
	list.AddStack(*python)
	list.AddStack(*php)
	list.AddStack(*java)

	receivedLength := len(list.GetStacks())
	if receivedLength != expectedLength {
		t.Errorf("value error. Expected %v, but received %v.", expectedLength, receivedLength)
	}

	output := list.GetStacks()
	for _, stack := range output {
		check := 0
		if stack.Name == goLang.Name {
			check++
		}
		if stack.Name == python.Name {
			check++
		}
		if stack.Name == php.Name {
			check++
		}
		if stack.Name == java.Name {
			check++
		}
		if check == 0 {
			t.Errorf("The stack %v wasnt found in the stack list", stack.Name)
		}
	}
}
