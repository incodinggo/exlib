package searcher

import (
	"sync"
)

var (
	article = Article{}
)

type Attrib struct {
	Id        string
	Author    string
	Title     string
	Image     string
	Type      string
	Content   string
	View      int64
	Timestamp int64
	Text      string
	Tags      []string
}
type Article struct {
	sync.RWMutex
	Data map[string]Attrib
}

func init() {
	article.Data = map[string]Attrib{}
}
func (art *Article) Set(key string, data Attrib) {
	art.Lock()
	defer art.Unlock()
	art.Data[key] = data
}

func (art *Article) Get(key string) (data Attrib) {
	data = art.Data[key]
	return
}
