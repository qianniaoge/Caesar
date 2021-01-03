
DATE= `date '+%Y-%m-%d %H:%M:%S'`
VERSION="v0.0.1"


.PHONY : version
version :
	@echo "caesar" ${VERSION}
	@echo ${DATE}

#构建可执行文件，在deployments目录下面
.PHONY : build
build:
	@echo "build开始"
	@chmod +x ./scripts/build.sh && ./scripts/build.sh
	@echo "build结束"

#生成dockerfile
.PHONY : docker
docker :
	@echo "生成dockerfile开始"
	@chmod +x ./scripts/dockerfile.sh && ./scripts/dockerfile.sh
	@echo "生成Dockerfile结束"