module github.com/freightcms/organizations/web

go 1.24.2

require (
	github.com/freightcms/organizations/db v0.0.0-20250319134210-79a6e808531e
	github.com/freightcms/organizations/models v0.0.0-20250319134210-79a6e808531e
	github.com/labstack/echo/v4 v4.13.4
)

require (
	github.com/freightcms/common/models v1.1.0 // indirect
	github.com/freightcms/locations/models v1.1.0 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
)

replace github.com/freightcms/organizations/db => ../db

replace github.com/freightcms/organizations/models => ../models
