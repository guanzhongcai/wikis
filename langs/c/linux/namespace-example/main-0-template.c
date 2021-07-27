// main-0-template.c
#define _GNU_SOURCE
#include <sys/types.h>
#include <sys/wait.h>
#include <stdio.h>
#include <sched.h>
#include <signal.h>
#include <unistd.h>

#define STACK_SIZE (1024*1024)

static char child_stack[STACK_SIZE];
char* const child_args[] = {
	"/bin/bash",
	NULL
};

int child_main(void* arg) {
	printf("- World !\n");
	execv(child_args[0], child_args);
	printf("Ooops\n");
	return 1;
}

int main() {
	printf("- Hello ?\n");
	int child_pid = clone(child_main, child_stack + STACK_SIZE, SIGCHLD, NULL);
	waitpid(child_pid, NULL, 0);
	return 0;
}

// gcc -Wall main.c -o ns && ./ns

/*
Linux的3.12内核支持6种Namespace：
UTS: hostname（本文介绍）
IPC: 进程间通信 （之后的文章会讲到）
PID: "chroot"进程树（之后的文章会讲到）
NS: 挂载点，首次登陆Linux（之后的文章会讲到）
NET: 网络访问，包括接口（之后的文章会讲到）
USER: 将本地的虚拟user-id映射到真实的user-id（之后的文章会讲到）


Linux命名空间学习教程: http://dockerone.com/article/76
*/