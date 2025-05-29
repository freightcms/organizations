module github.com/freightcms/organizations/db

go 1.23.4

replace github.com/freightcms/organizations/models => ../models

require github.com/freightcms/organizations/models v0.0.0-00010101000000-000000000000

require (
	github.com/freightcms/common/models v1.1.0 // indirect
	github.com/freightcms/locations/models v1.1.0 // indirect
)
