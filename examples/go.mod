module github.com/kjhickman/bnet-client/examples

go 1.24.4

require (
	github.com/joho/godotenv v1.5.1
	github.com/kjhickman/bnet-client v0.0.0-00010101000000-000000000000
)

require golang.org/x/oauth2 v0.30.0 // indirect

replace github.com/kjhickman/bnet-client => ..
