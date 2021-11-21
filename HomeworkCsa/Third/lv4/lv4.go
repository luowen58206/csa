package lv4

import (
	"context"
	"fmt"
	"time"
)

func Lv4()  {
	ctx,cancel := context.WithCancel(context.Background())
	go dealRequest(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("it is time to stop all sub goroutine")
	cancel()

}

func dealRequest(ctx context.Context)  {
	go writeRedis(ctx)
	go writeDatabase(ctx)
	for  {
		select {
		case <-ctx.Done():
		fmt.Println("dealRequest done")
			return
		default:
			fmt.Println("dealRequest running")
			time.Sleep(2 * time.Second)
		}
	}
}

func writeRedis(ctx context.Context)  {
	for  {
		select {
		case <-ctx.Done():
			fmt.Println("writeRedis done")
			return
		default:
			fmt.Println("writeRedis running")
			time.Sleep(2 * time.Second)
		}
	}
}

func writeDatabase(ctx context.Context)  {
	for  {
		select {
		case <-ctx.Done():
			fmt.Println("writeDatabase done")
			return
		default:
			fmt.Println("writeDatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}
