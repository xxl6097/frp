package frpc

import (
	"embed"

	"github.com/xxl6097/frp/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
