//go:build examples
// +build examples

// MIT License
//
// Copyright (c) 2024 Aidenwork
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"github.com/aidencompsci/ento"
)

type Component1 struct{ value int }
type Component2 struct{ value int }
type Component3 struct{ value int }

// Create a System which will operate on queries
// Use only Public fields for queries

type System struct {
	// Construct the first query
	Query1 ento.Query[struct {
		// Query components must be embedded pointers and Public
		*Component1 `ento:"required"`
	}]

	// Second query
	Query2 ento.Query[struct {
		*Component3 `ento:"required"`
	}]

	// unused by the ECS, make non-query fields as desired
	calls int
}

func (self *System) Update() {
	// Iterate over the entities in the queries using go's `range` and `iter`
	for e, v := range self.Query1.Iter() {
		// .. do work ..
		for e2, v2 := range self.Query2.Iter() {
			// .. do work ..
			for e3, v3 := range self.Query1.Iter() {
				// .. do work ..
			}
		}
	}

}

func main() {
	// Create the world and register components
	world := ento.NewWorldBuilder().
		// Use "zero-values" as values are ignored when registering
		WithSparseComponents(Component1{}, Component2{}, Component3{}).
		// Pre-allocate space for 256 entities (world can grow beyond that automatically)
		Build(256)

	// Add entities
	world.AddEntity(Component1{1})
	world.AddEntity(Component1{2}, Component2{3})
	world.AddEntity(Component3{4})

	// Add system
	world.AddSystems(&System{})

	// Update all systems in the world a single time
	world.Update()
}
