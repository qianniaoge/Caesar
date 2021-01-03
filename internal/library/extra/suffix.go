package extra

import (
	"strings"

	"Caesar/internal/relation"
	"Caesar/pkg/utils"
)

func CheckSuffix(paths []string) []string {
	var newPathSlice []string
	for _, v := range paths {
		// 判断是否是动态文件
		matchedFIle := utils.MatchDynamic(v)
		if matchedFIle {
			// index.php~
			//each := v + "~"
			each := v + "~"
			newPathSlice = append(newPathSlice, each)

			for _, k := range relation.Engine.DirectoryDirSuffix {
				for _, m := range relation.Engine.SuffixSymbol {
					newPath := v + m + k
					newPathSlice = append(newPathSlice, newPath)

				}

			}

		}

		// 判断是否是目录
		//matchedDir, _ := regexp.MatchString(`\..{3,4}$`, v)
		matchedDir := utils.MatchDir(v)
		if matchedDir {
			for _, l := range relation.Engine.DynamicFileSuffix {

				// 下面的逻辑if判断是为了解决:
				// http://127.0.0.1/admin -> 200
				// http://127.0.0.1/admin/ -> 404
				if strings.HasSuffix(v, "/") {
					dirPath := strings.TrimSuffix(v, "/")
					newPathSlice = append(newPathSlice, dirPath)

				} else {
					dirPath := v + "/"
					newPathSlice = append(newPathSlice, dirPath)
				}

				newPath := strings.TrimSuffix(v, "/") + "." + l
				newPathSlice = append(newPathSlice, newPath)
			}

		}

	}

	return utils.RemoveDuplicateElement(newPathSlice)

}
