package model

type AddFavoriteReq struct {
	Name string `json:"name"`
	Mid  int64  `json:"mid"`
}

type AddFavoriteResp struct {
	ID int64 `json:"id"`
}

