package golang

import "github.com/discordianfish/imagen"

type Config struct {
	imagen.Config
	Labels map[string]string
}
