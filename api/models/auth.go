package models

type HasAccessModel struct {
	ID               string `json:"id"`
	ProjectID        string `json:"platform_id"`
	ClientPlatformID string `json:"client_platform_id"`
	ClientTypeID     string `json:"client_type_id"`
	UserID           string `json:"user_id"`
	UserName         string `json:"user_name"`
	RoleID           string `json:"role_id"`
	IP               string `json:"ip"`
	Data             string `json:"data"`
	ExpiresAt        string `json:"expires_at"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	Inn              string `json:"inn"`
	OtherInn         string `json:"other_inn"`
	Pinfl            string `json:"pinfl"`
	IsAdditional     string `json:"is_additional"`
	IsSpecial        bool   `json:"is_special"`
	Login            string `json:"login"`
	FullName         string `json:"full_name"`
	BranchCode       string `json:"branch_code"`
	BranchName       string `json:"branch_name"`
}

type LoginWithSign struct {
	Sign     string            `json:"sign"`
	Pinfl    string            `json:"pinfl"`
	Name     string            `json:"name"`
	IsPhis   bool              `json:"isphis"`
	Info     map[string]string `json:"info"`
	IsMobile bool              `json:"is_mobile"`
}

type ResetPassword struct {
	Password string `json:"password"`
}
type UpdatePhoto struct {
	PhotoUrl string `json:"photourl"`
}

type PhotoUrl struct {
	PhotoUrl     string `json:"photo_url"`
	PhotoFullUrl string `json:"photo_full_url"`
}
