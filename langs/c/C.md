# C语言

C 语言诞生于 1972 年

C 语言有哪些特性呢？

1. C 语言是一个静态弱类型语言，在使用变量时需要声明变量类型，但是类型间可以有隐式转换；
2. 不同的变量类型可以用结构体（struct）组合在一起，以此来声明新的数据类型；
3. C 语言可以用 typedef 关键字来定义类型的别名，以此来达到变量类型的抽象；
4. C 语言是一个有结构化程序设计、具有变量作用域以及递归功能的过程式语言；
5. C 语言传递参数一般是以值传递，也可以传递指针；
6. 通过指针，C 语言可以容易地对内存进行低级控制，然而这加大了编程复杂度；
7. 编译预处理让 C 语言的编译更具有弹性，比如跨平台。

C 语言的这些特性，可以让程序员在微观层面写出非常精细和精确的编程操作，让程序员可以在底层和系统细节上非常自由、灵活和精准地控制代码。



## 值得探究的C泛型代码

### 一个泛型的示例 - swap 函数

```c
void swap(void* x, void* y, size_t size)
{
  char tmp[size];
  memcpy(tmp, y, size);
  memcpy(y, x, size);
  memcpy(x, tmp, size);
}
```

- 函数接口中增加了一个size参数。
- 函数的实现中使用了memcpy()函数。
- 函数的实现中使用了一个temp[size]数组。

带来的问题：

1. 新增的size参数，使用的memcpy内存拷贝以及一个 buffer，这增加了编程的复杂度。这就是 C 语言的类型抽象所带来的复杂度的提升。
2. 想交换两个字符串数组，类型是char*，那么，我的swap()函数的x和y参数是不是要用void**了？这样一来，接口就没法定义了。



### 宏定义的泛型：

```c
#define swap(x, y, size) {\
  char temp[size]; \
  memcpy(temp, &y, size); \
  memcpy(&y,   &x, size); \
  memcpy(&x, temp, size); \
}

#define swap(x, y, size) { \
	char temp[size]; \
	memcpy(temp, &y, size); \
	memcpy(&y, &x, size); \
	memcpy(&x, temp, size);
}
```

但用宏带来的问题就是编译器做字符串替换，因为宏是做字符串替换，所以会导致代码膨胀，导致编译出的执行文件比较大。



### min和max的宏替换

```c
#define min(x, y) ((x) > (y) ? (y): (x))
```

其中一个最大的问题，就是有可能会有重复执行的问题。如：

1. min(i++, j++)
2. min(foo(), bar())



### C语言版search

```c
int search(void* a, size_t size, void* target, 
  size_t elem_size, int(*cmpFn)(void*, void*) )
{
  for(int i=0; i<size; i++) {
    if ( cmpFn (a + elem_size * i, target) == 0 ) {
      return i;
    }
  }
  return -1;
}
```



## C语言特点

C 语言的伟大之处在于——使用 C 语言的程序员在高级语言的特性之上还能简单地做任何底层上的微观控制。

C 语言是高级语言中的汇编语言。

编程语言的发展方向，C语言本来就是开发unix系统的语言，处理业务非其所长。

在编程这个世界中，更多的编程工作是解决业务上的问题，而不是计算机的问题，所以，我们需要更为贴近业务、更为抽象的语言。



## 参考资料

- [30 | 编程范式游记（1）- 起源](https://time.geekbang.org/column/article/301)