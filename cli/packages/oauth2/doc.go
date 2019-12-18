// OAuth 2.0 server library for the Go programming language
//
//     package main
//     import (
//         "net/http"
//         "github.com/2637309949/dolphin/cli/packages/oauth2/manage"
//         "github.com/2637309949/dolphin/cli/packages/oauth2/server"
//         "github.com/2637309949/dolphin/cli/packages/oauth2/store"
//     )
//     func main() {
//         manager := manage.NewDefaultManager()
//         manager.MustTokenStorage(store.NewMemoryTokenStore())
//         manager.MapClientStorage(store.NewTestClientStore())
//         srv := server.NewDefaultServer(manager)
//         http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
//             srv.HandleAuthorizeRequest(w, r)
//         })
//         http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
//             srv.HandleTokenRequest(w, r)
//         })
//         http.ListenAndServe(":9096", nil)
//     }

package oauth2
