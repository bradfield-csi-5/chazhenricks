package main

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("check magic number", func(t *testing.T) {

		got, gotErr := LoadData()
		if gotErr != nil {
			t.Errorf("error reading file %v", gotErr)
		}
		magic := got[:4]
		wantHex := "d4c3b2a1"
		want, _ := hex.DecodeString(wantHex)

		if !bytes.Equal(magic, want) {
			t.Errorf("got %x - want %x", magic, want)
		}
	})
}

func TestParseGlobalHeader(t *testing.T) {
	data, _ := LoadData()
	got := ParseGlobalHeader(data)

	t.Run("Check magic number", func(t *testing.T) {

		var want uint32 = 0xa1b2c3d4

		if got.Magic != want {
			t.Errorf("got %x - want %x", got, want)
		}
	})
	t.Run("Check major/minor version", func(t *testing.T) {
		var major uint16 = 2
		var minor uint16 = 4

		if got.MajorV != major {
			t.Errorf("got %x - want %x", got.MajorV, major)
		}
		if got.MinorV != minor {
			t.Errorf("got %x - want %x", got.MinorV, minor)
		}
	})
	t.Run("Test Zero Values", func(t *testing.T) {
		if got.TimeZoneOffset != 0 {
			t.Errorf("got %x - want %x", got.TimeZoneOffset, 0)
		}
		if got.Accuracy != 0 {
			t.Errorf("got %x - want %x", got.Accuracy, 0)
		}
	})

	t.Run("Test LinkLayer/Snapshot", func(t *testing.T) {
		if got.LinkLayerHeaderType != 1 {
			t.Errorf("got %x - want %x", got.LinkLayerHeaderType, 1)
		}
		if got.SnapshotLength != 1514 {
			t.Errorf("got %d - want %d", got.SnapshotLength, 1514)
		}
	})
}

func TestCountPackets(t *testing.T) {
	data, _ := LoadData()

  got := CountPackets(data[24:])
	want := 99

	if got != want {
		t.Errorf("got %d - want %d", got, want)
	}
}

