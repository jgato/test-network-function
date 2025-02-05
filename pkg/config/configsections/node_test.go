// Copyright (C) 2020-2021 Red Hat, Inc.
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, write to the Free Software Foundation, Inc.,
// 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

package configsections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMaster(t *testing.T) {
	testCases := []struct {
		nodeLabel      string
		expectedOutput bool
	}{
		{
			expectedOutput: true,
			nodeLabel:      MasterLabel,
		},
		{
			expectedOutput: false,
			nodeLabel:      WorkerLabel,
		},
		{
			expectedOutput: false,
			nodeLabel:      "",
		},
	}

	for _, tc := range testCases {
		testNode := Node{
			Name:   "mastertest",
			Labels: []string{tc.nodeLabel},
		}
		assert.Equal(t, tc.expectedOutput, testNode.IsMaster())
	}
}

func TestIsWorker(t *testing.T) {
	testCases := []struct {
		nodeLabel      string
		expectedOutput bool
	}{
		{
			expectedOutput: true,
			nodeLabel:      WorkerLabel,
		},
		{
			expectedOutput: false,
			nodeLabel:      MasterLabel,
		},
		{
			expectedOutput: false,
			nodeLabel:      "",
		},
	}

	for _, tc := range testCases {
		testNode := Node{
			Name:   "workertest",
			Labels: []string{tc.nodeLabel},
		}
		assert.Equal(t, tc.expectedOutput, testNode.IsWorker())
	}
}
