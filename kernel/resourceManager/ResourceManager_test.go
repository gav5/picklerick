package resourceManager

import (
  "testing"
  "fmt"
)

var newTests = []struct{
  capacity int
}{
  {capacity: 4},
  {capacity: 1024},
  {capacity: 2048},
}

func TestNew(t *testing.T) {
  for _, tt := range newTests {
    rm := New(tt.capacity)

    // make sure it returned something!
    if rm == nil {
      t.Errorf("New(%d) => nil\n", tt.capacity)
    }

    // make sure the available resources are correct
    availableLen := rm.available.Len()
    if availableLen != tt.capacity {
      t.Errorf(
        "New(%d) => %d resources, expected %d resources\n",
        tt.capacity, availableLen, tt.capacity,
      )
    }
    for r := 0; r < tt.capacity; r++ {
      // make sure all resources are available
      count := 0
      for e := rm.available.Front(); e != nil; e = e.Next() {
        if r == e.Value.(int) {
          count++
        }
      }
      if count < 1 {
        t.Errorf(
          "New(%d) did not make resource %d available\n",
          tt.capacity, r,
        )
      } else if count > 1 {
        t.Errorf(
          "New(%d) made %d instances of resource %d available\n",
          tt.capacity, count, r,
        )
      }
    }
  }
}

var claimTests = []struct{
  capacity int
  taken int
  quantity int
  resourceCount int
  err error
}{
  {
    capacity: 4, taken: 0, quantity: 4,
    resourceCount: 4, err: nil,
  },
  {
    capacity: 4, taken: 2, quantity: 2,
    resourceCount: 2, err: nil,
  },
  {
    capacity: 4, taken: 1, quantity: 4,
    resourceCount: 0, err: ClaimTooLargeError{
      claimSize: 4,
      available: 3,
      capacity: 4,
    },
  },
}

func TestClaim(t *testing.T) {
  for _, tt := range claimTests {
    rm := simulatedManager(tt.capacity, tt.taken)
    // make the call to Claim for the given quantity
    resources, err := rm.Claim(tt.quantity)
    // used to identify the test case
    claimStr := fmt.Sprintf(
      "Claim(%d) [cap %d/%d]",
      tt.quantity, tt.taken, tt.capacity,
    )
    // ensure the resources have the right quantity
    resourceCount := len(resources)
    if resourceCount != tt.resourceCount {
      t.Errorf(
        "%s:\nwant %d resources\ngot %d resources: %v\n",
        claimStr, tt.resourceCount, resourceCount, resources,
      )
    }
    // ensure the right error was returned
    if err != tt.err {
      t.Errorf(
        "%s:\nwant error %v\ngot error %v\n",
        claimStr, tt.err, err,
      )
    }
  }
}

var releaseTests = []struct{
  capacity int
  taken int
  resources []int
  quantity int
  err error
}{
  {
    capacity: 4, taken: 4, resources: []int{0},
    quantity: 1, err: nil,
  },
  {
    capacity: 4, taken: 1, resources: []int{0},
    quantity: 1, err: nil,
  },
  {
    capacity: 4, taken: 4, resources: []int{4},
    quantity: 0, err: ReleaseOutOfBoundsError{
      resource: 4,
      available: 0,
      capacity: 4,
    },
  },
  {
    capacity: 4, taken: 4, resources: []int{-1},
    quantity: 0, err: ReleaseOutOfBoundsError{
      resource: -1,
      available: 0,
      capacity: 4,
    },
  },
  {
    capacity: 4, taken: 0, resources: []int{0},
    quantity: 0, err: ReleaseAlreadyUsedError{
      resource: 0,
      available: 4,
      capacity: 4,
    },
  },
}

func TestRelease(t *testing.T) {
  for _, tt := range releaseTests {
    rm := simulatedManager(tt.capacity, tt.taken)
    availBefore := rm.available.Len()
    // make the call to Release for the given resources
    err := rm.Release(tt.resources)
    // used to identify the test case
    releaseStr := fmt.Sprintf(
      "Release(x%d) [cap %d/%d]",
      len(tt.resources), tt.taken, tt.capacity,
    )
    // ensure the right error (or lack therof) is returned
    if err != tt.err {
      t.Errorf(
        "%s:\ngot\t%v\nwant\t%v\n",
        releaseStr, err, tt.err,
      )
    }
    // ensure the right number of resources was taken out
    // this is done by checking the availability before vs after
    quantity := rm.available.Len() - availBefore
    if quantity != tt.quantity {
      t.Errorf(
        "%s:\nreleased %d resources\nexpected %d\n",
        releaseStr, quantity, tt.quantity,
      )
    }
  }
}

func simulatedManager(capacity int, taken int) *ResourceManager {
  // make a new resource manager and remove the appropriate resources
  // (for each not available from capacity, remove one from the list)
  rm := New(capacity)
  for i := 0; i < taken; i++ {
    e := rm.available.Front()
    rm.available.Remove(e)
  }
  return rm
}
