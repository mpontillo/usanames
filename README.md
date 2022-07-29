## `usanames` library

This is a Go library written for a single purpose: to generate "realistic"
human names.

It uses public U.S. Census and Social Security Administration data to populate
its data set.

### Example

```
$ ./build.sh
$ bin/usanamegen naive --count 10
Kylyn JENNA
Kristofer DURRETTE
Florencia TROVILLION
Khade LENAU
Jalai BRZYCKI
Talulah TIMBERS
Christine SIMPKIN
Folasade KOTHAPALLI
Irys FACIO
Abdulhakeem PEARO
```

The fact that last names are all-uppercase is an artifact of the U.S. Census
data set.
