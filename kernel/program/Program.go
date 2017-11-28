package program

import (
  "regexp"

  "../../util"
)

// Program describes a parsed program from the program file.
type Program struct {
  JobID uint8
  NumberOfWords uint8
  PriorityNumber uint8
  Instructions []uint32
  InputBufferSize uint8
  OutputBufferSize uint8
  TempBufferSize uint8
  DataBlock []uint32
}

// Make builds a program from the given match data.
func Make(matchData []string) (Program, error) {
  a := Program{}
  var err error

  // parse JobID
  a.JobID, err = util.HexExtract8(matchData[1])
  if err != nil {
    return a, err
  }

  // parse NumberOfWords
  a.NumberOfWords, err = util.HexExtract8(matchData[2])
  if err != nil {
    return a, err
  }

  // parse PriorityNumber
  a.PriorityNumber, err = util.HexExtract8(matchData[3])
  if err != nil {
    return a, err
  }

  // parse Instructions
  inRe := regexp.MustCompile("[[:xdigit:]]{8}")
  splitIn := inRe.FindAllString(matchData[4], -1)
  a.Instructions, err = util.HexExtractArrayFixed32(splitIn)
  if err != nil {
    return a, err
  }

  // parse InputBufferSize
  a.InputBufferSize, err = util.HexExtract8(matchData[5])
  if err != nil {
    return a, err
  }

  // parse OutputBufferSize
  a.OutputBufferSize, err = util.HexExtract8(matchData[6])
  if err != nil {
    return a, err
  }

  // parse TempBufferSize
  a.TempBufferSize, err = util.HexExtract8(matchData[7])
  if err != nil {
    return a, err
  }

  // parse dataBlock
  blRe := regexp.MustCompile("[[:xdigit:]]{8}")
  splitBlock := blRe.FindAllString(matchData[8], -1)
  a.DataBlock, err = util.HexExtractArrayFixed32(splitBlock)
  if err != nil {
    return a, err
  }

  return a, nil
}

// Sleep makes a sleep program.
// This program has only one NOP instruction.
func Sleep() Program {
  return Program{
    JobID: 0x00,
    NumberOfWords: 1,
    PriorityNumber: 0,
    InputBufferSize: 0,
    OutputBufferSize: 0,
    TempBufferSize: 0,
    DataBlock: []uint32{},
    Instructions: []uint32{0x13000000},
  }
}
