module github.com/freightcms/organizations

go 1.24.2

require (
	github.com/dotenv-org/godotenvvault v0.6.0
	github.com/freightcms/logging v0.0.0-20250526023031-ace946d39537
	github.com/freightcms/organizations/db v0.0.0-20250319134210-79a6e808531e
	github.com/freightcms/organizations/db/mongodb v0.0.0-20250319134210-79a6e808531e
	github.com/freightcms/organizations/web v0.0.0-20250319134210-79a6e808531e
	github.com/labstack/echo/v4 v4.13.4
	go.mongodb.org/mongo-driver v1.17.3
)

require (
	github.com/freightcms/common/models v1.1.0 // indirect
	github.com/freightcms/locations/models v1.1.0 // indirect
	github.com/freightcms/organizations/models v0.0.0-20250319134210-79a6e808531e // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/graphql-go/graphql v0.8.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	golang.org/x/time v0.11.0 // indirect
)

replace github.com/freightcms/organizations/models => ./models

replace github.com/freightcms/organizations/db => ./db

replace github.com/freightcms/organizations/web => ./web

replace github.com/freightcms/organizations/db/mongodb => ./db/mongodb
