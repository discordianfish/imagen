package imagen

type Source struct {
	Origin string
	Ref    string
}

type Base struct {
	Version string
}

type Config struct {
	Base   Base
	Source Source
	Labels map[string]string
}
