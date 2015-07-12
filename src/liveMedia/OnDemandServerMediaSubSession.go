package liveMedia

import (
	"fmt"
	. "groupsock"
)

type OnDemandServerMediaSubSession struct {
	ServerMediaSubSession
	sdpLines        string
	portNumForSDP   int
	initialPortNum  int
	lastStreamToken interface{}
}

type StreamParameter struct {
	isMulticast     bool
	clientRTPPort   int
	clientRTCPPort  int
	serverRTPPort   int
	serverRTCPPort  int
	destinationTTL  uint
	destinationAddr string
	streamToken     interface{}
}

func (this *OnDemandServerMediaSubSession) InitOnDemandServerMediaSubSession(isubsession IServerMediaSubSession) {
	this.InitServerMediaSubSession(isubsession)
}

func (this *OnDemandServerMediaSubSession) SDPLines() string {
	if this.sdpLines == "" {
		inputSource := this.isubsession.CreateNewStreamSource()

		rtpPayloadType := 96 + this.TrackNumber() - 1
		dummyGroupSock := NewGroupSock(0)
		dummyRTPSink := this.isubsession.CreateNewRTPSink(dummyGroupSock, rtpPayloadType)

		this.setSDPLinesFromRTPSink(dummyRTPSink, inputSource, 500)
	}

	return this.sdpLines
}

func (this *OnDemandServerMediaSubSession) getStreamParameters(rtpChannelId, rtcpChannelId uint) *StreamParameter {
	//streamBitrate := 500
	//serverPortNum := this.initialPortNum

	//serverRTPPort  := serverPortNum
	//serverRTCPPort := serverPortNum + 1

	//rtpGroupSock  := NewGroupSock(serverRTPPort)
	//rtcpGroupSock := NewGroupSock(serverRTCPPort)

	//rtpSink := ""
	//mediaSource := ""
	//udpSink := NewBasicUDPSink(rtpGroupSock)

	//this.lastStreamToken = NewStreamState(serverRTPPort, serverRTCPPort, rtpSink, udpSink, streamBitrate, mediaSource, rtpGroupSock, rtcpGroupSock)

	sp := new(StreamParameter)
	return sp
}

func (this *OnDemandServerMediaSubSession) setSDPLinesFromRTPSink(rtpSink IRTPSink, inputSource IFramedSource, estBitrate uint) {
	if rtpSink == nil {
		return
	}

	rtpPayloadType := rtpSink.RtpPayloadType()
	mediaType := rtpSink.SdpMediaType()
	rtpmapLine := rtpSink.RtpmapLine()
	rangeLine := "" //this.rangeSDPLine()

	auxSDPLine := "" //this.getAuxSDPLine()
	if auxSDPLine == "" {
		auxSDPLine = ""
	}

	ipAddr := "0.0.0.0"

	sdpFmt := "m=%s %d RTP/AVP %d\r\n" +
		"c=IN IP4 %s\r\n" +
		"b=AS:%d\r\n" +
		"%s" +
		"%s" +
		"%s" +
		"a=control:%s\r\n"

	this.sdpLines = fmt.Sprintf(sdpFmt,
		mediaType,
		this.portNumForSDP,
		rtpPayloadType,
		ipAddr,
		estBitrate,
		rtpmapLine,
		rangeLine,
		auxSDPLine,
		this.TrackId())
}

func (this *OnDemandServerMediaSubSession) getAuxSDPLine(rtpSink *RTPSink) string {
	if rtpSink == nil {
		return ""
	} else {
		return rtpSink.AuxSDPLine()
	}
}

func (this *OnDemandServerMediaSubSession) startStream() {
	//streamState.startPlaying()
}

func (this *OnDemandServerMediaSubSession) pauseStream() {
	//streamState.pause()
}

func (this *OnDemandServerMediaSubSession) deleteStream() {
	//streamState.endPlaying()
}
