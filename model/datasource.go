package model

type DataSource interface{
	ImageLoader
}

type ImageLoader interface {
	AllImages(offset, limit int) []Image
	Image(id ID) Image
}