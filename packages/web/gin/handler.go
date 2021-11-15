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

// set defined TODO
func (c *handler) set(numIn reflect.Value, ctx *gin.Context) reflect.Value {
	ct := numIn.Elem().Field(0)
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
	return numIn
}

// handlerFunc defined TODO
func (c *handler) handlerFunc(h core.HandlerFunc) func(*gin.Context) {
	return func(ctx *gin.Context) {
		fv := reflect.ValueOf(h)
		ft := fv.Type()
		in0 := ft.In(0)
		isPtr := false
		if in0.Kind() == reflect.Ptr {
			isPtr, in0 = true, in0.Elem()
		}
		in := reflect.New(in0).Elem()
		_, isContext := in.Interface().(core.Context)
		_, isGinContext := in.Interface().(*gin.Context)
		switch true {
		case in0.String() == "core.Context":
			fv.Call([]reflect.Value{reflect.ValueOf(NewContext(ctx))})
		case isGinContext:
			fv.Call([]reflect.Value{reflect.ValueOf(ctx)})
		case isContext:
			in := c.set(reflect.New(in0), ctx)
			if !isPtr {
				in = in.Elem()
			}
			fv.Call([]reflect.Value{in})
		}
	}
}

// Handler defined TODO
func (c *handler) Handle(httpMethod string, relativePath string, handlers ...core.HandlerFunc) {
	hfc := []gin.HandlerFunc{}
	for _, hlr := range handlers {
		fv := reflect.ValueOf(hlr)
		ft := fv.Type()
		if ft.NumIn() != 1 {
			panic("invalid handler: wrong numIn")
		}
		in0 := ft.In(0)
		in := reflect.New(in0).Elem()
		_, isContext := in.Interface().(core.Context)
		_, isGinContext := in.Interface().(*gin.Context)
		switch true {
		case isGinContext:
			hfc = append(hfc, c.handlerFunc(hlr))
		case isContext:
			hfc = append(hfc, c.handlerFunc(hlr))
		case in0.String() == "core.Context":
			hfc = append(hfc, c.handlerFunc(hlr))
		default:
			panic("invalid handler context: should extend core.Context")
		}
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
