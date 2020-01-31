package mutex

type Mutex struct {
	lock map[string]chan int
}

func NewMutex() *Mutex {
	m := &Mutex{}
	m.lock = make(map[string]chan int)
	return m
}

func (m *Mutex) NewItem(key string) {
	m.lock[key] = make(chan int)
}

func (m *Mutex) Wait(key string) {
	m.lock[key] <- 1
}

func (m *Mutex) Signal(key string) {
	<-m.lock[key]
}
