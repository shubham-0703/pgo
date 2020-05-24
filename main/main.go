package main

 

import(
    "fmt"
    "log"
    "os"
    "path/filepath"
)
var (
    file  *os.File
    file2 *os.File
    err   error
)

 

func create(name string) {
    _,err := os.Stat(name)
    if err == nil {
        fmt.Printf("File exists\n")
        return

 

    }
    file, err = os.Create(name)

 

    if err != nil {
        panic(err)
    }
    log.Println("File Created")
    //    log.Println(file)
    //file.Close()
}
func rename(name, rname string) {
    err = os.Rename(name, rname)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("File Renamed")
}

 


func delete(name string) {
    err = os.Remove(name)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("%s Deleted", name)
}
func get(dir string) error {
    return filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
        if e != nil {
            return e
        }

 

        // check if it is a regular file (not dir)
        if info.Mode().IsRegular() {
            fmt.Println("file name:", info.Name())
            
        }
        return nil
    })
}

 


func main() {
    //create("File1.txt")
    //rename("File1.txt","Demo1.txt")
    //delete("Demo1.txt")
    get("File.txt")
}