# routinelocal patch
协程局部存储器 补丁
这个方案是通过给golang打补丁,添加 runtime.GetRoutineLocal 和 runtime.SetRoutineLocal 方法来支持协程局部存储的
##步骤
  1.linux下安装go环境.(脚本需要golang的依赖包)
  2.运行 patch_go.sh脚本,执行期间会下载墙外资源，请设置https_prox变量或者手动下载后，删除脚本的下载行.
  3.将编译好的golang文件夹下 bin 路径添加到 PATH
  4.设置GOROOT和GOPATH环境变量,完毕

