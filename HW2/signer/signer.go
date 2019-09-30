package main

import (
	"strconv"
	"strings"
	"sort"
	"sync"
)


const (
	maxElems = 100
)

func ExecutePipeline(jobs... job) {
	prevOut := make(chan interface{}, maxElems)

	for _, j := range jobs {
		out := make(chan interface{}, maxElems)

		go func(j job, in, out chan interface{}){
			j(in, out)
			close(out)
		}(j, prevOut, out)

		prevOut = out
	}

	<-prevOut
}

var mu sync.Mutex

func evalMd5(data string) string {
	mu.Lock()
	defer mu.Unlock()

	return DataSignerMd5(data)
}

func oneSingleHash(data string) string {
	var wg sync.WaitGroup
	var left string
	var right string

	wg.Add(2)

	go func(wg *sync.WaitGroup){
		left = DataSignerCrc32(data)
		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup){
		right = DataSignerCrc32(evalMd5(data))
		wg.Done()
	}(&wg)

	wg.Wait()
	return left + "~" + right
}

func SingleHash(in, out chan interface{}) {
	var wg sync.WaitGroup

	for data := range in {
		d := strconv.Itoa(data.(int))

		wg.Add(1)

		go func(wg *sync.WaitGroup){
			out <- oneSingleHash(d)
			wg.Done()
		}(&wg)
	}

	wg.Wait()
}

func oneMultiHash(src string) string {
	var wg sync.WaitGroup

	crcs := make([]string, 6)

	for i := 0; i < 6; i++ {
		wg.Add(1)

		go func(crcIdx int, wg *sync.WaitGroup){
			hashSrc := strconv.Itoa(crcIdx) + src
			crcs[crcIdx] = DataSignerCrc32(hashSrc)
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()
	res := strings.Join(crcs, "")

	return res
}

func MultiHash(in, out chan interface{}) {
	var wg sync.WaitGroup

	i := 0
	for data := range in {
		i++
		d := data.(string)
		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			out <- oneMultiHash(d)
			wg.Done()
		}(&wg)
	}

	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	results := make([]string, 0)

	for data := range in {
		d := data.(string)
		results = append(results, d)
	}

	sort.Strings(results)
	result := strings.Join(results, "_")

	out <- result
}
