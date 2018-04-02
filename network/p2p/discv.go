// Copyright 2018 The go-hpb Authors
// This file is part of the go-hpb.
//
// The go-hpb is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-hpb is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-hpb. If not, see <http://www.gnu.org/licenses/>.

package p2p

import (
	"runtime/debug"

	"github.com/hpb-project/ghpb/network/p2p/discover"
	"github.com/hpb-project/ghpb/common/log"
)

// NodeType is node type used in peer management
type NodeType uint8
const(
	NtUnknown  NodeType = 0x00
	NtPublic   NodeType = 0x10
	NtCommitt  NodeType = 0x20
	NtPrecomm  NodeType = 0x30
	NtAccess   NodeType = 0x40
	NtLight    NodeType = 0x50
)
// NodeType to string
func (pt NodeType)String() string {
	switch pt {
	case NtUnknown:
		return "unknown-node"
	case NtPublic:
		return "public-node"
	case NtCommitt:
		return "committee-node"
	case NtPrecomm:
		return "pre-committee-node"
	case NtAccess:
		return "access-node"
	case NtLight:
		return "light-node"
	}
	return "unknown-node"
}

func (pt NodeType)Uint8() uint8 {
	return uint8(pt)
}

// Convert NodeType to discover role
func (pt NodeType)ToDiscv() uint8 {
	disc := discover.UnKnowRole
	switch pt {
	case NtUnknown:
		disc = discover.UnKnowRole
	case NtPublic:
		disc = discover.BootRole
	case NtCommitt:
		disc = discover.HpRole
	case NtPrecomm:
		disc = discover.PreRole
	case NtAccess:
		disc = discover.AccessRole
	case NtLight:
		disc = discover.LightRole
	}
	return disc
}

// Convert uint8 to NodeType
func Uint8ToNodeType(discNt uint8) NodeType {
	switch discNt {
	case discover.BootRole:
		return NtPublic
	case discover.HpRole:
		return NtCommitt
	case discover.PreRole:
		return NtPrecomm
	case discover.AccessRole:
		return NtAccess
	case discover.LightRole:
		return NtLight
	}
	log.Debug("NodeType unknown ","uint8",discNt, "stack",debug.Stack())
	return NtUnknown
}

// Convert uint8 to NodeType
func StrToNodeType(role string) NodeType {
	switch {
	case role == "public":
		return NtPublic
	case role == "committee":
		return NtCommitt
	case role == "pre-committee":
		return NtPrecomm
	case role == "access":
		return NtAccess
	case role == "light":
		return NtLight
	}
	//log.Debug("NodeType unknown ","role",role, "stack",debug.Stack())
	return NtUnknown
}



