package main

type KubeweeklyContent struct {
	Title string `yaml:"title"`
	Date string `yaml:"date"`
	Source string `yaml:"source"`
	Data []KubeweeklyContentData
}

type KubeweeklyContentData struct {
	Title string `yaml:"title"`
	Link string `yaml:"link"`
	Type string `yaml:"type"`
}

type KubeweeklyContentList struct {
	ContentLists []ContentList
}

type ContentList struct{
	Content string `yaml:"content"`
	Date string `yaml:"date"`
	Status string `yaml:"status"`
	Tags []Tag
}

type Tag struct {
	string
}