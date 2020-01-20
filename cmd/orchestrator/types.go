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


