package myreplication

const (
	_COM_SLEEP               = 0x00
	_COM_QUIT                = 0x01
	_COM_INIT_DB             = 0x02
	_COM_QUERY               = 0x03
	_COM_FIELD_LIST          = 0x04
	_COM_CREATE_DB           = 0x05
	_COM_DROP_DB             = 0x06
	_COM_REFRESH             = 0x07
	_COM_SHUTDOWN            = 0x08
	_COM_STATISTICS          = 0x09
	_COM_PROCESS_INFO        = 0x0a
	_COM_CONNECT             = 0x0b
	_COM_PROCESS_KILL        = 0x0c
	_COM_DEBUG               = 0x0d
	_COM_PING                = 0x0e
	_COM_TIME                = 0x0f
	_COM_DELAYED_INSERT      = 0x10
	_COM_CHANGE_USER         = 0x11
	_COM_BINLOG_DUMP         = 0x12
	_COM_TABLE_DUMP          = 0x13
	_COM_CONNECT_OUT         = 0x14
	_COM_REGISTER_SLAVE      = 0x15
	_COM_STMT_PREPARE        = 0x16
	_COM_STMT_EXECUTE        = 0x17
	_COM_STMT_SEND_LONG_DATA = 0x18
	_COM_STMT_CLOSE          = 0x19
	_COM_STMT_RESET          = 0x1a
	_COM_SET_OPTION          = 0x1b
	_COM_STMT_FETCH          = 0x1c
	_COM_DAEMON              = 0x1d
	_COM_BINLOG_DUMP_GTID    = 0x1e
	_COM_RESET_CONNECTION    = 0x1f

	MYSQL_TYPE_DECIMAL     = 0x00
	MYSQL_TYPE_TINY        = 0x01
	MYSQL_TYPE_SHORT       = 0x02
	MYSQL_TYPE_LONG        = 0x03
	MYSQL_TYPE_FLOAT       = 0x04
	MYSQL_TYPE_DOUBLE      = 0x05
	MYSQL_TYPE_NULL        = 0x06
	MYSQL_TYPE_TIMESTAMP   = 0x07
	MYSQL_TYPE_LONGLONG    = 0x08
	MYSQL_TYPE_INT24       = 0x09
	MYSQL_TYPE_DATE        = 0x0a
	MYSQL_TYPE_TIME        = 0x0b
	MYSQL_TYPE_DATETIME    = 0x0c
	MYSQL_TYPE_YEAR        = 0x0d
	MYSQL_TYPE_NEWDATE     = 0x0e
	MYSQL_TYPE_VARCHAR     = 0x0f
	MYSQL_TYPE_BIT         = 0x10
	MYSQL_TYPE_TIMESTAMP2  = 0x11
	MYSQL_TYPE_DATETIME2   = 0x12
	MYSQL_TYPE_TIME2       = 0x13
	MYSQL_TYPE_NEWDECIMAL  = 0xf6
	MYSQL_TYPE_ENUM        = 0xf7
	MYSQL_TYPE_SET         = 0xf8
	MYSQL_TYPE_TINY_BLOB   = 0xf9
	MYSQL_TYPE_MEDIUM_BLOB = 0xfa
	MYSQL_TYPE_LONG_BLOB   = 0xfb
	MYSQL_TYPE_BLOB        = 0xfc
	MYSQL_TYPE_VAR_STRING  = 0xfd
	MYSQL_TYPE_STRING      = 0xfe
	MYSQL_TYPE_GEOMETRY    = 0xff

	_MYSQL_EOF = 0xFE
	_MYSQL_OK  = 0x00
	_MYSQL_ERR = 0xFF

	_UNKNOWN_EVENT            = 0x00
	_START_EVENT_V3           = 0x01
	_QUERY_EVENT              = 0x02
	_STOP_EVENT               = 0x03
	_ROTATE_EVENT             = 0x04
	_INTVAR_EVENT             = 0x05
	_LOAD_EVENT               = 0x06
	_SLAVE_EVENT              = 0x07
	_CREATE_FILE_EVENT        = 0x08
	_APPEND_BLOCK_EVENT       = 0x09
	_EXEC_LOAD_EVENT          = 0x0a
	_DELETE_FILE_EVENT        = 0x0b
	_NEW_LOAD_EVENT           = 0x0c
	_RAND_EVENT               = 0x0d
	_USER_VAR_EVENT           = 0x0e
	_FORMAT_DESCRIPTION_EVENT = 0x0f
	_XID_EVENT                = 0x10
	_BEGIN_LOAD_QUERY_EVENT   = 0x11
	_EXECUTE_LOAD_QUERY_EVENT = 0x12
	_TABLE_MAP_EVENT          = 0x13
	_WRITE_ROWS_EVENTv0       = 0x14
	_UPDATE_ROWS_EVENTv0      = 0x15
	_DELETE_ROWS_EVENTv0      = 0x16
	_WRITE_ROWS_EVENTv1       = 0x17
	_UPDATE_ROWS_EVENTv1      = 0x18
	_DELETE_ROWS_EVENTv1      = 0x19
	_INCIDENT_EVENT           = 0x1a
	_HEARTBEAT_EVENT          = 0x1b
	_IGNORABLE_EVENT          = 0x1c
	_ROWS_QUERY_EVENT         = 0x1d
	_WRITE_ROWS_EVENTv2       = 0x1e
	_UPDATE_ROWS_EVENTv2      = 0x1f
	_DELETE_ROWS_EVENTv2      = 0x20
	_GTID_EVENT               = 0x21
	_ANONYMOUS_GTID_EVENT     = 0x22
	_PREVIOUS_GTIDS_EVENT     = 0x23

	_FORMAT_DESCRIPTION_LENGTH_QUERY_POSITION    = 1
	_FORMAT_DESCRIPTION_LENGTH_DELETEV1_POSITION = 22
	_FORMAT_DESCRIPTION_LENGTH_UPDATEV1_POSITION = 23
	_FORMAT_DESCRIPTION_LENGTH_WRITEV1_POSITION  = 24

	INVALID_INT_EVENT    = 0x00
	LAST_INSERT_ID_EVENT = 0x01
	INSERT_ID_EVENT      = 0x02

	MYSQL_FLAG_COLUMN_NON_NULLABLE = 0x01
	MYSQL_FLAG_COLUMN_PRIMARY_KEY  = 0x02
	MYSQL_FLAG_COLUMN_UNIQUE_KEY   = 0x04
	MYSQL_FLAG_COLUMN_MULTIPLE_KEY = 0x08
	MYSQL_FLAG_COLUMN_BLOB         = 0x10
	MYSQL_FLAG_COLUMN_UNSIGNED     = 0x20
	MYSQL_FLAG_COLUMN_ZERO_FILL    = 0x40
	MYSQL_FLAG_COLUMN_BINARY       = 0x80

	_DIGITS_PER_INTEGER = 9
)
