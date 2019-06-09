package memo5

// Func is cachable function
type Func func(string) (interface{}, error)

// Memo is cache
type Memo struct {
	requests chan request
}

type result struct {
	value interface{}
	err   error
}

type request struct {
	key      string
	response chan result
}

type entry struct {
	ready  chan struct{}
	result result
}

// New creates new cache
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

// Get returns value of Func
func (m *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	m.requests <- request{key: key, response: response}
	res := <-response
	return res.value, res.err
}

func (m *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range m.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.result.value, e.result.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.result
}
