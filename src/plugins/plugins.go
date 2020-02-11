package plugins

import (
	"Vtb_Record/src/plugins/monitor"
	"Vtb_Record/src/plugins/structUtils"
	"Vtb_Record/src/utils"
	"time"
)

type LiveStatus struct {
	isLive bool
	video  *structUtils.VideoInfo
}
type LiveTrace func(monitor monitor.VideoMonitor, usersConfig utils.UsersConfig) *LiveStatus

func GetLiveStatus(monitor monitor.VideoMonitor, usersConfig utils.UsersConfig) *LiveStatus {
	return &LiveStatus{
		isLive: monitor.CheckLive(usersConfig),
		video:  monitor.CreateVideo(usersConfig),
	}
}

func StartMonitor(monitor monitor.VideoMonitor, usersConfig utils.UsersConfig) {
	LiveStatus := &LiveStatus{video: &structUtils.VideoInfo{}}
	ticker := time.NewTicker(time.Second * time.Duration(utils.Config.CheckSec))
	for {
		p := &ProcessVideo{liveTrace: GetLiveStatus, monitor: monitor}
                liveStatus := GetLiveStatus(monitor, usersConfig)
		if liveStatus.isLive == true &&
			(LiveStatus.video.Title != liveStatus.video.Title || LiveStatus.video.Target != liveStatus.video.Target) {
			p.liveStatus = liveStatus
			LiveStatus = liveStatus
			go p.StartProcessVideo()
		}
		<-ticker.C
	}
}
