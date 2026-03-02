package resources



import (
	_ "embed"
)

var (
	//go:embed arial.ttf
	DefaultFont []byte

	//go:embed Aircraft.png
	Aircraft []byte

	//go:embed Desert.png
	Desert []byte

	//go:embed download.png
	Player []byte
)
