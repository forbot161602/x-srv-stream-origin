package xvstrm

type StreamURLPathInfo struct {
	StreamID string
	FileName string
}

func NewStreamPathInfo(streamID, fileName string) *StreamURLPathInfo {
	info := &StreamURLPathInfo{
		StreamID: streamID,
		FileName: fileName,
	}
	return info
}
