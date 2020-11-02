# lockfile

```
go get github.com/jsxxzy/lockfile
```

example

```go
package main

import (
	"fmt"
	"os"

	"github.com/jsxxzy/lockfile"
)

// ExitFunc 退出
func ExitFunc() {
	os.Exit(0)
}

func main() {
	var lockfileName = "inet"
	run, code, _ := lockfile.NewSingleApp(lockfileName)
	if code == lockfile.AppRunOtherProcess {
		fmt.Println("进程同时存在了")
		return
	}
	run.Free(ExitFunc)
	for {

	}
}
```

# 参考

- https://github.com/jdxcode/golock

- https://github.com/nightlyone/lockfile

- https://gist.github.com/biezhi/74bfe20f9758210c1be18c64e6992a37

- https://stackoverflow.com/a/18110518

- https://stackoverflow.com/a/49560701

- https://www.jianshu.com/p/ae72ad58ecb6

- https://golangtc.com/t/56342908b09ecc3ac5000052

- https://www.janbar.top/index.php/2017/09/19/51.html

- https://linkscue.com/posts/2018-09-07-golang-flock-example/

- https://www.cnblogs.com/pingyeaa/p/11418527.html