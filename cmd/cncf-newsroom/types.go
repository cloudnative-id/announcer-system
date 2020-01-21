package main

type ContentCNCF struct {
	content []ContentCNCFList `yaml:"content"`
}

type ContentCNCFList struct {
	title string `yaml:"title"`
	url string `yaml:"url"`
	kind string `yaml:"kind"`
	isDelivered bool `yaml:"isdelivered"`
}