package resourceManager

import "fmt"

// ReleaseAlreadyUsedError describes when a resource is already available.
// (i.e. this shouldn't have been in someone's possession to begin with!)
type ReleaseAlreadyUsedError struct {
  resource int
  available int
  capacity int
}

func (err ReleaseAlreadyUsedError) Error() string {
  return fmt.Sprintf(
    "Released item %v was already available! (Capacity %d/%d)\n",
    err.resource, err.available, err.capacity,
  )
}
