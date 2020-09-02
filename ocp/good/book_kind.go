package good

type BookKind string

var (
	NOVEL   BookKind = "小説"
	COMICS  BookKind = "漫画"
	HISTORY BookKind = "歴史"
)

func (k BookKind) String() string {
	return string(k)
}
