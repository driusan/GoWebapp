package HTMLPageRenderer

type JSPage struct {
	Title    string
	JSFiles  []string
	CSSFiles []string
}

func (r JSPage) GetTitle() string {
	return r.Title
}
func (r JSPage) GetJSFiles() []string {
	return r.JSFiles
}
func (r JSPage) GetCSSFiles() []string {
	return r.CSSFiles
}

func (r JSPage) GetBody() string {
	return "<div id=\"workspace\"></div>"
}
