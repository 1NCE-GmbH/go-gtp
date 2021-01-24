// Copyright 2019-2021 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file

package gtpv2

import "testing"

func TestCauseStr(t *testing.T) {
	tests := []struct {
		name string
		cause uint8
		want string
	} {
		{
			name: "LocalDetach",
			cause: CauseLocalDetach,
			want: "LocalDetach",
		},
		{
			name: "QoSParameterMismatch",
			cause: CauseQoSParameterMismatch,
			want: "QoSParameterMismatch",
		},
		{
			name: "RequestAccepted",
			cause: CauseRequestAccepted,
			want: "RequestAccepted",
		},
		{
			name: "InvalidOverallLengthOfTheTriggeredResponseMessageAndAPiggybackedInitialMessage",
			cause: CauseInvalidOverallLengthOfTheTriggeredResponseMessageAndAPiggybackedInitialMessage,
			want: "InvalidOverallLengthOfTheTriggeredResponseMessageAndAPiggybackedInitialMessage",
		},
		{
			name: "Unknown code",
			cause: 15,
			want: "15",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CauseStr(tt.cause); got != tt.want {
				t.Errorf("CauseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
