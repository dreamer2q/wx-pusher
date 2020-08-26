package helper

import (
	"fmt"
	"wx-pusher/config"
)

func ShowUrl(id string) string {
	return fmt.Sprintf("%s/show?id=%s", config.Api.Domain, id)
}
