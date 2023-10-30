package model

type LogSettings struct {
	FileNameFormat string
}

type ApiSettings struct {
	Host     string
	BasePath string
}

type AppSettings struct {
	Log LogSettings
	Api ApiSettings
}
