package benchmarks

import (
	"Caesar/internal/library"
	"Caesar/internal/library/cores"
	"Caesar/internal/library/director"
	"Caesar/internal/relation"
	"net/http"
	"sync"
	"testing"
)

func BenchmarkFastHttpRequest(b *testing.B) {

	var tags []map[string]string

	tags = append(tags, map[string]string{"common": "/Users/null/go/src/Caesar/assets/directory/common.json"})

	paths := cores.ReadDict([]string{"common"}, tags)
	ThreadsChan := make(chan struct{}, 12)
	wg := &sync.WaitGroup{}
	length := len(paths)
	var finishChan = make(chan struct{}, length)
	//println(length)
	// 启动进度条goroutine
	go library.HeartProgress(wg, finishChan, "demo")

	b.ResetTimer()
	for _, v := range paths {
		ThreadsChan <- struct{}{}
		wg.Add(1)

		go func(t relation.TagPath) {
			director.FastHttpRequest("http://27.211.65.98:8081"+t.Path, http.MethodGet, "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.16; rv:84.0) Gecko/20100101 Firefox/84.0")
			wg.Done()
			<-ThreadsChan
			finishChan <- struct{}{}

		}(v)

	}

	wg.Wait()

}

func BenchmarkStandRequest(b *testing.B) {

	var tags []map[string]string

	tags = append(tags, map[string]string{"common": "/Users/null/go/src/Caesar/assets/directory/common.json"})

	paths := cores.ReadDict([]string{"common"}, tags)
	ThreadsChan := make(chan struct{}, 12)
	wg := &sync.WaitGroup{}
	length := len(paths)
	var finishChan = make(chan struct{}, length)
	//println(length)
	// 启动进度条goroutine
	go library.HeartProgress(wg, finishChan, "demo")

	b.ResetTimer()
	for _, v := range paths {
		ThreadsChan <- struct{}{}
		wg.Add(1)

		go func(t relation.TagPath) {
			director.GenerateNormalGet("http://27.211.65.98:8081" + t.Path)
			wg.Done()
			<-ThreadsChan
			finishChan <- struct{}{}

		}(v)

	}

	wg.Wait()

}
