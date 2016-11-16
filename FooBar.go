package main
import (
"ethos/syscall"
"ethos/ethos"
ethosLog "ethos/log"
"ethos/efmt" 
)
func main () {
me := syscall.GetUser()
path := "/user/" + me + "/myDir/"
status := ethosLog.RedirectToLog("myProgram")
if status != syscall.StatusOk {
efmt.Fprintf(syscall.Stderr, "Error opening %v: %v\n", path, 
status)
syscall.Exit(syscall.StatusOk)
}

data := MyType { "foo", "bar" }
fd, status := ethos.OpenDirectoryPath(path)
data.Write(fd)
data.WriteVar(path +"foobar")
efmt.Fprint(syscall.Stderr, "this will print in the log")
}