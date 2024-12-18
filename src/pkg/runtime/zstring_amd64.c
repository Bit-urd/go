// auto generated by go tool dist
// goos=linux goarch=amd64

#include "runtime.h"
#include "arch_GOARCH.h"
#include "malloc.h"

#line 1284 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
String runtime·emptystring; 
#line 1286 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
int32 
runtime·findnull ( byte *s ) 
{ 
int32 l; 
#line 1291 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
if ( s == nil ) 
return 0; 
for ( l=0; s[l]!=0; l++ ) 
; 
return l; 
} 
#line 1298 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
int32 
runtime·findnullw ( uint16 *s ) 
{ 
int32 l; 
#line 1303 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
if ( s == nil ) 
return 0; 
for ( l=0; s[l]!=0; l++ ) 
; 
return l; 
} 
#line 1310 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
uint32 runtime·maxstring = 256; 
#line 1312 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
static String 
gostringsize ( int32 l ) 
{ 
String s; 
uint32 ms; 
#line 1318 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
if ( l == 0 ) 
return runtime·emptystring; 
#line 1321 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
s.str = runtime·mallocgc ( l+1 , FlagNoPointers , 1 , 0 ) ; 
s.len = l; 
s.str[l] = 0; 
for ( ;; ) { 
ms = runtime·maxstring; 
if ( ( uint32 ) l <= ms || runtime·cas ( &runtime·maxstring , ms , ( uint32 ) l ) ) 
break; 
} 
return s; 
} 
#line 1332 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
String 
runtime·gostring ( byte *str ) 
{ 
int32 l; 
String s; 
#line 1338 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
l = runtime·findnull ( str ) ; 
s = gostringsize ( l ) ; 
runtime·memmove ( s.str , str , l ) ; 
return s; 
} 
#line 1344 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
String 
runtime·gostringn ( byte *str , int32 l ) 
{ 
String s; 
#line 1349 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
s = gostringsize ( l ) ; 
runtime·memmove ( s.str , str , l ) ; 
return s; 
} 
#line 1354 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
Slice 
runtime·gobytes ( byte *p , int32 n ) 
{ 
Slice sl; 
#line 1359 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
sl.array = runtime·mallocgc ( n , FlagNoPointers , 1 , 0 ) ; 
sl.len = n; 
sl.cap = n; 
runtime·memmove ( sl.array , p , n ) ; 
return sl; 
} 
#line 1366 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
String 
runtime·gostringnocopy ( byte *str ) 
{ 
String s; 
#line 1371 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
s.str = str; 
s.len = runtime·findnull ( str ) ; 
return s; 
} 
#line 1376 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
String 
runtime·gostringw ( uint16 *str ) 
{ 
int32 n1 , n2 , i; 
byte buf[8]; 
String s; 
#line 1383 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
n1 = 0; 
for ( i=0; str[i]; i++ ) 
n1 += runtime·runetochar ( buf , str[i] ) ; 
s = gostringsize ( n1+4 ) ; 
n2 = 0; 
for ( i=0; str[i]; i++ ) { 
#line 1390 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
if ( n2 >= n1 ) 
break; 
n2 += runtime·runetochar ( s.str+n2 , str[i] ) ; 
} 
s.len = n2; 
s.str[s.len] = 0; 
return s; 
} 
#line 1399 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
String 
runtime·catstring ( String s1 , String s2 ) 
{ 
String s3; 
#line 1404 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
if ( s1.len == 0 ) 
return s2; 
if ( s2.len == 0 ) 
return s1; 
#line 1409 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
s3 = gostringsize ( s1.len + s2.len ) ; 
runtime·memmove ( s3.str , s1.str , s1.len ) ; 
runtime·memmove ( s3.str+s1.len , s2.str , s2.len ) ; 
return s3; 
} 
#line 1415 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
static String 
concatstring ( int32 n , String *s ) 
{ 
int32 i , l; 
String out; 
#line 1421 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
l = 0; 
for ( i=0; i<n; i++ ) { 
if ( l + s[i].len < l ) 
runtime·throw ( "string concatenation too long" ) ; 
l += s[i].len; 
} 
#line 1428 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
out = gostringsize ( l ) ; 
l = 0; 
for ( i=0; i<n; i++ ) { 
runtime·memmove ( out.str+l , s[i].str , s[i].len ) ; 
l += s[i].len; 
} 
return out; 
} 
#line 1437 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
#pragma textflag 7 
void
runtime·concatstring(int32 n, String s1)
{
#line 1440 "/home/nch/workspace/go/src/pkg/runtime/string.goc"

	(&s1)[n] = concatstring(n, &s1);
}

#line 1444 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
static int32 
cmpstring ( String s1 , String s2 ) 
{ 
uint32 i , l; 
byte c1 , c2; 
#line 1450 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
l = s1.len; 
if ( s2.len < l ) 
l = s2.len; 
for ( i=0; i<l; i++ ) { 
c1 = s1.str[i]; 
c2 = s2.str[i]; 
if ( c1 < c2 ) 
return -1; 
if ( c1 > c2 ) 
return +1; 
} 
if ( s1.len < s2.len ) 
return -1; 
if ( s1.len > s2.len ) 
return +1; 
return 0; 
} 
void
runtime·cmpstring(String s1, String s2, int32 v)
{
#line 1468 "/home/nch/workspace/go/src/pkg/runtime/string.goc"

	v = cmpstring(s1, s2);
	FLUSH(&v);
}

#line 1472 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
int32 
runtime·strcmp ( byte *s1 , byte *s2 ) 
{ 
uint32 i; 
byte c1 , c2; 
#line 1478 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
for ( i=0;; i++ ) { 
c1 = s1[i]; 
c2 = s2[i]; 
if ( c1 < c2 ) 
return -1; 
if ( c1 > c2 ) 
return +1; 
if ( c1 == 0 ) 
return 0; 
} 
} 
#line 1490 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
byte* 
runtime·strstr ( byte *s1 , byte *s2 ) 
{ 
byte *sp1 , *sp2; 
#line 1495 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
if ( *s2 == 0 ) 
return s1; 
for ( ; *s1; s1++ ) { 
if ( *s1 != *s2 ) 
continue; 
sp1 = s1; 
sp2 = s2; 
for ( ;; ) { 
if ( *sp2 == 0 ) 
return s1; 
if ( *sp1++ != *sp2++ ) 
break; 
} 
} 
return nil; 
} 
void
runtime·slicestring(String si, int32 lindex, int32 hindex, String so)
{
#line 1512 "/home/nch/workspace/go/src/pkg/runtime/string.goc"

	int32 l;

	if(lindex < 0 || lindex > si.len ||
	   hindex < lindex || hindex > si.len) {
	   	runtime·panicslice();
	}

	l = hindex-lindex;
	so.str = si.str + lindex;
	so.len = l;
	FLUSH(&so);
}
void
runtime·slicestring1(String si, int32 lindex, uint32, String so)
{
#line 1525 "/home/nch/workspace/go/src/pkg/runtime/string.goc"

	int32 l;

	if(lindex < 0 || lindex > si.len) {
		runtime·panicslice();
	}

	l = si.len-lindex;
	so.str = si.str + lindex;
	so.len = l;
	FLUSH(&so);
}
void
runtime·intstring(int64 v, String s)
{
#line 1537 "/home/nch/workspace/go/src/pkg/runtime/string.goc"

	s = gostringsize(8);
	s.len = runtime·runetochar(s.str, v);
	s.str[s.len] = 0;
	FLUSH(&s);
}
void
runtime·slicebytetostring(Slice b, String s)
{
#line 1543 "/home/nch/workspace/go/src/pkg/runtime/string.goc"

	s = gostringsize(b.len);
	runtime·memmove(s.str, b.array, s.len);
	FLUSH(&s);
}
void
runtime·stringtoslicebyte(String s, Slice b)
{
#line 1548 "/home/nch/workspace/go/src/pkg/runtime/string.goc"

	b.array = runtime·mallocgc(s.len, FlagNoPointers, 1, 0);
	b.len = s.len;
	b.cap = s.len;
	runtime·memmove(b.array, s.str, s.len);
	FLUSH(&b);
}
void
runtime·slicerunetostring(Slice b, String s)
{
#line 1555 "/home/nch/workspace/go/src/pkg/runtime/string.goc"

	int32 siz1, siz2, i;
	int32 *a;
	byte dum[8];

	a = (int32*)b.array;
	siz1 = 0;
	for(i=0; i<b.len; i++) {
		siz1 += runtime·runetochar(dum, a[i]);
	}

	s = gostringsize(siz1+4);
	siz2 = 0;
	for(i=0; i<b.len; i++) {
		// check for race
		if(siz2 >= siz1)
			break;
		siz2 += runtime·runetochar(s.str+siz2, a[i]);
	}
	s.len = siz2;
	s.str[s.len] = 0;
	FLUSH(&s);
}
void
runtime·stringtoslicerune(String s, Slice b)
{
#line 1578 "/home/nch/workspace/go/src/pkg/runtime/string.goc"

	int32 n;
	int32 dum, *r;
	uint8 *p, *ep;

	// two passes.
	// unlike slicerunetostring, no race because strings are immutable.
	p = s.str;
	ep = s.str+s.len;
	n = 0;
	while(p < ep) {
		p += runtime·charntorune(&dum, p, ep-p);
		n++;
	}

	b.array = runtime·mallocgc(n*sizeof(r[0]), FlagNoPointers, 1, 0);   // src/pkg/runtime/zmalloc_amd64.c:20
	b.len = n;
	b.cap = n;
	p = s.str;
	r = (int32*)b.array;
	while(p < ep)
		p += runtime·charntorune(r++, p, ep-p);
	FLUSH(&b);
}

#line 1602 "/home/nch/workspace/go/src/pkg/runtime/string.goc"
enum 
{ 
Runeself = 0x80 , 
} ; 
void
runtime·stringiter(String s, int32 k, uint32, int32 retk)
{
#line 1607 "/home/nch/workspace/go/src/pkg/runtime/string.goc"

	int32 l;

	if(k >= s.len) {
		// retk=0 is end of iteration
		retk = 0;
		goto out;
	}

	l = s.str[k];
	if(l < Runeself) {
		retk = k+1;
		goto out;
	}

	// multi-char rune
	retk = k + runtime·charntorune(&l, s.str+k, s.len-k);

out:
	FLUSH(&retk);
}
void
runtime·stringiter2(String s, int32 k, uint32, int32 retk, int32 retv)
{
#line 1628 "/home/nch/workspace/go/src/pkg/runtime/string.goc"

	if(k >= s.len) {
		// retk=0 is end of iteration
		retk = 0;
		retv = 0;
		goto out;
	}

	retv = s.str[k];
	if(retv < Runeself) {
		retk = k+1;
		goto out;
	}

	// multi-char rune
	retk = k + runtime·charntorune(&retv, s.str+k, s.len-k);

out:
	FLUSH(&retk);
	FLUSH(&retv);
}
