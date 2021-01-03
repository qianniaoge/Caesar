package audit

import (
	"strings"
	"time"

	"Caesar/internal/library"
	"Caesar/internal/library/cores"
	"Caesar/internal/relation"
	"Caesar/pkg/record"
	"Caesar/pkg/utils"
)

func start(threads int, delay int) {
	// 开始运行任务
	// 设置并发数
	// 设置延时
	relation.Engine.TimeSleep = delay

	// 如果设置延时，线程自动降为1
	if delay == 0 {
		relation.Engine.Threads = threads
	} else {
		relation.Engine.Threads = 1
	}
}

func StartSensitiveFoundFromAddress(target, flag string, threads, delay int) {
	start(threads, delay)

	var targetList []string
	var flagList []string

	targetList = library.GetTargets(target)
	flagList = strings.Split(flag, ",")

	startTime := time.Now()

	record.Logger.Debug("The project start at " + startTime.Format("2006-01-02 15:04:05"))

	for _, v := range targetList {
		if !relation.Engine.Silence {
			println()
		}
		targetTime := time.Now()
		record.Logger.Debug("The target " + v + " start at " + targetTime.Format("2006-01-02 15:04:05"))
		// 主进程
		cores.Start(v, true, cores.ReadDict(flagList, relation.Paths.Dict))
		record.Logger.Debug("The target " + v + " end at " + time.Now().Format("2006-01-02 15:04:05"))
		record.Logger.Debug("The target " + v + " cost time is  " + time.Since(targetTime).String())

		if !relation.Engine.Silence {
			println()
		}

	}

	// 导出结果
	cores.Export(relation.Engine.CollectAssets)

	record.Logger.Debug("The project end at " + time.Now().Format("2006-01-02 15:04:05"))
	record.Logger.Debug("All time is  " + time.Since(startTime).String())

	return
}

func StartSensitiveFoundFromText(text, flag string, threads, delay int) {
	start(threads, delay)

	startTime := time.Now()
	record.Logger.Debug("The project start at " + startTime.Format("2006-01-02 15:04:05"))

	if !relation.Engine.Silence {
		println()
	}

	info := utils.ReadFile(text)
	if len(utils.ReadFile(text)) == 0 {
		println("Open File Errors " + text)
		return
	}

	cores.Start(info, false, cores.ReadDict(strings.Split(flag, ","), relation.Paths.Dict))

	// 导出结果
	cores.Export(relation.Engine.CollectAssets)

	record.Logger.Debug("The project end at " + time.Now().Format("2006-01-02 15:04:05"))
	record.Logger.Debug("All time is  " + time.Since(startTime).String())

	return
}
