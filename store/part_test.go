package store

import (
	"bytes"
	"testing"

	"github.com/intob/rocketkv/util"
)

func getTestPart(blocks int) Part {
	partId, _ := util.RandomId()
	p := NewPart(partId)
	for i := 0; i < blocks; i++ {
		// make block with random ID
		blockId, _ := util.RandomId()
		block := NewBlock(blockId)
		// fill with blocks*256 random slots
		for s := 0; s < blocks*256; s++ {
			slotId, _ := util.RandomId()
			block.Slots[getName(slotId)] = Slot{Value: slotId}
		}
		p.Blocks[getNumber(blockId)] = block
	}
	return p
}

// Tests that calling getClosestBlock always returns
// the same block.
func TestGetClosestBlock(t *testing.T) {
	part := getTestPart(16)
	keyHash := util.HashKey("test")
	clCtl := part.getClosestBlock(keyHash)
	for i := 0; i < len(part.Blocks); i++ {
		clCur := part.getClosestBlock(keyHash)
		if !bytes.Equal(clCtl.Id, clCur.Id) {
			t.FailNow()
		}
	}

}

func BenchmarkGetClosestBlock(b *testing.B) {
	part := getTestPart(16)
	keyHash := util.HashKey("test")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		part.getClosestBlock(keyHash)
	}
}

func BenchmarkListKeys(b *testing.B) {
	part := getTestPart(16)
	out := make(chan string)
	go func() {
		for {
			<-out
		}
	}()
	for i := 0; i < b.N; i++ {
		part.listKeys("", out)
	}
}