package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	list := newList()

	list.insert(10)
	list.insert(2)
	list.insert(3)
	list.insert(20)
	list.insert(32)

	fmt.Printf("list = %s\n", list)

	list.insert(109)
	list.insert(2)
	list.insert(50)

	fmt.Printf("list = %s\n", list)

	nd1 := list.find(49)
	fmt.Printf("nd1 = %s\n", nd1)

	nd2 := list.find(50)
	fmt.Printf("nd2 = %s\n", nd2)
}

const (
	MAX_LEVEL = 32
	PROP      = 0.25
)

type list struct {
	header *node
}

type node struct {
	key     int
	forward []*node
}

func newNode() *node {
	return &node{
		forward: make([]*node, MAX_LEVEL),
	}
}
func (c *node) String() string {
	return fmt.Sprintf("%d", c.key)
}

func newList() *list {
	return &list{
		header: newNode(),
	}
}
func randomLevel() int {
	i := 1
	for rand.Float64() < PROP && i < MAX_LEVEL {
		i++
	}
	return i
}

func (c *list) insert(key int) {
	var update [MAX_LEVEL]*node
	x := c.header
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		// x at the rightmost element that is smaller than `key`
		for x.forward[i] != nil && x.forward[i].key < key { // here it is a `for`,not `if`
			x = x.forward[i]
		}
		update[i] = x // has x.forward[i].key >= key
	}

	lvl := randomLevel()
	nd := newNode()
	nd.key = key
	for i := 0; i < lvl; i++ {
		nd.forward[i] = update[i].forward[i]
		update[i].forward[i] = nd
	}
}
func (c *list) delete(key int) bool {
	return false
}
func (c *list) find(key int) *node {
	x := c.header
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		// traverse to the last element smaller than key
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
	}
	if x.forward[0] != nil && x.forward[0].key == key {
		return x
	}
	return nil
}
func (c *list) String() string {
	lvls := make([]string, MAX_LEVEL)
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		var layer []string
		for x := c.header; x.forward[i] != nil; x = x.forward[i] {
			layer = append(layer, fmt.Sprintf("%s", x.forward[i]))
		}
		lvls[MAX_LEVEL-1-i] = fmt.Sprintf("[%2d]%s", i+1, strings.Join(layer, "   "))
	}

	return strings.Join(lvls, "\n")
}
