package novel

//Novel information
type Novel struct {
	Title    string
	Slug     string
	URL      string
	Chapters Chapters
}

//Chapter information
type Chapter struct {
	Title   string
	Number  int
	URL     string
	Updated string
}

//Chapters list
type Chapters []Chapter

//Novels list
type Novels []Novel

//Sources struct
type Sources struct {
	Title  string
	URL    string
	Novels Novels
}
