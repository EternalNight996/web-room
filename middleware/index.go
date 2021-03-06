// middleware/index.go
package middleware

import "sync"

var once sync.Once

func init() {
	once.Do(func() {
		initJWT()
		initCORS()
		initUserInfo()
	})
}
