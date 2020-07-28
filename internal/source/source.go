package source

//InSource interface
type InSource interface {
	New(name string)
	Remove()
	Update()
	List() string
	Novel()
}

//InNovel interface
type InNovel interface {
	New(name string)
	List()
	Update()
	Download(chapter string)
	Delete(chapter string)
	Remove()
}
