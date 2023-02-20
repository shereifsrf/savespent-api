package dao

import "sync"

// initialise dao
func Initialise() {
	wg := &sync.WaitGroup{}
	wg.Add(3)

	go worker(connectMySQL, wg)
	go worker(connectMongoDB, wg)
	go worker(connectRedis, wg)

	wg.Wait()
}

// method to add wg.done() to waitgroup for initialisation functions
func worker(fn func(), wg *sync.WaitGroup) {
	fn()
	wg.Done()
}
