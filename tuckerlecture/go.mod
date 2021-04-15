module GolangStudy

go 1.16

require (
	GolangStudy/tuckerlecture/accounts v1.16.0
	GolangStudy/tuckerlecture/datastruct v1.16.0
)

replace (
	GolangStudy/tuckerlecture/accounts v1.16.0 => ./accounts
	GolangStudy/tuckerlecture/datastruct v1.16.0 => ./datastruct
)
