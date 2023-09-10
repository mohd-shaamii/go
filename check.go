package main
import(
    "fmt"
    "sync"
    "time"
)

func main(){
    var wg sync.WaitGroup
    start:=make(chan struct{})
    num:=5
    
    for i:=1;i<=num;i++{
        wg.Add(1)
        go func(id int){
            defer wg.Done()
            <-start
            fmt.Printf("Worker %d is starting to work\n",id)
            time.Sleep(time.Second)
            
            fmt.Printf("Worker %d has reached checkpoint\n",id)
            time.Sleep(time.Second)
            time.Sleep(time.Second)
            
            fmt.Printf("Worker %d resuming the work\n",id)
            time.Sleep(time.Second)
            
            
        }(i)
    }
    fmt.Println("Starting workers.....")
    close(start)
    wg.Wait()
    fmt.Println("All workers completed their work")
}

// Starting workers.....
// Worker 1 is starting to work
// Worker 5 is starting to work
// Worker 3 is starting to work
// Worker 4 is starting to work
// Worker 2 is starting to work
// Worker 2 has reached checkpoint
// Worker 1 has reached checkpoint
// Worker 3 has reached checkpoint
// Worker 4 has reached checkpoint
// Worker 5 has reached checkpoint
// Worker 2 resuming the work
// Worker 1 resuming the work
// Worker 4 resuming the work
// Worker 3 resuming the work
// Worker 5 resuming the work
// All workers completed their work
