## 使用


后来使用才发现，这个是sftp功能，性能不行，传输速度为200+KB。

```go
package main

import (
	"fmt"
	"go-scp/scp"
)

func main() {
	scpinfo1 := scp.Scpinfo{User: "root", Password: "123456", Host: "192.168.5.203", Port: 22, LocalFilePath: "./go.mod", RemoteDir: "/backup/qunhui/"}
	fmt.Printf("scpinfo1: %v\n", scpinfo1)
	scpinfo1.Sendfile()
	scpinfo2 := scp.Scpinfo{User: "root", Password: "123456", Host: "192.168.5.203", Port: 22, LocalFilePath: "./", RemoteDir: "/backup/qunhui/CentOS-7-x86_64-Vagrant-2004_01.VirtualBox.box"}
	// scp.Receivefile()
	scpinfo2.Receivefile()
}

```

```go
package main

import (
	"fmt"
	"go-scp/scp"
	"os"
)

func main() {

	str, _ := os.Getwd()
	localfilepath := str + "/" + "go.sum"
	fmt.Printf("localfilepath: %v\n", localfilepath)
	str = fmt.Sprintf("sshpass -p 'password123456' scp %s  root@192.168.5.203:/volume1/devops/", localfilepath)
	fmt.Printf("str: %v\n", str)

	command := str
	scp.Runshell(command)
}

```
## 拉取
```bash
go get github.com/qq516249940/go-scp@v1.0.3
```
