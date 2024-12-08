```bash
[root@ph src]# ./make.bash
+ set -e
+ '[' '!' -f run.bash ']'
+ case "$(uname)" in
  ++ uname
+ ld --version
+ grep 'gold.* 2\.20'
+ for se_mount in /selinux /sys/fs/selinux
+ '[' -d /selinux -a -f /selinux/booleans/allow_execstack -a -x /usr/sbin/selinuxenabled ']'
+ for se_mount in /selinux /sys/fs/selinux
+ '[' -d /sys/fs/selinux -a -f /sys/fs/selinux/booleans/allow_execstack -a -x /usr/sbin/selinuxenabled ']'
+ echo '# Building C bootstrap tool.'
# Building C bootstrap tool.
+ echo cmd/dist
  cmd/dist
  ++ cd ..
  ++ pwd
+ export GOROOT=/home/nch/workspace/go
+ GOROOT=/home/nch/workspace/go
+ GOROOT_FINAL=/home/nch/workspace/go
+ DEFGOROOT='-DGOROOT_FINAL="/home/nch/workspace/go"'
+ mflag=
+ case "$GOHOSTARCH" in
+ gcc -O2 -Wall -Werror -ggdb -o cmd/dist/dist -Icmd/dist '-DGOROOT_FINAL="/home/nch/workspace/go"' cmd/dist/buf.c cmd/dist/build.c cmd/dist/buildgc.c cmd/dist/buildruntime.c cmd/dist/goc2c.c cmd/dist/main.c cmd/dist/unix.c cmd/dist/windows.c
  ++ ./cmd/dist/dist env -p
+ eval 'GOROOT="/home/nch/workspace/go"' 'GOBIN="/home/nch/workspace/go/bin"' 'GOARCH="amd64"' 'GOOS="linux"' 'GOHOSTARCH="amd64"' 'GOHOSTOS="linux"' 'GOTOOLDIR="/home/nch/workspace/go/pkg/tool/linux_amd64"' 'GOCHAR="6"' 'PATH="/home/nch/workspace/go/bin:/home/idss/dbaudit-2.2.5/jdk8u333/bin:/home/dbaudit-2.2.5/jdk8u333/bin:/opt/idss/phantom/run/mysql/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/home/dbaudit-2.2.5/jdk8u333/jre/lib:/home/dbaudit-2.2.5/jdk8u333/bin:/home/idss/dbaudit-2.2.5/jdk8u333/jre/lib:/home/idss/dbaudit-2.2.5/jdk8u333/bin:/home/nch/workspace/go/bin:/home/nch/workspace/gpath/bin:/root/bin:/home/nch/workspace/go/bin:/home/nch/workspace/gpath/bin"'
  ++ GOROOT=/home/nch/workspace/go
  ++ GOBIN=/home/nch/workspace/go/bin
  ++ GOARCH=amd64
  ++ GOOS=linux
  ++ GOHOSTARCH=amd64
  ++ GOHOSTOS=linux
  ++ GOTOOLDIR=/home/nch/workspace/go/pkg/tool/linux_amd64
  ++ GOCHAR=6
  ++ PATH=/home/nch/workspace/go/bin:/home/idss/dbaudit-2.2.5/jdk8u333/bin:/home/dbaudit-2.2.5/jdk8u333/bin:/opt/idss/phantom/run/mysql/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/home/dbaudit-2.2.5/jdk8u333/jre/lib:/home/dbaudit-2.2.5/jdk8u333/bin:/home/idss/dbaudit-2.2.5/jdk8u333/jre/lib:/home/idss/dbaudit-2.2.5/jdk8u333/bin:/home/nch/workspace/go/bin:/home/nch/workspace/gpath/bin:/root/bin:/home/nch/workspace/go/bin:/home/nch/workspace/gpath/bin
+ echo

+ '[' '' = --dist-tool ']'
+ echo '# Building compilers and Go bootstrap tool for host, linux/amd64.'
# Building compilers and Go bootstrap tool for host, linux/amd64.
+ buildall=-a
+ '[' '' = --no-clean ']'
+ ./cmd/dist/dist bootstrap -a -v
  lib9
  libbio
  libmach
  misc/pprof
  cmd/addr2line
  cmd/cov
  cmd/nm
  cmd/objdump
  cmd/pack
  cmd/prof
  cmd/cc
  cmd/gc
  cmd/6l
  cmd/6a
  cmd/6c
  cmd/6g
  pkg/runtime
  pkg/errors
  pkg/sync/atomic
  pkg/sync
  pkg/io
  pkg/unicode
  pkg/unicode/utf8
  pkg/unicode/utf16
  pkg/bytes
  pkg/math
  pkg/strings
  pkg/strconv
  pkg/bufio
  pkg/sort
  pkg/container/heap
  pkg/encoding/base64
  pkg/syscall
  pkg/time
  pkg/os
  pkg/reflect
  pkg/fmt
  pkg/encoding/json
  pkg/flag
  pkg/path/filepath
  pkg/path
  pkg/io/ioutil
  pkg/log
  pkg/regexp/syntax
  pkg/regexp
  pkg/go/token
  pkg/go/scanner
  pkg/go/ast
  pkg/go/parser
  pkg/os/exec
  pkg/net/url
  pkg/text/template/parse
  pkg/text/template
  pkg/go/doc
  pkg/go/build
  cmd/go
+ mv cmd/dist/dist /home/nch/workspace/go/pkg/tool/linux_amd64/dist
+ /home/nch/workspace/go/pkg/tool/linux_amd64/go_bootstrap clean -i std
+ echo

+ '[' amd64 '!=' amd64 -o linux '!=' linux ']'
+ echo '# Building packages and commands for linux/amd64.'
# Building packages and commands for linux/amd64.
+ /home/nch/workspace/go/pkg/tool/linux_amd64/go_bootstrap install -gcflags '' -ldflags '' -v std
  runtime
  errors
  sync/atomic
  unicode
  unicode/utf8
  math
  unicode/utf16
  crypto/subtle
  container/list
  container/ring
  image/color
  sync
  io
  syscall
  hash
  crypto/cipher
  crypto
  hash/crc32
  crypto/hmac
  hash/adler32
  crypto/md5
  crypto/sha1
  crypto/sha256
  crypto/sha512
  hash/crc64
  hash/fnv
  bytes
  strings
  path
  bufio
  text/tabwriter
  html
  time
  os
  strconv
  sort
  math/rand
  math/cmplx
  os/signal
  path/filepath
  container/heap
  compress/bzip2
  reflect
  regexp/syntax
  io/ioutil
  net/url
  os/exec
  encoding/base64
  crypto/aes
  crypto/rc4
  encoding/ascii85
  encoding/base32
  image
  encoding/pem
  regexp
  image/draw
  image/jpeg
  fmt
  encoding/binary
  debug/dwarf
  crypto/des
  index/suffixarray
  flag
  go/token
  text/template/parse
  log
  debug/elf
  debug/macho
  debug/pe
  go/scanner
  encoding/json
  encoding/xml
  compress/flate
  math/big
  go/ast
  text/template
  mime
  runtime/debug
  encoding/gob
  runtime/pprof
  text/scanner
  html/template
  cmd/yacc
  archive/tar
  go/doc
  go/parser
  go/printer
  crypto/elliptic
  crypto/rand
  crypto/rsa
  crypto/dsa
  compress/gzip
  encoding/asn1
  archive/zip
  go/build
  cmd/cgo
  cmd/fix
  cmd/gofmt
  cmd/vet
  crypto/x509/pkix
  compress/lzw
  compress/zlib
  crypto/x509
  cmd/api
  crypto/ecdsa
  database/sql/driver
  debug/gosym
  encoding/csv
  database/sql
  encoding/hex
  image/gif
  image/png
  testing
  testing/iotest
  testing/quick
  runtime/cgo
  net
  os/user
  crypto/tls
  net/textproto
  log/syslog
  mime/multipart
  net/mail
  net/http
  net/smtp
  cmd/go
  expvar
  net/http/pprof
  net/http/cgi
  net/http/httptest
  net/http/httputil
  net/rpc
  cmd/godoc
  net/http/fcgi
  net/rpc/jsonrpc
+ echo

+ '[' '' '!=' --no-banner ']'
+ /home/nch/workspace/go/pkg/tool/linux_amd64/dist banner

---
Installed Go for linux/amd64 in /home/nch/workspace/go
Installed commands in /home/nch/workspace/go/bin
[root@ph src]#
```