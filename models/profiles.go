package models

type KopiaProfile struct {
	Username            string `json:"username"`
	PasswordHashVersion int    `json:"passwordHashVersion"`
	PasswordHash        string `json:"passwordHash"`
}

type Profile struct {
	Username string `json:"username"`
	User     string `json:"name"`
	Hostname string `json:"host"`
}

type ProfilesResponse struct {
	Profiles []*Profile `json:"profiles"`
}

type AddProfileRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UpdateProfilePasswordRequest struct {
	Password string `json:"password"`
}
