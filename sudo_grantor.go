package main

import (
        "fmt"
        "os/exec"
//      "sync"
        "log"
	"flag"
)

func main() {
	hostPtr := flag.String("host", "dummy", "the hostname")
	userPtr := flag.String("user", "dummy", "the user")
	actionPtr := flag.Bool("delete", false, "pass -delete=true if you want to remove the sudo access")
	flag.Parse()

        hosts := [] string {*hostPtr}
        for _,element := range hosts {
                add := "/usr/bin/ssh -o StrictHostKeyChecking=no -i /etc/fk-ops-fai/keys/provisional " + element + " \"sudo echo -e '" + *userPtr + " ALL = (ALL) NOPASSWD: ALL' >> /etc/sudoers.d/temp; sudo chmod 440 /etc/sudoers.d/temp \" ";

                delete := "/usr/bin/ssh -o StrictHostKeyChecking=no -i /etc/fk-ops-fai/keys/provisional " + element + " \"sudo sed -i /'" + *userPtr + "/d' /etc/sudoers.d/* \" ";

if *actionPtr == true {
                output, err := exec.Command("sh", "-c", delete).Output()
                if err !=nil {
                	log.Fatal(err) } else {
			fmt.Println("Success.") }
		fmt.Println(string(output))
		} else {
			output1, err := exec.Command("sh", "-c", delete).Output()
		if err !=nil {
                        log.Fatal(err) } else {
                        fmt.Println("Checked for existing user. Passed.") }
                fmt.Println(string(output))
			output, err := exec.Command("sh", "-c", add).Output()
                if err !=nil {
                        log.Fatal(err) } else {
                        fmt.Println("Success.") }
                fmt.Println(string(output))
                }
}
}
