# persian-cal
A simple calendar for use in projects that need to work with the solar calendar, such as converting the Gregorian calendar to the solar calendar and vice versa and formatting the date as desired in the CLI programs.

## Quick Start

``` code
import	"majidzarephysics/persian-cal"
```

you can get jalali time from system unix time

``` go code
date, _ := jalali.Unix(1608116340)
```

or you can convert georgian date to jalali directly

``` go code
date, _ := jalali.ToJalali(2003,10,25)
```
