// Code generated by "stringer -type=actionType"; DO NOT EDIT

package builder

import "fmt"

const _actionType_name = "BuildPush"

var _actionType_index = [...]uint8{0, 5, 9}

func (i actionType) String() string {
	if i < 0 || i >= actionType(len(_actionType_index)-1) {
		return fmt.Sprintf("actionType(%d)", i)
	}
	return _actionType_name[_actionType_index[i]:_actionType_index[i+1]]
}