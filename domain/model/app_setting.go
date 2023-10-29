package model

type LogSettings struct {
	FileNameFormat string
}

type ApiSettings struct {
	Host string
	Port string
}

type AppSettings struct {
	Log LogSettings
	Api ApiSettings
}
