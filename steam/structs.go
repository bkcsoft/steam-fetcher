package steam

// App holds App-data
type App struct {
	AppID int    `json:"appid"`
	Name  string `json:"name"`
}

// AppList holds a list of Apps
type AppList struct {
	Apps []App `json:"apps"`
}

// AppListResponse is a wrapper because steam doesn't know how to API...
type AppListResponse struct {
	AppList AppList `json:"applist"`
}
