package loader

import (
  "io/ioutil"
  "regexp"

  "../program"
)

const contentRegexp = "(?mi)// *job ([[:xdigit:]]{1,2}) ([[:xdigit:]]{1,2}) ([[:xdigit:]]{1,2})[\n\r]+((?:0x[[:xdigit:]]{8} *[\n\r]+)+)// *data ([[:xdigit:]]{1,2}) ([[:xdigit:]]{1,2}) ([[:xdigit:]]{1,2})[\n\r]+((?:0x[[:xdigit:]]{8} *[\n\r]+){44})// *end"

func readFile(path string) (string, error) {
  data, err := ioutil.ReadFile(path)
  if err != nil {
    return "", err
  }
  return string(data), nil
}

func parse(content string) ([]program.Program, error) {
  fileRe := regexp.MustCompile(contentRegexp)
  matches := fileRe.FindAllStringSubmatch(content, -1)
  parsedContent := make([]program.Program, len(matches))
  for index, m := range matches {
    a, err := program.Make(m)
    if err != nil {
      return nil, err
    }
    parsedContent[index] = a
  }
  return parsedContent, nil
}
