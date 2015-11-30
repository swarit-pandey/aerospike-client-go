// Copyright 2013-2015 Aerospike, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aerospike

import (
	// . "github.com/aerospike/aerospike-client-go/logger"

	. "github.com/aerospike/aerospike-client-go/types"
	Buffer "github.com/aerospike/aerospike-client-go/utils/buffer"
)

type queryRecordCommand struct {
	*queryCommand
}

func newQueryRecordCommand(node *Node, policy *QueryPolicy, statement *Statement, recordset *Recordset) *queryRecordCommand {
	return &queryRecordCommand{
		queryCommand: newQueryCommand(node, policy, statement, recordset),
	}
}

func (cmd *queryRecordCommand) parseRecordResults(ifc command, receiveSize int) (bool, error) {
	// Read/parse remaining message bytes one record at a time.
	cmd.dataOffset = 0

	for cmd.dataOffset < receiveSize {
		if err := cmd.readBytes(int(_MSG_REMAINING_HEADER_SIZE)); err != nil {
			cmd.recordset.Errors <- newNodeError(cmd.node, err)
			return false, err
		}
		resultCode := ResultCode(cmd.dataBuffer[5] & 0xFF)

		if resultCode != 0 {
			if resultCode == KEY_NOT_FOUND_ERROR {
				// consume the rest of the input buffer from the socket
				if cmd.dataOffset < receiveSize {
					if err := cmd.readBytes(receiveSize - cmd.dataOffset); err != nil {
						return false, err
					}
				}
				return false, nil
			}
			err := NewAerospikeError(resultCode)
			cmd.recordset.Errors <- newNodeError(cmd.node, err)
			return false, err
		}

		info3 := int(cmd.dataBuffer[3])

		// If cmd is the end marker of the response, do not proceed further
		if (info3 & _INFO3_LAST) == _INFO3_LAST {
			return false, nil
		}

		generation := Buffer.BytesToUint32(cmd.dataBuffer, 6)
		expiration := TTL(Buffer.BytesToUint32(cmd.dataBuffer, 10))
		fieldCount := int(Buffer.BytesToUint16(cmd.dataBuffer, 18))
		opCount := int(Buffer.BytesToUint16(cmd.dataBuffer, 20))

		key, err := cmd.parseKey(fieldCount)
		if err != nil {
			cmd.recordset.Errors <- newNodeError(cmd.node, err)
			return false, err
		}

		// Parse bins.
		var bins BinMap

		for i := 0; i < opCount; i++ {
			if err := cmd.readBytes(8); err != nil {
				cmd.recordset.Errors <- newNodeError(cmd.node, err)
				return false, err
			}

			opSize := int(Buffer.BytesToUint32(cmd.dataBuffer, 0))
			particleType := int(cmd.dataBuffer[5])
			nameSize := int(cmd.dataBuffer[7])

			if err := cmd.readBytes(nameSize); err != nil {
				cmd.recordset.Errors <- newNodeError(cmd.node, err)
				return false, err
			}
			name := string(cmd.dataBuffer[:nameSize])

			particleBytesSize := int((opSize - (4 + nameSize)))
			if err = cmd.readBytes(particleBytesSize); err != nil {
				cmd.recordset.Errors <- newNodeError(cmd.node, err)
				return false, err
			}
			value, err := bytesToParticle(particleType, cmd.dataBuffer, 0, particleBytesSize)
			if err != nil {
				cmd.recordset.Errors <- newNodeError(cmd.node, err)
				return false, err
			}

			if bins == nil {
				bins = make(BinMap, opCount)
			}
			bins[name] = value
		}

		// If the channel is full and it blocks, we don't want this command to
		// block forever, or panic in case the channel is closed in the meantime.
		select {
		// send back the result on the async channel
		case cmd.recordset.Records <- newRecord(cmd.node, key, bins, generation, expiration):
		case <-cmd.recordset.cancelled:
			return false, NewAerospikeError(SCAN_TERMINATED)
		}
	}

	return true, nil
}

func (cmd *queryRecordCommand) Execute() error {
	defer cmd.recordset.signalEnd()
	return cmd.execute(cmd)
}
