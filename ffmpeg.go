package lust

import (
	"github.com/giorgisio/goav/avformat"
	"github.com/godcong/go-trait"
)

var log = trait.InitGlobalZapSugar()

func Transfer() {
	filename := "/mnt/d/video/周杰伦唱歌贼难听.mp4"

	// Register all formats and codecs
	avformat.AvRegisterAll()

	ctx := avformat.AvformatAllocContext()

	// Open video file
	if avformat.AvformatOpenInput(&ctx, filename, nil, nil) != 0 {
		log.Info("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctx.AvformatFindStreamInfo(nil) < 0 {
		log.Info("Error: Couldn't find stream information.")

		// Close input file and free context
		ctx.AvformatCloseInput()
		return
	}

	for key, value := range ctx.Streams() {
		log.Info(key, value)
		parameters := value.CodecParameters()
		log.Info("type:", parameters.AvCodecGetType() == avformat.AVMEDIA_TYPE_VIDEO)
		log.Infof("width:%d", parameters.AvCodecGetWidth())
		log.Infof("height:%d", parameters.AvCodecGetHeight())
	}

}
