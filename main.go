// Michael Harris
// COP4520 PA1 Problem 2, Version 4
// Multithreaded demonstration of the Dining Philosophers problem
// Number of diners is passed as the first argument through the command prompt

package main

import "os"
import "fmt"
import "sync"
import "time"
import "strconv"

var diners int = 5

var table sync.WaitGroup

func StartEating(seatIndex int, leftHand *sync.Mutex, rightHand *sync.Mutex) {
    // kick off the loop to satiate the philosophers' endless hunger
    for ;; {
        // once hungry, grab each chopstick and start to eat
        fmt.Printf("%d is now hungry.\n", seatIndex + 1)
        leftHand.Lock()
        rightHand.Lock()
        fmt.Printf("%d is now eating.\n", seatIndex + 1)

        // after eating for a while, put the chopsticks down
        time.Sleep(time.Millisecond * 5)
        leftHand.Unlock()
        rightHand.Unlock()

        // go back to thinking for a while
        fmt.Printf("%d is now thinking.\n", seatIndex + 1)
        time.Sleep(time.Millisecond * 5)
    }
    // close the waitgroup once finished, no way to bind this to keystroke 'n'
    // without changing terminal type, so this line never gets touched in this implementation
    table.Done()
}

func main() {
    temp := os.Args[1]
    diners, _ = strconv.Atoi(temp)
    // add the amount of diners to the waitgroup to auto-manage resource locks
    table.Add(diners)

    // create a mutex for the first fork and assign it for the first run
    newFork := &sync.Mutex{}
    forkRight := newFork

    // kick off a process for each diner, starting at position +1. each fork gets a mutex
    for i := 1; i < diners; i++ {
        forkLeft := &sync.Mutex{}
        go StartEating(i, forkLeft, forkRight)
        forkRight = forkLeft
    }

    // once everyone else is eating, the final process can begin with the final 2 forks
    go StartEating(0, newFork, forkRight)

    // keeps main running as long as there are hungry philosophers
    table.Wait()
}
