//go:build tuneparam
// +build tuneparam

/*	Copyright (c) 2019, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package sorty

// MaxLenIns is the default maximum slice length for insertion sort.
var MaxLenIns = 100

// MaxLenInsFC is the maximum slice length for insertion sort when
// sorting strings or calling [Sort]().
var MaxLenInsFC = 40

// MaxLenRec is the maximum slice length for recursion when there is goroutine
// quota. So MaxLenRec+1 is the minimum slice length for new sorting goroutines.
var MaxLenRec = 496
