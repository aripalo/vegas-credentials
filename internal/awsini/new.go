package awsini

func NewSourceProfile(iniConfig any, name string) (SourceProfile, error) {
	source := SourceProfile{Name: name}
	return source, load(iniConfig, source.Name, &source)
}

func NewTargetProfile(iniConfig any, name string) (TargetProfile, error) {
	target := TargetProfile{Name: name}
	return target, load(iniConfig, target.Name, &target)
}
