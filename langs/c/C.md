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



## const char*..

**const char \*ptr==char const \*ptr; 可以直接改变指针指向，但不能直接改变指针指向的值；\*ptr=\*ss;**

**char \*const ptr; 可以直接改变指针指向的值，但不能直接改变指针指向;ptr[0]='s';**

```c
int main()
{
    char str[] = "hello world";
    char sec[] = "code world";
 
    const char *ptr1 = str;
    cout << ptr1 << endl;
    strcpy(str,"hi world");
    cout << ptr1 << endl;
    ptr1 = sec;//直接改变指针指向
    cout << ptr1 << endl;
    sec[0] = 'o';
    cout << ptr1 << endl;
    ptr1[0] = 'a';//直接改变指针指向的值,报错
 
 
    char ss[] = "good game";
    char *const ptr2 = ss;
    cout << ptr2 << endl;
    ptr2[0] ='a';//直接改变指针指向的值
    cout << ptr2 << endl;
    strcpy(ptr2, "last");
    cout << ptr2 << endl;
    ss[0] = 'z';
    cout << ptr2 << endl;
    ptr2 = sec;//直接改变指针指向,报错
    system("pause");
}
```



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



## exec函数族

exec函数族提供了一个在进程中启动另一个程序执行的方法。它可以根据指定的文件名或目录名找到可执行文件，并用它来取代原调用进程的**数据段、代码段和堆栈段**，在执行完之后，原调用进程的内容除了进程号外，其他全部被新的进程替换了。

使用exec函数族主要有两种情况：

(1)当进程认为自己不能再为系统和用户做出任何贡献时，就可以调用exec函数族中的任意一个函数让自己重生。

(2)如果一个进程想执行另一个程序，那么它就可以调用fork函数新建一个进程，然后调用exec函数族中的任意一个函数，这样看起来就像通过执行应用程序而产生了一个新进程(这种情况非常普遍)。

```c
// exec函数的原型如下(l=list，v=vector)：
int execl(const char * path，const char * arg，…)；
int execle(const char * path，const char * arg，char * const envp[])；
int execlp(const char * file，const char * arg，…)；
int execv(const char * path，char * const argv[])；
int execve(const char * path，char * const argv[]，char * const envp[])；
int execvp(const char * file，char * const argv[])；
```

**参数说明：**

**`path`**：要执行的程序路径。可以是绝对路径或者是相对路径。在execv、execve、execl和execle这4个函数中，使用带路径名的文件名作为参数。

**`file`**：要执行的程序名称。如果该参数中包含“/”字符，则视为路径名直接执行；否则视为单独的文件名，系统将根据PATH环境变量指定的路径顺序搜索指定的文件。

**`argv`**：命令行参数的矢量数组。

**`envp`**：带有该参数的exec函数可以在调用时指定一个环境变量数组。其他不带该参数的exec函数则使用调用进程的环境变量。

**`arg`**：程序的第0个参数，即程序名自身。相当于argv[O]。

**`…`**：命令行参数列表。调用相应程序时有多少[命令行参数](https://baike.baidu.com/item/命令行参数/3206082)，就需要有多少个输入参数项。注意：在使用此类函数时，在所有命令行参数的最后应该增加一个空的参数项(NULL)，表明命令行参数结束。

[返回值](https://baike.baidu.com/item/返回值/9629649)：一1表明调用exec失败，无返回表明调用成功。 [1] 



## fcntl

在多数`unix`系统中，当多个进程/线程同时编辑一个文件时，该文件的最后状态取决于最后一个写该文件的进程。

文件锁（也叫记录锁）的作用是，当一个进程读写文件的某部分时，其他进程就无法修改同一文件区域。

能够实现文件锁的函数主要有2个：`flock`和`fcntl`。

> [文件锁的使用浅析](https://blog.csdn.net/guotianqing/article/details/80044087)



通过fcntl可以改变已打开的文件性质。

```c
#include <fcntl.h>

// 定义函数 
int fcntl(int fd, int cmd, .../* arg */);

// fcntl()针对(文件)描述符提供控制.参数fd 是被参数cmd操作的描述符.
// 针对cmd的值,fcntl能够接受第三个参数
int fcntl(int fd, int cmd, long arg);

int fcntl(int fd, int cmd, struct flock *lock);
```



fcntl针对描述符提供控制。参数fd是被参数cmd操作的描述符。针对[cmd](https://baike.baidu.com/item/cmd/1193011)的值，fcntl能够接受第三个参数int arg。

fcntl()用来操作[文件描述符](https://baike.baidu.com/item/文件描述符)的一些特性。fcntl 不仅可以施加建议性锁，还可以施加强制锁。同时，fcntl还能对文件的某一记录进行上锁，也就是记录锁。

```c
// F_GETFL 取得文件描述符状态旗标，此旗标为open（）的参数flags。
int flags = fcntl(socket, F_GETFL, 0);

// F_SETFL 设置文件描述符状态旗标，参数arg为新旗标，但只允许O_APPEND、O_NONBLOCK和O_ASYNC位的改变，其他位的改变将不受影响。
/* 设置为非阻塞*/
if (fcntl(socket_descriptor, F_SETFL, flags | O_NONBLOCK) < 0)
{
/* Handle error */
}
/* 设置为阻塞 */
if ((flags = fcntl(sock_descriptor, F_SETFL, 0)) < 0)
{
/* Handle error */
}
```

另外cmd的两个常用取值，做文件锁：

```c
// F_GETLK 取得文件锁定的状态。

// F_SETLK 设置文件锁定的状态。此时flcok 结构的l_type 值必须是F_RDLCK、F_WRLCK或F_UNLCK。如果无法建立锁定，则返回-1，错误代码为EACCES 或EAGAIN。

int _try_flock(int fd, int type, int whence, int start, int len)
{
  struct flock fl;
  fl.l_type = type;
  fl.l_whence = whence;
  fl.l_start = start;
  fl.l_len = len;
  
  return fcntl(fd, F_SETFL, &fl);
}
```

其中flock结构为：

```c
struct flock
{
  short int l_type; // 三种状态：F_RDLCK 建立一个供读取用的锁定；F_WRLCK 建立一个供写入用的锁定；F_UNLCK 删除之前建立的锁定
  short int l_whence; // 三种方式: SEEK_SET 以文件开头为锁定的起始位置; SEEK_CUR 以目前文件读写位置为锁定的起始位置; SEEK_END 以文件结尾为锁定的起始位置。
  off_t l_start; // 表示相对l_whence位置的偏移量，两者一起确定锁定区域的开始位置。
  off_t l_len; // 表示锁定区域的长度，如果为0表示从起点(由l_whence和 l_start决定的开始位置)开始直到最大可能偏移量为止。即不管在后面增加多少数据都在锁的范围内。
  off_t l_pid;
};

// 返回值 成功返回依赖于cmd的值，若有错误则返回-1，错误原因存于errno.
```



## 共享内存

要使用共享内存要执行以下几步：

1. 发起一个系统调用，让系统帮你生产一块内存，或者取得一块已经存在的内存来使用。
2. 把内存attach到当前进程，让当前进程可以使用。大家都知道，我们在进程中访问的是虚拟内存地址，系统会把它映射到物理内存中。如果没有这一步，第1步创建的内存就不能在当前进程访问。
3. 这时就可以对内存进程读写操作了。
4. 进程结束的时候要把上面attach的内存给释放。

系统调用（英语：system call），又称为系统呼叫，指运行在使用者空间的程序向操作系统内核请求需要更高权限运行的服务。系统调用提供用户程序与操作系统之间的接口。

`SYS_SHMGET`: 创建或者取得共享内存。
`SYS_SHMAT`: 将共享内存attach到当前进程空间。
`SYS_SHMDT`: 将共享内存从当前进程中deattach。

```c
int shmget(key_t key, size_t size, int shmflg);  
void *shmat(int shm_id, const void *shm_addr, int shmflg); 
int shmdt(const void *shmaddr);
```

golang不提供使用共享内存来通信，golang中通过cgo来调c语言来实现的

```go
func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
```

```go
shmid, _, err := syscall.Syscall(syscall.SYS_SHMGET, 2, 4, IpcCreate|0600)
```



## 参考资料

- [30 | 编程范式游记（1）- 起源](https://time.geekbang.org/column/article/301)

- [exec函数族-百度百科](https://baike.baidu.com/item/exec%E5%87%BD%E6%95%B0%E6%97%8F/3489348?fromtitle=EXEC&fromid=9077756&fr=aladdin)

- [fctl-百度百科](https://baike.baidu.com/item/fcntl/6860021?fr=aladdin)

- [Golang直接操作共享内存](https://studygolang.com/articles/10203)

