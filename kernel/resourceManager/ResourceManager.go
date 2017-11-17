package resourceManager

import (
  "container/list"
)

// ResourceManager keeps track of available resources.
type ResourceManager struct {
  capacity int
  available *list.List
}

// New makes a new resource manager with a given capacity.
func New(capacity int) *ResourceManager {
  // allocate and initialize the resource manager
  rm := new(ResourceManager)
  rm.capacity = capacity
  rm.available = list.New()

  // make sure each unit is accessible
  for i := 0; i < capacity; i++ {
    rm.available.PushBack(i)
  }
  return rm
}

// Claim claims a certain number of resources.
func (rm *ResourceManager) Claim(quantity int) ([]int, error) {
  // make sure what is being claimed is not too much
  // otherwise, an error needs to be returned to let the caller know
  quantityAvailable := rm.QuantityAvailable()
  if quantity > quantityAvailable {
    return nil, ClaimTooLargeError{
      claimSize: quantity,
      available: quantityAvailable,
      capacity: rm.capacity,
    }
  }
  // since the caller isn't asking for too much, the request should be granted!
  // just give the caller the first items off the available list
  retval := make([]int, 0, quantity)
  for i := 0; i < quantity; i++ {
    // get the next item in the list, add it to retval, and remove it
    // this ensures the caller gets the item and it is not going to others
    front := rm.available.Front()
    retval = append(retval, front.Value.(int))
    rm.available.Remove(front)
  }
  return retval, nil
}

// Release gives back resources to the resource manager.
func (rm *ResourceManager) Release(resources []int) error {
  // make sure the resources are all valid
  // otherwise an error should be returned to the caller
  for _, r := range resources {
    if err := rm.validateResourceRelease(r); err != nil {
      // if an error was yielded for this, the resource was invalid
      // that error needs to be forwarded to the caller for analysis
      return err
    }
  }
  // since everything is valid, we can give those resources back!
  // this is done by adding these items to the available list
  for _, r := range resources {
    rm.available.PushBack(r)
  }
  return nil
}

// QuantityAvailable returns the number of available resources.
func (rm ResourceManager) QuantityAvailable() int {
  return rm.available.Len()
}

func (rm ResourceManager) validateResourceRelease(resource int) error {
  if !rm.resourceInBounds(resource) {
    // resource was not in-bounds, so return an error!
    return ReleaseOutOfBoundsError{
      resource: resource,
      available: rm.QuantityAvailable(),
      capacity: rm.capacity,
    }
  } else if rm.resourceAlreadyAvailable(resource) {
    // resource has already been used, so return an error!
    return ReleaseAlreadyUsedError{
      resource: resource,
      available: rm.QuantityAvailable(),
      capacity: rm.capacity,
    }
  }
  return nil
}

func (rm ResourceManager) resourceInBounds(resource int) bool {
  if resource < 0 {
    // negative values are out-of-bounds (of course!)
    return false
  } else if resource >= rm.capacity {
    // a value greater than or equal to the capacity is too large
    // (note this is 0-based indexing, so yes it's equal to as well)
    return false
  }
  // if both tests pass, then we are good!
  return true
}

func (rm ResourceManager) resourceAlreadyAvailable(resource int) bool {
  for e := rm.available.Front(); e != nil; e = e.Next() {
    if resource == e.Value.(int) {
      // this resource has already been made available
      return true
    }
  }
  return false
}
