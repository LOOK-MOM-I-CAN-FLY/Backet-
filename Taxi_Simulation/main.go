package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

func getTaxi(ctx context.Context, serviceName string, resulst chan string) {
	time.Sleep(3 * time.Second)
	for {
		select {
		case <-ctx.Done():
			log.Printf("Stopped the search in %q (%v)", serviceName, ctx.Err())
			return
		default:
			if rand.Float64() > 0.75 {
				resulst <- serviceName
				return
			}
			continue
		}
	}
}

func main() {
	var (
		resultCh    = make(chan string)
		ctx, cancel = context.WithCancel(context.Background())
		services    = []string{"Uber", "Delimobil", "Moscvich", "Yandex Go"}
		wg          sync.WaitGroup
		winner      string
	)
	for _, service := range services {
		wg.Add(1)
		go func() {
			getTaxi(ctx, service, resultCh)
			wg.Done()
		}()
	}
	go func() {
		winner = <-resultCh
		cancel()
	}()
	wg.Wait()
	log.Printf("found car in %q", winner)

}
