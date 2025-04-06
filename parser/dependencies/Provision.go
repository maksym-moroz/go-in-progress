package dependencies

func ProvideDependencies() *Dependencies {
	options := getOptions()
	return NewProvision(options...)
}

func getOptions() []Option {
	return []Option{
		WithStrategy(StandardStrategy),
		WithTimeout(StandardTimeout),
		WithContent(StandardContent),
	}
}
