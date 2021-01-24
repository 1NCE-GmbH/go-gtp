// Copyright 2019-2021 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file

package gtpv1

import "testing"

func TestCauseStr(t *testing.T) {
	tests := []struct {
		name string
		cause uint8
		want string
	} {
		{
			name: "QoSParameterMismatch",
			cause: ReqCauseQoSParameterMismatch,
			want: "QoSParameterMismatch",
		},
		{
			name: "RequestAccepted",
			cause: ResCauseRequestAccepted,
			want: "RequestAccepted",
		},
		{
			name: "NonExistent",
			cause: ResCauseNonExistent,
			want: "NonExistent",
		},
		{
			name: "APNRestrictionTypeIncompatibilityWithCurrentlyActivePDPContexts",
			cause: ResCauseAPNRestrictionTypeIncompatibilityWithCurrentlyActivePDPContexts,
			want: "APNRestrictionTypeIncompatibilityWithCurrentlyActivePDPContexts",
		},
		{
			name: "Unknown code",
			cause: 142, // code for future use/reserved
			want: "142",
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
