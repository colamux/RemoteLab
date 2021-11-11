package user

// TODO: need updating the fields
type UserInfo struct {
	Name   string `db:"name" json:"name"`
	Number uint64 `db:"number" json: "number"`
	Class  uint64 `db:"class" json:"class"`
	Vm     string `db:"vm" json:"vm"`
	Id     uint64 `db:"id" json:"id"`
}

const (
	_ = iota
	number
	phone
	email
)

type UserAuth struct {
	Id           uint64 `json:"id"`
	UId          uint64 `json:"uid"`
	IdentityType string `db:"identity_type" json:"identity_type"`
	Identifier   string `json:"identifier"`
	Password     string `json:"password"`
}
