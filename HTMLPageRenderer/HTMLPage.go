package HTMLPageRenderer

type HTMLPage interface {
	GetTitle() string
	GetJSFiles() []string
	GetCSSFiles() []string
	GetBody() string
}
