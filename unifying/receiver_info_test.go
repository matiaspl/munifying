package unifying

import (
	"strings"
	"testing"
)

func TestGetSetInfoMissingReceiverSerial(t *testing.T) {
	u := LocalUSBDongle{sndQueue: make(chan USBReport, 8), rcvQueue: make(chan USBReport, 2)}
	page := make([]byte, USB_REPORT_TYPE_HIDPP_LONG_PAYLOAD_LEN)
	page[0], page[1] = byte(DONGLE_HIDPP_REGISTER_PAIRING_INFORMATION), 2
	u.rcvQueue <- &HidPPMsg{USB_REPORT_TYPE_HIDPP_LONG, 0xff, HIDPP_MSG_ID_GET_LONG_REGISTER_RSP, page}
	// A handheld notification arrives while waiting for the receiver's serial.
	u.rcvQueue <- &HidPPMsg{USB_REPORT_TYPE_HIDPP_LONG, 1, 4, make([]byte, len(page))}
	if _, err := u.GetSetInfo(); err == nil || !strings.Contains(err.Error(), "dongle serial") {
		t.Fatalf("expected a receiver serial error, got %v", err)
	}
}
