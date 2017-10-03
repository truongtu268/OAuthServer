package Dtos

type QueryOAuthDto struct {
	ClientId     string   `json:"client_id" validate:"required"`
	CallBack     string   `json:"redirect_uri" validate:"required"`
	ResponseType string   `json:"response_type" validate:"required"`
	Scope        []string `json:"scope" validate:"required"`
	State        string   `json:"state" validate:"required"`
	Query        string   `json:"query"`
}
