package main

import (
	"fmt"
	"os"

	"github.com/iesreza/go-watcher"
)

func main() {
	params := watcher.ParseArgs(os.Args)
	fmt.Println("Initialize Auto Build Tool ...")
	w := watcher.MustRegisterWatcher(params)

	r := watcher.NewRunner()

	// wait for build and run the binary with given params
	go r.Run(params)
	b := watcher.NewBuilder(w, r)

	// build given package
	go b.Build(params)

	// listen for further changes
	go w.Watch()

	r.Wait()
}
