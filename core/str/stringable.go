package str

import "travel-agency/util"

type Printable struct {

}

//转成 JSON
func (current Printable) ToString() string {
	return util.ObjectToStr(current)
}