package route

/*
  @Author : lanyulei
  @Desc :
*/

type Route struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

const (
	Unregistered = "unregistered"
	Invalid      = "invalid"
)
