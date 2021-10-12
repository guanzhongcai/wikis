# VS Code

## 插件安装

- C/C++
- cmake
- cmake tools



## CodeLLDB

A native debugger powered by LLDB. Debug C++, Rust and other compiled languages.

CodeLLDB离线下载链接：

https://github.com/vadimcn/vscode-lldb/releases

或vscode的扩展中搜索：codelldb

Here's a minimal debug configuration to get you started:

```json
{
    "name": "Launch",
    "type": "lldb",
    "request": "launch",
    "program": "${workspaceFolder}/<my program>",
    "args": ["-arg1", "-arg2"],
}
```

See Also: 

https://www.bilibili.com/video/BV1U741157Rd/?spm_id_from=333.788.recommend_more_video.1



## MinGW

MinGW，是*Minimalist* *[GNU](https://baike.baidu.com/item/GNU)**for Windows*的缩写。它是一个可自由使用和自由发布的Windows特定[头文件](https://baike.baidu.com/item/头文件/10978258)和使用GNU工具集导入库的集合，允许你在[GNU](https://baike.baidu.com/item/GNU)/[Linux](https://baike.baidu.com/item/Linux)和[Windows](https://baike.baidu.com/item/Windows)平台生成本地的[Windows程序](https://baike.baidu.com/item/Windows程序/15644576)而不需要第三方C运行时（C Runtime）库。

MinGW 是一组包含文件和端口库，其功能是允许控制台模式的程序使用微软的标准C运行时（C Runtime）库（[MSVCRT.DLL](https://baike.baidu.com/item/MSVCRT.DLL)）,该库在所有的 NT OS 上有效，在所有的 [Windows 95](https://baike.baidu.com/item/Windows 95)发行版以上的 Windows OS 有效，使用基本运行时，你可以使用 GCC 写控制台模式的符合美国标准化组织（ANSI）程序，可以使用微软提供的 C 运行时（C Runtime）扩展，与基本运行时相结合，就可以有充分的权利既使用 CRT（C Runtime）又使用 Windows[API](https://baike.baidu.com/item/API)功能。