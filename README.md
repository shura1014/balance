# balance

负载均衡器 支持 轮询，随机，权重（在随机的基础上增大概率）

# 快速使用
定一个一个结构体，表示一个实例
```go
type Server struct {
	addr   string
	valid  bool
	weight int
}

func (s *Server) IsValid() bool {
	return s.valid
}

func (s *Server) GetWeight() int {
	return s.weight
}
```


## 无任何有效节点测试
```go
func NoInstance() {
	b := balance.GetBalance[*Server](balance.RoundRobin_)
	b.InitNodes(&Server{addr: "127.0.0.1:8081"}, &Server{addr: "127.0.0.1:8082"})
	for i := 0; i < 3; i++ {
		instance, err := b.Instance()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(instance.addr)
		}
	}
}


500:无可用实例
500:无可用实例
500:无可用实例
```

## 部分节点可用

```go
// Usable 部分可用情况
func Usable() {
	b := balance.GetBalance[*Server](balance.RoundRobin_)
	b.InitNodes(&Server{addr: "127.0.0.1:8081", valid: true}, &Server{addr: "127.0.0.1:8082"})
	for i := 0; i < 3; i++ {
		instance, err := b.Instance()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(instance.addr)
		}
	}
}

127.0.0.1:8081
127.0.0.1:8081
127.0.0.1:8081
```

## 轮询
```go
func RoundRobin() {
	b := balance.GetBalance[*Server](balance.RoundRobin_)
	b.InitNodes(&Server{addr: "127.0.0.1:8081", valid: true}, &Server{addr: "127.0.0.1:8082", valid: true})
	for i := 0; i < 3; i++ {
		instance, err := b.Instance()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(instance.addr)
		}
	}
}

127.0.0.1:8081
127.0.0.1:8082
127.0.0.1:8081
```

### 随机

```go
func Random() {
	b := balance.GetBalance[*Server](balance.Random_)
	b.InitNodes(&Server{addr: "127.0.0.1:8081", valid: true}, &Server{addr: "127.0.0.1:8082", valid: true})
	for i := 0; i < 6; i++ {
		instance, err := b.Instance()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(instance.addr)
		}
	}
}

127.0.0.1:8081
127.0.0.1:8082
127.0.0.1:8082
127.0.0.1:8082
127.0.0.1:8082
127.0.0.1:8081
```

# 权重

（本质也是随机，只不过随机的概率大一些）次数少的情况下，并不一定看的出效果

```go
func Weight() {
	b := balance.NewWeight[*Server]()
	b.InitNodes(&Server{addr: "127.0.0.1:8081", valid: true, weight: 2}, &Server{addr: "127.0.0.1:8082", valid: true, weight: 4})
	for i := 0; i < 15; i++ {
		instance, err := b.Instance()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(instance.addr)
		}
	}
}


127.0.0.1:8082
127.0.0.1:8082
127.0.0.1:8081
127.0.0.1:8082
127.0.0.1:8082
127.0.0.1:8081
127.0.0.1:8081
127.0.0.1:8082
127.0.0.1:8082
127.0.0.1:8082
127.0.0.1:8082
127.0.0.1:8082
127.0.0.1:8081
127.0.0.1:8081
127.0.0.1:8082

其中 8082命中10次,8081命中5次完全符合


127.0.0.1:8081
127.0.0.1:8082
127.0.0.1:8082
127.0.0.1:8081
127.0.0.1:8081
127.0.0.1:8082
127.0.0.1:8081
127.0.0.1:8082
127.0.0.1:8081
127.0.0.1:8082
127.0.0.1:8082
127.0.0.1:8082
127.0.0.1:8082
127.0.0.1:8081
127.0.0.1:8082
其中 8082命中9次,8081命中6次基本符合
```