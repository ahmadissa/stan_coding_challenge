package jsonData

//Request structure that the web service is expecting
type Request struct {
	ExpectedCode int           `json:"expectedCode"`
	Payload      []RequestShow `json:"payload"`
	Skip         int           `json:"skip"`
	Take         int           `json:"take"`
	TotalRecords int           `json:"totalRecords"`
}

//RequestShow is the show used in Request
type RequestShow struct {
	Country      string `json:"country,omitempty"`
	Description  string `json:"description,omitempty"`
	Drm          bool   `json:"drm,omitempty"`
	EpisodeCount int    `json:"episodeCount,omitempty"`
	Genre        string `json:"genre,omitempty"`
	Image        struct {
		ShowImage string `json:"showImage"`
	} `json:"image,omitempty"`
	Language      string      `json:"language,omitempty"`
	NextEpisode   interface{} `json:"nextEpisode,omitempty"`
	PrimaryColour string      `json:"primaryColour,omitempty"`
	Seasons       []struct {
		Slug string `json:"slug"`
	} `json:"seasons,omitempty"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	TvChannel string `json:"tvChannel"`
}

//ResponseOK structure that the web service should send when success
type ResponseOK struct {
	Response []ResponseShow `json:"response"`
}

//ResponseShow is the show used in ResponseOK
type ResponseShow struct {
	Image string `json:"image"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

//ResponseError structure that the web service should send when fail
type ResponseError struct {
	Error string `json:"error"`
}
