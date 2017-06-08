package main

import (

        "fmt"
        "os"
        "strconv"
        "strings"
        "bufio"
        "time"
)

var counter float64 = 0
var prev_bytes float64 = 0
var flag bool = true

func ReadFile(filename string) ([]string, error){
        
        f,err := os.Open(filename)
        
        if err != nil {
                return []string{""}, err
        }
        
        defer f.Close()

        var lines []string

        read := bufio.NewReader(f)

        for {

                line, err := read.ReadString('\n')
                if err != nil {
                        break
                }

                lines = append(lines, strings.Trim(line,"\n"))

        }

        return lines, nil

}

func Fromfile() (float64){

        lines, _ := ReadFile("/proc/net/dev")
        
        var port string
        var data []string
        var rec_bytes float64
        var Average_rec_bytes float64 = 1
        
        for _ , line := range lines {
                fields := strings.Split(line, ":")
                if len(fields) < 2 {
                        continue
                }

                port = strings.TrimSpace(fields[0])
                data = strings.Fields(strings.TrimSpace(fields[1]))

                rec,err:=  strconv.ParseInt(data[0], 10, 64)
                

                if err != nil {
                        break
                }
                received := float64(rec)
                rec_bytes = rec_bytes + received
                     
        }
       

        rec_bytes = rec_bytes - prev_bytes
        
        if flag {
                prev_bytes = rec_bytes
                flag=false        
                return 0
        }


        Average_rec_bytes = Average_rec_bytes + rec_bytes
        counter = counter + 1
        
        return Average_rec_bytes/counter

}

func RepeatStat(d time.Duration, f func(time.Time)) {

        for n := range time.Tick(d) {
                f(n)
        }
}

func PrintStat(t time.Time){

        Throughput_B := Fromfile()
        Throughput_Mb := (Throughput_B*8)/1000000

        fmt.Printf("Average Throughput: %.6f Mbps\n",Throughput_Mb)
}

func main(){

        Fromfile()
        RepeatStat(1000*time.Millisecond, PrintStat)
}