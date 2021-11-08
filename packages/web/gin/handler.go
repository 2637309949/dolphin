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

func (c *handler) injectContext(in reflect.Value, ctx *gin.Context) {
	ct := in.Elem().FieldByName("Context")
	if ct.CanSet() {
		switch ct.Interface().(type) {
		case *gin.Context:
			ct.Set(reflect.ValueOf(ctx))
		case *Context:
			ct.Set(reflect.ValueOf(NewContext(ctx)))
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
			hfc = append(hfc, gin.HandlerFunc(handlers[i].(func(*gin.Context))))
			continue
		default:
			panic("invalid handler context: should extend core.Context")
		}

		hfc = append(hfc, func(v reflect.Value) func(ct *gin.Context) {
			ft := v.Type()
			fti := ft.In(0)
			isPtr := false
			if fti.Kind() == reflect.Ptr {
				fti = fti.Elem()
				isPtr = true
			}
			_, isContext := in.Interface().(core.Context)
			return func(ctx *gin.Context) {
				switch true {
				case fti.String() == "core.Context":
					v.Call([]reflect.Value{reflect.ValueOf(NewContext(ctx))})
				case isContext:
					in := reflect.New(fti)
					c.injectContext(in, ctx)
					if isPtr {
						v.Call([]reflect.Value{in})
					} else {
						v.Call([]reflect.Value{in.Elem()})
					}
				}
			}
		}(fv))
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
