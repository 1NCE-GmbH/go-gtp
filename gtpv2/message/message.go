// Copyright 2019-2021 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

/*
Package message provides encoding/decoding feature of GTPv2 protocol.
*/
package message

import (
	"fmt"
	"strconv"
)

// Message Type definitions.
const (
	_ uint8 = iota
	MsgTypeEchoRequest
	MsgTypeEchoResponse
	MsgTypeVersionNotSupportedIndication
	MsgTypeDirectTransferRequest
	MsgTypeDirectTransferResponse
	MsgTypeNotificationRequest
	MsgTypeNotificationResponse
	_
	_
	_
	_
	_
	_
	_
	_
	_ // 8-16: Reserved for S101 interface
	MsgTypeRIMInformationTransfer
	_
	_
	_
	_
	_
	_
	_ // 18-24: Reserved for S121 interface
	MsgTypeSRVCCPsToCsRequest
	MsgTypeSRVCCPsToCsResponse
	MsgTypeSRVCCPsToCsCompleteNotification
	MsgTypeSRVCCPsToCsCompleteAcknowledge
	MsgTypeSRVCCPsToCsCancelNotification
	MsgTypeSRVCCPsToCsCancelAcknowledge
	MsgTypeSRVCCCsToPsRequest
	MsgTypeCreateSessionRequest
	MsgTypeCreateSessionResponse
	MsgTypeModifyBearerRequest
	MsgTypeModifyBearerResponse
	MsgTypeDeleteSessionRequest
	MsgTypeDeleteSessionResponse
	MsgTypeChangeNotificationRequest
	MsgTypeChangeNotificationResponse
	MsgTypeRemoteUEReportNotification
	MsgTypeRemoteUEReportAcknowledge
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_ // 42-63: Reserved for S4/S11, S5/S8 interfaces
	MsgTypeModifyBearerCommand
	MsgTypeModifyBearerFailureIndication
	MsgTypeDeleteBearerCommand
	MsgTypeDeleteBearerFailureIndication
	MsgTypeBearerResourceCommand
	MsgTypeBearerResourceFailureIndication
	MsgTypeDownlinkDataNotificationFailureIndication
	MsgTypeTraceSessionActivation
	MsgTypeTraceSessionDeactivation
	MsgTypeStopPagingIndication
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_ // 74-94: Reserved for GTPv2 non-specific interfaces
	MsgTypeCreateBearerRequest
	MsgTypeCreateBearerResponse
	MsgTypeUpdateBearerRequest
	MsgTypeUpdateBearerResponse
	MsgTypeDeleteBearerRequest
	MsgTypeDeleteBearerResponse
	MsgTypeDeletePDNConnectionSetRequest
	MsgTypeDeletePDNConnectionSetResponse
	MsgTypePGWDownlinkTriggeringNotification
	MsgTypePGWDownlinkTriggeringAcknowledge
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_ // 105-127: Reserved for S5, S4/S11 interfaces
	MsgTypeIdentificationRequest
	MsgTypeIdentificationResponse
	MsgTypeContextRequest
	MsgTypeContextResponse
	MsgTypeContextAcknowledge
	MsgTypeForwardRelocationRequest
	MsgTypeForwardRelocationResponse
	MsgTypeForwardRelocationCompleteNotification
	MsgTypeForwardRelocationCompleteAcknowledge
	MsgTypeForwardAccessContextNotification
	MsgTypeForwardAccessContextAcknowledge
	MsgTypeRelocationCancelRequest
	MsgTypeRelocationCancelResponse
	MsgTypeConfigurationTransferTunnel
	_
	_
	_
	_
	_
	_
	_ // 142-148: Reserved for S3/S10/S16 interfaces
	MsgTypeDetachNotification
	MsgTypeDetachAcknowledge
	MsgTypeCSPagingIndication
	MsgTypeRANInformationRelay
	MsgTypeAlertMMENotification
	MsgTypeAlertMMEAcknowledge
	MsgTypeUEActivityNotification
	MsgTypeUEActivityAcknowledge
	MsgTypeISRStatusIndication
	MsgTypeUERegistrationQueryRequest
	MsgTypeUERegistrationQueryResponse
	MsgTypeCreateForwardingTunnelRequest
	MsgTypeCreateForwardingTunnelResponse
	MsgTypeSuspendNotification
	MsgTypeSuspendAcknowledge
	MsgTypeResumeNotification
	MsgTypeResumeAcknowledge
	MsgTypeCreateIndirectDataForwardingTunnelRequest
	MsgTypeCreateIndirectDataForwardingTunnelResponse
	MsgTypeDeleteIndirectDataForwardingTunnelRequest
	MsgTypeDeleteIndirectDataForwardingTunnelResponse
	MsgTypeReleaseAccessBearersRequest
	MsgTypeReleaseAccessBearersResponse
	_
	_
	_
	_ // 172-175: Reserved for S4/S11 interfaces
	MsgTypeDownlinkDataNotification
	MsgTypeDownlinkDataNotificationAcknowledge
	_
	MsgTypePGWRestartNotification
	MsgTypePGWRestartNotificationAcknowledge
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_ // 181-199: Reserved for S4 interface
	MsgTypeUpdatePDNConnectionSetRequest
	MsgTypeUpdatePDNConnectionSetResponse
	_
	_
	_
	_
	_
	_
	_
	_
	_ // 202-210: Reserved for S5/S8 interfaces
	MsgTypeModifyAccessBearersRequest
	MsgTypeModifyAccessBearersResponse
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_ // 213-230: Reserved for S11 interface
	MsgTypeMBMSSessionStartRequest
	MsgTypeMBMSSessionStartResponse
	MsgTypeMBMSSessionUpdateRequest
	MsgTypeMBMSSessionUpdateResponse
	MsgTypeMBMSSessionStopRequest
	MsgTypeMBMSSessionStopResponse
	_
	_
	_ // 237-239: Reserved for Sm/Sn interface
	MsgTypeSRVCCCsToPsResponse
	MsgTypeSRVCCCsToPsCompleteNotification
	MsgTypeSRVCCCsToPsCompleteAcknowledge
	MsgTypeSRVCCCsToPsCancelNotification
	MsgTypeSRVCCCsToPsCancelAcknowledge
	_
	_
	_ // 245-247: Reserved for Sv interface
	_
	_
	_
	_
	_
	_
	_
	_ // 248-255: Reserved for others
)

// MsgTypeStr converts gtpv2 message type ID to more readable string representation.
func MsgTypeStr(msgType uint8) string {
	msgTypeStr := ""
	switch msgType {
	case MsgTypeEchoRequest:
		msgTypeStr = "EchoRequest"
	case MsgTypeEchoResponse:
		msgTypeStr = "EchoResponse"
	case MsgTypeVersionNotSupportedIndication:
		msgTypeStr = "VersionNotSupportedIndication"
	case MsgTypeDirectTransferRequest:
		msgTypeStr = "DirectTransferRequest"
	case MsgTypeDirectTransferResponse:
		msgTypeStr = "DirectTransferResponse"
	case MsgTypeNotificationRequest:
		msgTypeStr = "NotificationRequest"
	case MsgTypeNotificationResponse:
		msgTypeStr = "NotificationResponse"
	// 8-16: Reserved for S101 interface
	case MsgTypeRIMInformationTransfer:
		msgTypeStr = "RIMInformationTransfer"
	// 18-24: Reserved for S121 interface
	case MsgTypeSRVCCPsToCsRequest:
		msgTypeStr = "SRVCCPsToCsRequest"
	case MsgTypeSRVCCPsToCsResponse:
		msgTypeStr = "SRVCCPsToCsResponse"
	case MsgTypeSRVCCPsToCsCompleteNotification:
		msgTypeStr = "SRVCCPsToCsCompleteNotification"
	case MsgTypeSRVCCPsToCsCompleteAcknowledge:
		msgTypeStr = "SRVCCPsToCsCompleteAcknowledge"
	case MsgTypeSRVCCPsToCsCancelNotification:
		msgTypeStr = "SRVCCPsToCsCancelNotification"
	case MsgTypeSRVCCPsToCsCancelAcknowledge:
		msgTypeStr = "SRVCCPsToCsCancelAcknowledge"
	case MsgTypeSRVCCCsToPsRequest:
		msgTypeStr = "SRVCCCsToPsRequest"
	case MsgTypeCreateSessionRequest:
		msgTypeStr = "CreateSessionRequest"
	case MsgTypeCreateSessionResponse:
		msgTypeStr = "CreateSessionResponse"
	case MsgTypeModifyBearerRequest:
		msgTypeStr = "ModifyBearerRequest"
	case MsgTypeModifyBearerResponse:
		msgTypeStr = "ModifyBearerResponse"
	case MsgTypeDeleteSessionRequest:
		msgTypeStr = "ModifyBearerResponse"
	case MsgTypeDeleteSessionResponse:
		msgTypeStr = "DeleteSessionResponse"
	case MsgTypeChangeNotificationRequest:
		msgTypeStr = "ChangeNotificationRequest"
	case MsgTypeChangeNotificationResponse:
		msgTypeStr = "ChangeNotificationResponse"
	case MsgTypeRemoteUEReportNotification:
		msgTypeStr = "RemoteUEReportNotification"
	case MsgTypeRemoteUEReportAcknowledge:
		msgTypeStr = "RemoteUEReportAcknowledge"
	// 42-63: Reserved for S4/S11, S5/S8 interfaces
	case MsgTypeModifyBearerCommand:
		msgTypeStr = "ModifyBearerCommand"
	case MsgTypeModifyBearerFailureIndication:
		msgTypeStr = "ModifyBearerFailureIndication"
	case MsgTypeDeleteBearerCommand:
		msgTypeStr = "DeleteBearerCommand"
	case MsgTypeDeleteBearerFailureIndication:
		msgTypeStr = "DeleteBearerFailureIndication"
	case MsgTypeBearerResourceCommand:
		msgTypeStr = "BearerResourceCommand"
	case MsgTypeBearerResourceFailureIndication:
		msgTypeStr = "BearerResourceFailureIndication"
	case MsgTypeDownlinkDataNotificationFailureIndication:
		msgTypeStr = "DownlinkDataNotificationFailureIndication"
	case MsgTypeTraceSessionActivation:
		msgTypeStr = "TraceSessionActivation"
	case MsgTypeTraceSessionDeactivation:
		msgTypeStr = "TraceSessionDeactivation"
	case MsgTypeStopPagingIndication:
		msgTypeStr = "StopPagingIndication"
	// 74-94: Reserved for GTPv2 non-specific interfaces
	case MsgTypeCreateBearerRequest:
		msgTypeStr = "CreateBearerRequest"
	case MsgTypeCreateBearerResponse:
		msgTypeStr = "CreateBearerResponse"
	case MsgTypeUpdateBearerRequest:
		msgTypeStr = "UpdateBearerRequest"
	case MsgTypeUpdateBearerResponse:
		msgTypeStr = "UpdateBearerResponse"
	case MsgTypeDeleteBearerRequest:
		msgTypeStr = "DeleteBearerRequest"
	case MsgTypeDeleteBearerResponse:
		msgTypeStr = "DeleteBearerResponse"
	case MsgTypeDeletePDNConnectionSetRequest:
		msgTypeStr = "DeletePDNConnectionSetRequest"
	case MsgTypeDeletePDNConnectionSetResponse:
		msgTypeStr = "DeletePDNConnectionSetResponse"
	case MsgTypePGWDownlinkTriggeringNotification:
		msgTypeStr = "PGWDownlinkTriggeringNotification"
	case MsgTypePGWDownlinkTriggeringAcknowledge:
		msgTypeStr = "PGWDownlinkTriggeringAcknowledge"
	// 105-127: Reserved for S5, S4/S11 interfaces
	case MsgTypeIdentificationRequest:
		msgTypeStr = "IdentificationRequest"
	case MsgTypeIdentificationResponse:
		msgTypeStr = "IdentificationResponse"
	case MsgTypeContextRequest:
		msgTypeStr = "ContextRequest"
	case MsgTypeContextResponse:
		msgTypeStr = "ContextResponse"
	case MsgTypeContextAcknowledge:
		msgTypeStr = "ContextAcknowledge"
	case MsgTypeForwardRelocationRequest:
		msgTypeStr = "ForwardRelocationRequest"
	case MsgTypeForwardRelocationResponse:
		msgTypeStr = "ForwardRelocationResponse"
	case MsgTypeForwardRelocationCompleteNotification:
		msgTypeStr = "ForwardRelocationCompleteNotification"
	case MsgTypeForwardRelocationCompleteAcknowledge:
		msgTypeStr = "ForwardRelocationCompleteAcknowledge"
	case MsgTypeForwardAccessContextNotification:
		msgTypeStr = "ForwardAccessContextNotification"
	case MsgTypeForwardAccessContextAcknowledge:
		msgTypeStr = "ForwardAccessContextAcknowledge"
	case MsgTypeRelocationCancelRequest:
		msgTypeStr = "RelocationCancelRequest"
	case MsgTypeRelocationCancelResponse:
		msgTypeStr = "RelocationCancelResponse"
	case MsgTypeConfigurationTransferTunnel:
		msgTypeStr = "ConfigurationTransferTunnel"
	// 142-148: Reserved for S3/S10/S16 interfaces
	case MsgTypeDetachNotification:
		msgTypeStr = "DetachNotification"
	case MsgTypeDetachAcknowledge:
		msgTypeStr = "DetachAcknowledge"
	case MsgTypeCSPagingIndication:
		msgTypeStr = "CSPagingIndication"
	case MsgTypeRANInformationRelay:
		msgTypeStr = "RANInformationRelay"
	case MsgTypeAlertMMENotification:
		msgTypeStr = "AlertMMENotification"
	case MsgTypeAlertMMEAcknowledge:
		msgTypeStr = "AlertMMEAcknowledge"
	case MsgTypeUEActivityNotification:
		msgTypeStr = "UEActivityNotification"
	case MsgTypeUEActivityAcknowledge:
		msgTypeStr = "UEActivityAcknowledge"
	case MsgTypeISRStatusIndication:
		msgTypeStr = "ISRStatusIndication"
	case MsgTypeUERegistrationQueryRequest:
		msgTypeStr = "UERegistrationQueryRequest"
	case MsgTypeUERegistrationQueryResponse:
		msgTypeStr = "UERegistrationQueryResponse"
	case MsgTypeCreateForwardingTunnelRequest:
		msgTypeStr = "CreateForwardingTunnelRequest"
	case MsgTypeCreateForwardingTunnelResponse:
		msgTypeStr = "CreateForwardingTunnelResponse"
	case MsgTypeSuspendNotification:
		msgTypeStr = "SuspendNotification"
	case MsgTypeSuspendAcknowledge:
		msgTypeStr = "SuspendAcknowledge"
	case MsgTypeResumeNotification:
		msgTypeStr = "ResumeNotification"
	case MsgTypeResumeAcknowledge:
		msgTypeStr = "ResumeAcknowledge"
	case MsgTypeCreateIndirectDataForwardingTunnelRequest:
		msgTypeStr = "CreateIndirectDataForwardingTunnelRequest"
	case MsgTypeCreateIndirectDataForwardingTunnelResponse:
		msgTypeStr = "CreateIndirectDataForwardingTunnelResponse"
	case MsgTypeDeleteIndirectDataForwardingTunnelRequest:
		msgTypeStr = "DeleteIndirectDataForwardingTunnelRequest"
	case MsgTypeDeleteIndirectDataForwardingTunnelResponse:
		msgTypeStr = "DeleteIndirectDataForwardingTunnelResponse"
	case MsgTypeReleaseAccessBearersRequest:
		msgTypeStr = "ReleaseAccessBearersRequest"
	case MsgTypeReleaseAccessBearersResponse:
		msgTypeStr = "ReleaseAccessBearersResponse"
		// 172-175: Reserved for S4/S11 interfaces
	case MsgTypeDownlinkDataNotification:
		msgTypeStr = "DownlinkDataNotification"
	case MsgTypeDownlinkDataNotificationAcknowledge:
		msgTypeStr = "DownlinkDataNotificationAcknowledge"
	case MsgTypePGWRestartNotification:
		msgTypeStr = "PGWRestartNotification"
	case MsgTypePGWRestartNotificationAcknowledge:
		msgTypeStr = "PGWRestartNotificationAcknowledge"
		// 181-199: Reserved for S4 interface
	case MsgTypeUpdatePDNConnectionSetRequest:
		msgTypeStr = "UpdatePDNConnectionSetRequest"
	case MsgTypeUpdatePDNConnectionSetResponse:
		msgTypeStr = "UpdatePDNConnectionSetResponse"
		// 202-210: Reserved for S5/S8 interfaces
	case MsgTypeModifyAccessBearersRequest:
		msgTypeStr = "ModifyAccessBearersRequest"
	case MsgTypeModifyAccessBearersResponse:
		msgTypeStr = "ModifyAccessBearersResponse"
	// 213-230: Reserved for S11 interface
	case MsgTypeMBMSSessionStartRequest:
		msgTypeStr = "MBMSSessionStartRequest"
	case MsgTypeMBMSSessionStartResponse:
		msgTypeStr = "MBMSSessionStartResponse"
	case MsgTypeMBMSSessionUpdateRequest:
		msgTypeStr = "MBMSSessionUpdateRequest"
	case MsgTypeMBMSSessionUpdateResponse:
		msgTypeStr = "MBMSSessionUpdateResponse"
	case MsgTypeMBMSSessionStopRequest:
		msgTypeStr = "MBMSSessionStopRequest"
	case MsgTypeMBMSSessionStopResponse:
		msgTypeStr = "MBMSSessionStopResponse"
	// 237-239: Reserved for Sm/Sn interface
	case MsgTypeSRVCCCsToPsResponse:
		msgTypeStr = "SRVCCCsToPsResponse"
	case MsgTypeSRVCCCsToPsCompleteNotification:
		msgTypeStr = "SRVCCCsToPsCompleteNotification"
	case MsgTypeSRVCCCsToPsCompleteAcknowledge:
		msgTypeStr = "SRVCCCsToPsCompleteAcknowledge"
	case MsgTypeSRVCCCsToPsCancelNotification:
		msgTypeStr = "SRVCCCsToPsCancelNotification"
	case MsgTypeSRVCCCsToPsCancelAcknowledge:
		msgTypeStr = "SRVCCCsToPsCancelAcknowledge"
		// 245-247: Reserved for Sv interface
		// 248-255: Reserved for others
	default:
		msgTypeStr = strconv.FormatUint(uint64(msgType), 10)
	}
	return msgTypeStr
}

// Message is an interface that defines GTPv2 message.
type Message interface {
	MarshalTo([]byte) error
	UnmarshalBinary(b []byte) error
	MarshalLen() int
	Version() int
	MessageType() uint8
	MessageTypeName() string
	TEID() uint32
	SetTEID(uint32)
	Sequence() uint32
	SetSequenceNumber(uint32)

	// deprecated
	SerializeTo([]byte) error
	DecodeFromBytes(b []byte) error
}

// Marshal returns the byte sequence generated from a Message instance.
// Better to use MarshalXxx instead if you know the name of message to be serialized.
func Marshal(m Message) ([]byte, error) {
	b := make([]byte, m.MarshalLen())
	if err := m.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// Parse decodes the given bytes as Message.
func Parse(b []byte) (Message, error) {
	var m Message

	switch b[1] {
	case MsgTypeEchoRequest:
		m = &EchoRequest{}
	case MsgTypeEchoResponse:
		m = &EchoResponse{}
	case MsgTypeVersionNotSupportedIndication:
		m = &VersionNotSupportedIndication{}
	case MsgTypeCreateSessionRequest:
		m = &CreateSessionRequest{}
	case MsgTypeCreateSessionResponse:
		m = &CreateSessionResponse{}
	case MsgTypeDeleteSessionRequest:
		m = &DeleteSessionRequest{}
	case MsgTypeDeleteSessionResponse:
		m = &DeleteSessionResponse{}
	case MsgTypeModifyBearerCommand:
		m = &ModifyBearerCommand{}
	case MsgTypeModifyBearerFailureIndication:
		m = &ModifyBearerFailureIndication{}
	case MsgTypeDeleteBearerCommand:
		m = &DeleteBearerCommand{}
	case MsgTypeDeleteBearerFailureIndication:
		m = &DeleteBearerFailureIndication{}
	case MsgTypeDeleteBearerRequest:
		m = &DeleteBearerRequest{}
	case MsgTypeCreateBearerRequest:
		m = &CreateBearerRequest{}
	case MsgTypeCreateBearerResponse:
		m = &CreateBearerResponse{}
	case MsgTypeDeleteBearerResponse:
		m = &DeleteBearerResponse{}
	case MsgTypeModifyBearerRequest:
		m = &ModifyBearerRequest{}
	case MsgTypeModifyBearerResponse:
		m = &ModifyBearerResponse{}
	case MsgTypeContextRequest:
		m = &ContextRequest{}
	case MsgTypeContextResponse:
		m = &ContextResponse{}
	case MsgTypeContextAcknowledge:
		m = &ContextAcknowledge{}
	case MsgTypeReleaseAccessBearersRequest:
		m = &ReleaseAccessBearersRequest{}
	case MsgTypeReleaseAccessBearersResponse:
		m = &ReleaseAccessBearersResponse{}
	case MsgTypeStopPagingIndication:
		m = &StopPagingIndication{}
	case MsgTypeModifyAccessBearersRequest:
		m = &ModifyAccessBearersRequest{}
	case MsgTypeModifyAccessBearersResponse:
		m = &ModifyAccessBearersResponse{}
	case MsgTypeDeletePDNConnectionSetRequest:
		m = &DeletePDNConnectionSetRequest{}
	case MsgTypeDeletePDNConnectionSetResponse:
		m = &DeletePDNConnectionSetResponse{}
	case MsgTypeUpdatePDNConnectionSetRequest:
		m = &UpdatePDNConnectionSetRequest{}
	case MsgTypeUpdatePDNConnectionSetResponse:
		m = &UpdatePDNConnectionSetResponse{}
	case MsgTypePGWRestartNotification:
		m = &PGWRestartNotification{}
	case MsgTypePGWRestartNotificationAcknowledge:
		m = &PGWRestartNotificationAcknowledge{}
	case MsgTypeDetachNotification:
		m = &DetachNotification{}
	case MsgTypeDetachAcknowledge:
		m = &DetachAcknowledge{}
	case MsgTypeDownlinkDataNotification:
		m = &DownlinkDataNotification{}
	case MsgTypeDownlinkDataNotificationAcknowledge:
		m = &DownlinkDataNotificationAcknowledge{}
	case MsgTypeDownlinkDataNotificationFailureIndication:
		m = &DownlinkDataNotificationFailureIndication{}
	default:
		m = &Generic{}
	}

	if err := m.UnmarshalBinary(b); err != nil {
		return nil, fmt.Errorf("failed to decode GTPv2 Message: %w", err)
	}
	return m, nil
}
