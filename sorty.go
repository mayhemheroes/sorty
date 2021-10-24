/*	Copyright (c) 2019, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

// Package sorty is a type-specific, fast, efficient, concurrent / parallel QuickSort
// implementation. It is in-place and does not require extra memory. You can call:
//  sorty.SortSlice(native_slice) // []int32, []float64 etc. in ascending order
//  sorty.SortLen(len_slice)      // []string or [][]T 'by length' in ascending order
//  sorty.Sort(n, lesswap)        // lesswap() based
package sorty

// MaxGor is the maximum concurrent goroutines (including caller) used for sorting
// per Sort*() call. MaxGor can be changed live, even during an ongoing Sort*() call.
// MaxGor=1 (or a short input) yields single-goroutine sorting: no goroutines or
// channel will be created by sorty.
var MaxGor uint32 = 3

func init() {
	if !(4097 > MaxGor && MaxGor > 0 && MaxLenRec > 2*MaxLenIns &&
		MaxLenIns > MaxLenInsFC && MaxLenInsFC > 9) {
		panic("sorty: check your MaxGor/MaxLen* values")
	}
}

// mid-point
func mid(l, h int) int {
	return int(uint(l+h) >> 1)
}

// Search returns lowest integer k in [0,n) where fn(k) is true, assuming:
//  fn(k) => fn(k+1)
// If there is no such k, it returns n. It can be used to locate an element
// in a sorted collection.
func Search(n int, fn func(int) bool) int {
	l, h := 0, n

	for l < h {
		m := mid(l, h)

		if fn(m) {
			h = m
		} else {
			l = m + 1
		}
	}
	return l
}

// synchronization variables for [g]long*()
type syncVar struct {
	ngr  uint32   // number of sorting goroutines
	done chan int // end signal
}
