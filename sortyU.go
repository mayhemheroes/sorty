package sorty

// Concurrent Sorting
// Author: Serhat Şevki Dinçer, jfcgaussATgmail

import "sync/atomic"

// uint array to be sorted
var arU []uint

// IsSortedU checks if ar is sorted in ascending order.
func IsSortedU(ar []uint) bool {
	for i := len(ar) - 1; i > 0; i-- {
		if ar[i] < ar[i-1] {
			return false
		}
	}
	return true
}

func forSortU(ar []uint) {
	for h := len(ar) - 1; h > 0; h-- {
		for l := h - 1; l >= 0; l-- {
			if ar[h] < ar[l] {
				ar[l], ar[h] = ar[h], ar[l]
			}
		}
	}
}

// given vl <= vh, inserts pv in the middle
// returns vl <= pv <= vh
func ipU(pv, vl, vh uint) (a, b, c uint, r int) {
	if pv > vh {
		vh, pv = pv, vh
		r = 1
	} else if pv < vl {
		vl, pv = pv, vl
		r = -1
	}
	return vl, pv, vh, r
}

// return pivot as median of five scattered values
func medianU(l, h int) uint {
	// lo, med, hi
	m := mean(l, h)
	vl, pv, vh := arU[l], arU[m], arU[h]

	// intermediates
	a, b := mean(l, m), mean(m, h)
	va, vb := arU[a], arU[b]

	// put lo, med, hi in order
	if vh < vl {
		vl, vh = vh, vl
	}
	vl, pv, vh, _ = ipU(pv, vl, vh)

	// update pivot with intermediates
	if vb < va {
		va, vb = vb, va
	}
	va, pv, vb, r := ipU(pv, va, vb)

	// if pivot was out of [va, vb]
	if r == 1 {
		vl, va, pv, _ = ipU(vl, va, pv)
	} else if r == -1 {
		pv, vb, vh, _ = ipU(vh, pv, vb)
	}

	// here: vl <= va <= pv <= vb <= vh
	arU[l], arU[m], arU[h] = vl, pv, vh
	arU[a], arU[b] = va, vb
	return pv
}

var ngU, mxU uint32 // number of sorting goroutines, max limit
var doneU = make(chan bool, 1)

// SortU concurrently sorts ar in ascending order. Should not be called by multiple goroutines at the same time.
// mx is the maximum number of goroutines used for sorting simultaneously, saturated to [2, 65535].
func SortU(ar []uint, mx uint32) {
	if len(ar) < S {
		forSortU(ar)
		return
	}

	mxU = sat(mx)
	arU = ar

	ngU = 1 // count self
	gsrtU(0, len(arU)-1)
	<-doneU

	arU = nil
}

func gsrtU(lo, hi int) {
	srtU(lo, hi)

	if atomic.AddUint32(&ngU, ^uint32(0)) == 0 { // decrease goroutine counter
		doneU <- false // we are the last, all done
	}
}

// assumes hi-lo >= S-1
func srtU(lo, hi int) {
	var l, h int
start:
	l, h = lo+1, hi-1 // medianU handles lo,hi positions

	for pv := medianU(lo, hi); l <= h; {
		swap := true
		if arU[h] >= pv { // extend ranges in balance
			h--
			swap = false
		}
		if arU[l] <= pv {
			l++
			swap = false
		}

		if swap {
			arU[l], arU[h] = arU[h], arU[l]
			h--
			l++
		}
	}

	if h-lo < hi-l {
		h, hi = hi, h // [lo,h] is the bigger range
		l, lo = lo, l
	}

	if hi-l >= S-1 { // two big ranges?

		if ngU >= mxU { // max number of goroutines? not atomic but good enough
			srtU(l, hi) // start a recursive (slave) sort on the smaller range
			hi = h
			goto start
		}

		atomic.AddUint32(&ngU, 1) // increase goroutine counter
		go gsrtU(lo, h)           // start a goroutine on the bigger range
		lo = l
		goto start
	}

	forSortU(arU[l : hi+1])

	if h-lo < S-1 { // two small ranges?
		forSortU(arU[lo : h+1])
		return
	}

	hi = h
	goto start
}
