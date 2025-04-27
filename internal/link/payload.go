package link

type LinkCreateRequest struct { // Это то что приходит по апи
	Url string `json:"url" validate:"required,url"`
}

type LinkUpdateRequest struct { // Это то что приходит по апи

	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

type GetLinksResponce struct {
	Links []Link `json:"links"`
	Count int64  `json:"count"`
}
