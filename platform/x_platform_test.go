package main

import (
	"encoding/json"
	"net/http/httptest"
)

func XTestSysMenuPage(t testingT) {
	httpTools.Get("/api/sys/menu/page",
		func(w *httptest.ResponseRecorder) {
			if w.Code != 200 {
				t.Errorf("SysUserLogin = %v want %v", w.Code, 200)
			}
			ret := Response{}
			json.Unmarshal(w.Body.Bytes(), &ret)
			if ret.Code != 401 {
				t.Errorf("SysMenuPage = %v want %v", ret.Code, 200)
			}
		},
	)
}
