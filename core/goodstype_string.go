// Code generated by "stringer -type=GoodsType"; DO NOT EDIT.

package core

import "strconv"

const _GoodsType_name = "GoodsTypeMilkGoodsTypeDoor"

var _GoodsType_index = [...]uint8{0, 13, 26}

func (i GoodsType) String() string {
	if i >= GoodsType(len(_GoodsType_index)-1) {
		return "GoodsType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _GoodsType_name[_GoodsType_index[i]:_GoodsType_index[i+1]]
}
