package db

type RClient struct {
	Addr string
	Pass string
	Db   int
}
type Mksresource struct {
	Created   string
	Active    string
	Failed    string
	Completed string
	Deleted   string
}

type Application struct {
	Mkstask, Mkspipelinerun, Mkstaskrun Mksresource
}
