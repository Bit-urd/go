// auto generated by go tool dist
// goos=linux goarch=amd64

#include "runtime.h"
#include "defs_GOOS_GOARCH.h"
#include "os_GOOS.h"
#include "arch_GOARCH.h"
#include "malloc.h"

#line 1661 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
static Timers timers; 
static void addtimer ( Timer* ) ; 
static bool deltimer ( Timer* ) ; 
void
time·Sleep(int64 ns)
{
#line 1671 "/home/nch/workspace/go/src/pkg/runtime/time.goc"

	g->status = Gwaiting;
	g->waitreason = "sleep";
	runtime·tsleep(ns);
}
void
time·startTimer(Timer* t)
{
#line 1678 "/home/nch/workspace/go/src/pkg/runtime/time.goc"

	addtimer(t);
}
void
time·stopTimer(Timer* t, bool stopped)
{
#line 1684 "/home/nch/workspace/go/src/pkg/runtime/time.goc"

	stopped = deltimer(t);
	FLUSH(&stopped);
}

#line 1690 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
static void timerproc ( void ) ; 
static void siftup ( int32 ) ; 
static void siftdown ( int32 ) ; 
#line 1695 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
static void 
ready ( int64 now , Eface e ) 
{ 
USED ( now ) ; 
#line 1700 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
runtime·ready ( e.data ) ; 
} 
#line 1705 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
void 
runtime·tsleep ( int64 ns ) 
{ 
Timer t; 
#line 1710 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
if ( ns <= 0 ) 
return; 
#line 1713 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
t.when = runtime·nanotime ( ) + ns; 
t.period = 0; 
t.f = ready; 
t.arg.data = g; 
addtimer ( &t ) ; 
runtime·gosched ( ) ; 
} 
#line 1723 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
static void 
addtimer ( Timer *t ) 
{ 
int32 n; 
Timer **nt; 
#line 1729 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
runtime·lock ( &timers ) ; 
if ( timers.len >= timers.cap ) { 
#line 1732 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
n = 16; 
if ( n <= timers.cap ) 
n = timers.cap*3 / 2; 
nt = runtime·malloc ( n*sizeof nt[0] ) ; 
runtime·memmove ( nt , timers.t , timers.len*sizeof nt[0] ) ; 
runtime·free ( timers.t ) ; 
timers.t = nt; 
timers.cap = n; 
} 
t->i = timers.len++; 
timers.t[t->i] = t; 
siftup ( t->i ) ; 
if ( t->i == 0 ) { 
#line 1746 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
if ( timers.sleeping ) { 
timers.sleeping = false; 
runtime·notewakeup ( &timers.waitnote ) ; 
} 
if ( timers.rescheduling ) { 
timers.rescheduling = false; 
runtime·ready ( timers.timerproc ) ; 
} 
} 
if ( timers.timerproc == nil ) 
timers.timerproc = runtime·newproc1 ( ( byte* ) timerproc , nil , 0 , 0 , addtimer ) ; 
runtime·unlock ( &timers ) ; 
} 
#line 1763 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
static bool 
deltimer ( Timer *t ) 
{ 
int32 i; 
#line 1768 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
runtime·lock ( &timers ) ; 
#line 1773 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
i = t->i; 
if ( i < 0 || i >= timers.len || timers.t[i] != t ) { 
runtime·unlock ( &timers ) ; 
return false; 
} 
#line 1779 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
timers.len--; 
if ( i == timers.len ) { 
timers.t[i] = nil; 
} else { 
timers.t[i] = timers.t[timers.len]; 
timers.t[timers.len] = nil; 
timers.t[i]->i = i; 
siftup ( i ) ; 
siftdown ( i ) ; 
} 
runtime·unlock ( &timers ) ; 
return true; 
} 
#line 1797 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
static void 
timerproc ( void ) 
{ 
int64 delta , now; 
Timer *t; 
void ( *f ) ( int64 , Eface ) ; 
Eface arg; 
#line 1805 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
for ( ;; ) { 
runtime·lock ( &timers ) ; 
now = runtime·nanotime ( ) ; 
for ( ;; ) { 
if ( timers.len == 0 ) { 
delta = -1; 
break; 
} 
t = timers.t[0]; 
delta = t->when - now; 
if ( delta > 0 ) 
break; 
if ( t->period > 0 ) { 
#line 1819 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
t->when += t->period * ( 1 + -delta/t->period ) ; 
siftdown ( 0 ) ; 
} else { 
#line 1823 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
timers.t[0] = timers.t[--timers.len]; 
timers.t[0]->i = 0; 
siftdown ( 0 ) ; 
t->i = -1; 
} 
f = t->f; 
arg = t->arg; 
runtime·unlock ( &timers ) ; 
f ( now , arg ) ; 
runtime·lock ( &timers ) ; 
} 
if ( delta < 0 ) { 
#line 1836 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
timers.rescheduling = true; 
g->status = Gwaiting; 
g->waitreason = "timer goroutine (idle)" ; 
runtime·unlock ( &timers ) ; 
runtime·gosched ( ) ; 
continue; 
} 
#line 1844 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
timers.sleeping = true; 
runtime·noteclear ( &timers.waitnote ) ; 
runtime·unlock ( &timers ) ; 
runtime·entersyscall ( ) ; 
runtime·notetsleep ( &timers.waitnote , delta ) ; 
runtime·exitsyscall ( ) ; 
} 
} 
#line 1855 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
static void 
siftup ( int32 i ) 
{ 
int32 p; 
Timer **t , *tmp; 
#line 1861 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
t = timers.t; 
while ( i > 0 ) { 
p = ( i-1 ) /2; 
if ( t[i]->when >= t[p]->when ) 
break; 
tmp = t[i]; 
t[i] = t[p]; 
t[p] = tmp; 
t[i]->i = i; 
t[p]->i = p; 
i = p; 
} 
} 
#line 1875 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
static void 
siftdown ( int32 i ) 
{ 
int32 c , len; 
Timer **t , *tmp; 
#line 1881 "/home/nch/workspace/go/src/pkg/runtime/time.goc"
t = timers.t; 
len = timers.len; 
for ( ;; ) { 
c = i*2 + 1; 
if ( c >= len ) { 
break; 
} 
if ( c+1 < len && t[c+1]->when < t[c]->when ) 
c++; 
if ( t[c]->when >= t[i]->when ) 
break; 
tmp = t[i]; 
t[i] = t[c]; 
t[c] = tmp; 
t[i]->i = i; 
t[c]->i = c; 
i = c; 
} 
} 