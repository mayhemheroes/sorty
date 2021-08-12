// +build tuneparam

/*	Copyright (c) 2019, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package sorty

var (
	// Mli is the maximum array length for insertion sort in
	// Sort*() except SortS() and Sort().
	Mli = 100
	// Hmli is the maximum array length for insertion sort in SortS() and Sort().
	Hmli = 40

	// Mlr is the maximum array length for recursion when there is available goroutines.
	// So Mlr+1 is the minimum array length for new sorting goroutines.
	Mlr = 496
)