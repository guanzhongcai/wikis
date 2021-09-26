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





## 参考资料 

- [【技术】手把手教你写CMake一条龙教程——421施公队Clang出品](https://www.bilibili.com/video/BV16V411k7eF?from=search&seid=12475052695208406196&spm_id_from=333.337.0.0)

