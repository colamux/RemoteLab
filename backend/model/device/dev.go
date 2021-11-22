package device

// const ()

type device struct {
	Id     uint64 `json:"id"`
	Uuid   string `json:"uuid"`
	Status string `json:"status"`
}

func GetStatus()