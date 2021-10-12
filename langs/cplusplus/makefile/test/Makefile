# 1、wildcard : 扩展通配符
# 2、notdir ： 去除路径
# 3、patsubst ：替换通配符
# https://blog.csdn.net/srw11/article/details/7516712

src=$(wildcard *.c ./sub/*.c)
dir=$(notdir $(src))
obj=$(patsubst %.c,%.o,$(dir))

all:
	@echo src: $(src)
	@echo dir: $(dir)
	@echo obj: $(obj)
