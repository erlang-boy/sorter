package main

import "bufio"
import "flag"
import "fmt"
import "io"
import "os"
import "strconv"
import "time"

import "algorithm/bubblesort"
import "algorithm/qsort"


var infile  *string = flag.String("i","unsorted.dat","infile")
var outfile *string = flag.String("o","sorted.dat","outfile") 
var algorithm *string = flag.String("a","sqort","Sort")


func readValue(infile string) (value [] int, err error) {
       file,err := os.Open(infile)
       if err != nil {
          fmt.Println("open file failed: ", infile)
       }
         
       defer file.Close()
        br := bufio.NewReader(file)
       values= make([] int, 0)
       for {
         line, isPrefix, err := br.ReadLine()
         str := string(line)
         value, err1 := strconv.Atoi(str)
         values = append(values,value)
       }
       return
}

func writeValue(values [] int, outfile string) error {
      file, err := os.Create(outfile)
       if err != nil {
          fmt.Println("write file failed: ", outfile)
       }
      defer file.Close()
      for _, value := range values {
         str := strconv.Itoa(value) 
         file.WriteString(str + "\n")
      }
      return nil 
}

func main() {
   flag.Parse()
    
    if infile != nil {
      fmt.Println("infile = ", *infile)
    }

    values,err := readValues(*infile)
    if err == nil {
      t1 := time.Now()
       switch *algorithm { // here, maybe use interface 
          case "qsort":
                qsort.QuickSort(values)
          case "bubblesort":
                bubblesort.BubbleSort(values)
          default:
              fmt.Println("Not supported")
       } 
      t2 := time.Now()
      fmt.Println("the sorting costs",t2.Sub(t1),"to complete.")
      writeValues(value, *outfile)
    }    
    else {
       fmt.Println(err)
    }
}
