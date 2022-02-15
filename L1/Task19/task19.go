package main

import (
	"errors"
	"fmt"
)

func main() {
	str:="Начало End"
	st:=NewStack()
	
	for _, v := range str {
		st.Push(string(v))
	}

	for st.Empty() {
		s,_:=st.Pop()
		fmt.Print(s)
	}
	fmt.Println()
}

type Stack struct {
	Head *StackElement
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(key string) {
	 se:=NewStackElement(key)
	 s.Head, se.Next=se,s.Head
}

func (s *Stack) Pop() (string, error) {
	out, err:= s.Top()
	
	if err == nil {
		s.Head= s.Head.Next
	}

	return out, err
}

func (s Stack) Top() (string, error) {
	var err error

	if s.Empty() {
		return s.Head.Value, err
	}

	err=errors.New("Стек пустой")
	return "", err
}

func (s Stack) Empty() bool {
	if s.Head==nil {
		return false
	}
	return true
}

type StackElement struct {
	Next  *StackElement
	Value string
}

func NewStackElement (key string) *StackElement{
	return &StackElement{Value: key}
}