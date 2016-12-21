package model

type DataSource interface{
	ImageLoader
}

type ImageLoader interface {
	Image(id ID) Image
}