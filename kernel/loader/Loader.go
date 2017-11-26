package loader

import (
  "../program"
)

// Load loads files at the given path and returns an array of programs.
func Load(path string) ([]program.Program, error) {
  // read from the file
  content, err := readFile(path)
  if err != nil {
    return nil, err
  }
  prgAry, err := parse(content)
  if err != nil {
    return nil, err
  }
  return prgAry, nil
}
