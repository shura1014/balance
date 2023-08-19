package balance

import (
	"github.com/shura1014/common/goerr"
	"sync"
)

// RoundRobin 轮询
type RoundRobin[V Node] struct {
	lock  sync.Mutex
	nodes []V
	next  int
}

func (b *RoundRobin[V]) Instance() (node V, err error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	for i := 0; i < len(b.nodes); i++ {
		next := (b.next) % len(b.nodes)
		b.next = next + 1
		node = b.nodes[next]
		if node.IsValid() {
			return
		}
	}
	err = goerr.Text("无可用实例")
	return
}

func (b *RoundRobin[V]) InitNodes(nodes ...V) {
	b.lock.Lock()
	b.nodes = append(b.nodes, nodes...)
	b.next = 0
	b.lock.Unlock()

}

func (b *RoundRobin[V]) Name() string {
	return RoundRobin_
}
