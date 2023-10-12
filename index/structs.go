// package index for building index from xkcd web comics.
package index

const (
	URL           = "https://xkcd.com"
	lastComicsNum = 2840
)

type Comics struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

type Index struct {
	Items []*Comics
	File  string
}
