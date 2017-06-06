package nmp

const (
	NMP_OP_READ      = 0
	NMP_OP_READ_RSP  = 1
	NMP_OP_WRITE     = 2
	NMP_OP_WRITE_RSP = 3
)

const (
	NMP_ERR_OK       = 0
	NMP_ERR_EUNKNOWN = 1
	NMP_ERR_ENOMEM   = 2
	NMP_ERR_EINVAL   = 3
	NMP_ERR_ETIMEOUT = 4
	NMP_ERR_ENOENT   = 5
)

// First 64 groups are reserved for system level newtmgr commands.
// Per-user commands are then defined after group 64.

const (
	NMP_GROUP_DEFAULT = 0
	NMP_GROUP_IMAGE   = 1
	NMP_GROUP_STAT    = 2
	NMP_GROUP_CONFIG  = 3
	NMP_GROUP_LOG     = 4
	NMP_GROUP_CRASH   = 5
	NMP_GROUP_SPLIT   = 6
	NMP_GROUP_RUN     = 7
	NMP_GROUP_FS      = 8
	NMP_GROUP_PERUSER = 64
)

// Default group (0).
const (
	NMP_ID_DEF_ECHO           = 0
	NMP_ID_DEF_CONS_ECHO_CTRL = 1
	NMP_ID_DEF_TASKSTAT       = 2
	NMP_ID_DEF_MPSTAT         = 3
	NMP_ID_DEF_DATETIME_STR   = 4
	NMP_ID_DEF_RESET          = 5
)

// Image group (1).
const (
	NMP_ID_IMAGE_STATE    = 0
	NMP_ID_IMAGE_UPLOAD   = 1
	NMP_ID_IMAGE_CORELIST = 3
	NMP_ID_IMAGE_CORELOAD = 4
	NMP_ID_IMAGE_ERASE    = 5
)

// Stat group (2).
const (
	NMP_ID_STAT_READ = 0
	NMP_ID_STAT_LIST = 1
)

// Config group (3).
const (
	NMP_ID_CONFIG_VAL = 0
)

// Log group (4).
const (
	NMP_ID_LOG_SHOW        = 0
	NMP_ID_LOG_CLEAR       = 1
	NMP_ID_LOG_APPEND      = 2
	NMP_ID_LOG_MODULE_LIST = 3
	NMP_ID_LOG_LEVEL_LIST  = 4
	NMP_ID_LOG_LIST        = 5
)

// Crash group (5).
const (
	NMP_ID_CRASH_TRIGGER = 0
)

// Run group (7).
const (
	NMP_ID_RUN_TEST = 0
	NMP_ID_RUN_LIST = 1
)

// File system group (8).
const (
	NMP_ID_FS_FILE = 0
)
