# Makefile

.PHONY: cclean

# 源文件目录、头文件目录和后缀名
SRC_DIR := ./c_src
INC_DIR := ./c_header
SRC_EXT := .c
HDR_EXT := .h

# 获取源文件列表和头文件列表
SRCS := $(wildcard $(SRC_DIR)/*$(SRC_EXT))
HDRS := $(wildcard $(INC_DIR)/*$(HDR_EXT))

# 清理规则：删除除源文件和头文件之外的所有编译产物
cclean:
	find $(SRC_DIR) ! -name "*$(SRC_EXT)" ! -name "*$(HDR_EXT)" -type f -delete
