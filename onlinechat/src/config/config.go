package config

// Net Server config
var SERVER_ADDR string = "127.0.0.1:8090"
var TLS_SERVER_ADDR string = "127.0.0.1:8091"
var BUFFER_SIZE = 4096

// TLS socket cert
var TLS_PUB_FILE string = ""
var TLS_RPI_FILE string = ""

var CHANNEL_BUF_SIZE = 1 << 20

// Protocol type
// ProtocolBuffer(PB), json or other proto.
var PROTO_TYPE string = "PB"

// database config
var DB_TYPE string = "mysql"
var DB_CONFIG string = "root:huangsu@tcp(127.0.0.1:3306)/golangonlinechat"

// session config
// memory, file, database, redis?
var SESSION_TYPE = "memory"

// Expired time for user
var EXPIRED_TIME_SEC = 10 * 60 // 10min

// Default online users, base on the problem
var DEFAULT_ONLINE_USER = 10240
