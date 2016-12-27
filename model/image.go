package model

import "encoding/json"

type image struct {
	ID   ID
	Name string
}

type Image struct {
	image
}

func from(img Image) Image {
	return Image{image{
		ID:   img.ID(),
		Name: img.Name(),
	}}
}

func (i Image) ID() ID {
	return i.image.ID
}

func (i Image) WithID(id ID) Image {
	img := from(i)
	img.image.ID = id
	return img
}

func (i Image) Name() string {
	return i.image.Name
}

func (i Image) WithName(name string) Image {
	img := from(i)
	img.image.Name = name
	return img
}

func (i *Image) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &i.image)
}

func (i Image) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.image)
}
