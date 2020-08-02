package novel

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
}

//Novel information
type Novel struct {
	Author     string   `json:"author"`
	Title      string   `json:"title"`
	Slug       string   `json:"slug"`
	URL        string   `json:"url"`
	Summary    string   `json:"summary"`
	Tags       []string `json:"tags"`
	Categories []string `json:"categories"`
	Chapters   Chapters `json:"chapters"`
	Cover      string   `json:"cover"`
	Complete   bool     `json:"complete"`
}

//Chapter information
type Chapter struct {
	Title          string `json:"title"`
	Number         int    `json:"number"`
	OriginalNumber string `json:"original-number"`
	URL            string `json:"url"`
	Updated        string `json:"updated"`
	Content        string `json:"content"`
}

//Chapters list
type Chapters []Chapter

//Novels list
type Novels []Novel

//Sources struct
type Sources struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Novels Novels `json:"novels"`
}

func (n *Novel) GetTitle() string {
	return n.Title
}
