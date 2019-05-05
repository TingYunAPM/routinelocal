# routinelocal
协程局部存储器
## 安装

- 运行
```
go get github.com/TingYunAPM/routinelocal
```

## 接口说明
- ### 引用: 
```
import "github.com/TingYunAPM/go/routinelocal"
```
- ### 接口类 routine.Storage:
1. 方法: Storage.Get() uintptr :获取取局部存储对象指针
2. 方法: Storage.Set(p uintptr) : 替换局部存储对象指针

- ### 方法 routinelocal.Get() Storage : 返回局部存储对象的实例

## 使用
参考 test

## 说明:
 这个项目提供两种实现:
   1. 使用runtime.Stack,解析其中的goid, + 全局 map + 读写锁. 实现协程局部存储
        缺点: runtime.Stack和读写锁可能对性能会有损耗.
   2. native文件夹下提供了修改golang源代码的补丁。
        在runtime package 下增加GetRoutineLocal 和 SetRoutineLocal方法,
		通过打补丁,重新编译golang,获得协程局部存储的支持。

## 使用注意事项:
   在协程结束前,请务必调用 Storage.Set(uintptr(0)) 释放掉协程对应的局部存储对象指针，否则可能引发内存泄漏

