package xvview

import (
	"github.com/forbot161602/x-srv-stream-origin/source/core/base/xvconst"
	"github.com/forbot161602/x-srv-stream-origin/source/kernel/stream/xvstrm"
)

var ViewStreamFileHandlers = []Handler{
	ViewStreamFileHandler,
}

func ViewStreamFileHandler(ctx *Context) {
	flow := &ViewStreamFileHandlerFlow{}
	flow.Initiate(ctx)
	// flow.IntakePathInfo()
	// flow.SetResult()
}

type ViewStreamFileHandlerFlow struct {
	APISIXFlow
}

func (flow *ViewStreamFileHandlerFlow) IntakePathInfo() {
	if flow.HasError() {
		return
	}

	streamID := flow.GetParam("streamID")
	fileName := flow.GetParam("fileName")
	pathInfo := xvstrm.NewStreamPathInfo(streamID, fileName)
	flow.Expose(xvconst.FlowKeyPathInfo, pathInfo)
}

func (flow *ViewStreamFileHandlerFlow) IntakeStreamInfo() {
	if flow.HasError() {
		return
	}
}

// func temp() {
// 	regex := regexp.MustCompile("^.+_(?P<suffix_number>\\d+)$")
// 	matches := regex.FindStringSubmatch("Gochuumon_wa_Usagi_Desu_ka_season_1_episode_1_720p_0012")
// 	fmt.Println(regex.SubexpNames()[1], mathces[1])
// }

// 1. intake streamID, fileName
// 2. read streamID/_info.json  !ok 404        HLSAssetInfo  AssetHLSInfo
// 3. type & protocol -> HLSAssetInfoGuide
// 4. make subject     HLSAssetSubject
// 4.1 by file name directly   ok ->>
// 4.2 by match each regex     !ok 404
// 5. read streamID/_info__{subject}.json      HLSAssetSubjectInfo
// 6.1 variant -> read file directly ->>
// 6.2 measure, file ordinal <= total_measures, extension right?  !ok 404
// 6.3 read file directly ->>
