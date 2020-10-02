package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "os/exec"
    "strconv"
    "io"
    "bytes"
    "strings"
    "github.com/cheggaaa/pb/v3"
)

const alpha = "abcdefghijklmnopqrstuvwxyz0123456789"

func isalpha(s string) bool {
    for _, char := range s {
        if !strings.Contains(alpha, strings.ToLower(string(char))) {
            return false
        }
    }
    return true
}

func isdir(path string)(bool) {
    fileInfo, err := os.Stat(path)
    if err != nil {
        return false
    }
    return fileInfo.IsDir()
}

func lineCounter(r io.Reader) (int, error) {
    buf := make([]byte, 32*1024)
    count := 0
    lineSep := []byte{'\n'}

    for {
        c, err := r.Read(buf)
        count += bytes.Count(buf[:c], lineSep)

        switch {
        case err == io.EOF:
            return count, nil

        case err != nil:
            return count, err
        }
    }
}

func sort(text string) bool {
    if text != "" {
        letter1 := strings.ToLower(string([] rune(text)[0]))
        if isalpha(letter1) {
            if !isdir("data/" + letter1) {
                file, err := os.OpenFile("data/" + letter1, os.O_APPEND | os.O_WRONLY, 0644);
                if err != nil {
                    log.Println(err)
                }
                defer file.Close()
                if _, err := file.WriteString(text + "\n");
                err != nil {
                    log.Fatal(err)
                }
                return true
            } else {
                letter2 := strings.ToLower(string([] rune(text)[1]))
                if isalpha(letter2) {
                    if !isdir("data/" + letter1 + "/" + letter2) {
                        file, err := os.OpenFile("data/" + letter1 + "/" + letter2, os.O_APPEND | os.O_WRONLY, 0644);
                        if err != nil {
                            log.Println(err)
                        }
                        defer file.Close()
                        if _, err := file.WriteString(text + "\n");
                        err != nil {
                            log.Fatal(err)
                        }
                        return true
                    } else {
                        letter3 := strings.ToLower(string([] rune(text)[2]))
                        if isalpha(letter3) {
                            if !isdir("data/" + letter1 + "/" + letter2 + "/" + letter3) {
                                file, err := os.OpenFile("data/" + letter1 + "/" + letter2 + "/" + letter3, os.O_APPEND | os.O_WRONLY, 0644);
                                if err != nil {
                                    log.Println(err)
                                }
                                defer file.Close()
                                if _, err := file.WriteString(text + "\n");
                                err != nil {
                                    log.Fatal(err)
                                }
                                return true
                            }
                        } else {
                            //letter 3 is symbol
                            file, err := os.OpenFile("data/" + letter1 + "/" + letter2 + "/symbols", os.O_APPEND | os.O_WRONLY, 0644);
                            if err != nil {
                                log.Println(err)
                            }
                            defer file.Close()
                            if _, err := file.WriteString(text + "\n");
                            err != nil {
                                log.Fatal(err)
                            }
                            return true
                        }
                    }
                } else {
                    //letter 2 is symbol
                    file, err := os.OpenFile("data/" + letter1 + "/symbols", os.O_APPEND | os.O_WRONLY, 0644);
                    if err != nil {
                        log.Println(err)
                    }
                    defer file.Close()
                    if _, err := file.WriteString(text + "\n");
                    err != nil {
                        log.Fatal(err)
                    }
                    return true
                }
            }
        } else {
            //letter 1 is symbol
            file, err := os.OpenFile("data/symbols", os.O_APPEND | os.O_WRONLY, 0644);
            if err != nil {
                log.Println(err)
            }
            defer file.Close()
            if _, err := file.WriteString(text + "\n");
            err != nil {
                log.Fatal(err)
            }
            return true
        }

    } else {
        fmt.Println("Incorrect Format!")
        return false
    }
    return false
}


func main() {
    //Look for ARG
    if len(os.Args) > 1 {
        if os.Args[1] == "-h" {
            fmt.Println("[*] Usage ")
            fmt.Println("   Add breach files in the directory inputbreach and run sorter.py")
            fmt.Println("   The breach has to be in the format:")
            fmt.Println("       adress1@domain.com:password1")
            fmt.Println("       adress2@domain.com:password2")
            fmt.Println("       ...")
            fmt.Println("   Once finished, use ./query.py to search for lines")
            fmt.Println(" ")
            fmt.Println("[*] Parameters:")
            fmt.Println("   \"-D\" : Deletes input file or files after being imported")
            os.Exit(0)
        }
    }


    //Verify checksum
    pidgo := strconv.Itoa(os.Getpid())
    execfile := "./scripts/checksum.sh"
    //fmt.Println(pid)
    cmd := &exec.Cmd {
        Path: execfile,
        Args: []string{execfile, pidgo},
        Stdout: os.Stdout,
        Stderr: os.Stdout,
    }

    cmd.Start();
    cmd.Wait()


    //Counting lines
    breach := "inputbreach/breach.txt"
    fmt.Println("[*] Checking number of lines for " + breach + ".")
    fl, err := os.Open(breach)
    linecount, err := lineCounter(fl)
    fmt.Println(strconv.Itoa(linecount))
    fmt.Println("  ")
    fmt.Println("  ")



    //Loop through file
    fmt.Println("[*] Sorting " + breach + ".")
    bar := pb.StartNew(linecount)
    file, erro := os.Open(breach)
    if erro != nil {
        log.Fatal(err)
    }
    defer file.Close()

    line := "init"

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line = scanner.Text()
        if line != ""{
            if sort(line) == true {
                bar.Increment()
            }else{
                fmt.Println(line)
            }    
        }else{
            bar.Increment()
        }
    }

    if erro := scanner.Err();
    erro != nil {
        log.Fatal(err)
    }

    bar.Finish()
    //check -d
    if len(os.Args) > 1 {
        if os.Args[1] == "-D" {
            execfile = "scripts/shalog.sh"
            cmd2 := &exec.Cmd {
                Path: execfile,
                Args: []string{execfile, "-D"},
                Stdout: os.Stdout,
                Stderr: os.Stdout,
            }
            cmd2.Start();
            cmd2.Wait()
        }else{
            execfile = "scripts/shalog.sh"
            cmd2 := &exec.Cmd {
                Path: execfile,
                Args: []string{execfile, "-F"},
                Stdout: os.Stdout,
                Stderr: os.Stdout,
            }
            cmd2.Start();
            cmd2.Wait()
        }
    }else{
        execfile = "scripts/shalog.sh"
        cmd2 := &exec.Cmd {
                Path: execfile,
                Args: []string{execfile, "-F"},
                Stdout: os.Stdout,
                Stderr: os.Stdout,
            }
        cmd2.Start();
        cmd2.Wait()
    }

}