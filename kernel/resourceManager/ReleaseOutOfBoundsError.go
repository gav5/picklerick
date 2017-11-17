package resourceManager

import "fmt"

// ReleaseOutOfBoundsError describes when a resource is out of bounds.
// (i.e. what's given back is negative or greater than or equal to capacity)
type ReleaseOutOfBoundsError struct {
  resource int
  available int
  capacity int
}

func (err ReleaseOutOfBoundsError) Error() string {
  return fmt.Sprintf(
    "Released item %v is out of bounds! (Capacity %d/%d)\n",
    err.resource, err.available, err.capacity,
  )
}
