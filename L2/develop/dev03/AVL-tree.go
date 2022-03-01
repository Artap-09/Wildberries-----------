package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Tree struct {
	Root *Node
}

// NewTree - создает новую структур Tree с вершиной с задномы значением
func NewTree() *Tree {
	return &Tree{nil}
}

// Insert - метод для добавления значений.
func (t *Tree) Insert(line []string, sf sortFunc) error {
	var err error

	if t.Root == nil {
		t.Root = NewNode(line)
		return err
	}

	nodeInsert, ok := t.Root.Find(line, sf)

	if ok {
		err = fmt.Errorf("значение \"%v\" уже существует", line)
		return err
	}

	newNode:=NewNode(line)

	if sf.Sort(sf.idx,nodeInsert,newNode) { // Смотрим левого ребенка
		nodeInsert.LeftChild = newNode
		nodeInsert.LeftChild.Parent = nodeInsert

	} else { // Тоже самое но для правого ребенка
		nodeInsert.RightChild = newNode
		nodeInsert.RightChild.Parent = nodeInsert
	}
	nodeInsert.Balancing(sf)
	return err
}

func (t Tree) Println() []string {
	return t.Root.Show()
}

func (t Tree) Find(line []string,sf sortFunc) {
	if node, ok := t.Root.Find(line,sf); ok {
		fmt.Println(node.Line)
	} else {
		fmt.Println("Нету")
	}
}

type Node struct {
	Parent     *Node
	LeftChild  *Node
	RightChild *Node
	Line       []string
	Rank       int
}

// NewNode - создает структуру Node с заданым значением.
func NewNode(line []string) *Node {
	return &Node{
		Line: line,
		Rank: 1,
	}
}

// Find - метод структуры Node. Ищет line.
func (n *Node) Find(line []string,sf sortFunc) (*Node, bool) {
	str1 := strings.Join(n.Line, " ")
	str2 := strings.Join(line, " ")
	if strings.Compare(str1, str2) == 0 { // Ищем совпадение
		return n, true
	}

	newNode:=NewNode(line)

	if sf.Sort(sf.idx,n,newNode) { // Смотрим левого ребенка
		if n.LeftChild != nil { // Если он есть
			return n.LeftChild.Find(line,sf) // ВЫполняем поиск по левому ребенку
		}
	} else { // Тоже самое но для правого ребенка
		if n.RightChild != nil {
			return n.RightChild.Find(line,sf)
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
func (n *Node) Balancing(sf sortFunc) {
	var dif int
	n.UpdateRank()
	if n.LeftChild != nil && n.RightChild != nil {
		dif = n.RightChild.Rank - n.LeftChild.Rank
	} else {
		if n.LeftChild == nil {
			dif = n.RightChild.Rank
		} else {
			dif = n.LeftChild.Rank
		}
	}

	if dif > 1 && dif < -1 {
		if dif > 0 { // Значит правая ветка больше левой
			if n.RightChild.LeftChild == nil || (n.RightChild.RightChild != nil && n.RightChild.LeftChild.Rank <= n.RightChild.RightChild.Rank) {
				n.Parent, n.RightChild.Parent = n.RightChild, n.Parent
				if n.Parent.Parent != nil {
					if sf.Sort(sf.idx, n, n.Parent.Parent) {
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
					if sf.Sort(sf.idx, n, n.Parent.Parent) {
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
					if sf.Sort(sf.idx, n, n.Parent.Parent) {
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
					if sf.Sort(sf.idx, n, n.Parent.Parent) {
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
			n.Parent.Parent.Balancing(sf)
		}
		return
	}
	if n.Parent != nil {
		n.Parent.Balancing(sf)
	}
}

func (n Node) Show() []string {
	var str []string

	if n.LeftChild != nil {
		str = append(str, n.LeftChild.Show()...)
	}

	str = append(str, strings.Join(n.Line, " "))

	if n.RightChild != nil {
		str = append(str, n.RightChild.Show()...)
	}

	return str
}

type sortFunc struct {
	idx  int
	Sort func(idx int, n1, n2 *Node) bool
}

// Сортировка по символам
func SortChar(idx int, n1, n2 *Node) bool {
	result := strings.Compare(n1.Line[idx], n2.Line[idx])
	/*if len(n1.Line[idx]) > len(n2.Line[idx]) {
		for i, v := range n2.Line[idx] {
			if v < rune(n1.Line[idx][i]) {
				return true
			}
		}
	} else {
		for i, v := range n1.Line[idx] {
			if v > rune(n2.Line[idx][i]) {
				return true
			}
		}
	}*/

	return result > 0
}

// Сортировка по числу
func SortInt(idx int, n1, n2 *Node) bool {
	nInt, err := strconv.Atoi(n1.Line[idx])
	if err != nil {
		log.Fatalln(err)
	}

	nodeInt, err := strconv.Atoi(n2.Line[idx])
	if err != nil {
		log.Fatalln(err)
	}

	
	return nInt > nodeInt 
}
