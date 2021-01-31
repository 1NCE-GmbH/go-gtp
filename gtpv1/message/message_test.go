package message

import "testing"

func TestMsgTypeStr(t *testing.T) {
	tests := []struct {
		name    string
		msgType uint8
		want    string
	}{
		{
			name:    "EchoRequest",
			msgType: MsgTypeEchoRequest,
			want:    "EchoRequest",
		},
		{
			name:    "DataRecordTransferResponse",
			msgType: MsgTypeDataRecordTransferResponse,
			want:    "DataRecordTransferResponse",
		},
		{
			name:    "Unknown msg type",
			msgType: 14,
			want:    "14",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MsgTypeStr(tt.msgType); got != tt.want {
				t.Errorf("MsgTypeStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
