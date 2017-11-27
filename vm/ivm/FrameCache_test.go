package ivm

import "testing"

var cache1 = FrameCache{
  0x00: Frame{0xC050004C, 0x4B060000, 0x4B000000, 0x4B010000},
  0x01: Frame{0x4B020000, 0x4B030001, 0x4F07009C, 0xC1270000},
  0x02: Frame{0x4C070004, 0x4C060001, 0x05320000, 0xC1070000},
  0x03: Frame{0x4C070004, 0x4C060001, 0x04230000, 0x04300000},
  0x04: Frame{0x10658000, 0x56810028, 0x92000000, 0x00000000},
  0x05: Frame{0x0000000A, 0x00000000, 0x00000000, 0x00000000},
  0x06: Frame{0x00000000, 0x00000000, 0x00000000, 0x00000000},
  0x07: Frame{0x00000000, 0x00000000, 0x00000000, 0x00000000},
  0x08: Frame{0x00000000, 0x00000000, 0x00000000, 0x00000000},
  0x13: Frame{0x00000000, 0x00000000, 0x00000000, 0x00000000},
}

var frameCacheAddressFetchWordTests = []struct{
  addr Address
  word Word
  fc FrameCache
}{
  {addr: 0x0000, word: 0xC050004C, fc: cache1},
  {addr: 0x0004, word: 0x4B060000, fc: cache1},
  {addr: 0x0008, word: 0x4B000000, fc: cache1},
  {addr: 0x000C, word: 0x4B010000, fc: cache1},
  {addr: 0x0010, word: 0x4B020000, fc: cache1},
  {addr: 0x0014, word: 0x4B030001, fc: cache1},
  {addr: 0x0018, word: 0x4F07009C, fc: cache1},
  {addr: 0x001C, word: 0xC1270000, fc: cache1},
  {addr: 0x0020, word: 0x4C070004, fc: cache1},
  {addr: 0x0024, word: 0x4C060001, fc: cache1},
  {addr: 0x0028, word: 0x05320000, fc: cache1},
  {addr: 0x002C, word: 0xC1070000, fc: cache1},
  {addr: 0x0030, word: 0x4C070004, fc: cache1},
  {addr: 0x0034, word: 0x4C060001, fc: cache1},
  {addr: 0x0038, word: 0x04230000, fc: cache1},
  {addr: 0x003C, word: 0x04300000, fc: cache1},
  {addr: 0x0040, word: 0x10658000, fc: cache1},
  {addr: 0x0044, word: 0x56810028, fc: cache1},
  {addr: 0x0048, word: 0x92000000, fc: cache1},
}

func TestFrameCacheAddressFetchWord(t *testing.T) {
  for _, tt := range frameCacheAddressFetchWordTests {
    word := tt.fc.AddressFetchWord(tt.addr)
    if word != tt.word {
      t.Errorf(
        "FrameCache.AddressFetchWord(%v) => %v, expected %v",
        tt.addr, word, tt.word,
      )
    }
  }
}
