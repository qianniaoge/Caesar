package library

import (
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"Caesar/internal/relation"

	"github.com/cheggaaa/pb/v3"
)

func HeartProgress(wg *sync.WaitGroup, finished chan struct{}, msg string) {
	/*
		心跳函数，主要用来打印进度条
	*/
	wg.Add(1)
	defer wg.Done()

	total := cap(finished)

	tmpl := `{{ red "'` + msg + `'" }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{speed . | rndcolor }} {{percent .}} {{string . "my_green_string" | green}} {{string . "my_blue_string" | blue}}`
	//start bar based on our template
	bar := pb.ProgressBarTemplate(tmpl).Start64(int64(total))
	// set values for string elements
	bar.Set("my_green_string", "count")
	//	Set("my_blue_string", "blue")

	if relation.Engine.Silence {
		// file, err := os.Create(filepath.Join(relation.Paths.Result, "progress.log"))
		file, err := os.OpenFile(filepath.Join(relation.Paths.Result, "progress.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			bar.SetWriter(nil)
		} else {
			bar.SetWriter(file)
		}

		defer func() { _ = file.Close() }()

	}

	for true {

		if relation.Engine.StopFlag {
			bar.Set("my_green_string", strconv.Itoa(total)+"/"+strconv.Itoa(total))
			bar.SetCurrent(int64(total))
			bar.Finish()

			close(finished)
			break

		}

		if len(finished) == total {
			bar.Set("my_green_string", strconv.Itoa(total)+"/"+strconv.Itoa(total))
			bar.SetCurrent(int64(total))
			bar.Finish()
			close(finished)
			break
		}
		bar.Set("my_green_string", strconv.Itoa(len(finished))+"/"+strconv.Itoa(total))
		bar.SetCurrent(int64(len(finished)))
		time.Sleep(400 * time.Millisecond)

	}

}
