package controller

import (
	"strconv"
	"syscall/js"
)

func Update(v float64, element []js.Value) {

	scrollStar := strconv.FormatFloat(v*0.25, 'f', -1, 64) + "px"
	element[0].Get("style").Set("left", scrollStar)

	scrollMoon := strconv.FormatFloat(v*1.05, 'f', -1, 64) + "px"
	element[1].Get("style").Set("top", scrollMoon)

	scrollMountain_Behind := strconv.FormatFloat(v*0.5, 'f', -1, 64) + "px"
	element[2].Get("style").Set("top", scrollMountain_Behind)

	scrollMountain_Front := strconv.FormatFloat(v*0, 'f', -1, 64) + "px"
	element[5].Get("style").Set("top", scrollMountain_Front)

	scrollText := strconv.FormatFloat(v*4, 'f', -1, 64) + "px"
	element[3].Get("style").Set("marginRight", scrollText)

	scrollText = strconv.FormatFloat(v*1.5, 'f', -1, 64) + "px"
	element[3].Get("style").Set("marginTop", scrollText)

	scrollBtn := strconv.FormatFloat(v*1.5, 'f', -1, 64) + "px"
	element[4].Get("style").Set("marginTop", scrollBtn)

	scrollHeader := strconv.FormatFloat(v*0.5, 'f', -1, 64) + "px"
	element[6].Get("style").Set("top", scrollHeader)

}
