package main

import (
	"fmt"
	"sync"
)

const NumResources = 10

type sharedResource struct {
	resource     func() string // accesses the shared resource and returns its result
	resourceLock *sync.Mutex   // protects concurrent access to resource
}

func resourceAccessor(num int, r *sharedResource) {
	r.resourceLock.Lock()
	defer r.resourceLock.Unlock()
	fmt.Sprintf("resource accessed by accessor %d: %s", num, r.resource())
}

func main() {

}
