module GolangStudy/SideProject/cmd/httpserver

go 1.18

require (
	GolangStudy/SideProject/pkg/https v1.0.0
	GolangStudy/SideProject/pkg/xmlparser v1.0.0
)

require github.com/golang/gddo v0.0.0-20210115222349-20d68f94ee1f // indirect

replace (
	GolangStudy/SideProject/pkg/https v1.0.0 => ../../pkg/https
	GolangStudy/SideProject/pkg/xmlparser v1.0.0 => ../../pkg/xmlparser
)
