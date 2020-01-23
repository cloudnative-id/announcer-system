package main

type ContentCNCF struct {
	Content []ContentCNCFList `yaml:"content"`
}

type ContentCNCFList struct {
	Title string `yaml:"title"`
	Url string `yaml:"url"`
	Kind string `yaml:"kind"`
	IsDelivered bool `yaml:"isDelivered"`
}