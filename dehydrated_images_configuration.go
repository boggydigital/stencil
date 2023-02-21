package stencil

type DehydratedImagesConfiguration struct {
	dehydratedImageProperty string
}

func (dic *DehydratedImagesConfiguration) GetDehydratedImageProperty() string {
	return dic.dehydratedImageProperty
}
