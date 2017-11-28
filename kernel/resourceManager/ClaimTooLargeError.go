package resourceManager

import "fmt"

// ClaimTooLargeError describes where the claim is too large.
type ClaimTooLargeError struct {
  claimSize int
  available int
  capacity int
}

func (err ClaimTooLargeError) Error() string {
  return fmt.Sprintf(
    "Claim of %d is too large! (Capacity %d/%d)\n",
    err.claimSize, err.available, err.capacity,
  )
}
