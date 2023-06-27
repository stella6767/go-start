package main

import "fmt"

/**
interface는 다른 interface에 상속(embed)할 수 있다.

재밌네 ㅋ 근데 맘에 안 드네.. 명시적으로 구현받지 않는다라..
덕 타이핑이 이런 의미구나.. 존나 맘에 안 드는데.

*/

type Wiki interface {
	Parser(string) string
}

// MediaWiki 엔진 구현을 위한 구조체
type MediaWiki struct {
	Type string
}

// MediaWiki 파서 실제 구현
func (m MediaWiki) Parser(a string) string {
	return "Moniwiki " + a
}

// 마크다운 엔진 구현을 위한 구조체
type MarkDown struct {
	Type string
}

// 마크다운 파서 실제 구현
func (m MarkDown) Parser(a string) string {
	return "Markdown " + a
}

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	md := MarkDown{Type: "MarkDown"}
	me := MediaWiki{Type: "MoniWiki"}

	// 위키엔진 저장을 위한 map 자료구조를 만들었다.
	WikiEngine := make(map[string]Wiki, 2)
	WikiEngine["markdown"] = md
	WikiEngine["mediawiki"] = me

	// MarkDown과 MediaWiki는 서로다른 구조체다.
	// 하지만 Wiki interface로 서로 연결이 됐다.
	for _, value := range WikiEngine {
		fmt.Println(value.Parser("Text Data"))
	}

	names := []string{"stanley", "david", "oscar"}
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}
	PrintAll(vals)
}
