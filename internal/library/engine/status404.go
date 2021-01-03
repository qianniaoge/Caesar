package engine

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"strconv"
	"sync"
	"time"

	"Caesar/internal/library"
	"Caesar/internal/library/director"
	"Caesar/internal/library/extra"
	"Caesar/internal/relation"
	"Caesar/pkg/record"
	"Caesar/pkg/utils"
)

type target404 struct {
	request  RequestInfo
	response ResponseInfo
	opts     ServerOpt

	application
}

func (t *target404) AlphaFuzz() {

	var wg = &sync.WaitGroup{}
	var mu = &sync.Mutex{}
	var ThreadsChan chan struct{}
	var length = len(t.opts.Paths)

	// 获取原子锁,用来计算连接超时的次数
	counter := extra.NewCounter()

	ThreadsChan = make(chan struct{}, t.opts.Threads)
	finishChan := make(chan struct{}, length)
	threadSlice := library.NewSlice()

	// 启动进度条goroutine
	go library.HeartProgress(wg, finishChan, "Alpha")

	for _, v := range t.opts.Paths {

		if relation.Engine.TimeSleep > 0 {
			time.Sleep(time.Duration(relation.Engine.TimeSleep) * time.Second)
		}

		if !relation.Engine.StopFlag {
			ThreadsChan <- struct{}{}
			wg.Add(1)

			go func(v relation.TagPath) {
				var code int
				var body []byte
				var err error

				targetAddress := t.request.Address + v.Path

				//code, _, body, err := director.GenerateGet(targetAddress, true)

				if len(t.request.Proxy) > 0 || len(t.request.Header) > 0 || len(t.request.Cookies) > 0 || t.request.Timeout > 3 || len(t.request.Body) > 0 {
					code, _, body, err = director.UnitTest(targetAddress, t.request.Method, utils.GetRandomElement(t.request.UserAgent), t.request.Header, t.request.Cookies, t.request.Proxy, t.request.Timeout, t.request.Body)

				} else {
					code, _, body, err = director.FastHttpRequest(targetAddress, t.request.Method, utils.GetRandomElement(t.request.UserAgent))
				}
				// code, _, body, err := director.UnitTest(targetAddress, t.request.Method, utils.GetRandomElement(t.request.UserAgent), t.request.Header, t.request.Cookies, t.request.Proxy, t.request.Timeout, t.request.Body)

				if err != nil {

					//超时时的处理
					counter.AddErr()
					record.Logger.Error(t.request.Address+" ", err.Error())

					if counter.CountErr() >= relation.Engine.TimeoutCount && !relation.Engine.StopFlag {
						mu.Lock()
						relation.Engine.StopFlag = true
						mu.Unlock()
					}

					wg.Done()
					<-ThreadsChan
					if !relation.Engine.StopFlag {
						finishChan <- struct{}{}
					}

					return
				}

				if utils.IntInSlice(code, relation.Engine.StatusCode) {
					var title = ""

					if utils.MatchDynamic(v.Path) {
						title = utils.MatchTitle(string(body))
					}

					record.Logger.Info(targetAddress + " " + strconv.Itoa(code) + " " + title)

					each := relation.ResultPtah{
						Code:    code,
						Address: targetAddress,
						Title:   title,
						Length:  len(body),
					}

					element := relation.StorePath{
						TagPath:    v,
						ResultPtah: each,
					}

					threadSlice.Add(element)
				}

				wg.Done()
				<-ThreadsChan
				if !relation.Engine.StopFlag {
					finishChan <- struct{}{}
				}
				return

			}(v)
		} else {

			break
		}
	}

	wg.Wait()

	if relation.Engine.StopFlag {
		relation.Engine.StopFlag = false
		record.Logger.Error(t.request.Address + " Connect timeout too many, stop")
	}

	t.Store = threadSlice.Get()

	for _, v := range threadSlice.Get() {
		t.Results = append(t.Results, v.ResultPtah)
	}
}

func (t *target404) BetaFuzz() {

	var suffixString []string

	// 用来避免一些不同路径返回相同信息的状况
	var existsMap = map[string]bool{}

	for _, v := range t.Store {
		suffixString = append(suffixString, v.Path)

		firstExists := strconv.Itoa(v.Code) + strconv.Itoa(v.Length)
		//existsList = append(existsList, firstExists)
		if _, ok := existsMap[firstExists]; !ok {
			existsMap[firstExists] = true
		}
	}

	suffixSlice := extra.CheckSuffix(suffixString)

	if !relation.Engine.StopFlag {
		length := len(suffixSlice)
		finishChan := make(chan struct{}, length)
		wg := &sync.WaitGroup{}

		// 启动进度条goroutine
		go library.HeartProgress(wg, finishChan, "Beta")

		for _, v := range suffixSlice {

			if relation.Engine.TimeSleep > 0 {
				time.Sleep(time.Duration(relation.Engine.TimeSleep) * time.Second)
			}

			targetAddress := t.request.Address + v
			//code, _, body, _ := director.GenerateNormalGet(targetAddress)
			//code, _, body, _ := director.GenerateGet(targetAddress, true)
			code, _, body, _ := director.UnitTest(targetAddress, t.request.Method, utils.GetRandomElement(t.request.UserAgent), t.request.Header, t.request.Cookies, t.request.Proxy, t.request.Timeout, t.request.Body)

			if utils.IntInSlice(code, relation.Engine.StatusCode) {
				record.Logger.Info(targetAddress + " " + strconv.Itoa(code) + "  " + strconv.Itoa(len(body)))

				// 避免 http://127.0.0.1/admin和 http://127.0.0.1/admin/返回结果相同的问题
				//exists := strings.TrimSuffix(targetAddress, "/") + strconv.Itoa(code) + strconv.Itoa(len(body))

				firstExists := strconv.Itoa(code) + strconv.Itoa(len(body))

				if _, ok := existsMap[firstExists]; !ok {

					result := relation.ResultPtah{
						Code:    code,
						Address: targetAddress,
						Title:   " ",
						Length:  len(body),
					}

					t.Results = append(t.Results, result)

					existsMap[firstExists] = true

				}

			}

			// 计算进度
			finishChan <- struct{}{}

		}

		wg.Wait()
	}

	//if !relation.Engine.StopFlag {
	//	tmpl := `{{ red "Beat" }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{speed . | rndcolor }} {{percent .}} {{string . "my_green_string" | green}} {{string . "my_blue_string" | blue}}`
	//
	//	bar := pb.ProgressBarTemplate(tmpl).Start64(int64(len(suffixSlice)))
	//
	//	if relation.Engine.Silence {
	//		file, err := os.OpenFile(filepath.Join(relation.Paths.Result, "progress.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	//		if err != nil {
	//			bar.SetWriter(nil)
	//		} else {
	//			bar.SetWriter(file)
	//		}
	//		defer func() { _ = file.Close() }()
	//
	//	}
	//
	//	for _, v := range suffixSlice {
	//
	//		if relation.Engine.TimeSleep > 0 {
	//			time.Sleep(time.Duration(relation.Engine.TimeSleep) * time.Second)
	//		}
	//
	//		targetAddress := t.request.Address + v
	//		//code, _, body, _ := director.GenerateNormalGet(targetAddress)
	//		//code, _, body, _ := director.GenerateGet(targetAddress, true)
	//		code, _, body, _ := director.UnitTest(targetAddress, t.request.Method, utils.GetRandomElement(t.request.UserAgent), t.request.Header, t.request.Cookies, t.request.Proxy, t.request.Timeout, t.request.Body)
	//
	//		if utils.IntInSlice(code, relation.Engine.StatusCode) {
	//			record.Logger.Info(targetAddress + " " + strconv.Itoa(code) + "  " + strconv.Itoa(len(body)))
	//
	//			// 避免 http://127.0.0.1/admin和 http://127.0.0.1/admin/返回结果相同的问题
	//			//exists := strings.TrimSuffix(targetAddress, "/") + strconv.Itoa(code) + strconv.Itoa(len(body))
	//
	//			firstExists := strconv.Itoa(code) + strconv.Itoa(len(body))
	//
	//			if _, ok := existsMap[firstExists]; !ok {
	//
	//				result := relation.ResultPtah{
	//					Code:    code,
	//					Address: targetAddress,
	//					Title:   " ",
	//					Length:  len(body),
	//				}
	//
	//				t.Results = append(t.Results, result)
	//
	//				existsMap[firstExists] = true
	//
	//			}
	//
	//		}
	//
	//		bar.Increment()
	//		bar.Set("my_green_string", strconv.FormatInt(bar.Current(), 10)+"/"+strconv.FormatInt(bar.Total(), 10))
	//
	//	}
	//
	//	bar.Set("my_green_string", strconv.FormatInt(bar.Current(), 10)+"/"+strconv.FormatInt(bar.Total(), 10))
	//	bar.Finish()
	//} else {
	//	return
	//}

}

func (t *target404) Aftermath() {

	// 将每个json数据按照Hits进行排序
	sort.Slice(t.Results,
		func(i, j int) bool {
			return t.Results[i].Length > t.Results[j].Length
		})
	relation.Engine.CollectAssets[t.request.Address] = t.Results

	var results = make(map[string][]string)

	// 进行hits+1操作
	for _, v := range t.Store {
		results[v.Tag] = append(results[v.Tag], v.Path)
	}

	for key, value := range results {
		var mapJson []relation.EachPath
		var newJson []relation.EachPath
		dictPath := utils.GetValueFromList(key, relation.Paths.Dict)
		bytes, err := ioutil.ReadFile(dictPath)
		if err != nil {
			record.Logger.Error(dictPath + " open failed")
		}
		if err1 := json.Unmarshal(bytes, &mapJson); err1 != nil {
			record.Logger.Error("Write json " + dictPath + " failed")
		}

		// 用来给hits加1的的地方
		for _, m := range mapJson {
			if utils.StringInSlice(m.Path, value) {
				m.Hits += 1
				newJson = append(newJson, m)

			} else {
				newJson = append(newJson, m)
			}
		}

		// 序列化数据
		result, errMarshall := utils.CustomMarshal(newJson)

		if errMarshall != nil {
			record.Logger.Error(errMarshall.Error())
			return
		}

		if err := ioutil.WriteFile(dictPath, []byte(result), 0644); err != nil {
			record.Logger.Error("Write file " + dictPath + " error!")
			return
		}
	}

	return
}

func New404(req RequestInfo, resp ResponseInfo, opts ServerOpt) *target404 {

	return &target404{
		request:  req,
		response: resp,
		opts:     opts,

		application: application{
			Store:   []relation.StorePath{},
			Results: []relation.ResultPtah{},
		},
	}

}
