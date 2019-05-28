package lust

import (
	"github.com/godcong/go-trait"
	"github.com/godcong/goav"
)

var log = trait.InitGlobalZapSugar()

func Transfer() {
	filename := "/mnt/d/video/周杰伦唱歌贼难听.mp4"

	// Register all formats and codecs
	goav.AVRegisterAll()

	ctx := goav.AVFormatAllocContext()

	// Open video file
	if goav.AVFormatOpenInput(&ctx, filename, nil, nil) != 0 {
		log.Info("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctx.AVFormatFindStreamInfo(nil) < 0 {
		log.Info("Error: Couldn't find stream information.")

		// Close input file and free context
		ctx.AVFormatCloseInput()
		return
	}

	for key, value := range ctx.Streams() {
		log.Info(key, value)
		parameters := value.CodecParameters()
		log.Info("type:", parameters.AVCodecGetType() == goav.AVMediaTypeVideo)
		log.Infof("width:%d", parameters.AVCodecGetWidth())
		log.Infof("height:%d", parameters.AVCodecGetHeight())
	}

}
