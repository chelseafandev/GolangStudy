module GolangStudy/SideProject

go 1.16

require (
	GolangStudy/SideProject/https v1.0.0
	GolangStudy/SideProject/xmlparser v1.0.0
)

replace (
	GolangStudy/SideProject/https v1.0.0 => ./https
	GolangStudy/SideProject/xmlparser v1.0.0 => ./xmlparser
)
