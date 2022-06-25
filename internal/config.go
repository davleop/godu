package godu

/* File Flags (struct dir -> flags) */
const FF_DIR    = 0x01
const FF_FILE   = 0x02
const FF_ERR    = 0x04  // error while reading this item
const FF_OTHFS  = 0x08  // excluded because it was another filesystem
const FF_EXL    = 0x10  // excluded using exclude patterns
const FF_SERR   = 0x20  // error in subdirectory
const FF_HLNKC  = 0x40  // hard link candidate (file with st_nlink > 1)
const FF_BSEL   = 0x80  // selected
const FF_EXT    = 0x100 // extended struct available
const FF_KERNFS = 0x200 // excluded because it was a Linux pseudo filesystem
const FF_FRMLNK = 0x400 // excluded because it was a firmlink

/* Program states */
const ST_CALC   = 0
const ST_BROWSE = 1
const ST_DEL    = 2
const ST_HELP   = 3
const ST_SHELL  = 4
const ST_QUIT   = 5

type Dir struct {
  size int64
  asize int64
  ino uint64
  dev uint64
  parent *Dir
  next *Dir
  prev *Dir
  sub *Dir
  hlnk *Dir
  items int32
  flags uint16
  name []byte
}

type DirExt struct {
  mtime uint64
  uid int32
  gid int32
  mode uint32
}

