//go:build darwin || freebsd || netbsd || dragonfly || openbsd

package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Youtch/rclone/fs/accounting"
)

// SigInfoHandler creates SigInfo handler
func SigInfoHandler() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINFO)
	go func() {
		for range signals {
			log.Printf("%v\n", accounting.GlobalStats())
		}
	}()
}
