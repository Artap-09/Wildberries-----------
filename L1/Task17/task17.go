package main

import (
	"fmt"
)

func main() {
	array := [10]int{55, 90, 74, 20, 16, 46, 43, 59, 2, 99}
	tree:=NewTree(array[0])
	for i := 1; i < 10; i++ {
		tree.Insert(array[i])
	}
	tree.Println()
	tree.Find(43)
	tree.Find(1)
}

type Tree struct {
	Root *Node
}

// NewTree - создает новую структур Tree с вершиной с задномы значением
func NewTree(val int) *Tree {
	return &Tree{NewNode(val)}
}

// Insert - метод для добавления значений.
func (t *Tree) Insert(val int) error {
	var err error

	nodeInsert, ok := t.Root.Find(val)

	if ok {
		err = fmt.Errorf("Значение \"%d\" уже существует", val)
		return err
	}

	if nodeInsert.Value > val { // Смотрим левого ребенка
		nodeInsert.LeftChild = NewNode(val)
		nodeInsert.LeftChild.Parent = nodeInsert

	} else { // Тоже самое но для правого ребенка
		nodeInsert.RightChild = NewNode(val)
		nodeInsert.RightChild.Parent = nodeInsert
	}
	nodeInsert.Balancing()
	return err
}

func (t Tree) Println (){
	fmt.Println(t.Root.Show())
}

func (t Tree) Find(val int) {
	if node, ok:=t.Root.Find(val); ok {
		fmt.Println(node.Value)
	} else {
		fmt.Println("Нету")
	}
}

type Node struct {
	Parent     *Node
	LeftChild  *Node
	RightChild *Node
	Value      int
	Rank       int
}

// NewNode - создает структуру Node с заданым значением.
func NewNode(val int) *Node {
	return &Node{Value: val, Rank: 1}
}

// Find - метод структуры Node. Ищет val.
func (n *Node) Find(val int) (*Node, bool) {
	if n.Value == val { // Ищем совпадение
		return n, true
	}

	if n.Value > val { // Смотрим левого ребенка
		if n.LeftChild != nil { // Если он есть
			return n.LeftChild.Find(val) // ВЫполняем поиск по левому ребенку
		}
	} else { // Тоже самое но для правого ребенка
		if n.RightChild != nil {
			return n.RightChild.Find(val)
		}
	}
	return n, false
}

// UpdateRank - обновляет rank в узле.
func (n *Node) UpdateRank() {
	if n.LeftChild != nil && n.RightChild != nil {
		if n.LeftChild.Rank > n.RightChild.Rank {
			n.Rank = n.LeftChild.Rank + 1
			return
		} else {
			n.Rank = n.RightChild.Rank + 1
			return
		}
	}
	if n.LeftChild != nil {
		n.Rank = n.LeftChild.Rank + 1
		return
	}

	if n.RightChild != nil {
		n.Rank = n.RightChild.Rank + 1
		return
	}
}

// Balancing - поддерживает соблюдения требований к Node что бы существовало АВЛ дерево.
func (n *Node) Balancing() {
	var dif int
	n.UpdateRank()
	if n.LeftChild!=nil && n.RightChild!=nil{
		dif = n.RightChild.Rank - n.LeftChild.Rank
	} else {
		if n.LeftChild==nil{
			dif=n.RightChild.Rank
		} else {
			dif=n.LeftChild.Rank
		}
	}

	if dif > 1 && dif < -1 {
		if dif > 0 { // Значит правая ветка больше левой
			if n.RightChild.LeftChild == nil || (n.RightChild.RightChild != nil && n.RightChild.LeftChild.Rank <= n.RightChild.RightChild.Rank) {
				n.Parent, n.RightChild.Parent = n.RightChild, n.Parent
				if n.Parent.Parent != nil {
					if n.Parent.Parent.Value < n.Value {
						n.Parent.Parent.RightChild = n.Parent
					} else {
						n.Parent.Parent.LeftChild = n.Parent
					}
				}
				n.Parent.LeftChild, n.RightChild = n, n.Parent.LeftChild
				if n.RightChild != nil {
					n.RightChild.Parent = n
				}
			} else {
				n.Parent, n.RightChild.LeftChild.Parent = n.RightChild.LeftChild, n.Parent
				if n.Parent.Parent != nil {
					if n.Parent.Parent.Value < n.Value {
						n.Parent.Parent.RightChild = n.Parent
					} else {
						n.Parent.Parent.LeftChild = n.Parent
					}
				}
				n.RightChild.LeftChild, n.Parent.RightChild = n.Parent.RightChild, n.RightChild
				n.RightChild.Parent = n.Parent
				if n.RightChild.LeftChild != nil {
					n.RightChild.LeftChild.Parent = n.RightChild
				}
				n.RightChild, n.Parent.LeftChild = n.Parent.LeftChild, n
				if n.RightChild != nil {
					n.RightChild.Parent = n
				}
			}
		} else {
			if n.LeftChild.RightChild == nil || (n.LeftChild.RightChild != nil && n.LeftChild.RightChild.Rank <= n.LeftChild.LeftChild.Rank) {
				n.Parent, n.LeftChild.Parent = n.LeftChild, n.Parent
				if n.Parent.Parent != nil {
					if n.Parent.Parent.Value < n.Value {
						n.Parent.Parent.RightChild = n.Parent
					} else {
						n.Parent.Parent.LeftChild = n.Parent
					}
				}
				n.Parent.RightChild, n.LeftChild = n, n.Parent.RightChild
				if n.LeftChild != nil {
					n.LeftChild.Parent = n
				}
			} else {
				n.Parent, n.LeftChild.RightChild.Parent = n.LeftChild.RightChild, n.Parent
				if n.Parent.Parent != nil {
					if n.Parent.Parent.Value < n.Value {
						n.Parent.Parent.RightChild = n.Parent
					} else {
						n.Parent.Parent.LeftChild = n.Parent
					}
				}
				n.LeftChild.RightChild, n.Parent.LeftChild = n.Parent.LeftChild, n.LeftChild
				n.LeftChild.Parent = n.Parent
				if n.LeftChild.RightChild != nil {
					n.LeftChild.RightChild.Parent = n.LeftChild
				}
				n.LeftChild, n.Parent.RightChild = n.Parent.RightChild, n
				if n.LeftChild != nil {
					n.LeftChild.Parent = n
				}
			}
		}
		if n.Parent.Parent != nil {
			n.Parent.Parent.Balancing()
		}
		return
	}
	if n.Parent != nil{
		n.Parent.Balancing()
	}
}

func (n Node) Show() string {
	var str string

	if n.LeftChild != nil {
		str+=n.LeftChild.Show()
	}

	str+=fmt.Sprintf("%d ",n.Value)

	if n.RightChild!= nil {
		str+=n.RightChild.Show()
	}

	return fmt.Sprintf("%s",str)
}
