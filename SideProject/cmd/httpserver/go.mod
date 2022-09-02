module server

go 1.18

require (
	https v1.0.0
	xmlparser v1.0.0
)

require github.com/golang/gddo v0.0.0-20210115222349-20d68f94ee1f // indirect

replace (
	https v1.0.0 => ../../internal/https
	xmlparser v1.0.0 => ../../internal/xmlparser
)
