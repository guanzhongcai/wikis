# CMake



## gcc

预处理、编译、汇编、链接四个过程：

```bash
# pre-processing
gcc -E main.c -o main.i

# compiling 汇编语言文件的缺省扩展名 .s
gcc -S main.i -o main.s

# assembling 目标文件的缺省扩展名 .o
gcc -c main.s -o main.o

# linking
gcc main.o -o main

$ ./main
```



## Makefile

目录文件：

```bash
$ ls
fun.c fun.h main.c Makefile
```

`Makefile`文件内容：

```makefile
obj=main.o fun.o
main:$(obj)
	gcc -o main $(obj)
%.o:%.c
	gcc -c $<
clean:
	rm *.o
	rm main
```

执行：

```bash
$ make
gcc -c main.c
gcc -c fun.c
gcc -o main main.o fun.o
```



## pkg-config

pkg-config是一个linux下的命令，用于获得某一个库/模块的所有编译相关的信息。

pkg-config 是编译应用程序和库时使用的辅助工具。例如，它可以帮助您在命令行中插入正确的编译器选项，以便应用程序可以使用：

```bash
gcc -o test test.c `pkg-config --libs --cflags glib-2.0`
```

而不是在何处找到 glib（或其他库）的硬编码值。例如，它与语言无关，因此可用于定义文档工具的位置。

### pkg-config的信息从哪里来？

很简单，有2种路径：
第一种：取系统的/usr/lib下的所有*.pc文件。
第二种：PKG_CONFIG_PATH环境变量所指向的路径下的所有*.pc文件。



Mac中的pc位置：

```bash
$ ls /usr/lib/pkgconfig
apr-1.pc        apr-util-1.pc   libiodbc.pc     libpcre.pc      libpcreposix.pc
```

pc文件内容：

```bash
$ more apr-1.pc 
prefix=/usr
exec_prefix=${prefix}
libdir=${exec_prefix}/lib
APR_MAJOR_VERSION=1
includedir=${prefix}/include/apr-${APR_MAJOR_VERSION}

Name: APR
Description: The Apache Portable Runtime library
Version: 1.5.2
Libs: -L${libdir} -lapr-${APR_MAJOR_VERSION} -lpthread
Cflags: -DDARWIN -DSIGPROCMASK_SETS_THREAD_MASK  -I${includedir}
```

一目了然，就是存了所有`apr-1`的头文件/库的路径信息。



## CMake

### synopsis/概要

```bash
# 生成项目构建系统
 cmake [<options>] <path-to-source>
 cmake [<options>] <path-to-existing-build>
 cmake [<options>] -S <path-to-source> -B <path-to-build>

# 构建项目
 cmake --build <dir> [<options>] [-- <build-tool-options>]

# 安装项目
 cmake --install <目录> [<选项>]

# 打开项目
 cmake --open <目录>

# 运行脚本
 cmake [{-D <var>=<value>}...] -P <cmake-script-file>

# 运行命令行工具
 cmake -E <命令> [<选项>]

# 运行查找包工具
 cmake --find-package [<选项>]

# 查看帮助
 cmake --help[-<topic>]
```





## 参考资料 

- [【技术】手把手教你写CMake一条龙教程——421施公队Clang出品](https://www.bilibili.com/video/BV16V411k7eF?from=search&seid=12475052695208406196&spm_id_from=333.337.0.0)

- [https://cmake.org/](https://cmake.org/cmake/help/v3.21/manual/cmake.1.html#build-a-project)

- [pkg-config](https://www.freedesktop.org/wiki/Software/pkg-config/)

- [cmake使用示例与整理总结](https://github.com/carl-wang-cn/demo/tree/master/cmake)

