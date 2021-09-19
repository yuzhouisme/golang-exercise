package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "sync"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	eg, _ := errgroup.WithContext(ctx)

	// 模拟执行任务
	jobCount := 3
	done := make(chan bool, jobCount)
	for i := jobCount; i > 0; i-- {
		j := i
		eg.Go(func() error {
			fmt.Printf("Job#%v started, will sleep\n", j)
			time.Sleep(time.Second * time.Duration(10*j))

			select {
			case done <- true:
				fmt.Printf("Job%v done\n", j)
			case <-ctx.Done():
				fmt.Printf("Canceled Job#%v\n", j)
				return ctx.Err()
			}
			return nil
		})
	}

	// 处理操作系统中断信号
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 任务全部结束的信号
	// f := make(chan bool)
	// results := []bool{}

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监听程序被中断退出...")
				return ctx.Err()
			case <-sigs:
				// 接收 ctrl+c 信号处理
				fmt.Println("System interrupt")
				stopServer(cancel)
				// case <-time.After(time.Second * 30):
				// 假设30秒没结束，就退出
				// fmt.Println("Handle timeout")
				// stopServer(cancel)
				// case result := <-done:
				// 	results = append(results, result)
				// 	if len(results) == 3 {
				// 		fmt.Println("all done")
				// 		f <- true
				// 	}
			}
		}
	})

	// <-f

	if err := eg.Wait(); err == nil {
		fmt.Println("Job all done.")
	} else {
		fmt.Printf("Get errored : %v\n", err)
	}

	fmt.Println("Exit.")
}

func stopServer(cancel context.CancelFunc) {
	if cancel != nil {
		fmt.Println("Call context cancel func")
		cancel()
	}
}
