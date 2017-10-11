package prog

import (
	"io/ioutil"
	"regexp"

	"../util"
)

// ParseFile parses a file and returns a Program
func ParseFile(filename string) ([]Program, error) {
	content, fileErr := ioutil.ReadFile(filename)
	if fileErr != nil {
		return nil, fileErr
	}
	return parseContent(string(content))
}

func parseContent(content string) ([]Program, error) {
	matches := matchesForContent(content)
	programs := make([]Program, len(matches))
	for index, m := range matches {
		p, perr := parseProgram(m)
		if perr != nil {
			return []Program{}, perr
		}
		programs[index] = p
	}
	return programs, nil
}

func matchesForContent(content string) [][]string {
	fileRe := regexp.MustCompile("(?mi)// *job ([[:xdigit:]]{1,2}) ([[:xdigit:]]{1,2}) ([[:xdigit:]]{1,2})[\n\r]+((?:0x[[:xdigit:]]{8} *[\n\r]+)+)// *data ([[:xdigit:]]{1,2}) ([[:xdigit:]]{1,2}) ([[:xdigit:]]{1,2})[\n\r]+((?:0x[[:xdigit:]]{8} *[\n\r]+)+)// *end")
	matches := fileRe.FindAllStringSubmatch(content, -1)
	return matches
}

func parseProgram(matchData []string) (Program, error) {
	j, jerr := parseJob(matchData)
	if jerr != nil {
		return Program{}, jerr
	}
	d, derr := parseData(matchData)
	if derr != nil {
		return Program{}, derr
	}
	return Program{Job: j, Data: d}, nil
}

func parseJob(matchData []string) (Job, error) {
	id, idErr := util.HexExtract8(matchData[1])
	if idErr != nil {
		return Job{}, idErr
	}
	nw, nwErr := util.HexExtract8(matchData[2])
	if nwErr != nil {
		return Job{}, nwErr
	}
	pn, pnErr := util.HexExtract8(matchData[3])
	if pnErr != nil {
		return Job{}, pnErr
	}
	inRe := regexp.MustCompile("[[:xdigit:]]{8}")
	splitIn := inRe.FindAllString(matchData[4], -1)
	in, inErr := util.HexExtractArrayFixed32(splitIn)
	if inErr != nil {
		return Job{}, inErr
	}
	return Job{
		ID:             id,
		NumberOfWords:  nw,
		PriorityNumber: pn,
		Instructions:   in,
	}, nil
}

func parseData(matchData []string) (Data, error) {
	inSize, inSizeErr := util.HexExtract8(matchData[5])
	if inSizeErr != nil {
		return Data{}, inSizeErr
	}
	outSize, outSizeErr := util.HexExtract8(matchData[6])
	if outSizeErr != nil {
		return Data{}, outSizeErr
	}
	tempSize, tempSizeErr := util.HexExtract8(matchData[7])
	if tempSizeErr != nil {
		return Data{}, tempSizeErr
	}
	blRe := regexp.MustCompile("[[:xdigit:]]{8}")
	splitBlock := blRe.FindAllString(matchData[8], -1)
	dataBlock, dataBlockErr := util.HexExtractArrayFixed32(splitBlock)
	if dataBlockErr != nil {
		return Data{}, dataBlockErr
	}
	return Data{
		InputBufferSize:  inSize,
		OutputBufferSize: outSize,
		TempBufferSize:   tempSize,
		DataBlock:        dataBlock,
	}, nil
}
