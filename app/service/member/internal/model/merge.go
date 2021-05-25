package model

type Merge struct {
	// mid
	Mid int64 `json:"mid" `
	Kid int64 `json:"kid" `
	// business 业务
	Business string `json:"-"`
}
