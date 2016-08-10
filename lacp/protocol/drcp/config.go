//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

// config.go
package drcp

import (
	"fmt"
	//"sync"
	"errors"
	"l2/lacp/protocol/lacp"
	"l2/lacp/protocol/utils"
	"net"
)

const (
	DRNI_PORTAL_SYSTEM_ID_MIN = 1
	DRNI_PORTAL_SYSTEM_ID_MAX = 2 // only support two portal system
)

const DRCPConfigModuleStr = "DRCP Config"

// 802.1.AX-2014 7.4.1.1 Distributed Relay Attributes GET-SET
type DistrubtedRelayConfig struct {
	// GET-SET
	DrniName                               string
	DrniPortalAddress                      string
	DrniPortalPriority                     uint16
	DrniThreePortalSystem                  bool
	DrniPortalSystemNumber                 uint8
	DrniIntraPortalLinkList                [3]uint32
	DrniAggregator                         uint32
	DrniConvAdminGateway                   [4096][3]uint8
	DrniNeighborAdminConvGatewayListDigest [16]uint8
	DrniNeighborAdminConvPortListDigest    [16]uint8
	DrniGatewayAlgorithm                   string
	DrniNeighborAdminGatewayAlgorithm      string
	DrniNeighborAdminPortAlgorithm         string
	DrniNeighborAdminDRCPState             string
	DrniEncapMethod                        string
	DrniIPLEncapMap                        [16]uint32
	DrniNetEncapMap                        [16]uint32
	DrniPortConversationControl            bool
	DrniIntraPortalPortProtocolDA          string
}

// Conversations are typically related to the various service types to which
// traffic is associated with.  If portList is empty is is assumed to be Gateway
// Algorithm, otherwise it is a Port Algorithm
// 802.1 AX-2014 8.1
// Therefore, a Conversation Identifier (or Conversation ID) is defined as a value in the range 0 through 4095.
// By administrative means, every possible conversation is assigned to a single Conversation ID value for each
// supported Conversation ID type. More than one conversation can be assigned to a Conversation ID. It is not
// necessary that every Conversation ID value have any conversations assigned to it. In this standard, several
// types of Conversation ID are specified for different uses.
type DRConversationConfig struct {
	DrniName   string
	Idtype     GatewayAlgorithm
	Isid       uint32
	Cvlan      uint16
	Svlan      uint16
	Bvid       uint16
	Psuedowire uint32
	PortList   []int32
}

func (d *DistrubtedRelayConfig) GetKey() string {
	return d.DrniName
}

// DistrubtedRelayConfigParamCheck will validate the config from the user after it has
// been translated to something the Lacp module expects.  Thus if translation
// layer fails it should produce an invalid value.  The error returned
// will be translated to model values
func DistrubtedRelayConfigParamCheck(mlag *DistrubtedRelayConfig) error {

	_, err := net.ParseMAC(mlag.DrniPortalAddress)
	if err != nil {
		return errors.New(fmt.Sprintln("ERROR Portal System MAC Supplied must be in the format of 00:00:00:00:00:00 rcvd:", mlag.DrniPortalAddress))
	}

	linkcnt := 0
	for _, ippid := range mlag.DrniIntraPortalLinkList {
		if ippid > 0 {
			if _, ok := utils.PortConfigMap[int32(ippid)]; !ok {
				return errors.New(fmt.Sprintln("ERROR Invalid Intra Portal Link Port Id supplied", ippid, utils.PortConfigMap))
			}
		} else {
			linkcnt++
		}
	}
	if linkcnt == 3 {
		return errors.New("ERROR Invalid Intra Portal Link Port required")
	}

	if mlag.DrniThreePortalSystem {
		return errors.New(fmt.Sprintln("ERROR Only support a 2 Portal System"))
	}

	if mlag.DrniPortalSystemNumber < DRNI_PORTAL_SYSTEM_ID_MIN ||
		mlag.DrniPortalSystemNumber > DRNI_PORTAL_SYSTEM_ID_MAX {
		return errors.New(fmt.Sprintln("ERROR Invalid Portal System Number must be between 1 and ", DRNI_PORTAL_SYSTEM_ID_MAX))
	}

	validPortGatewayAlgorithms := map[string]bool{
		"00:80:C2:01": true,
		"00:80:C2:02": true,
		"00:80:C2:03": true,
		"00:80:C2:04": true,
		"00:80:C2:05": true,
	}

	if _, ok := validPortGatewayAlgorithms[mlag.DrniGatewayAlgorithm]; !ok {
		return errors.New(fmt.Sprintln("ERROR Invalid Gateway Algorithm supplied must be in the format 00:80:C2:XX where XX is 1-5 the value of the algorithm ", mlag.DrniGatewayAlgorithm))
	}

	if _, ok := validPortGatewayAlgorithms[mlag.DrniNeighborAdminGatewayAlgorithm]; !ok {
		return errors.New(fmt.Sprintln("ERROR Invalid Neighbor Gateway Algorithm supplied must be in the format 00:80:C2:XX where XX is 1-5 the value of the algorithm ", mlag.DrniNeighborAdminGatewayAlgorithm))
	}

	if _, ok := validPortGatewayAlgorithms[mlag.DrniNeighborAdminPortAlgorithm]; !ok {
		return errors.New(fmt.Sprintln("ERROR Invalid Neighbor Port Algorithm supplied must be in the format 00:80:C2:XX where XX is 1-5 the value of the algorithm ", mlag.DrniNeighborAdminPortAlgorithm))
	}

	validEncapStrings := map[string]bool{
		"00:80:C2:00": true, // seperate physical or lag link
		"00:80:C2:01": true, // shared by time
		"00:80:C2:02": true, // shared by tag
	}

	if _, ok := validEncapStrings[mlag.DrniEncapMethod]; !ok {
		return errors.New(fmt.Sprintln("ERROR Invalid Encap Method supplied must be in the format 00:80:C2:XX where XX is 0-2 the value of the encap method ", mlag.DrniEncapMethod))
	}

	if mlag.DrniPortConversationControl {
		return errors.New(fmt.Sprintln("ERROR Invalid Port Conversation Control is always false as the Home Gateway Vector is controlled by protocol ", mlag.DrniPortConversationControl))
	}

	_, err = net.ParseMAC(mlag.DrniIntraPortalPortProtocolDA)
	if err != nil {
		return errors.New(fmt.Sprintln("ERROR Invalid Port Protocol DA invalid format must be 00:00:00:00:00:00 rcvd: ", mlag.DrniIntraPortalPortProtocolDA))
	}

	// only going to support this address
	if mlag.DrniIntraPortalPortProtocolDA != "01:80:C2:00:00:03" {
		return errors.New(fmt.Sprintln("ERROR Invalid Port Protocol DA only support 01:80:C2:00:00:03 rcvd: ", mlag.DrniIntraPortalPortProtocolDA))
	}

	return nil
}

// CreateDistributedRelay will create the distributed relay then attach
// the Aggregator to the Distributed Relay
func CreateDistributedRelay(cfg *DistrubtedRelayConfig) {

	dr := NewDistributedRelay(cfg)
	// aggregator must exist for the protocol to really make sense
	if dr != nil {
		AttachAggregatorToDistributedRelay(int(dr.DrniAggregator))
	}
}

// DeleteDistributedRelay will detach the distributed relay from the aggregator
// and delete the distributed relay instance
func DeleteDistributedRelay(name string) {

	dr, ok := DistributedRelayDB[name]
	if ok {
		DetachAggregatorFromDistributedRelay(int(dr.DrniAggregator))
		dr.DeleteDistributedRelay()
	}
}

// AttachAggregatorToDistributedRelay: will attach the aggregator and start the Distributed
// relay protocol for the given dr if this agg is associated with a DR
func AttachAggregatorToDistributedRelay(aggId int) {
	for _, dr := range DistributedRelayDBList {
		if dr.DrniAggregator == int32(aggId) &&
			dr.a == nil {
			var a *lacp.LaAggregator
			if lacp.LaFindAggById(aggId, &a) {
				dr.a = a

				// set port and gateway info and digest
				dr.setTimeSharingPortAndGatwewayDigest()

				// lets update the aggregator parameters
				// configured ports
				for _, pId := range a.PortNumList {
					var p *lacp.LaAggPort
					if lacp.LaFindPortById(pId, &p) {
						dr.PrevAggregatorId = p.ActorAdmin.System.Actor_System
						dr.PrevAggregatorPriority = p.ActorAdmin.System.Actor_System_priority
						// assign the new values to the aggregator
						lacp.SetLaAggPortSystemInfo(
							uint16(pId),
							fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
								dr.DrniPortalAddr[0],
								dr.DrniPortalAddr[1],
								dr.DrniPortalAddr[2],
								dr.DrniPortalAddr[3],
								dr.DrniPortalAddr[4],
								dr.DrniPortalAddr[5]),
							dr.DrniPortalPriority)
					}
				}

				dr.BEGIN(false)
				// start the IPP links
				for _, ipp := range dr.Ipplinks {
					ipp.DRCPEnabled = true
					ipp.BEGIN(false)
				}
			}
		}
	}
}

// DetachCreatedAggregatorFromDistributedRelay: will detach the aggregator and stop the Distributed
// relay protocol for the given dr if since this aggregator is no longer attached
func DetachAggregatorFromDistributedRelay(aggId int) {
	for _, dr := range DistributedRelayDBList {
		if dr.DrniAggregator == int32(aggId) &&
			dr.a != nil {
			var a *lacp.LaAggregator
			if lacp.LaFindAggById(aggId, &a) {

				// lets update the aggregator parameters
				// configured ports
				for _, pId := range a.PortNumList {
					lacp.SetLaAggPortSystemInfo(
						uint16(pId),
						fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
							dr.PrevAggregatorId[0],
							dr.PrevAggregatorId[1],
							dr.PrevAggregatorId[2],
							dr.PrevAggregatorId[3],
							dr.PrevAggregatorId[4],
							dr.PrevAggregatorId[5]),
						dr.PrevAggregatorPriority)
				}
			}
			dr.a = nil
		}
	}
}
