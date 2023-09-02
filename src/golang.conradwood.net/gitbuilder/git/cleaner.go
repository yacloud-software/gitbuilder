package git

import (
	"fmt"
	"golang.conradwood.net/go-easyops/utils"
	"time"
)

func cleaner() {
	for {
		clean()
		time.Sleep(60 * time.Second)
	}
}
func clean() {
	wd_lock.Lock()
	defer wd_lock.Unlock()
	var nr []*LocalRepo
	for _, r := range repos {
		if r.inuse {
			nr = append(nr, r)
			continue
		}
		if time.Since(r.released) < time.Duration(5)*time.Minute {
			nr = append(nr, r)
			continue
		}
		err := utils.RemoveAll(r.workdir)
		if err != nil {
			if err != nil {
				fmt.Printf("failed to delete: %s\n", err)
				nr = append(nr, r)
				continue
			}
		}
		fmt.Printf("Deleted \"%s\"\n", r.workdir)
	}
	repos = nr
}
