// Copyright 2023 The klaytn Authors
// This file is part of the klaytn library.
//
// The klaytn library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The klaytn library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the klaytn library. If not, see <http://www.gnu.org/licenses/>.

package bls

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerify(t *testing.T) {
	sk, _ := RandKey()
	pk := sk.PublicKey()

	msg := make([]byte, 32)
	sig := Sign(sk, msg)

	assert.True(t, Verify(sig, msg, pk))
}

func TestAggregateVerify(t *testing.T) {
	sk1, _ := RandKey()
	sk2, _ := RandKey()

	pkb1 := sk1.PublicKey().Marshal()
	pkb2 := sk2.PublicKey().Marshal()

	msg := make([]byte, 32)
	sigb1 := Sign(sk1, msg).Marshal()
	sigb2 := Sign(sk2, msg).Marshal()

	apk, _ := AggregatePublicKeysFromBytes([][]byte{pkb1, pkb2})
	asig, _ := AggregateSignaturesFromBytes([][]byte{sigb1, sigb2})
	assert.True(t, Verify(asig, msg, apk))
}
