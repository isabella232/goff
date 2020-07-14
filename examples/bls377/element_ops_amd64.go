// Copyright 2020 ConsenSys AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by goff (v0.3.0) DO NOT EDIT

// Package bls377 contains field arithmetic operations
package bls377

// -------------------------------------------------------------------------------------------------
// Declarations

//go:noescape
func reduceElement(res *Element)

//go:noescape
func addElement(res, x, y *Element)

//go:noescape
func doubleElement(res, x *Element)

//go:noescape
func _fromMontADXElement(res *Element)

//go:noescape
func _mulADXElement(res, x, y *Element)

//go:noescape
func _squareADXElement(res, x *Element)

//go:noescape
func subElement(res, x, y *Element)

// -------------------------------------------------------------------------------------------------
// APIs

// Add z = x + y mod q
func (z *Element) Add(x, y *Element) *Element {
	addElement(z, x, y)
	return z
}

// Double z = x + x mod q, aka Lsh 1
func (z *Element) Double(x *Element) *Element {
	doubleElement(z, x)
	return z
}

// FromMont converts z in place (i.e. mutates) from Montgomery to regular representation
// sets and returns z = z * 1
func (z *Element) FromMont() *Element {
	_fromMontADXElement(z)
	return z
}

// Sub  z = x - y mod q
func (z *Element) Sub(x, y *Element) *Element {
	subElement(z, x, y)
	return z
}

// Mul z = x * y mod q
// see https://hackmd.io/@zkteam/modular_multiplication
func (z *Element) Mul(x, y *Element) *Element {
	_mulADXElement(z, x, y)
	return z
}

// Square z = x * x mod q
// see https://hackmd.io/@zkteam/modular_multiplication
func (z *Element) Square(x *Element) *Element {

	_squareADXElement(z, x)

	return z
}
