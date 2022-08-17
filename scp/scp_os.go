package scp

import (
	"fmt"
	"log"
	"os/exec"
)

func Runshell(command string) {
	//c := exec.Command("cmd", "/C", "mysqldump -uroot -p654321 -h127.0.0.1 -P3306 db_cow > d:/db_cow.sql")
	// c := exec.Command("bash", "-c", "ls -al")

	// if err := c.Run(); err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	cmd := exec.Command("bash", "-c", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
