package waf

import (
	"sync"
	"time"

	"Caesar/internal/library"
	"Caesar/internal/library/cores"
	"Caesar/internal/library/director"
	"Caesar/pkg/record"
)

func StartWafCheck(target string, threads, delay int) {

	targetList := library.GetTargets(target)
	wg := &sync.WaitGroup{}

	var ThreadsChan chan struct{}
	var finishChan chan string

	// 设置线程

	if delay != 0 {
		threads = 1

	} else {
		if threads > len(targetList) {
			threads = len(targetList)
		}
	}

	ThreadsChan = make(chan struct{}, threads)
	finishChan = make(chan string, len(targetList))

	startTime := time.Now()

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		// 心跳线程 用来关闭数据chan
		defer wg.Done()

		for {
			if len(finishChan) == len(targetList) {
				close(finishChan)
				return
			}
			time.Sleep(time.Second * 1)
		}

	}(wg)

	record.Logger.Debug("The project start at " + startTime.Format("2006-01-02 15:04:05"))

	for _, v := range targetList {

		ThreadsChan <- struct{}{}
		wg.Add(1)

		// 主进程
		go func(address string) {
			var msg string

			// 检查目标连通性
			if _, _, _, err := director.GenerateGet(address, true); err != nil {
				msg = address + " not connect"
			} else {
				if cores.CheckWaf(address) {
					msg = address + " has waf"
				} else {
					msg = address + " no waf"
				}
			}

			finishChan <- msg

			wg.Done()
			<-ThreadsChan
		}(v)

	}

	wg.Wait()

	for v := range finishChan {
		println(v)
	}

	record.Logger.Debug("The task  end at " + time.Now().Format("2006-01-02 15:04:05"))
	record.Logger.Debug("The task  cost time is  " + time.Since(startTime).String())

}
