// Exercise 8.9: Write a version of du that computes and periodically displays separate totals for
// each of the root directories.
package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type result struct {
	index, value int64
}

// !+
func main() {
	// ...determine roots...

	//!-
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//!+
	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan result)
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(i, root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	//!-

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))
loop:
	for {
		select {
		case res, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nf, nb := res.index, res.value
			nfiles[nf]++
			nbytes[nf] += nb
		case <-tick:
			for i, n := range nfiles {
				printDiskUsage(n, nbytes[i])
			}
		}
	}

	for i, n := range nfiles {
		printDiskUsage(n, nbytes[i])
	}
	//!+
	// ...select loop...
}

//!-

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
// !+walkDir
func walkDir(index int, dir string, n *sync.WaitGroup, fileSizes chan<- result) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(index, subdir, n, fileSizes)
		} else {
			fileSizes <- result{int64(index), entry.Size()}
		}
	}
}

//!-walkDir

// !+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
