// main-3-pid.c

#define _GNU_SOURCE
#include <sys/types.h>
#include <sys/wait.h>
#include <stdio.h>
#include <sched.h>
#include <signal.h>
#include <unistd.h>
#define STACK_SIZE (1024 * 1024)
// sync primitive
int checkpoint[2];
static char child_stack[STACK_SIZE];
char* const child_args[] = {
	"/bin/bash",
	NULL
};

int child_main(void* arg) {
	char c;
	// init sync primitive
	close(checkpoint[1]);
	// wait...
	read(checkpoint[0], &c, 1);
	printf(" - [%5d] World !\n", getpid());
	sethostname("In Namespace", 12);
	execv(child_args[0], child_args);
	printf("Ooops\n");
	return 1;
}

int main() {
	// init sync primitive
	pipe(checkpoint);
	printf(" - [%5d] Hello ?\n", getpid());
	int child_pid = clone(child_main, child_stack+STACK_SIZE,
	  CLONE_NEWUTS | CLONE_NEWIPC | CLONE_NEWPID | SIGCHLD, NULL);
	// further init here (nothing yet)
	// signal "done"
	close(checkpoint[1]);
	waitpid(child_pid, NULL, 0);
	return 0;
}



/**
 * Docker核心解决的问题是利用LXC来实现类似VM的功能，从而利用更加节省的硬件资源提供给用户更多的计算资源。而 LXC所实现的隔离性主要是来自内核的命名空间, 其中pid、net、ipc、mnt、uts 等命名空间将容器的进程、网络、消息、文件系统和hostname 隔离开。本文是Linux命名空间系列教程的第三篇，重点介绍PID命名空间。DockerOne在撸代码的基础上进行了校对和整理。
 * 
 * 
 * 
 * */