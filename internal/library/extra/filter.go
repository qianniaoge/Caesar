package extra

import (
	"Caesar/internal/relation"
	"Caesar/pkg/utils"
)

func GetFilterPath(paths []relation.TagPath, mvc bool, Top int) []relation.TagPath {
	var newPaths []relation.TagPath

	if mvc {
		// 如果是mvc模式的网站，则排除类似.jsp, .php, .zip的后缀路径
		for _, v := range paths {
			if matchedDir := utils.MatchDir(v.Path); matchedDir {
				newPaths = append(newPaths, v)

			}
		}
	} else {
		newPaths = paths
	}

	if Top != 0 {
		return newPaths[0:Top]
	} else {
		return newPaths
	}

}
