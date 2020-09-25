package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type TF10 struct {
	V1  int64
	V2  int64
	V3  int64
	V4  int64
	V5  int64
	V6  int64
	V7  int64
	V8  int64
	V9  int64
	V10 int64
}

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
	const m40 = 1000000 * 10

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

	vsf10 := make([]TF10, m40)
	for idx := range vsf10 {
		vsf10[idx].V1 = rand.Int63n(rnd)
		vsf10[idx].V2 = rand.Int63n(rnd)
		vsf10[idx].V3 = rand.Int63n(rnd)
		vsf10[idx].V4 = rand.Int63n(rnd)
		vsf10[idx].V5 = rand.Int63n(rnd)
		vsf10[idx].V6 = rand.Int63n(rnd)
		vsf10[idx].V7 = rand.Int63n(rnd)
		vsf10[idx].V8 = rand.Int63n(rnd)
		vsf10[idx].V9 = rand.Int63n(rnd)
		vsf10[idx].V10 = rand.Int63n(rnd)
	}

	vspf10 := make([]*TF10, m40)
	for idx := range vspf10 {
		vspf10[idx] = &TF10{
			V1:  vsf10[idx].V1,
			V2:  vsf10[idx].V2,
			V3:  vsf10[idx].V3,
			V4:  vsf10[idx].V4,
			V5:  vsf10[idx].V5,
			V6:  vsf10[idx].V6,
			V7:  vsf10[idx].V7,
			V8:  vsf10[idx].V8,
			V9:  vsf10[idx].V9,
			V10: vsf10[idx].V10,
		}
	}

	res, diff := countValuesF10(t, vsf10)
	fmt.Printf("vf10 min %v res %v\n", diff, res)

	res2, diff := countPtrsF10(t, vspf10)
	fmt.Printf("pf10 min %v res %v\n", diff, res)

	require.Equal(t, res, res2)

	res, diff = countValuesF4(t, vsf4)
	fmt.Printf("vf4 min %v res %v\n", diff, res)

	res2, diff = countPtrsF4(t, vspf4)
	fmt.Printf("pf4 min %v res %v\n", diff, res)

	require.Equal(t, res, res2)

	res, diff = countValuesF3(t, vsf3)
	fmt.Printf("vf3 min %v res %v\n", diff, res)

	res2, diff = countPtrsF3(t, vspf3)
	fmt.Printf("pf3 min %v res %v\n", diff, res)

	require.Equal(t, res, res2)

	res, diff = countValuesF2(t, vsf2)
	fmt.Printf("vf2 min %v res %v\n", diff, res)

	res2, diff = countPtrsF2(t, vspf2)
	fmt.Printf("pf2 min %v res %v\n", diff, res)

	require.Equal(t, res, res2)

	res, diff = countValuesF1(t, vsf1)
	fmt.Printf("vf1 min %v res %v\n", diff, res)

	res2, diff = countPtrsF1(t, vspf1)
	fmt.Printf("pf1 min %v res %v\n", diff, res)

	require.Equal(t, res, res2)
}

func countValuesF10(t *testing.T, vs []TF10) (int, time.Duration) {

	var res int
	start := time.Now()

	for idx := 0; idx < len(vs); idx++ {
		res += countBiggerThenValuesF10(vs[idx], 100)
	}

	diff := time.Since(start)
	return res, diff
}

func countValuesF4(t *testing.T, vs []TF4) (int, time.Duration) {

	var res int
	start := time.Now()

	for idx := 0; idx < len(vs); idx++ {
		res += countBiggerThenValuesF4(vs[idx], 100)
	}

	diff := time.Since(start)
	return res, diff
}

func countValuesF3(t *testing.T, vs []TF3) (int, time.Duration) {

	var res int
	start := time.Now()

	for idx := 0; idx < len(vs); idx++ {
		res += countBiggerThenValuesF3(vs[idx], 100)
	}

	diff := time.Since(start)
	return res, diff
}

func countValuesF2(t *testing.T, vs []TF2) (int, time.Duration) {

	var res int
	start := time.Now()

	for idx := 0; idx < len(vs); idx++ {
		res += countBiggerThenValuesF2(vs[idx], 100)
	}

	diff := time.Since(start)
	return res, diff
}

func countValuesF1(t *testing.T, vs []TF1) (int, time.Duration) {

	var res int
	start := time.Now()

	for idx := 0; idx < len(vs); idx++ {
		res += countBiggerThenValuesF1(vs[idx], 100)
	}

	diff := time.Since(start)
	return res, diff
}

func countPtrsF10(t *testing.T, vs []*TF10) (int, time.Duration) {

	var res int
	start := time.Now()

	for idx := 0; idx < len(vs); idx++ {
		res += countBiggerThenPtrF10(vs[idx], 100)
	}

	diff := time.Since(start)
	return res, diff
}

func countPtrsF4(t *testing.T, vs []*TF4) (int, time.Duration) {

	var res int
	start := time.Now()

	for idx := 0; idx < len(vs); idx++ {
		res += countBiggerThenPtrF4(vs[idx], 100)
	}

	diff := time.Since(start)
	return res, diff
}

func countPtrsF3(t *testing.T, vs []*TF3) (int, time.Duration) {

	var res int
	start := time.Now()

	for idx := 0; idx < len(vs); idx++ {
		res += countBiggerThenPtrF3(vs[idx], 100)
	}

	diff := time.Since(start)
	return res, diff
}

func countPtrsF2(t *testing.T, vs []*TF2) (int, time.Duration) {

	var res int
	start := time.Now()

	for idx := 0; idx < len(vs); idx++ {
		res += countBiggerThenPtrF2(vs[idx], 100)
	}

	diff := time.Since(start)
	return res, diff
}

func countPtrsF1(t *testing.T, vs []*TF1) (int, time.Duration) {

	var res int
	start := time.Now()

	for idx := 0; idx < len(vs); idx++ {
		res += countBiggerThenPtrF1(vs[idx], 100)
	}

	diff := time.Since(start)
	return res, diff
}

func countBiggerThenValuesF10(v TF10, biggerThen int64) int {

	var count int
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
	if v.V5 > biggerThen {
		count++
	}
	if v.V6 > biggerThen {
		count++
	}
	if v.V7 > biggerThen {

		count++
	}
	if v.V8 > biggerThen {
		count++
	}
	if v.V9 > biggerThen {
		count++
	}
	if v.V10 > biggerThen {
		count++
	}
	return count
}

func countBiggerThenValuesF4(v TF4, biggerThen int64) int {

	var count int
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
	return count
}

func countBiggerThenValuesF3(v TF3, biggerThen int64) int {

	var count int
	if v.V1 > biggerThen {
		count++
	}
	if v.V2 > biggerThen {
		count++
	}
	if v.V3 > biggerThen {
		count++
	}
	return count
}

func countBiggerThenValuesF2(v TF2, biggerThen int64) int {

	var count int
	if v.V1 > biggerThen {
		count++
	}
	if v.V2 > biggerThen {
		count++
	}
	return count
}

func countBiggerThenValuesF1(v TF1, biggerThen int64) int {

	var count int
	if v.V1 > biggerThen {
		count++
	}
	return count
}

func countBiggerThenPtrF10(v *TF10, biggerThen int64) int {

	var count int
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
	if v.V5 > biggerThen {
		count++
	}
	if v.V6 > biggerThen {
		count++
	}
	if v.V7 > biggerThen {
		count++
	}
	if v.V8 > biggerThen {
		count++
	}
	if v.V9 > biggerThen {
		count++
	}
	if v.V10 > biggerThen {
		count++
	}
	return count
}

func countBiggerThenPtrF4(v *TF4, biggerThen int64) int {

	var count int
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
	return count
}

func countBiggerThenPtrF3(v *TF3, biggerThen int64) int {

	var count int
	if v.V1 > biggerThen {
		count++
	}
	if v.V2 > biggerThen {
		count++
	}
	if v.V3 > biggerThen {
		count++
	}
	return count
}

func countBiggerThenPtrF2(v *TF2, biggerThen int64) int {

	var count int
	if v.V1 > biggerThen {
		count++
	}
	if v.V2 > biggerThen {
		count++
	}
	return count
}

func countBiggerThenPtrF1(v *TF1, biggerThen int64) int {

	var count int
	if v.V1 > biggerThen {
		count++
	}
	return count
}
