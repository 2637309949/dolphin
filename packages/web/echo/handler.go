package echo

import (
	"net/http"
	"reflect"

	"github.com/2637309949/dolphin/packages/web/core"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

var _ core.Handler = &handler{}

type handler struct {
	*echo.Echo
}

// Handler defined TODO
func (c *handler) Handler() http.Handler {
	return c.Echo
}

// Handler defined TODO
func (c *handler) Handle(httpMethod string, relativePath string, handlers ...core.HandlerFunc) {
	hfc := []echo.HandlerFunc{}
	hfcMiddles := []echo.MiddlewareFunc{}
	for i := range handlers {
		fv := reflect.ValueOf(handlers[i])
		ft := fv.Type()
		if ft.NumIn() != 1 {
			panic("invalid handler: wrong numIn")
		}

		in := reflect.New(ft.In(0)).Elem()
		_, isContext := in.Interface().(core.Context)
		_, isEchoContext := in.Interface().(echo.Context)
		switch true {
		case isContext, ft.In(0).String() == "core.Context":
		case isEchoContext:
			hfc = append(hfc, echo.HandlerFunc(handlers[i].(func(echo.Context) error)))
			hfcMiddles = append(hfcMiddles, func(hf echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					return echo.HandlerFunc(handlers[i].(func(echo.Context) error))(c)
				}
			})
			continue
		default:
			panic("invalid handler context: should extend core.Context")
		}
		hfc = append(hfc, func(v reflect.Value) func(ct echo.Context) error {
			ft := v.Type()
			fti := ft.In(0)
			isPtr := false
			if fti.Kind() == reflect.Ptr {
				fti = fti.Elem()
				isPtr = true
			}
			_, isContext := in.Interface().(core.Context)
			return func(ctx echo.Context) error {
				switch true {
				case fti.String() == "core.Context":
					v.Call([]reflect.Value{reflect.ValueOf(NewContext(ctx))})
				case isContext:
					in := reflect.New(fti)
					ct := in.Elem().FieldByName("Context")
					if ct.CanSet() {
						ct.Set(reflect.ValueOf(NewContext(ctx)))
						if isPtr {
							v.Call([]reflect.Value{in})
						} else {
							v.Call([]reflect.Value{in.Elem()})
						}
					}
				}
				return nil
			}
		}(fv))

		hfcMiddles = append(hfcMiddles, func(v reflect.Value) func(hf echo.HandlerFunc) echo.HandlerFunc {
			ft := v.Type()
			fti := ft.In(0)
			isPtr := false
			if fti.Kind() == reflect.Ptr {
				fti = fti.Elem()
				isPtr = true
			}
			_, isContext := in.Interface().(core.Context)
			return func(hf echo.HandlerFunc) echo.HandlerFunc {
				return func(ctx echo.Context) error {
					switch true {
					case fti.String() == "core.Context":
						v.Call([]reflect.Value{reflect.ValueOf(NewContext(ctx, func() { hf(ctx) }))})
					case isContext:
						in := reflect.New(fti)
						ct := in.Elem().FieldByName("Context")
						if ct.CanSet() {
							ct.Set(reflect.ValueOf(NewContext(ctx, func() { hf(ctx) })))
							if isPtr {
								v.Call([]reflect.Value{in})
							} else {
								v.Call([]reflect.Value{in.Elem()})
							}
						}
					}
					return nil
				}
			}
		}(fv))
	}
	for i := range hfc {
		hfcMiddles = append(hfcMiddles, func(hf echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				hfc[i](c)
				return hfc[i](c)
			}
		})
	}

	// remove first hander
	if len(handlers) > 1 && len(hfcMiddles) > 0 {
		hfcMiddles = hfcMiddles[1:]
	}
	c.Echo.Add(httpMethod, relativePath, hfc[0], hfcMiddles...)
}

func NewHandler() core.Handler {
	return &handler{echo.New()}
}

func init() {
	gin.SetMode(gin.ReleaseMode)
	core.SetHandler(NewHandler())
}
