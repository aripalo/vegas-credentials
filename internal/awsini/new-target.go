package awsini

func NewTargetProfile(iniConfig any, name string) (TargetProfile, error) {
	target := TargetProfile{Name: name}
	return target, load(iniConfig, target.Name, &target)
}
