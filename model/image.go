package model

type Image struct {
	id   ID
	name string
}

func from(img Image) Image {
	return Image{
		id: img.ID(),
		name: img.Name(),
	}
}

func (i Image) ID() ID {
	return i.id;
}

func (i Image) WithID(id ID) Image {
	img := from(i)
	img.id = id
	return img
}

func (i Image) Name() string {
	return i.name;
}

func (i Image) WithName(name string) Image {
	img := from(i)
	img.name = name
	return img
}
