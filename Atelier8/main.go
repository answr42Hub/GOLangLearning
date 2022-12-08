package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const MAX_TOILETTES = 5

func main() {
	//no1()
	//no2()
	//no3()
	no4()
}

func no1() {
	test()
	var compte uint32 = 0
	var wg sync.WaitGroup
	for i := 1; i < 1001; i++ {
		wg.Add(1)
		go func(compte *uint32, wg *sync.WaitGroup, nb int) {
			defer wg.Done()
			if (nb%3) == 0 || (nb%5) == 0 {
				atomic.AddUint32(compte, 1)
			}

		}(&compte, &wg, i)
	}
	wg.Wait()
	fmt.Println(compte)
}

func test() {
	compte := 0
	for i := 1; i <= 1000; i++ {
		if (i%3) == 0 || (i%5) == 0 {
			compte++
		}
	}
	fmt.Println(compte)
}

func no2() {
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(nb int, wg *sync.WaitGroup, m *sync.Mutex) {
			routine := "Goroutine #" + strconv.Itoa(nb)
			defer wg.Done()
			m.Lock()
			for _, c := range routine {
				time.Sleep(time.Millisecond * 100)
				fmt.Printf("%c", c)
			}
			m.Unlock()
			fmt.Println()
		}(i, &wg, &m)
	}
	wg.Wait()
}

func no3() {
	var wg sync.WaitGroup
	var m sync.RWMutex

	var nombre uint32 = 90
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(nombre *uint32, wg *sync.WaitGroup, m *sync.RWMutex) {
			defer wg.Done()
			m.RLock()
			time.Sleep(time.Millisecond * 2000)
			fmt.Println("Je regarde le nomnre " + strconv.Itoa(int(*nombre)))
			m.RUnlock()
		}(&nombre, &wg, &m)
	}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(nombre *uint32, wg *sync.WaitGroup, m *sync.RWMutex) {
			defer wg.Done()
			m.Lock()
			time.Sleep(time.Millisecond * 2000)
			*nombre++
			fmt.Println("J'incrémente le nombre")
			m.Unlock()
		}(&nombre, &wg, &m)
	}
	wg.Wait()
}

func no4() {
	var toilettes uint32 = MAX_TOILETTES
	var m sync.Mutex
	var wg sync.WaitGroup

	rand.Seed(time.Now().Unix())

	for etudiant := 1; etudiant <= 100; etudiant++ {
		wg.Add(1)
		go func(id int, toilettes *uint32, m *sync.Mutex, wg *sync.WaitGroup) {
			time.Sleep(time.Second * time.Duration(rand.Intn(10)))
			etudiantEntre(id, toilettes, m, wg)
		}(etudiant, &toilettes, &m, &wg)
	}
	wg.Wait()
}

func etudiantEntre(id int, toilettes *uint32, m *sync.Mutex, wg *sync.WaitGroup, attente ...bool) {
	m.Lock()
	if *toilettes > 0 {
		*toilettes--
		fmt.Println("Etudiant #" + strconv.Itoa(id) + " entre dans les toilettes")
		m.Unlock()
		time.Sleep(time.Millisecond * 2000)
		m.Lock()
		*toilettes++
		fmt.Println("Etudiant #" + strconv.Itoa(id) + " sort des toilettes")
		m.Unlock()
	} else {
		if len(attente) > 0 && attente[0] {
			fmt.Println("Etudiant #" + strconv.Itoa(id) + " attend dans la file")
			m.Unlock()
			etudiantEntre(id, toilettes, m, wg, true)
		} else {
			fmt.Println("Etudiant #" + strconv.Itoa(id) + " va voir ailleurs")
			m.Unlock()
		}
	}
}
