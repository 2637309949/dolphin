package gin

import (
	"net/http"
	"reflect"

	"github.com/2637309949/dolphin/packages/web/core"
	"github.com/gin-gonic/gin"
)

var _ core.Handler = &handler{}

type handler struct {
	*gin.Engine
}

// Handler defined TODO
func (c *handler) Handler() http.Handler {
	return c.Engine
}

func (c *handler) setContext(in reflect.Value, ctx *gin.Context) reflect.Value {
	ct := in.Elem().Field(0)
	switch ct.Interface().(type) {
	case *gin.Context: //gin
		ct.Set(reflect.ValueOf(ctx))
	case *Context: //web/Context
		ct.Set(reflect.ValueOf(NewContext(ctx)))
	default: //platform/Context
		cType := ct.Type().Elem()
		cValue := reflect.New(cType)
		cField := cValue.Elem().Field(0)
		cField.Set(reflect.ValueOf(NewContext(ctx)))
		ct.Set(cValue)
	}
	return in
}

func (c *handler) handlerFunc(h core.HandlerFunc) func(*gin.Context) {
	return func(ctx *gin.Context) {
		fv := reflect.ValueOf(h)
		ft := fv.Type()
		fti := ft.In(0)
		isPtr := false
		if fti.Kind() == reflect.Ptr {
			fti = fti.Elem()
			isPtr = true
		}
		in := reflect.New(fti).Elem()
		_, isContext := in.Interface().(core.Context)
		switch true {
		case fti.String() == "core.Context":
			fv.Call([]reflect.Value{reflect.ValueOf(NewContext(ctx))})
		case isContext:
			in := c.setContext(reflect.New(fti), ctx)
			if isPtr {
				fv.Call([]reflect.Value{in})
			} else {
				fv.Call([]reflect.Value{in.Elem()})
			}
		}
	}
}

// Handler defined TODO
func (c *handler) Handle(httpMethod string, relativePath string, handlers ...core.HandlerFunc) {
	hfc := []gin.HandlerFunc{}
	for i := range handlers {
		fv := reflect.ValueOf(handlers[i])
		ft := fv.Type()
		if ft.NumIn() != 1 {
			panic("invalid handler: wrong numIn")
		}
		in := reflect.New(ft.In(0)).Elem()
		_, isContext := in.Interface().(core.Context)
		_, isGinContext := in.Interface().(*gin.Context)
		switch true {
		case isContext, ft.In(0).String() == "core.Context":
		case isGinContext:
			hfc = append(hfc, handlers[i].(func(*gin.Context)))
			continue
		default:
			panic("invalid handler context: should extend core.Context")
		}
		hfc = append(hfc, c.handlerFunc(handlers[i]))
	}
	c.Engine.Handle(httpMethod, relativePath, hfc...)
}

func NewHandler() core.Handler {
	return &handler{gin.New()}
}

func init() {
	gin.SetMode(gin.ReleaseMode)
	core.SetHandler(NewHandler())
}
