package stencil

type DehydratedImagesConfiguration struct {
	listDehydratedImageProperty string
	itemDehydratedImageProperty string
}

func (dic *DehydratedImagesConfiguration) GetListDehydratedImageProperty() string {
	return dic.listDehydratedImageProperty
}

func (dic *DehydratedImagesConfiguration) GetItemDehydratedImageProperty() string {
	return dic.itemDehydratedImageProperty
}
