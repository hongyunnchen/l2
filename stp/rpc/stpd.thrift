namespace go stpd
typedef i32 int
typedef i16 uint16
struct Dot1dBaseDot1dBasePortEntry {
	1 : i32 Dot1dBaseType
	2 : i32 Dot1dBasePortIfIndex
	3 : i32 Dot1dBasePort
	4 : i32 Dot1dBasePortDelayExceededDiscards
	5 : i32 Dot1dBasePortMtuExceededDiscards
	6 : string Dot1dBaseBridgeAddress
	7 : i32 Dot1dBaseNumPorts
	8 : string Dot1dBasePortCircuit
}
struct Dot1dBaseDot1dBasePortEntryGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<Dot1dBaseDot1dBasePortEntry> Dot1dBaseDot1dBasePortEntryList
}
struct Dot1dTpDot1dTpPortEntry {
	1 : i32 Dot1dTpPortInFrames
	2 : i32 Dot1dTpPortMaxInfo
	3 : i32 Dot1dTpPortOutFrames
	4 : i32 Dot1dTpAgingTime
	5 : i32 Dot1dTpPort
	6 : i32 Dot1dTpPortInDiscards
	7 : i32 Dot1dTpLearnedEntryDiscards
}
struct Dot1dTpDot1dTpPortEntryGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<Dot1dTpDot1dTpPortEntry> Dot1dTpDot1dTpPortEntryList
}
struct Dot1dStaticDot1dStaticEntry {
	1 : string Dot1dStaticAddress
	2 : i32 Dot1dStaticReceivePort
	3 : string Dot1dStaticAllowedToGoTo
	4 : i32 Dot1dStaticStatus
}
struct Dot1dStpPortEntryStateCountersFsmStatesPortTimer {
	1 : i32 Dot1dStpPortPriority
	2 : string Dot1dStpPortDesignatedBridge
	3 : i64 TcInPkts
	4 : string PrxmPrevState
	5 : i64 StpOutPkts
	6 : i32 Dot1dStpBridgePortForwardDelay
	7 : string Dot1dStpPortDesignatedRoot
	8 : i32 BpduGuardDetected
	9 : i32 BpduGuard
	10 : i32 Dot1dStpPortAdminPathCost
	11 : string PstmCurrState
	12 : string PtimPrevState
	13 : string PrxmCurrState
	14 : i64 PvstInPkts
	15 : string PtimCurrState
	16 : i32 Dot1dStpPortOperPointToPoint
	17 : i32 Dot1dStpPortPathCost32
	18 : i32 EdgeDelayWhile
	19 : string PpmCurrState
	20 : string PtxmPrevState
	21 : i32 BridgeAssurance
	22 : i32 Dot1dStpPortForwardTransitions
	23 : i64 TcOutPkts
	24 : i32 Dot1dStpPort
	25 : string PrtmPrevState
	26 : string PimPrevState
	27 : string PrtmCurrState
	28 : i32 TcWhile
	29 : i32 HelloWhen
	30 : i64 PvstOutPkts
	31 : string TcmCurrState
	32 : i64 BpduInPkts
	33 : string PtxmCurrState
	34 : i32 Dot1dStpPortState
	35 : i32 BaWhile
	36 : string PstmPrevState
	37 : i32 MdelayWhile
	38 : i32 Dot1dStpBridgePortMaxAge
	39 : i64 BpduOutPkts
	40 : i32 BridgeAssuranceInconsistant
	41 : i32 RrWhile
	42 : i64 StpInPkts
	43 : i32 Dot1dBrgIfIndex
	44 : string PimCurrState
	45 : i32 Dot1dStpPortOperEdgePort
	46 : i32 Dot1dStpPortAdminPointToPoint
	47 : i32 Dot1dStpPortPathCost
	48 : i32 Dot1dStpPortEnable
	49 : string BdmPrevState
	50 : i32 RbWhile
	51 : string Dot1dStpPortDesignatedPort
	52 : i32 Dot1dStpPortAdminEdgePort
	53 : i32 RcvdInfoWhile
	54 : i64 RstpOutPkts
	55 : string PpmPrevState
	56 : string BdmCurrState
	57 : i32 Dot1dStpPortDesignatedCost
	58 : i32 Dot1dStpPortProtocolMigration
	59 : string TcmPrevState
	60 : i32 BpduGuardInterval
	61 : i32 Dot1dStpBridgePortHelloTime
	62 : i32 FdWhile
	63 : i64 RstpInPkts
}
struct Dot1dStpPortEntryStateCountersFsmStatesPortTimerGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<Dot1dStpPortEntryStateCountersFsmStatesPortTimer> Dot1dStpPortEntryStateCountersFsmStatesPortTimerList
}
struct Dot1dStpPortEntryConfig {
	1 : i32 Dot1dBrgIfIndex
	2 : i32 Dot1dStpPortPriority
	3 : i32 Dot1dStpPortAdminPathCost
	4 : i32 Dot1dStpPortProtocolMigration
	5 : i32 Dot1dStpPortAdminEdgePort
	6 : i32 BridgeAssurance
	7 : i32 Dot1dStpPortAdminPointToPoint
	8 : i32 BpduGuard
	9 : i32 Dot1dStpPort
	10 : i32 Dot1dStpPortPathCost
	11 : i32 BpduGuardInterval
	12 : i32 Dot1dStpPortEnable
	13 : i32 Dot1dStpPortPathCost32
}
struct Dot1dStpBridgeState {
	1 : i32 Dot1dBrgIfIndex
	2 : string Dot1dStpDesignatedRoot
	3 : i32 Dot1dStpBridgeForceVersion
	4 : i32 Dot1dStpRootCost
	5 : string Dot1dBridgeAddress
	6 : i32 Dot1dStpBridgeHelloTime
	7 : i32 Dot1dStpForwardDelay
	8 : i32 Dot1dStpHelloTime
	9 : i32 Dot1dStpProtocolSpecification
	10 : i32 Dot1dStpMaxAge
	11 : i32 Dot1dStpBridgeMaxAge
	12 : i32 Dot1dStpBridgeTxHoldCount
	13 : i32 Dot1dStpTimeSinceTopologyChange
	14 : i32 Dot1dStpRootPort
	15 : i32 Dot1dStpTopChanges
	16 : i32 Dot1dStpBridgeForwardDelay
	17 : i32 Dot1dStpPriority
	18 : i16 Dot1dStpVlan
	19 : i32 Dot1dStpHoldTime
}
struct Dot1dStpBridgeStateGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<Dot1dStpBridgeState> Dot1dStpBridgeStateList
}
struct Dot1dTpDot1dTpFdbEntry {
	1 : string Dot1dTpFdbAddress
	2 : i32 Dot1dTpFdbPort
	3 : i32 Dot1dTpLearnedEntryDiscards
	4 : i32 Dot1dTpFdbStatus
	5 : i32 Dot1dTpAgingTime
}
struct Dot1dTpDot1dTpFdbEntryGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<Dot1dTpDot1dTpFdbEntry> Dot1dTpDot1dTpFdbEntryList
}
struct Dot1dStpBridgeConfig {
	1 : i32 Dot1dStpBridgeForceVersion
	2 : string Dot1dBridgeAddress
	3 : i32 Dot1dStpBridgeHelloTime
	4 : i32 Dot1dStpBridgeMaxAge
	5 : i32 Dot1dStpBridgeTxHoldCount
	6 : i32 Dot1dStpBridgeForwardDelay
	7 : i32 Dot1dStpPriority
	8 : i16 Dot1dStpVlan
}
service STPDServices {
	Dot1dBaseDot1dBasePortEntryGetInfo GetBulkDot1dBaseDot1dBasePortEntry(1: int fromIndex, 2: int count);
	Dot1dTpDot1dTpPortEntryGetInfo GetBulkDot1dTpDot1dTpPortEntry(1: int fromIndex, 2: int count);
	bool CreateDot1dStaticDot1dStaticEntry(1: Dot1dStaticDot1dStaticEntry config);
	bool UpdateDot1dStaticDot1dStaticEntry(1: Dot1dStaticDot1dStaticEntry origconfig, 2: Dot1dStaticDot1dStaticEntry newconfig, 3: list<bool> attrset);
	bool DeleteDot1dStaticDot1dStaticEntry(1: Dot1dStaticDot1dStaticEntry config);

	Dot1dStpPortEntryStateCountersFsmStatesPortTimerGetInfo GetBulkDot1dStpPortEntryStateCountersFsmStatesPortTimer(1: int fromIndex, 2: int count);
	bool CreateDot1dStpPortEntryConfig(1: Dot1dStpPortEntryConfig config);
	bool UpdateDot1dStpPortEntryConfig(1: Dot1dStpPortEntryConfig origconfig, 2: Dot1dStpPortEntryConfig newconfig, 3: list<bool> attrset);
	bool DeleteDot1dStpPortEntryConfig(1: Dot1dStpPortEntryConfig config);

	Dot1dStpBridgeStateGetInfo GetBulkDot1dStpBridgeState(1: int fromIndex, 2: int count);
	Dot1dTpDot1dTpFdbEntryGetInfo GetBulkDot1dTpDot1dTpFdbEntry(1: int fromIndex, 2: int count);
	bool CreateDot1dStpBridgeConfig(1: Dot1dStpBridgeConfig config);
	bool UpdateDot1dStpBridgeConfig(1: Dot1dStpBridgeConfig origconfig, 2: Dot1dStpBridgeConfig newconfig, 3: list<bool> attrset);
	bool DeleteDot1dStpBridgeConfig(1: Dot1dStpBridgeConfig config);

}