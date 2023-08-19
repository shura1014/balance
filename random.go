package balance

import (
	"github.com/shura1014/common/goerr"
	"github.com/shura1014/common/random"
	"sync"
)

type Random[V Node] struct {
	lock  sync.Mutex
	nodes []V
}

func (b *Random[V]) Instance() (node V, err error) {
	index := random.Int(len(b.nodes) - 1)
	b.lock.Lock()
	defer b.lock.Unlock()
	for i := 0; i < len(b.nodes); i++ {
		node = b.nodes[(index+i)%len(b.nodes)]
		if node.IsValid() {
			return
		}
	}
	err = goerr.Text("无可用实例")
	return

}

func (b *Random[V]) InitNodes(nodes ...V) {
	b.lock.Lock()
	b.nodes = append(b.nodes, nodes...)
	b.lock.Unlock()
}

func (b *Random[V]) Name() string {
	return Random_
}
