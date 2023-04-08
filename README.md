# Scaliger

Library of date/time manipulation routines for practical astronomy.



- [Scaliger](#scaliger)
  - [Quick Start](#quick-start)
  - [Packages](#packages)
    - [julian](#julian)
    - [sidereal](#sidereal)
    - [deltat](#deltat)
    - [nutequ](#nutequ)
      - [Obliquity of the ecliptic](#obliquity-of-the-ecliptic)
      - [Nutation](#nutation)
  - [Caveats](#caveats)
    - [Civil vs. Astronomical year](#civil-vs-astronomical-year)
    - [Zero day](#zero-day)
    - [Gregorian calendar](#gregorian-calendar)
  - [How to contribute](#how-to-contribute)
  - [Sources](#sources)
  - [Footnotes](#footnotes)


## Quick Start

```console

$ go get github.com/skrushinsky/scaliger

```

## Packages

### julian

Most of the astronomical calculations are based on so called *Julian date*, which is
the number of days elapsed since mean UT noon of January 1st 4713 BC.

`CivilToJulian(date CivilDate) float64` converts calendar date into Julian days.

```go
date := CivilDate{Year: 837, Month: 4, Day: 10.3}
got := CivilToJulian(date) // 2026871.8
```
The calendar date is represented by `CivilDate` structure.

```go
type CivilDate struct {
	Year  int     // year, astronomical, negative for BC dates
	Month int     // month number, 1-12
	Day   float64 // day, fractional part represents hours
}
```

`JulianToCivil(jd float64) CivilDate` converts Julian days into the calendar date.

```go
jd := 2455197.5
date := JulianToCivil(2455197.5) // CivilDate{Year: 2010, Month: 1, Day: 1.0}
```

Other utilitity functions from the package are mostly used internally. Please, see the API docs.



### sidereal

*Sidereal Time* is reckoned by the daily transit of a fixed point in space (fixed with respect
to the distant stars), 24 hours of sidereal time elapsing between a successive transits.

`JulianToSidereal(jd float64, options SiderealOptions) float64` converts Julian date to
Sidereal Time.


```go
opts := SiderealOptions{Dpsi: -0.0043, Eps: 23.4443, Lng: 37.5833}
lst := JulianToSidereal(jd, opts) // 23.0370...
```

`SiderealOptions` structure controls type of the result.

```go
type SiderealOptions struct {
	Lng  float64 // geographical longitude, degrees, negative westwards
	Eps  float64 // obliquity of the ecliptic, degrees
	Dpsi float64 // nutation in longitude, degrees
}
```
If geographical longitude (`Lng`) is initialized, the function calculates *Local Sidereal Time*.
Without it *Greenwich Sidereal Time*[^1].

If obliquity of the ecliptic, `Eps`, and nutation in longitude, `Dpsi` are provided, then
the result is *apparent Sidereal Time*, or the Greenwich hour angle of the *true vernal equinox*.
For *Mean Sidereal Time*, referred to the *mean vernal point* [^2], leave these fields empty:

```go
opts := SiderealOptions{Lng: 37.5833}
...
```

To calculate *obliquity of the ecliptic* and *nutation*, use nutequ package.


### deltat

*DeltaT* indicates the difference between *UTC* (Greenwich Coordinated Time) and *TDT*
(Terrestrial Dynamic Time), which used to be called *'Ephemeris time'* in the last
century. While *UTC* is not a uniform time scale (it is occasionally adjusted, due to irregularities
in the Earth's rotation), *TDT* is a uniform time scale which is needed as an argument for
mathematical theories of celestial movements.

The exact value of the difference `DeltaT = TDT - UTC` can be deduced only from observations/
Approximate Delta-T in seconds for a given Julian Date may be obtained by `DeltaT(jd float64) float64`
function. To correct a date for *TDT*, simply add DeltaT seconds divided by `86400` (seconds per day).

```go
jd := 2459040.5  // Julian date for 2020-07-10
dt := DeltaT(jd) // 93.81 seconds
jde := jd + dt / 86400 // Dynamic time.
```

### nutequ

#### Obliquity of the ecliptic

*Obliquity of the ecliptic* is the angle between the eauator and the ecliptic.

To calculate *Mean Obliquity*, angle which the ecliptic makes with the mean equator,
use `MeanObliquity(jd float64) float64` function, where **jd** is Juliuan Date.
For better accuracy correct it for *TDT* (see [Dynamic Time](#dynamic-time--deltat-package)).

The formula used should give an accurate results (less than 1 srcsecond) within 2000
years around the epoch *J2000.0* (see *J.Meeus, Astronomical Algorithms, 2d edition, p.147*).

` TrueObliquity(jd, deps float64) float64` function calculates *True Obliquity*, where `deps`
is the *nutation in obliquity*.

#### Nutation

`Nutation(jd float64) (dpsi float64, deps float64)` calculates

 * `dpsi`, *nutation in longitude*, arc-degrees
 * `deps`, *nutation in obliquity*, arc-degrees



## Caveats

### Civil vs. Astronomical year

There is disagreement between astronomers and historians about how to count
the years preceding the year 1. Astronomers generally use zero-based system.
The year before the year +1, is the *year zero*, and the year preceding the
latter is the *year -1*. The year which the historians call *585 B.C.* is
actually the year *-584*.

### Zero day

Zero day is a special case of date: it indicates 12h UT of previous calendar
date. For instance, *1900 January 0.5* is often used instead of
*1899 December 31.5* to designate start of the astronomical epoch.

###  Gregorian calendar

_Civil calendar_ in most cases means _proleptic Gregorian calendar_. it is
assumed that Gregorian calendar started at *Oct. 4, 1582*, when it was first
adopted in several European countries. Many other countries still used the
older Julian calendar. In Soviet Russia, for instance, Gregorian system was
accepted on **Jan 26, 1918**. See
[Wiki article](https://en.wikipedia.org/wiki/Gregorian_calendar#Adoption_of_the_Gregorian_Calendar)


## How to contribute

You may contribute to the project by many different ways, starting from refining and correcting its documentation,
especially if you are a native English speaker, and ending with improving the code base. Any kind of testing and
suggestions are welcome.

You may follow the standard Github procedures or, in case you are not comfortable with them, just send your suggestions
to the author by email.

## Sources

The formulae were adopted from the following sources:

* _Peter Duffett-Smith, "Astronomy With Your Personal Computer", Cambridge University Press, 1997_
* _Jean Meeus, "Astronomical Algorithms", 2d edition, Willmann-Bell, 1998_
* _J.L.Lawrence, "Celestial Calculations", The MIT Press, 2018_


## Footnotes

[^1]: *Greenwich Sidereal Time* — Sidereal time at the Greenwich meridian, irrelevant to the observer's
position.

[^2]: *Mean vernal point* — intersection of the ecliptic of the date with the mean equator of the date.
