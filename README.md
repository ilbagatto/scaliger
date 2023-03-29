# Scaliger

Library of date/time manipulation routines for practical astronomy.


## Contents

### Julian Date

Most of the calculations are based on so called **Julian date (JD)**, which is
the number of days elapsed since mean UT noon of **January 1st 4713 BC**. This
system of time measurement is widely adopted by the astronomers.

Time is represented by fractional part of a day. For example, 7h30m UT
is `(7 + 30 / 60) / 24 = 0.3125`.

* **CivilToJulian(date CivilDate) float64** converts calendar date into Julian days
* **JulianToCivil(jd float64) CivilDate** Converts number of Julian days into the calendar date.

### Sidereal Time

*Sidereal Time* is reckoned by the daily transit of a fixed point in space
(fixed with respect to the distant stars), 24 hours of sidereal time elapsing
between an successive transits. The sidereal day is thus shorter than the
solar day by nearely 4 minutes, and although the solar and sidereal time
agree once a year, the difference between them grows systematically as the
months pass in the sense that sidereal time runs faster than solar time.
*Sidereal time* (ST) is used extensively by astronomers since it is the time
kept by the star.


### Dynamic Time

## Overview

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

