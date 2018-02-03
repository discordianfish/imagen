package imagen

type Source struct {
	Origin string
	Ref    string
}

type Config struct {
	Base   Base
	Source Source
}
