package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gojuniper/internal/httpx"
)

func main() {
	// 这个可执行程序演示：
	// - flag 解析命令行参数
	// - net/http 启动服务
	// - 使用 os/signal 做优雅退出
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()

	srv := &http.Server{
		Addr:              *addr,
		Handler:           httpx.NewMux(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
	}()

	fmt.Println("listening on", *addr)
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Fprintln(os.Stderr, "server error:", err)
		os.Exit(1)
	}
}
