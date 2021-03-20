// Copyright 2019-2021 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

/*
Package ie provides encoding/decoding feature of GTPv2 Information Elements.
*/
package ie

import (
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
)

// IE definitions.
const (
	_ uint8 = iota
	IMSI
	Cause
	Recovery
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
	_
	_
	_
	_
	_
	_
	_
	_
	_ // 4-34: Reserved for S101 interface
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
	_ // 35-50:  Reserved for S101 interface
	STNSR
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
	_ // 52-70: Reserved for Sv interface
	AccessPointName
	AggregateMaximumBitRate
	EPSBearerID
	IPAddress
	MobileEquipmentIdentity
	MSISDN
	Indication
	ProtocolConfigurationOptions
	PDNAddressAllocation
	BearerQoS
	FlowQoS
	RATType
	ServingNetwork
	BearerTFT
	TrafficAggregateDescription
	UserLocationInformation
	FullyQualifiedTEID
	TMSI
	GlobalCNID
	S103PDNDataForwardingInfo
	S1UDataForwarding
	DelayValue
	BearerContext
	ChargingID
	ChargingCharacteristics
	TraceInformation
	BearerFlags
	_
	PDNType
	ProcedureTransactionID
	_
	_
	MMContextGSMKeyAndTriplets
	MMContextUMTSKeyUsedCipherAndQuintuplets
	MMContextGSMKeyUsedCipherAndQuintuplets
	MMContextUMTSKeyAndQuintuplets
	MMContextEPSSecurityContextQuadrupletsAndQuintuplets
	MMContextUMTSKeyQuadrupletsAndQuintuplets
	PDNConnection
	PDUNumbers
	PacketTMSI
	PTMSISignature
	HopCounter
	UETimeZone
	TraceReference
	CompleteRequestMessage
	GUTI
	FContainer
	FCause
	PLMNID
	TargetIdentification
	_
	PacketFlowID
	RABContext
	SourceRNCPDCPContextInfo
	PortNumber
	APNRestriction
	SelectionMode
	SourceIdentification
	Reserved
	ChangeReportingAction
	FullyQualifiedCSID
	ChannelNeeded
	EMLPPPriority
	NodeType
	FullyQualifiedDomainName
	TI
	MBMSSessionDuration
	MBMSServiceArea
	MBMSSessionIdentifier
	MBMSFlowIdentifier
	MBMSIPMulticastDistribution
	MBMSDistributionAcknowledge
	RFSPIndex
	UserCSGInformation
	CSGInformationReportingAction
	CSGID
	CSGMembershipIndication
	ServiceIndicator
	DetachType
	LocalDistinguishedName
	NodeFeatures
	MBMSTimeToDataTransfer
	Throttling
	AllocationRetensionPriority
	EPCTimer
	SignallingPriorityIndication
	TMGI
	AdditionalMMContextForSRVCC
	AdditionalFlagsForSRVCC
	_
	MDTConfiguration
	AdditionalProtocolConfigurationOptions
	AbsoluteTimeofMBMSDataTransfer
	HeNBInformationReporting
	IPv4ConfigurationParameters
	ChangeToReportFlags
	ActionIndication
	TWANIdentifier
	ULITimestamp
	MBMSFlags
	RANNASCause
	CNOperatorSelectionEntity
	TrustedWLANModeIndication
	NodeNumber
	NodeIdentifier
	PresenceReportingAreaAction
	PresenceReportingAreaInformation
	TWANIdentifierTimestamp
	OverloadControlInformation
	LoadControlInformation
	Metric
	SequenceNumber
	APNAndRelativeCapacity
	WLANOffloadabilityIndication
	PagingAndServiceInformation
	IntegerNumber
	MillisecondTimeStamp
	MonitoringEventInformation
	ECGIList
	RemoteUEContext
	RemoteUserID
	RemoteUEIPinformation
	CIoTOptimizationsSupportIndication
	SCEFPDNConnection
	HeaderCompressionConfiguration
	ExtendedProtocolConfigurationOptions
	ServingPLMNRateControl
	Counter
	MappedUEUsageType
	SecondaryRATUsageDataReport
	UPFunctionSelectionIndicationFlags
	MaximumPacketLossRate
	APNRateControlStatus
	ExtendedTraceInformation
	MonitoringEventExtensionInformation
	AdditionalRRMPolicyIndex
	V2XContext
	PC5QoSParameters
	ServicesAuthorized
	BitRate
	PC5QoSFlow
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
	_ // 206-253: Spare for future use
	SpecialIETypeForIETypeExtension
	PrivateExtension
)

// IE is a GTPv2 Information Element.
type IE struct {
	Type     uint8
	Length   uint16
	instance uint8
	Payload  []byte
	ChildIEs []*IE
}

// New creates new IE.
func New(itype, ins uint8, data []byte) *IE {
	ie := &IE{
		Type:     itype,
		instance: ins & 0x0f,
		Payload:  data,
	}
	ie.SetLength()

	return ie
}

// SetInstance sets the instance.
func (i *IE) SetInstance(ins uint8) {
	i.instance = ins & 0x0f
}

// WithInstance sets the instance and returns IE.
func (i *IE) WithInstance(ins uint8) *IE {
	i.instance = ins & 0x0f
	return i
}

// Instance returns instance value in uint8
func (i *IE) Instance() uint8 {
	return i.instance & 0x0f
}

// Marshal returns the byte sequence generated from an IE instance.
func (i *IE) Marshal() ([]byte, error) {
	b := make([]byte, i.MarshalLen())
	if err := i.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (i *IE) MarshalTo(b []byte) error {
	l := len(b)
	if l < 4 {
		return io.ErrUnexpectedEOF
	}

	b[0] = i.Type
	binary.BigEndian.PutUint16(b[1:3], i.Length)
	b[3] = i.instance
	if i.IsGrouped() {
		offset := 4
		for _, ie := range i.ChildIEs {
			if l < offset+ie.MarshalLen() {
				break
			}

			if err := ie.MarshalTo(b[offset : offset+ie.MarshalLen()]); err != nil {
				return err
			}
			offset += ie.MarshalLen()
		}
		return nil
	}

	if l < i.MarshalLen() {
		return io.ErrUnexpectedEOF
	}

	copy(b[4:i.MarshalLen()], i.Payload)
	return nil
}

// Parse decodes given byte sequence as a GTPv2 Information Element.
func Parse(b []byte) (*IE, error) {
	ie := &IE{}
	if err := ie.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return ie, nil
}

// UnmarshalBinary sets the values retrieved from byte sequence in GTPv2 IE.
func (i *IE) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 5 {
		return io.ErrUnexpectedEOF
	}

	i.Type = b[0]
	i.Length = binary.BigEndian.Uint16(b[1:3])
	if int(i.Length) > l-4 {
		return ErrInvalidLength
	}

	i.instance = b[3]
	i.Payload = b[4 : 4+int(i.Length)]

	if i.IsGrouped() {
		var err error
		i.ChildIEs, err = ParseMultiIEs(i.Payload)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalLen returns field length in integer.
func (i *IE) MarshalLen() int {
	if i.IsGrouped() {
		l := 4
		for _, ie := range i.ChildIEs {
			l += ie.MarshalLen()
		}
		return l
	}
	return 4 + len(i.Payload)
}

// SetLength sets the length in Length field.
func (i *IE) SetLength() {
	if i.IsGrouped() {
		l := 0
		for _, ie := range i.ChildIEs {
			l += ie.MarshalLen()
		}
		i.Length = uint16(l)
	}
	i.Length = uint16(len(i.Payload))
}

// String returns the GTPv2 IE values in human readable format.
func (i *IE) String() string {
	return fmt.Sprintf("{Type: %d (%s), Length: %d, Instance: %#x, Payload: %#v}",
		i.Type,
		IETypeStr(i.Type),
		i.Length,
		i.Instance(),
		i.Payload,
	)
}

var grouped = []uint8{
	BearerContext,
	PDNConnection,
	OverloadControlInformation,
	LoadControlInformation,
	RemoteUEContext,
	SCEFPDNConnection,
	V2XContext,
	PC5QoSParameters,
}

// IsGrouped reports whether an IE is grouped type or not.
func (i *IE) IsGrouped() bool {
	for _, itype := range grouped {
		if i.Type == itype {
			return true
		}
	}
	return false
}

// Add adds variable number of IEs to a IE if the IE is grouped type and update length.
// Otherwise, this does nothing(no errors).
func (i *IE) Add(ies ...*IE) {
	if !i.IsGrouped() {
		return
	}

	i.Payload = nil
	i.ChildIEs = append(i.ChildIEs, ies...)
	for _, ie := range i.ChildIEs {
		serialized, err := ie.Marshal()
		if err != nil {
			continue
		}
		i.Payload = append(i.Payload, serialized...)
	}
	i.SetLength()
}

// Remove removes an IE looked up by type and instance.
func (i *IE) Remove(typ, instance uint8) {
	if !i.IsGrouped() {
		return
	}

	i.Payload = nil
	newChildren := make([]*IE, len(i.ChildIEs))
	idx := 0
	for _, ie := range i.ChildIEs {
		if ie.Type == typ && ie.Instance() == instance {
			newChildren = newChildren[:len(newChildren)-1]
			continue
		}
		newChildren[idx] = ie
		idx++

		serialized, err := ie.Marshal()
		if err != nil {
			continue
		}
		i.Payload = append(i.Payload, serialized...)
	}
	i.ChildIEs = newChildren
	i.SetLength()
}

// FindByType returns IE looked up by type and instance.
//
// The program may be slower when calling this method multiple times
// because this ranges over a ChildIEs each time it is called.
func (i *IE) FindByType(typ, instance uint8) (*IE, error) {
	if !i.IsGrouped() {
		return nil, ErrInvalidType
	}

	for _, ie := range i.ChildIEs {
		if ie.Type == typ && ie.Instance() == instance {
			return ie, nil
		}
	}
	return nil, ErrIENotFound
}

// ParseMultiIEs decodes multiple IEs at a time.
// This is easy and useful but slower than decoding one by one.
// When you don't know the number of IEs, this is the only way to decode them.
// See benchmarks in diameter_test.go for the detail.
func ParseMultiIEs(b []byte) ([]*IE, error) {
	var ies []*IE
	for {
		if len(b) == 0 {
			break
		}

		i, err := Parse(b)
		if err != nil {
			return nil, err
		}
		ies = append(ies, i)
		b = b[i.MarshalLen():]
	}
	return ies, nil
}

func newUint8ValIE(t, v uint8) *IE {
	return New(t, 0x00, []byte{v})
}

func newUint16ValIE(t uint8, v uint16) *IE {
	i := New(t, 0x00, make([]byte, 2))
	binary.BigEndian.PutUint16(i.Payload, v)
	return i
}

func newUint32ValIE(t uint8, v uint32) *IE {
	i := New(t, 0x00, make([]byte, 4))
	binary.BigEndian.PutUint32(i.Payload, v)
	return i
}

// unused for now.
// func newUint64ValIE(t uint8, v uint64) *IE {
// 	i := New(t, 0x00, make([]byte, 8))
// 	binary.BigEndian.PutUint64(i.Payload, v)
// 	return i
// }

func newStringIE(t uint8, v string) *IE {
	return New(t, 0x00, []byte(v))
}

func newGroupedIE(itype uint8, ies ...*IE) *IE {
	i := New(itype, 0x00, make([]byte, 0))
	i.ChildIEs = ies
	for _, ie := range i.ChildIEs {
		serialized, err := ie.Marshal()
		if err != nil {
			return nil
		}
		i.Payload = append(i.Payload, serialized...)
	}
	i.SetLength()

	return i
}

// IETypeStr returns string representation of passed GTPv2 IE type
func IETypeStr(ieType uint8) string {
	typeStr := ""
	switch ieType {
	case IMSI:
		typeStr = "IMSI"
	case Cause:
		typeStr = "Cause"
	case Recovery:
		typeStr = "Recovery"
	// 4-34: Reserved for S101 interface
	// 35-50:  Reserved for S101 interface
	case STNSR:
		typeStr = "STNSR"
	// 52-70: Reserved for Sv interface
	case AccessPointName:
		typeStr = "AccessPointName"
	case AggregateMaximumBitRate:
		typeStr = "AggregateMaximumBitRate"
	case EPSBearerID:
		typeStr = "EPSBearerID"
	case IPAddress:
		typeStr = "IPAddress"
	case MobileEquipmentIdentity:
		typeStr = "MobileEquipmentIdentity"
	case MSISDN:
		typeStr = "MSISDN"
	case Indication:
		typeStr = "Indication"
	case ProtocolConfigurationOptions:
		typeStr = "ProtocolConfigurationOptions"
	case PDNAddressAllocation:
		typeStr = "PDNAddressAllocation"
	case BearerQoS:
		typeStr = "BearerQoS"
	case FlowQoS:
		typeStr = "FlowQoS"
	case RATType:
		typeStr = "RATType"
	case ServingNetwork:
		typeStr = "ServingNetwork"
	case BearerTFT:
		typeStr = "BearerTFT"
	case TrafficAggregateDescription:
		typeStr = "TrafficAggregateDescription"
	case UserLocationInformation:
		typeStr = "UserLocationInformation"
	case FullyQualifiedTEID:
		typeStr = "FullyQualifiedTEID"
	case TMSI:
		typeStr = "TMSI"
	case GlobalCNID:
		typeStr = "GlobalCNID"
	case S103PDNDataForwardingInfo:
		typeStr = "S103PDNDataForwardingInfo"
	case S1UDataForwarding:
		typeStr = "S1UDataForwarding"
	case DelayValue:
		typeStr = "DelayValue"
	case BearerContext:
		typeStr = "BearerContext"
	case ChargingID:
		typeStr = "ChargingID"
	case ChargingCharacteristics:
		typeStr = "ChargingCharacteristics"
	case TraceInformation:
		typeStr = "TraceInformation"
	case BearerFlags:
		typeStr = "BearerFlags"
	case PDNType:
		typeStr = "PDNType"
	case ProcedureTransactionID:
		typeStr = "ProcedureTransactionID"
	case MMContextGSMKeyAndTriplets:
		typeStr = "MMContextGSMKeyAndTriplets"
	case MMContextUMTSKeyUsedCipherAndQuintuplets:
		typeStr = "MMContextUMTSKeyUsedCipherAndQuintuplets"
	case MMContextGSMKeyUsedCipherAndQuintuplets:
		typeStr = "MMContextGSMKeyUsedCipherAndQuintuplets"
	case MMContextUMTSKeyAndQuintuplets:
		typeStr = "MMContextUMTSKeyAndQuintuplets"
	case MMContextEPSSecurityContextQuadrupletsAndQuintuplets:
		typeStr = "MMContextEPSSecurityContextQuadrupletsAndQuintuplets"
	case MMContextUMTSKeyQuadrupletsAndQuintuplets:
		typeStr = "MMContextUMTSKeyQuadrupletsAndQuintuplets"
	case PDNConnection:
		typeStr = "PDNConnection"
	case PDUNumbers:
		typeStr = "PDUNumbers"
	case PacketTMSI:
		typeStr = "PacketTMSI"
	case PTMSISignature:
		typeStr = "PTMSISignature"
	case HopCounter:
		typeStr = "HopCounter"
	case UETimeZone:
		typeStr = "UETimeZone"
	case TraceReference:
		typeStr = "TraceReference"
	case CompleteRequestMessage:
		typeStr = "CompleteRequestMessage"
	case GUTI:
		typeStr = "GUTI"
	case FContainer:
		typeStr = "FContainer"
	case FCause:
		typeStr = "FCause"
	case PLMNID:
		typeStr = "PLMNID"
	case TargetIdentification:
		typeStr = "TargetIdentification"
	case PacketFlowID:
		typeStr = "PacketFlowID"
	case RABContext:
		typeStr = "RABContext"
	case SourceRNCPDCPContextInfo:
		typeStr = "SourceRNCPDCPContextInfo"
	case PortNumber:
		typeStr = "PortNumber"
	case APNRestriction:
		typeStr = "APNRestriction"
	case SelectionMode:
		typeStr = "SelectionMode"
	case SourceIdentification:
		typeStr = "SourceIdentification"
	case Reserved:
		typeStr = "Reserved"
	case ChangeReportingAction:
		typeStr = "ChangeReportingAction"
	case FullyQualifiedCSID:
		typeStr = "FullyQualifiedCSID"
	case ChannelNeeded:
		typeStr = "ChannelNeeded"
	case EMLPPPriority:
		typeStr = "EMLPPPriority"
	case NodeType:
		typeStr = "NodeType"
	case FullyQualifiedDomainName:
		typeStr = "FullyQualifiedDomainName"
	case TI:
		typeStr = "TI"
	case MBMSSessionDuration:
		typeStr = "MBMSSessionDuration"
	case MBMSServiceArea:
		typeStr = "MBMSServiceArea"
	case MBMSSessionIdentifier:
		typeStr = "MBMSSessionIdentifier"
	case MBMSFlowIdentifier:
		typeStr = "MBMSFlowIdentifier"
	case MBMSIPMulticastDistribution:
		typeStr = "MBMSIPMulticastDistribution"
	case MBMSDistributionAcknowledge:
		typeStr = "MBMSDistributionAcknowledge"
	case RFSPIndex:
		typeStr = "RFSPIndex"
	case UserCSGInformation:
		typeStr = "UserCSGInformation"
	case CSGInformationReportingAction:
		typeStr = "CSGInformationReportingAction"
	case CSGID:
		typeStr = "CSGID"
	case CSGMembershipIndication:
		typeStr = "CSGMembershipIndication"
	case ServiceIndicator:
		typeStr = "ServiceIndicator"
	case DetachType:
		typeStr = "DetachType"
	case LocalDistinguishedName:
		typeStr = "LocalDistinguishedName"
	case NodeFeatures:
		typeStr = "NodeFeatures"
	case MBMSTimeToDataTransfer:
		typeStr = "MBMSTimeToDataTransfer"
	case Throttling:
		typeStr = "Throttling"
	case AllocationRetensionPriority:
		typeStr = "AllocationRetensionPriority"
	case EPCTimer:
		typeStr = "EPCTimer"
	case SignallingPriorityIndication:
		typeStr = "SignallingPriorityIndication"
	case TMGI:
		typeStr = "TMGI"
	case AdditionalMMContextForSRVCC:
		typeStr = "AdditionalMMContextForSRVCC"
	case AdditionalFlagsForSRVCC:
		typeStr = "AdditionalFlagsForSRVCC"
	case MDTConfiguration:
		typeStr = "MDTConfiguration"
	case AdditionalProtocolConfigurationOptions:
		typeStr = "AdditionalProtocolConfigurationOptions"
	case AbsoluteTimeofMBMSDataTransfer:
		typeStr = "AbsoluteTimeofMBMSDataTransfer"
	case HeNBInformationReporting:
		typeStr = "HeNBInformationReporting"
	case IPv4ConfigurationParameters:
		typeStr = "IPv4ConfigurationParameters"
	case ChangeToReportFlags:
		typeStr = "ChangeToReportFlags"
	case ActionIndication:
		typeStr = "ActionIndication"
	case TWANIdentifier:
		typeStr = "TWANIdentifier"
	case ULITimestamp:
		typeStr = "ULITimestamp"
	case MBMSFlags:
		typeStr = "MBMSFlags"
	case RANNASCause:
		typeStr = "RANNASCause"
	case CNOperatorSelectionEntity:
		typeStr = "CNOperatorSelectionEntity"
	case TrustedWLANModeIndication:
		typeStr = "TrustedWLANModeIndication"
	case NodeNumber:
		typeStr = "NodeNumber"
	case NodeIdentifier:
		typeStr = "NodeIdentifier"
	case PresenceReportingAreaAction:
		typeStr = "PresenceReportingAreaAction"
	case PresenceReportingAreaInformation:
		typeStr = "PresenceReportingAreaInformation"
	case TWANIdentifierTimestamp:
		typeStr = "TWANIdentifierTimestamp"
	case OverloadControlInformation:
		typeStr = "OverloadControlInformation"
	case LoadControlInformation:
		typeStr = "LoadControlInformation"
	case Metric:
		typeStr = "Metric"
	case SequenceNumber:
		typeStr = "SequenceNumber"
	case APNAndRelativeCapacity:
		typeStr = "APNAndRelativeCapacity"
	case WLANOffloadabilityIndication:
		typeStr = "WLANOffloadabilityIndication"
	case PagingAndServiceInformation:
		typeStr = "PagingAndServiceInformation"
	case IntegerNumber:
		typeStr = "IntegerNumber"
	case MillisecondTimeStamp:
		typeStr = "MillisecondTimeStamp"
	case MonitoringEventInformation:
		typeStr = "MonitoringEventInformation"
	case ECGIList:
		typeStr = "ECGIList"
	case RemoteUEContext:
		typeStr = "RemoteUEContext"
	case RemoteUserID:
		typeStr = "RemoteUserID"
	case RemoteUEIPinformation:
		typeStr = "RemoteUEIPinformation"
	case CIoTOptimizationsSupportIndication:
		typeStr = "CIoTOptimizationsSupportIndication"
	case SCEFPDNConnection:
		typeStr = "SCEFPDNConnection"
	case HeaderCompressionConfiguration:
		typeStr = "HeaderCompressionConfiguration"
	case ExtendedProtocolConfigurationOptions:
		typeStr = "ExtendedProtocolConfigurationOptions"
	case ServingPLMNRateControl:
		typeStr = "ServingPLMNRateControl"
	case Counter:
		typeStr = "Counter"
	case MappedUEUsageType:
		typeStr = "MappedUEUsageType"
	case SecondaryRATUsageDataReport:
		typeStr = "SecondaryRATUsageDataReport"
	case UPFunctionSelectionIndicationFlags:
		typeStr = "UPFunctionSelectionIndicationFlags"
	case MaximumPacketLossRate:
		typeStr = "MaximumPacketLossRate"
	case APNRateControlStatus:
		typeStr = "APNRateControlStatus"
	case ExtendedTraceInformation:
		typeStr = "ExtendedTraceInformation"
	case MonitoringEventExtensionInformation:
		typeStr = "MonitoringEventExtensionInformation"
	case AdditionalRRMPolicyIndex:
		typeStr = "AdditionalRRMPolicyIndex"
	case V2XContext:
		typeStr = "V2XContext"
	case PC5QoSParameters:
		typeStr = "PC5QoSParameters"
	case ServicesAuthorized:
		typeStr = "ServicesAuthorized"
	case BitRate:
		typeStr = "BitRate"
	case PC5QoSFlow:
		typeStr = "PC5QoSFlow"
	// 206-253: Spare for future use
	// 206-253: Spare for future use
	case SpecialIETypeForIETypeExtension:
		typeStr = "SpecialIETypeForIETypeExtension"
	case PrivateExtension:
		typeStr = "PrivateExtension"
	default:
		typeStr = strconv.FormatUint(uint64(ieType), 10)
	}
	return typeStr
}
