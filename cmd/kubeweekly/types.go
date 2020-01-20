package main

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

type KubeweeklyContentList struct {
	ContentLists []ContentList `yaml:"contentList"`
}

type ContentList struct{
	Title string `yaml:"title"`
	Content string `yaml:"content"`
	Date string `yaml:"date"`
	Status struct{
		IsDelivered bool `yaml:"delivered"`
	} `yaml:"status"`
	Tags []string `yaml:"tags"`
}
