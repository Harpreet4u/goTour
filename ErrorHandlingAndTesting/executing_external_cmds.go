package main

import (
    f "fmt"
    "os"
    "os/exec"
)

func main() {
// Using os, starting a process

// Linux
    env := os.Environ()
    procAttr := &os.ProcAttr{
                    Env: env,
                    Files: []*os.File{
                            os.Stdin,
                            os.Stdout,
                            os.Stderr,
                    },
                    
                }
    pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
    if err != nil {
        f.Printf("Error %v starting process!", err)
        os.Exit(1)
    }
    f.Printf("the process id is: %v", pid)

// Using os/exec
    cmd := exec.Command("gedit")
    err = cmd.Run()
    if err != nil {
        f.Printf("Error %v executing command!", err)
        os.Exit(1)
    }
    f.Printf("The cmd is: %v", cmd)
}
