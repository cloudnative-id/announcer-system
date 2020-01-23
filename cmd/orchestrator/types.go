package main

type KubeweeklyContentList struct {
	ContentLists []struct{
		Title string `yaml:"title"`
		Content string `yaml:"content"`
		Date string `yaml:"date"`
		Status struct{
			IsDelivered bool `yaml:"delivered"`
		} `yaml:"status"`
		Tags []string `yaml:"tags"`
	}`yaml:"contentList"`
}

type KubeweeklyContent struct {
	Title string `yaml:"title"`
	Date string `yaml:"date"`
	Source string `yaml:"source"`
	Data []struct {
		Title string `yaml:"title"`
		Link string `yaml:"link"`
		Type string `yaml:"type"`
	} `yaml:"data"`
}

type MeetupEventList struct {
	EventLists []struct {
		Event string `yaml:"event"`
		City string `yaml:"city"`
		Number int `yaml:"number"`
		Status struct{
			IsDelivered bool `yaml:"delivered"`
		} `yaml:"status"`
		Tags []string `yaml:"tags"`
	} `yaml:"eventList"`
}

type MeetupEvent struct {
    Name string `yaml:"name"`
	Date string `yaml:"date"`
	Time string `yaml:"time"`
	Place string `yaml:"place"`
	City string `yaml:"city"`
	Sponsor string `yaml:"sponsor"`
	RegistrationURL string `yaml:"registrationURL"`
	PicturePath string `yaml:"picturePath"`
	Speakers []struct {
		Name string `yaml:"name"`
		Title string `yaml:"title"`
		Company string `yaml:"company"`
		Position string `yaml:"position"`
	} `yaml:"speaker"`
}

type PostMeetupEvent struct {
	EventLists []PostMeetupEventList `yaml:"eventList"`
}

type PostMeetupEventList struct {
	Name string `yaml:"name"`
	City string `yaml:"city"`
	Sponsor []struct {
		Name string `yaml:"name"`
		URL string `yaml:"URL"`
	} `yaml:"sponsor"`
	PicturePath string `yaml:"picturePath"`
	Status struct{
		IsDelivered bool `yaml:"delivered"`
	} `yaml:"status"`
	Tags []string `yaml:"tags"`
	Speakers []struct {
		Name string `yaml:"name"`
		Title string `yaml:"title"`
		Company string `yaml:"company"`
		Position string `yaml:"position"`
		VideoURL string `yaml:"videoURL"`
		SlideURL string `yaml:"slideURL"`
	} `yaml:"speaker"`
}

type ContentCNCF struct {
	Content []ContentCNCFList `yaml:"content"`
}

type ContentCNCFList struct {
	Title string `yaml:"title"`
	Url string `yaml:"url"`
	Kind string `yaml:"kind"`
	IsDelivered bool `yaml:"isDelivered"`
}
