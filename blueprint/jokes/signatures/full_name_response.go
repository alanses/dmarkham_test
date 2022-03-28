package signatures

type FullNameInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewFullNameInfo(firstName string, lastName string) *FullNameInfo {
	return &FullNameInfo{
		FirstName: firstName,
		LastName: lastName,
	}
}
