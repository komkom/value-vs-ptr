package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type TF4 struct {
	V1 int64
	V2 int64
	V3 int64
	V4 int64
}

type TF3 struct {
	V1 int64
	V2 int64
	V3 int64
}

type TF2 struct {
	V1 int64
	V2 int64
}

type TF1 struct {
	V1 int64
}

func TestValueTypeVsPtr(t *testing.T) {

	rand.Seed(42)

	rnd := int64(10000)
	const m40 = 1000000 * 20 //

	vsf1 := make([]TF1, m40)
	for idx := range vsf1 {
		vsf1[idx].V1 = rand.Int63n(rnd)
	}

	vspf1 := make([]*TF1, m40)
	for idx := range vspf1 {
		vspf1[idx] = &TF1{
			V1: vsf1[idx].V1,
		}
	}

	vsf2 := make([]TF2, m40)
	for idx := range vsf2 {
		vsf2[idx].V1 = rand.Int63n(rnd)
		vsf2[idx].V2 = rand.Int63n(rnd)
	}

	vspf2 := make([]*TF2, m40)
	for idx := range vspf2 {
		vspf2[idx] = &TF2{
			V1: vsf2[idx].V1,
			V2: vsf2[idx].V2,
		}
	}

	vsf3 := make([]TF3, m40)
	for idx := range vsf3 {
		vsf3[idx].V1 = rand.Int63n(rnd)
		vsf3[idx].V2 = rand.Int63n(rnd)
		vsf3[idx].V3 = rand.Int63n(rnd)
	}

	vspf3 := make([]*TF3, m40)
	for idx := range vspf3 {
		vspf3[idx] = &TF3{
			V1: vsf3[idx].V1,
			V2: vsf3[idx].V2,
			V3: vsf3[idx].V3,
		}
	}

	vsf4 := make([]TF4, m40)
	for idx := range vsf4 {
		vsf4[idx].V1 = rand.Int63n(rnd)
		vsf4[idx].V2 = rand.Int63n(rnd)
		vsf4[idx].V3 = rand.Int63n(rnd)
		vsf4[idx].V4 = rand.Int63n(rnd)
	}

	vspf4 := make([]*TF4, m40)
	for idx := range vspf4 {
		vspf4[idx] = &TF4{
			V1: vsf4[idx].V1,
			V2: vsf4[idx].V2,
			V3: vsf4[idx].V3,
			V4: vsf4[idx].V4,
		}
	}

	lastRes := -1
	var min time.Duration
	var idx int
	for i := 0; i < 100; i++ {

		w := 100 * (i + 1)

		res, diff := countValuesF4(t, vsf4, w)
		if min == 0 || diff < min {
			min = diff
			idx = i
		}

		if lastRes != -1 {
			require.Equal(t, lastRes, res)
		}
		lastRes = res
	}
	fmt.Printf("vf4 min %v idx %v\n", min, idx)

	lastRes = -1
	min = 0
	idx = 0
	for i := 0; i < 100; i++ {

		w := 100 * (i + 1)

		res, diff := countPtrsF4(t, vspf4, w)
		if min == 0 || diff < min {
			min = diff
			idx = i
		}

		if lastRes != -1 {
			require.Equal(t, lastRes, res)
		}
		lastRes = res
	}
	fmt.Printf("pf4 min %v idx %v\n", min, idx)

	lastRes = -1
	min = 0
	idx = 0
	for i := 0; i < 100; i++ {

		w := 100 * (i + 1)

		res, diff := countValuesF3(t, vsf3, w)
		if min == 0 || diff < min {
			min = diff
			idx = i
		}

		if lastRes != -1 {
			require.Equal(t, lastRes, res)
		}
		lastRes = res
	}
	fmt.Printf("vf3 min %v idx %v\n", min, idx)

	lastRes = -1
	min = 0
	idx = 0
	for i := 0; i < 100; i++ {

		w := 100 * (i + 1)

		res, diff := countPtrsF3(t, vspf3, w)
		if min == 0 || diff < min {
			min = diff
			idx = i
		}

		if lastRes != -1 {
			require.Equal(t, lastRes, res)
		}
		lastRes = res
	}
	fmt.Printf("pf3 min %v idx %v\n", min, idx)

	lastRes = -1
	min = 0
	idx = 0
	for i := 0; i < 100; i++ {

		w := 100 * (i + 1)

		res, diff := countValuesF2(t, vsf2, w)
		if min == 0 || diff < min {
			min = diff
			idx = i
		}

		if lastRes != -1 {
			require.Equal(t, lastRes, res)
		}
		lastRes = res
	}
	fmt.Printf("vf2 min %v idx %v\n", min, idx)

	lastRes = -1
	min = 0
	idx = 0
	for i := 0; i < 100; i++ {

		w := 100 * (i + 1)

		res, diff := countPtrsF2(t, vspf2, w)
		if min == 0 || diff < min {
			min = diff
			idx = i
		}

		if lastRes != -1 {
			require.Equal(t, lastRes, res)
		}
		lastRes = res
	}
	fmt.Printf("pf2 min %v idx %v\n", min, idx)

	lastRes = -1
	min = 0
	idx = 0
	for i := 0; i < 100; i++ {

		w := 100 * (i + 1)

		res, diff := countValuesF1(t, vsf1, w)
		if min == 0 || diff < min {
			min = diff
			idx = i
		}

		if lastRes != -1 {
			require.Equal(t, lastRes, res)
		}
		lastRes = res
	}
	fmt.Printf("vf1 min %v idx %v\n", min, idx)

	lastRes = -1
	min = 0
	idx = 0
	for i := 0; i < 100; i++ {

		w := 100 * (i + 1)

		res, diff := countPtrsF1(t, vspf1, w)
		if min == 0 || diff < min {
			min = diff
			idx = i
		}

		if lastRes != -1 {
			require.Equal(t, lastRes, res)
		}
		lastRes = res
	}
	fmt.Printf("pf1 min %v idx %v\n", min, idx)

}

func countValuesF4(t *testing.T, vs []TF4, w int) (int, time.Duration) {

	var pos int
	var res int
	var done bool
	var count int

	exp := len(vs) / w

	buf := make([]TF4, w)

	start := time.Now()
	for {
		if done {
			break
		}

		if pos+w > len(vs) {

			done = true
			w = len(vs) - pos
		}

		count++

		copy(buf, vs[pos:pos+w])

		res += countBiggerThenValuesF4(buf[:w], 100)
		pos += w
	}

	diff := time.Since(start)

	require.Equal(t, exp+1, count)
	return res, diff
}

func countValuesF3(t *testing.T, vs []TF3, w int) (int, time.Duration) {

	var pos int
	var res int
	var done bool
	var count int

	exp := len(vs) / w

	buf := make([]TF3, w)

	start := time.Now()
	for {
		if done {
			break
		}

		if pos+w > len(vs) {

			done = true
			w = len(vs) - pos
		}

		count++

		copy(buf, vs[pos:pos+w])

		res += countBiggerThenValuesF3(buf[:w], 100)
		pos += w
	}

	diff := time.Since(start)

	require.Equal(t, exp+1, count)
	return res, diff
}

func countValuesF2(t *testing.T, vs []TF2, w int) (int, time.Duration) {

	var pos int
	var res int
	var done bool
	var count int

	exp := len(vs) / w

	buf := make([]TF2, w)

	start := time.Now()
	for {
		if done {
			break
		}

		if pos+w > len(vs) {

			done = true
			w = len(vs) - pos
		}

		count++

		copy(buf, vs[pos:pos+w])

		res += countBiggerThenValuesF2(buf[:w], 100)
		pos += w
	}

	diff := time.Since(start)

	require.Equal(t, exp+1, count)
	return res, diff
}

func countValuesF1(t *testing.T, vs []TF1, w int) (int, time.Duration) {

	var pos int
	var res int
	var done bool
	var count int

	exp := len(vs) / w

	buf := make([]TF1, w)

	start := time.Now()
	for {
		if done {
			break
		}

		if pos+w > len(vs) {

			done = true
			w = len(vs) - pos
		}

		count++

		copy(buf, vs[pos:pos+w])

		res += countBiggerThenValuesF1(buf[:w], 100)
		pos += w
	}

	diff := time.Since(start)

	require.Equal(t, exp+1, count)
	return res, diff
}

func countPtrsF4(t *testing.T, vs []*TF4, w int) (int, time.Duration) {

	var pos int
	var res int
	var done bool
	var count int

	exp := len(vs) / w

	buf := make([]*TF4, w)

	start := time.Now()
	for {
		if done {
			break
		}

		if pos+w > len(vs) {

			done = true
			w = len(vs) - pos
		}

		count++

		copy(buf, vs[pos:pos+w])

		res += countBiggerThenPtrF4(buf[:w], 100)
		pos += w
	}

	diff := time.Since(start)

	require.Equal(t, exp+1, count)
	return res, diff
}

func countPtrsF3(t *testing.T, vs []*TF3, w int) (int, time.Duration) {

	var pos int
	var res int
	var done bool
	var count int

	exp := len(vs) / w

	buf := make([]*TF3, w)

	start := time.Now()
	for {
		if done {
			break
		}

		if pos+w > len(vs) {

			done = true
			w = len(vs) - pos
		}

		count++

		copy(buf, vs[pos:pos+w])

		res += countBiggerThenPtrF3(buf[:w], 100)
		pos += w
	}

	diff := time.Since(start)

	require.Equal(t, exp+1, count)
	return res, diff
}

func countPtrsF2(t *testing.T, vs []*TF2, w int) (int, time.Duration) {

	var pos int
	var res int
	var done bool
	var count int

	exp := len(vs) / w

	buf := make([]*TF2, w)

	start := time.Now()
	for {
		if done {
			break
		}

		if pos+w > len(vs) {

			done = true
			w = len(vs) - pos
		}

		count++

		copy(buf, vs[pos:pos+w])

		res += countBiggerThenPtrF2(buf[:w], 100)
		pos += w
	}

	diff := time.Since(start)

	require.Equal(t, exp+1, count)
	return res, diff
}

func countPtrsF1(t *testing.T, vs []*TF1, w int) (int, time.Duration) {

	var pos int
	var res int
	var done bool
	var count int

	exp := len(vs) / w

	buf := make([]*TF1, w)

	start := time.Now()
	for {
		if done {
			break
		}

		if pos+w > len(vs) {

			done = true
			w = len(vs) - pos
		}

		count++

		copy(buf, vs[pos:pos+w])

		res += countBiggerThenPtrF1(buf[:w], 100)
		pos += w
	}

	diff := time.Since(start)

	require.Equal(t, exp+1, count)
	return res, diff
}

func countBiggerThenValuesF4(s []TF4, biggerThen int64) int {

	var count int
	for _, v := range s {
		if v.V1 > biggerThen {
			count++
		}
		if v.V2 > biggerThen {
			count++
		}
		if v.V3 > biggerThen {
			count++
		}
		if v.V4 > biggerThen {
			count++
		}
	}
	return count
}

func countBiggerThenValuesF3(s []TF3, biggerThen int64) int {

	var count int
	for _, v := range s {
		if v.V1 > biggerThen {
			count++
		}
		if v.V2 > biggerThen {
			count++
		}
		if v.V3 > biggerThen {
			count++
		}
	}
	return count
}

func countBiggerThenValuesF2(s []TF2, biggerThen int64) int {

	var count int
	for _, v := range s {
		if v.V1 > biggerThen {
			count++
		}
		if v.V2 > biggerThen {
			count++
		}
	}
	return count
}

func countBiggerThenValuesF1(s []TF1, biggerThen int64) int {

	var count int
	for _, v := range s {
		if v.V1 > biggerThen {
			count++
		}
	}
	return count
}

func countBiggerThenPtrF4(s []*TF4, biggerThen int64) int {

	var count int
	for _, v := range s {
		if v.V1 > biggerThen {
			count++
		}
		if v.V2 > biggerThen {
			count++
		}
		if v.V3 > biggerThen {
			count++
		}
		if v.V4 > biggerThen {
			count++
		}
	}
	return count
}

func countBiggerThenPtrF3(s []*TF3, biggerThen int64) int {

	var count int
	for _, v := range s {
		if v.V1 > biggerThen {
			count++
		}
		if v.V2 > biggerThen {
			count++
		}
		if v.V3 > biggerThen {
			count++
		}
	}
	return count
}

func countBiggerThenPtrF2(s []*TF2, biggerThen int64) int {

	var count int
	for _, v := range s {
		if v.V1 > biggerThen {
			count++
		}
		if v.V2 > biggerThen {
			count++
		}
	}
	return count
}

func countBiggerThenPtrF1(s []*TF1, biggerThen int64) int {

	var count int
	for _, v := range s {
		if v.V1 > biggerThen {
			count++
		}
	}
	return count
}
