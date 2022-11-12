package infra

import (
	"net/http"
	"strings"
)

const (
	pathRoot      = "/"
	pathDelimiter = "/"
)

type tree struct {
	node *node
}

type node struct {
	label    string
	actions  map[string]*action //handle action of a label
	children map[string]*node   // next nodes
}

type action struct {
	handler http.Handler
}

type result struct {
	actions *action
}

func newResult() *result {
	return &result{}
}

func NewTree() *tree {
	return &tree{
		node: &node{
			label:    pathRoot,
			actions:  make(map[string]*action),
			children: make(map[string]*node),
		},
	}
}

func explodePath(path string) []string {
	//ルーティング
	// hoge/hoge/piyo -> ["hoge","hoge","piyo"]
	// hoge/hoge:user_id=hoge -> ["hoge","hoge:user_id=hoge"]
	s := strings.Split(path, pathDelimiter)
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func (t *tree) Insert(methods []string, path string, handler http.Handler) {
	curNode := t.node
	if path == pathRoot {
		//ルートの場合はメソッドで分岐
		curNode.label = path
		for _, method := range methods {
			curNode.actions[method] = &action{handler: handler}
		}
		return
	}
	// root pathでない場合は/hoge/hogeを分解する
	ep := explodePath(path)
	for i, p := range ep {

		//Goのmapはmapを参照したときに2つ目の返り値に値が存在したかどうかを判定するフラグが返る
		nextNode, ok := curNode.children[p]
		// 既に存在する場合はnodeを書き換え
		if ok {
			curNode = nextNode
		}
		if !ok {
			curNode.children[p] = &node{
				label:    p,
				actions:  make(map[string]*action),
				children: make(map[string]*node),
			}
			curNode = curNode.children[p]
		}
		if i == len(ep)-1 {
			curNode.label = p
			for _, method := range methods {
				curNode.actions[method] = &action{handler: handler}
			}
			break
		}
	}
}
