package assets

import "embed"

/*
  @Author : lanyulei
  @Desc :
*/

var (
	//go:embed assets/*
	StaticFs embed.FS

	//go:embed template/*
	TemplateFs embed.FS
)
