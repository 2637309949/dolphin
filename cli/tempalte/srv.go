package tempalte

// TmplSrv defined template
var TmplSrv = `// Code generated by dol build. Only Generate by tools if not existed.
// source: {{.Controller.Name}}.go

package app

import "fmt"

// HelloSrv defined srv
func (ctr *{{$.Controller.ToUpperCase $.Controller.Name}}) HelloSrv() {
	fmt.Println("hello")
}
`
