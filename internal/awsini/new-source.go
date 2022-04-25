package awsini

func NewSourceProfile(iniConfig any, name string) (SourceProfile, error) {
	source := SourceProfile{Name: name}
	return source, load(iniConfig, source.Name, &source)
}
