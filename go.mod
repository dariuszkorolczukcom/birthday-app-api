module github.com/dariuszkorolczukcom/birthday-app

require (
	github.com/aws/aws-lambda-go v1.6.0
	github.com/dariuszkorolczukcom/birthday-app-api/birthday/structs v0.0.0-00010101000000-000000000000
)

replace github.com/dariuszkorolczukcom/birthday-app-api/birthday/structs => ./birthday/structs

go 1.13
