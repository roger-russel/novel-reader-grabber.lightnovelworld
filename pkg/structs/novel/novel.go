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

//Tags list
type Tags []string

// Categories list
type Categories []string

//Novel information
type Novel struct {
	Author     string     `json:"author"`
	Title      string     `json:"title"`
	Slug       string     `json:"slug"`
	URL        string     `json:"url"`
	Summary    string     `json:"summary"`
	Tags       Tags       `json:"tags"`
	Categories Categories `json:"categories"`
	Chapters   Chapters   `json:"chapters"`
	Volumes    Volumes    `json:"volumes"`
	Cover      string     `json:"cover"`
	Complete   bool       `json:"complete"`
}

//Novels list
type Novels []Novel

//Volumes List
type Volumes []Volume

//Volume information
type Volume struct {
	Title    string    `json:"title"`
	Chapters *Chapters `json:"chapter"`
	Number   int       `json:"number"`
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

//Sources struct
type Sources struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Novels Novels `json:"novels"`
}
