package dto

type DocumentDto struct {
	Name string `json:"name"`
}

/*
type ListUserReq struct {
	Name      	string `json:"name,omitempty" form:"name"`
	Email 		string `json:`
}
*/

type ListDocumentReq struct {
	Name string `json:"name,omitempty" form:"name"`
}

type ListDocumentRes struct {
	Documents []*DocumentDto `json:"documents"`
}

type CreateDocumentReq struct {
	Name string `json:"name" validate:"required"`
}

type UpdateDocumentReq struct {
	Name string `json:"name,omitempty"`
}
