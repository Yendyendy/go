package main

import (
	"fmt"
	"math/rand"
)

// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

func main() {
	fmt.Println("Hello world")

	t1 := New(1)
	t2 := New(2)
	fmt.Println(t1, "\n", t2)

	same := Same(t1, t2)
	fmt.Println(same)
}

func Walk(t *Tree, c chan int) {

	if t == nil {
		return
	}
	Walk(t.Left, c)
	c <- t.Value
	Walk(t.Right, c)
}

func Same(t1 *Tree, t2 *Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	for i := 0; i < 10; i++ {
		x := <-c1
		y := <-c2
		if x != y {
			return false
		}
	}
	return true
}
