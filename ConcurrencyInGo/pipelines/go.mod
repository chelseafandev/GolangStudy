module GolangStudy/ConcurrencyInGo/pipelines

go 1.16

require (
GolangStudy/ConcurrencyInGo/pipelines/squaringnumbers v1.16.0
GolangStudy/ConcurrencyInGo/pipelines/fanoutfanin v1.16.0
GolangStudy/ConcurrencyInGo/pipelines/stoppingshort v1.16.0
GolangStudy/ConcurrencyInGo/pipelines/explicitcancellation v1.16.0
)

replace (
GolangStudy/ConcurrencyInGo/pipelines/squaringnumbers v1.16.0 => ./squaringnumbers
GolangStudy/ConcurrencyInGo/pipelines/fanoutfanin v1.16.0 => ./fanoutfanin
GolangStudy/ConcurrencyInGo/pipelines/stoppingshort v1.16.0 => ./stoppingshort
GolangStudy/ConcurrencyInGo/pipelines/explicitcancellation v1.16.0 => ./explicitcancellation
)