/// Determines the current status of the process
/// (ex: if it is running, waiting, etc)
enum Status {
  /// Process is currently being run in a processor
  Running,
  /// Process is ready to be run in a processor, but is not currently being run
  Ready,
  /// Process cannot be run in the processor for some reason
  Blocked,
  /// Process is freshly created (this should be the default)
  New
}

export default Status
