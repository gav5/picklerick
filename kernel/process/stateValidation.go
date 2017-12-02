package process

func validateTransition(o, n Status) bool {
  switch o {
  case New:
    return validateFromNew(n)
  case Ready:
    return validateFromReady(n)
  case Run:
    return validateFromRun(n)
  case Wait:
    return validateFromWait(n)
  case Terminated:
    return validateFromTerminated(n)
  default:
    return false
  }
}

func validateFromNew(n Status) bool {
  switch n {
  case New, Ready:
    return true
  default:
    return false
  }
}

func validateFromReady(n Status) bool {
  switch n {
  case Ready, Run:
    return true
  default:
    return false
  }
}

func validateFromRun(n Status) bool {
  switch n {
  case Run, Wait, Terminated:
    return true
  default:
    return false
  }
}

func validateFromWait(n Status) bool {
  switch n {
  case Wait, Ready:
    return true
  default:
    return false
  }
}

func validateFromTerminated(n Status) bool {
  switch n {
  case Terminated:
    return true
  default:
    return false
  }
}
