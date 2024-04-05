package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "time"
)

var scriptDir string

func init() {
    scriptDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
}

func runProgram(program string, script io.Reader) {
    env := make(map[string]string)
    var process *exec.Cmd
    var stdin io.WriteCloser
    skip := false

    scanner := bufio.NewScanner(script)
    for scanner.Scan() {
        line := scanner.Text()
        line = strings.TrimSpace(line)
        line = strings.ReplaceAll(line, "$ETTS_SD", scriptDir)

        if strings.HasPrefix(line, "END_SKIP") {
            fmt.Println("Skip Mode OFF")
            skip = false
            continue
        }

        if strings.HasPrefix(line, "#") || line == "" || skip {
            // fmt.Printf("S: %s\n", line)
            continue
        }

        if strings.HasPrefix(line, "START_SKIP") {
            fmt.Println("Skip Mode On")
            skip = true
            continue
        }

        if strings.HasPrefix(line, "ENV ") {
            parts := strings.SplitN(line, " ", 3)
            if len(parts) == 3 {
                key, value := parts[1], parts[2]
                env[key] = value
            }
        } else if line == "ENVCLEAR" {
            env = make(map[string]string)
        } else if line == "RESTART" {
            if process != nil {
                if stdin != nil {
                    stdin.Close()
                }
                process.Wait()
            }
            process = exec.Command(program)
            process.Env = os.Environ()
            for k, v := range env {
                process.Env = append(process.Env, fmt.Sprintf("%s=%s", k, v))
            }
            process.Stdout = os.Stdout
            process.Stderr = os.Stderr
            var err error
            stdin, err = process.StdinPipe()
            if err != nil {
                fmt.Printf("Error creating stdin pipe: %v\n", err)
                continue
            }
            err = process.Start()
            if err != nil {
                fmt.Printf("Error starting process: %v\n", err)
                stdin.Close()
                continue
            }
        } else if strings.HasPrefix(line, "DELAY ") {
            parts := strings.SplitN(line, " ", 2)
            if len(parts) == 2 {
                delay, err := time.ParseDuration(parts[1]+"s")
                if err == nil {
                    fmt.Printf("D: %v\n", delay)
                    time.Sleep(delay)
                } else {
					fmt.Printf("Error: %s\n", err.Error())
				}
            }
        } else {
            if process != nil && process.ProcessState == nil {
                fmt.Printf("> %s\n", line)
                _, err := fmt.Fprintln(stdin, line)
                if err != nil {
                    fmt.Printf("Error writing to stdin: %v\n", err)
                }
            }
        }
    }

    if process != nil {
        if stdin != nil {
            stdin.Close()
        }
        process.Wait()
    }
}

func main() {
    program := flag.String("p", "", "The program to run")
    scriptFile := flag.String("s", "", "The script file to read from")
    flag.Parse()

    if *program == "" {
        fmt.Println("Please provide a program to run using the -p flag")
        os.Exit(1)
    }

    var script io.Reader
    if *scriptFile != "" {
        file, err := os.Open(*scriptFile)
        if err != nil {
            fmt.Printf("Error opening script file: %v\n", err)
            os.Exit(1)
        }
        defer file.Close()
        script = file
    } else {
        script = os.Stdin
    }

    runProgram(*program, script)
}
