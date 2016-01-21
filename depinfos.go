package jk

//DependencyInfos is a slice of DependencyInfo
type DependencyInfos []*DependencyInfo

//GetPredRep returns the "rep" for the DependencyInfo
func (depi *DependencyInfo) GetPredRep() string {
	pname, ok := depi.Features["用言代表表記"]
	if !ok {
		return ""
	}
	vtype, ok := depi.Features["用言"]
	if !ok {
		return ""
	}

	return pname + ":" + vtype
}
