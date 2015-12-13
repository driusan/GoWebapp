package HTMLPageRenderer

type HTMLPage interface {
	GetJSFiles() []string
	GetCSSFiles() []string

	Render() string
}

