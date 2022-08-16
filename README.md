## 使用
```
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