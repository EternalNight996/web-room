// model/repo/token.go
package repo

// GetTokenInfo GET "/token/info" response object
type GetTokenInfo struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}
