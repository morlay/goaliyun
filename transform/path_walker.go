package transform

import (
	"fmt"
)

type PathWalker struct {
	path []interface{}
}

func (pw *PathWalker) Enter(i interface{}) {
	pw.path = append(pw.path, i)
}

func (pw *PathWalker) Exit() {
	pw.path = pw.path[:len(pw.path)-1]
}

func (pw *PathWalker) Paths() []interface{} {
	return pw.path
}

func (pw *PathWalker) Current() string {
	pathString := ""
	for i := 0; i < len(pw.path); i++ {
		if pathString != "" {
			pathString += "."
		}
		switch pw.path[i].(type) {
		case string:
			pathString += pw.path[i].(string)
		case int:
			pathString += fmt.Sprintf("%d", pw.path[i].(int))
		}
	}
	return pathString
}
