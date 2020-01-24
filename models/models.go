package models

// Kubeweekly Content
type KubeweeklyList struct {
	Data []KubeweeklyData `yaml:"kubeweekly"`
}

type KubeweeklyData struct{
	Title string `yaml:"title"`
	ContentFile string `yaml:"contentFile"`
	Date string `yaml:"date"`
	Status struct{
		IsDelivered bool `yaml:"delivered"`
	} `yaml:"status"`
	Tags []string `yaml:"tags"`
}

type KubeweeklyContent struct {
	Title string `yaml:"title"`
	Date string `yaml:"date"`
	Source string `yaml:"source"`
	Data []KubeweeklyContentData `yaml:"data"`
}

type KubeweeklyContentData struct {
	Title string `yaml:"title"`
	Link string `yaml:"link"`
	Type string `yaml:"type"`
}

// New Meetup Models
type NewMeetupList struct {
	Data []NewMeetupListData `yaml:"newMeetup"`
}

type NewMeetupListData struct {
	ContentFile string `yaml:"contentFile"`
	City string `yaml:"city"`
	Number int `yaml:"number"`
	Status struct{
		IsDelivered bool `yaml:"delivered"`
	} `yaml:"status"`
	Tags []string `yaml:"tags"`
}

type NewMeetupContent struct {
    Name string `yaml:"name"`
	Date string `yaml:"date"`
	Time string `yaml:"time"`
	Place string `yaml:"place"`
	City string `yaml:"city"`
	Sponsor string `yaml:"sponsor"`
	RegistrationURL string `yaml:"registrationURL"`
	PictureFile string `yaml:"pictureFile"`
	Speakers []struct {
		Name string `yaml:"name"`
		Title string `yaml:"title"`
		Company string `yaml:"company"`
		Position string `yaml:"position"`
	} `yaml:"speaker"`
}

// Post Meetup Models
type PostMeetupList struct {
	Content []PostMeetupContent `yaml:"postMeetup"`
}

type PostMeetupContent struct {
	Name string `yaml:"name"`
	City string `yaml:"city"`
	Sponsor []struct {
		Name string `yaml:"name"`
		URL string `yaml:"URL"`
	} `yaml:"sponsor"`
	PictureFile string `yaml:"pictureFile"`
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

// CNCF Newsroom Models
type NewsroomCNCFList struct {
	Content []NewsroomCNCFContent `yaml:"content"`
}

type NewsroomCNCFContent struct {
	Title string `yaml:"title"`
	Url string `yaml:"url"`
	Kind string `yaml:"kind"`
	IsDelivered bool `yaml:"isDelivered"`
}

// CNCF Webinar Models
type WebinarCNCFList struct {
	Content []WebinarCNCFContent `yaml:"content"`
}

type WebinarCNCFContent struct {
	Title string `yaml:"title"`
	Date string `yaml:"date"`
	Time string `yaml:"time"`
	Url string `yaml:"url"`
	IsDelivered bool `yaml:"isDelivered"`
}