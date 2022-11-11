package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	name         string
	think        bool
	eat          bool
	useLeftFork  bool
	useRightFork bool
	timesEaten   int
	doneEating   bool
}

type Fork struct {
	forkIsFree bool
}

var philosophers []Philosopher
var forks []Fork
var toFork []chan bool
var toPhilo []chan bool

func aPhilosopher(index int, leftFork int, rightFork int) {
	philosopher := &philosophers[index]
	ch_toLeft := toFork[leftFork]
	ch_toRight := toFork[rightFork]
	ch_fromLeft := toPhilo[leftFork]
	ch_fromRight := toPhilo[rightFork]
	for philosopher.timesEaten != 3 {
		ch_toLeft <- true
		if <-ch_fromLeft {
			philosopher.useLeftFork = true
		} else {
			continue
		}
		ch_toRight <- true
		if <-ch_fromRight {
			philosopher.useRightFork = true
		} else {
			philosopher.useLeftFork = false
			ch_toLeft <- false
			continue
		}
		if philosopher.useLeftFork == true && philosopher.useRightFork == true {
			philosopher.timesEaten++
			fmt.Println(philosopher.name+" has eaten ", philosopher.timesEaten)
			philosopher.useLeftFork = false
			philosopher.useRightFork = false
			ch_toLeft <- false
			ch_toRight <- false
			fmt.Println(philosopher.name + " thinking")
		}
	}
	fmt.Println(philosopher.name+" is done, times eaten: ", philosopher.timesEaten)
}

func aFork(index int, toFork chan bool, ToPhilo chan bool) {
	fork := &forks[index]
	for {
		request := <-toFork
		if request {
			if fork.forkIsFree {
				fork.forkIsFree = false
				ToPhilo <- true
			} else {
				ToPhilo <- false
			}
		} else {
			fork.forkIsFree = true
		}
	}
}

func main() {
	// Two parallel slices:
	// philosopher Bob
	//                     fork 0
	// philosopher Joe
	//                     fork 1
	// philosopher Ben
	//                     fork 2
	// philosopher Jack
	//                     fork 3
	// philosopher Steve
	//                     fork 4

	philosophers = append(philosophers, Philosopher{name: "Bob", think: true, eat: false, useLeftFork: false, useRightFork: false, timesEaten: 0, doneEating: false})
	philosophers = append(philosophers, Philosopher{name: "Joe", think: true, eat: false, useLeftFork: false, useRightFork: false, timesEaten: 0, doneEating: false})
	philosophers = append(philosophers, Philosopher{name: "Ben", think: true, eat: false, useLeftFork: false, useRightFork: false, timesEaten: 0, doneEating: false})
	philosophers = append(philosophers, Philosopher{name: "Jack", think: true, eat: false, useLeftFork: false, useRightFork: false, timesEaten: 0, doneEating: false})
	philosophers = append(philosophers, Philosopher{name: "Steve", think: true, eat: false, useLeftFork: false, useRightFork: false, timesEaten: 0, doneEating: false})

	forks = append(forks, Fork{forkIsFree: true})
	forks = append(forks, Fork{forkIsFree: true})
	forks = append(forks, Fork{forkIsFree: true})
	forks = append(forks, Fork{forkIsFree: true})
	forks = append(forks, Fork{forkIsFree: true})

	free0 := make(chan bool)
	free1 := make(chan bool)
	free2 := make(chan bool)
	free3 := make(chan bool)
	free4 := make(chan bool)

	toFork = append(toFork, free0)
	toFork = append(toFork, free1)
	toFork = append(toFork, free2)
	toFork = append(toFork, free3)
	toFork = append(toFork, free4)

	done0 := make(chan bool)
	done1 := make(chan bool)
	done2 := make(chan bool)
	done3 := make(chan bool)
	done4 := make(chan bool)

	toPhilo = append(toPhilo, done0)
	toPhilo = append(toPhilo, done1)
	toPhilo = append(toPhilo, done2)
	toPhilo = append(toPhilo, done3)
	toPhilo = append(toPhilo, done4)

	go aPhilosopher(0, 4, 0)
	go aPhilosopher(1, 0, 1)
	go aPhilosopher(2, 1, 2)
	go aPhilosopher(3, 2, 3)
	go aPhilosopher(4, 3, 4)

	go aFork(0, toFork[0], toPhilo[0])
	go aFork(1, toFork[1], toPhilo[1])
	go aFork(2, toFork[2], toPhilo[2])
	go aFork(3, toFork[3], toPhilo[3])
	go aFork(4, toFork[4], toPhilo[4])

	time.Sleep(1000 * time.Millisecond)
}
