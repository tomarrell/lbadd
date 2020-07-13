// Package profile implements profiling with generic events. An event can be
// defined by the user of this package and just needs to implement fmt.Stringer.
// This profiler works with string-based events. Use the profiler like this in
// your code.
//
//  ...
//  type Evt string
//  func (e Evt) String() string { return string(e) }
//  ...
//  const MyEvt = Evt("my expensive func")
//  ...
//  func main() {
//      prof = profile.NewProfiler()
//      MyExpensiveFunc()
//      fmt.Println(prof.Profile().String()) // will print the profile
//  }
//  ...
//  func MyExpensiveFunc() {
//      defer prof.Enter(MyEvt).Exit()
//      ...
//  }
//
// The above example will print a profile with one event, the respective
// timestamps and durations etc.
//
// NOTE: You don't need an actual profiler. All methods will also work on nil
// profilers, such as the following.
//
//  var prof *profile.Profiler
//  defer prof.Enter(MyEvt).Exit()
//
// Enter(...) will just create an empty Event, and Exit() will do nothing. This
// was implemented this way, so that no no-op implementation of a profiler is
// neccessary.
package profile