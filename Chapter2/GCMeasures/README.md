## Description of the example

This example shows that the garbage collector can run for varying amounts of time with different ways of storing large data.
Among the 4 small programs, the garbage collector works the fastest with the one where an ordinary slice was used.
And the difference in time is quite significant. moreover, it is not so much the accuracy and nature of the measurements that is important here,
but the difference in time.

## Approximate measurements of time

| Programm   | Time, seconds |
| -----------|:-------------:|
| sliceGC    | 1,6103779     |
| mapGC      | 8,4125357     |
| mapPtrGC   | 11,1446763    |
| mapSplitGC | 8,5930541     |
