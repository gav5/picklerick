import State from '../cpu/State'
import Status from './Status'

/// Used to keeps track of a given process in the system
export default class ControlBlock {

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  cpuid: any

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  programCounter: any

  /// Record of enviornment that is saved on interrupt
  state: State

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  codeSize: any

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  registers: any

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  schedule: any

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  accounts: any

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  memories: any

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  progeny: any

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  ptr: any

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  resources: any

  /// Determines the current status of the process
  /// (ex: if it is running, waiting, etc)
  status: Status = Status.New

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  statusInfo: any

  // TODO: add description of what this does
  // TODO: figure out what type this should be
  priority: any
}
