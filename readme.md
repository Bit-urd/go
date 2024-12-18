```bash
dlv exec a
> l  【是源代码】
> si   【是反汇编代码，逻辑可能有偏差】
> n
> so
> regs
> stack

../bin/go build -work -a a.go  #保留工作目录
```

```bash
(dlv) vars
syscall.initdone· = 0
syscall._zero = 0
syscall.signals = [32]string [...]
syscall.errors = [133]string [...]
syscall.envLock = sync.RWMutex {w: (*sync.Mutex)(0x4f0b58), writerSem: 0, readerSem: 0,...+2 more}
syscall.envOnce = sync.Once {m: (*sync.Mutex)(0x4f0b38), done: 0}
syscall.env = *hash<string,int> nil
math.initdone· = 0
math.pow10tab = [70]float64 [...]
reflect.initdone· = 0
reflect.dummy = struct { b bool; x interface {} } {b: false, x: (*interface {})(0x4df338)}
reflect.ptrMap = struct { sync.RWMutex; m map[*reflect.commonType]*reflect.ptrType } {sync.RWMutex: (*sync.RWMutex)(0x4df348), m: *hash<*reflect.commonType,*reflect.ptrType> nil}
reflect.kindNames = []string {array: (*string)(0x4dc310), len: 27, cap: 27}
os.Stderr = *os.File nil
syscall.Stderr = 2
syscall.Stdout = 1
syscall.Stdin = 0
os.Interrupt = os.Signal {tab: *runtime.itab nil, data: *nil}
os.initdone· = 0
os.ErrInvalid = error {tab: *runtime.itab nil, data: *nil}
os.Kill = os.Signal {tab: *runtime.itab nil, data: *nil}
os.ErrPermission = error {tab: *runtime.itab nil, data: *nil}
os.ErrNotExist = error {tab: *runtime.itab nil, data: *nil}
os.ErrExist = error {tab: *runtime.itab nil, data: *nil}
io.initdone· = 0
io.ErrClosedPipe = error {tab: *runtime.itab nil, data: *nil}
io.errOffset = error {tab: *runtime.itab nil, data: *nil}
io.errWhence = error {tab: *runtime.itab nil, data: *nil}
io.ErrShortWrite = error {tab: *runtime.itab nil, data: *nil}
io.ErrShortBuffer = error {tab: *runtime.itab nil, data: *nil}
strconv.initdone· = 0
strconv.isNotPrint32 = []uint16 {array: (*uint16)(0x4da188), len: 42, cap: 42}
strconv.isPrint32 = []uint32 {array: (*uint32)(0x4da560), len: 180, cap: 180}
strconv.isNotPrint16 = []uint16 {array: (*uint16)(0x4da450), len: 134, cap: 134}
strconv.isPrint16 = []uint16 {array: (*uint16)(0x4da830), len: 474, cap: 474}
strconv.shifts = [37]uint [...]
strconv.uint64pow10 = [20]uint64 [...]
strconv.powersOfTen = [87]strconv.extFloat [...]
strconv.smallPowersOfTen = [8]strconv.extFloat [...]
strconv.leftcheats = []strconv.leftCheat {array: (*strconv.leftCheat)(0x4dcbb8), len: 28, cap: 28}
strconv.float64info = strconv.floatInfo {mantbits: 52, expbits: 11, bias: -1023}
strconv.ErrRange = error {tab: *runtime.itab nil, data: *nil}
strconv.float32info = strconv.floatInfo {mantbits: 23, expbits: 8, bias: -127}
strconv.optimize = true
strconv.float32pow10 = []float32 {array: (*float32)(0x4da120), len: 11, cap: 11}
strconv.float64pow10 = []float64 {array: (*float64)(0x4da398), len: 23, cap: 23}
strconv.powtab = []int {array: (*int)(0x4da0c8), len: 9, cap: 9}
strconv.ErrSyntax = error {tab: *runtime.itab nil, data: *nil}
time.initdone· = 0
time.zoneDirs = []string {array: (*string)(0x4db5f0), len: 4, cap: 4}
time.badData = error {tab: *runtime.itab nil, data: *nil}
time.zoneinfo = string {str: *uint8 nil, len: 0}
time.localOnce = sync.Once {m: (*sync.Mutex)(0x4f0b48), done: 0}
time.localLoc = time.Location {name: (*string)(0x4df420), zone: (*[]time.zone)(0x4df430), tx: (*[]time.zoneTrans)(0x4df440),...+3 more}
time.daysBefore = [13]int32 [...]
time.utcLoc = time.Location {name: (*string)(0x4db630), zone: (*[]time.zone)(0x4db640), tx: (*[]time.zoneTrans)(0x4db650),...+3 more}
time.unitMap = *hash<string,float64> nil
time.errLeadingInt = error {tab: *runtime.itab nil, data: *nil}
time.Local = (*time.Location)(0x4df420)
time.UTC = (*time.Location)(0x4db630)
time.longDayNames = []string {array: (*string)(0x4db7e8), len: 7, cap: 7}
time.shortDayNames = []string {array: (*string)(0x4db858), len: 7, cap: 7}
time.longMonthNames = []string {array: (*string)(0x4dbb98), len: 13, cap: 13}
time.shortMonthNames = []string {array: (*string)(0x4dbac8), len: 13, cap: 13}
time.days = [7]string [...]
time.months = [12]string [...]
time.atoiError = error {tab: *runtime.itab nil, data: *nil}
time.errBad = error {tab: *runtime.itab nil, data: *nil}
Sending output to pager...
syscall.initdone· = 0
syscall._zero = 0
syscall.signals = [32]string [...]
syscall.errors = [133]string [...]
syscall.envLock = sync.RWMutex {w: (*sync.Mutex)(0x4f0b58), writerSem: 0, readerSem: 0,...+2 more}
syscall.envOnce = sync.Once {m: (*sync.Mutex)(0x4f0b38), done: 0}
syscall.env = *hash<string,int> nil
math.initdone· = 0
math.pow10tab = [70]float64 [...]
reflect.initdone· = 0
reflect.dummy = struct { b bool; x interface {} } {b: false, x: (*interface {})(0x4df338)}
reflect.ptrMap = struct { sync.RWMutex; m map[*reflect.commonType]*reflect.ptrType } {sync.RWMutex: (*sync.RWMutex)(0x4df348), m: *hash
<*reflect.commonType,*reflect.ptrType> nil}
reflect.kindNames = []string {array: (*string)(0x4dc310), len: 27, cap: 27}
os.Stderr = *os.File nil
syscall.Stderr = 2
syscall.Stdout = 1
syscall.Stdin = 0
os.Interrupt = os.Signal {tab: *runtime.itab nil, data: *nil}
os.initdone· = 0
os.ErrInvalid = error {tab: *runtime.itab nil, data: *nil}
os.Kill = os.Signal {tab: *runtime.itab nil, data: *nil}
os.ErrPermission = error {tab: *runtime.itab nil, data: *nil}
os.ErrNotExist = error {tab: *runtime.itab nil, data: *nil}
os.ErrExist = error {tab: *runtime.itab nil, data: *nil}
io.initdone· = 0
io.ErrClosedPipe = error {tab: *runtime.itab nil, data: *nil}
io.errOffset = error {tab: *runtime.itab nil, data: *nil}
io.errWhence = error {tab: *runtime.itab nil, data: *nil}
io.ErrShortWrite = error {tab: *runtime.itab nil, data: *nil}
io.ErrShortBuffer = error {tab: *runtime.itab nil, data: *nil}
strconv.initdone· = 0
strconv.isNotPrint32 = []uint16 {array: (*uint16)(0x4da188), len: 42, cap: 42}
strconv.isPrint32 = []uint32 {array: (*uint32)(0x4da560), len: 180, cap: 180}
strconv.isNotPrint16 = []uint16 {array: (*uint16)(0x4da450), len: 134, cap: 134}
strconv.isPrint16 = []uint16 {array: (*uint16)(0x4da830), len: 474, cap: 474}
strconv.shifts = [37]uint [...]
strconv.uint64pow10 = [20]uint64 [...]
strconv.powersOfTen = [87]strconv.extFloat [...]
strconv.smallPowersOfTen = [8]strconv.extFloat [...]
strconv.leftcheats = []strconv.leftCheat {array: (*strconv.leftCheat)(0x4dcbb8), len: 28, cap: 28}
strconv.float64info = strconv.floatInfo {mantbits: 52, expbits: 11, bias: -1023}
strconv.ErrRange = error {tab: *runtime.itab nil, data: *nil}
strconv.float32info = strconv.floatInfo {mantbits: 23, expbits: 8, bias: -127}
strconv.optimize = true
strconv.float32pow10 = []float32 {array: (*float32)(0x4da120), len: 11, cap: 11}
strconv.float64pow10 = []float64 {array: (*float64)(0x4da398), len: 23, cap: 23}
strconv.powtab = []int {array: (*int)(0x4da0c8), len: 9, cap: 9}
strconv.ErrSyntax = error {tab: *runtime.itab nil, data: *nil}
time.initdone· = 0
time.zoneDirs = []string {array: (*string)(0x4db5f0), len: 4, cap: 4}
time.badData = error {tab: *runtime.itab nil, data: *nil}
time.zoneinfo = string {str: *uint8 nil, len: 0}
time.localOnce = sync.Once {m: (*sync.Mutex)(0x4f0b48), done: 0}
time.localLoc = time.Location {name: (*string)(0x4df420), zone: (*[]time.zone)(0x4df430), tx: (*[]time.zoneTrans)(0x4df440),...+3 more}
time.daysBefore = [13]int32 [...]
time.utcLoc = time.Location {name: (*string)(0x4db630), zone: (*[]time.zone)(0x4db640), tx: (*[]time.zoneTrans)(0x4db650),...+3 more}
time.unitMap = *hash<string,float64> nil
time.errLeadingInt = error {tab: *runtime.itab nil, data: *nil}
time.Local = (*time.Location)(0x4df420)
time.UTC = (*time.Location)(0x4db630)
time.longDayNames = []string {array: (*string)(0x4db7e8), len: 7, cap: 7}
time.shortDayNames = []string {array: (*string)(0x4db858), len: 7, cap: 7}
time.longMonthNames = []string {array: (*string)(0x4dbb98), len: 13, cap: 13}
time.shortMonthNames = []string {array: (*string)(0x4dbac8), len: 13, cap: 13}
time.days = [7]string [...]
time.months = [12]string [...]
time.atoiError = error {tab: *runtime.itab nil, data: *nil}
time.errBad = error {tab: *runtime.itab nil, data: *nil}
fmt.complexBits = 0
fmt.floatBits = 0
fmt.initdone· = 0
fmt.uintptrBits = 0
fmt.intBits = 0
fmt.complexError = error {tab: *runtime.itab nil, data: *nil}
fmt.boolError = error {tab: *runtime.itab nil, data: *nil}
fmt.ssFree = *fmt.cache nil
fmt.space = [][2]uint16 {array: (*[2]uint16)(0x4da0f0), len: 11, cap: 11}
io.ErrUnexpectedEOF = error {tab: *runtime.itab nil, data: *nil}
io.EOF = error {tab: *runtime.itab nil, data: *nil}
os.Stdin = *os.File nil
fmt.extraBytes = []uint8 {array: (*uint8)(0x4da068), len: 9, cap: 9}
fmt.missingBytes = []uint8 {array: (*uint8)(0x4da058), len: 9, cap: 9}
fmt.noVerbBytes = []uint8 {array: (*uint8)(0x4da078), len: 10, cap: 10}
fmt.precBytes = []uint8 {array: (*uint8)(0x4da088), len: 11, cap: 11}
fmt.widthBytes = []uint8 {array: (*uint8)(0x4da098), len: 12, cap: 12}
fmt.nilParenBytes = []uint8 {array: (*uint8)(0x4da038), len: 5, cap: 5}
fmt.mapBytes = []uint8 {array: (*uint8)(0x4da010), len: 4, cap: 4}
fmt.panicBytes = []uint8 {array: (*uint8)(0x4da040), len: 7, cap: 7}
fmt.nilBytes = []uint8 {array: (*uint8)(0x4da008), len: 3, cap: 3}
fmt.commaSpaceBytes = []uint8 {array: (*uint8)(0x4da002), len: 2, cap: 2}
fmt.bytesBytes = []uint8 {array: (*uint8)(0x4da048), len: 7, cap: 7}
fmt.nilAngleBytes = []uint8 {array: (*uint8)(0x4da030), len: 5, cap: 5}
os.Stdout = *os.File nil
fmt.ppFree = *fmt.cache nil
fmt.irparenBytes = []uint8 {array: (*uint8)(0x4da004), len: 2, cap: 2}
fmt.falseBytes = []uint8 {array: (*uint8)(0x4da028), len: 5, cap: 5}
fmt.trueBytes = []uint8 {array: (*uint8)(0x4da00c), len: 4, cap: 4}
fmt.padSpaceBytes = []uint8 {array: *uint8 nil, len: 0, cap: 0}
fmt.padZeroBytes = []uint8 {array: *uint8 nil, len: 0, cap: 0}
syscall.envs = []string {array: *string nil, len: 0, cap: 0}
os.Args = []string {array: *string nil, len: 0, cap: 0}
runtime.g0 = runtime.g {stackguard: *uint8 nil, stackbase: *uint8 nil, _defer: *runtime._defer nil,...+27 more}
runtime.m0 = runtime.m {g0: *runtime.g nil, morepc: nil, moreargp: *nil,...+37 more}
runtime.checking = 0
runtime.class_to_transfercount = [61]int [...]
runtime.class_to_allocnpages = [61]int [...]
runtime.class_to_size = [61]int [...]
runtime.worldsema = 1
runtime.iscgo = 0
runtime.ncpu = 0
runtime.gcwaiting = 0
runtime.panicking = 0
runtime.singleproc = 0
runtime.gomaxprocs = 0
runtime.allm = *runtime.m nil
runtime.lastg = *runtime.g nil
runtime.allg = *runtime.g nil
runtime.emptystring = string {str: *uint8 nil, len: 0}
runtime.memStats = runtime.MemStats {Alloc: 0, TotalAlloc: 0, Sys: 0,...+24 more}
runtime.MemProfileRate = 524288
runtime.initdone· = 0
runtime.sizeof_C_MStats = 3696
runtime.algarray = [22]runtime.alg [...]
main.initdone· = 0

(dlv) p runtime.g0
runtime.g {
        stackguard: *uint8 nil,
        stackbase: *uint8 nil,
        _defer: *runtime._defer nil,
        _panic: *runtime._panic nil,
        sched: runtime.gobuf {
                sp: *uint8 nil,
                pc: *uint8 nil,
                g: *runtime.g nil,},
        gcstack: *uint8 nil,
        gcsp: *uint8 nil,
        gcguard: *uint8 nil,
        stack0: *uint8 nil,
        entry: *uint8 nil,
        alllink: *runtime.g nil,
        param: *nil,
        status: 0,
        goid: 0,
        selgen: 0,
        waitreason: *int8 nil,
        schedlink: *runtime.g nil,
        readyonstop: 0,
        ispanic: 0,
        m: *runtime.m nil,
        lockedm: *runtime.m nil,
        idlem: *runtime.m nil,
        sig: 0,
        writenbuf: 0,
        writebuf: *uint8 nil,
        sigcode0: 0,
        sigcode1: 0,
        sigpc: 0,
        gopc: 0,
        end: [0]uint64 [],}
```

```bash
https://hopehook.com/post/golang_assembly/
https://guidao.github.io/asm.html
```



```bash
(dlv) l
> [Breakpoint 1] C.runtime.gogocall() /home/nch/workspace/go/src/pkg/runtime/asm_amd64.s:120 (hits total:1) (PC: 0x416dd9)
   115:         JMP     BX
   116:
   117: // void gogocall(Gobuf*, void (*fn)(void))
   118: // restore state from Gobuf but then call fn.
   119: // (call fn, returning to state in Gobuf)
=> 120: TEXT runtime·gogocall(SB), 7, $0
   121:         MOVQ    16(SP), AX              // fn
   122:         MOVQ    8(SP), BX               // gobuf
   123:         MOVQ    gobuf_g(BX), DX
   124:         get_tls(CX)
   125:         MOVQ    DX, g(CX)
(dlv) stack
0  0x0000000000416dd9 in C.runtime.gogocall
   at /home/nch/workspace/go/src/pkg/runtime/asm_amd64.s:120
1  0x000000000040ede0 in C.schedule
   at /home/nch/workspace/go/src/pkg/runtime/proc.c:908
2  0x000000000040eb05 in C.runtime.mstart
   at /home/nch/workspace/go/src/pkg/runtime/proc.c:762
3  0x0000000000416d86 in C._rt0_amd64
   at /home/nch/workspace/go/src/pkg/runtime/asm_amd64.s:74
4  0x0000000000000000 in ???
   at :0
   error: NULL address
(truncated)



(dlv) stack
0  0x0000000000416df3 in C.runtime.gogocall
   at /home/nch/workspace/go/src/pkg/runtime/asm_amd64.s:127
1  0x000000000040ede0 in C.schedule
   at /home/nch/workspace/go/src/pkg/runtime/proc.c:908
2  0x000000000040eb05 in C.runtime.mstart
   at /home/nch/workspace/go/src/pkg/runtime/proc.c:762
3  0x0000000000416d86 in C._rt0_amd64
   at /home/nch/workspace/go/src/pkg/runtime/asm_amd64.s:74
4  0x0000000000000000 in ???
   at :0
   error: NULL address
(truncated)
(dlv) si
> C.runtime.gogocall() /home/nch/workspace/go/src/pkg/runtime/asm_amd64.s:128 (PC: 0x416df6)
        asm_amd64.s:122 0x416dde        488b5c2408              mov rbx, qword ptr [rsp+0x8]
        asm_amd64.s:123 0x416de3        488b5310                mov rdx, qword ptr [rbx+0x10]
        asm_amd64.s:125 0x416de7        6448891425f0ffffff      mov qword ptr fs:[0xfffffff0], rdx
        asm_amd64.s:126 0x416df0        488b0a                  mov rcx, qword ptr [rdx]
        asm_amd64.s:127 0x416df3        488b23                  mov rsp, qword ptr [rbx]
=>      asm_amd64.s:128 0x416df6        488b5b08                mov rbx, qword ptr [rbx+0x8]
        asm_amd64.s:129 0x416dfa        53                      push rbx
        asm_amd64.s:130 0x416dfb        ffe0                    jmp rax
(dlv) l
> C.runtime.gogocall() /home/nch/workspace/go/src/pkg/runtime/asm_amd64.s:128 (PC: 0x416df6)
   123:         MOVQ    gobuf_g(BX), DX
   124:         get_tls(CX)
   125:         MOVQ    DX, g(CX)
   126:         MOVQ    0(DX), CX       // make sure g != nil
   127:         MOVQ    gobuf_sp(BX), SP        // restore SP
=> 128:         MOVQ    gobuf_pc(BX), BX
   129:         PUSHQ   BX
   130:         JMP     AX
   131:         POPQ    BX      // not reached
   132:
   133: // void mcall(void (*fn)(G*))
(dlv) stack
0  0x0000000000416df6 in C.runtime.gogocall
   at /home/nch/workspace/go/src/pkg/runtime/asm_amd64.s:128
```



