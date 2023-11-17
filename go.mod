module goutils

go 1.21.4

require (
	github.com/huntelaar112/goutils/sched v0.0.0-20231117044315-e27db452652b
	github.com/huntelaar112/goutils/timeutils v0.0.0-20231117044315-e27db452652b
	github.com/huntelaar112/goutils/utils v0.0.0-20231117074037-158dfab3481d
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/google/uuid v1.4.0 // indirect
	github.com/rs/xid v1.5.0 // indirect
	github.com/sonnt85/gosutils/endec v0.0.0-20230927031609-2b3046a0b311 // indirect
	github.com/sonnt85/gosutils/sched v0.0.0-20230927031609-2b3046a0b311 // indirect
	github.com/sonnt85/gotimeutils v0.0.0-20221011032526-b825ab5ef455 // indirect
	golang.org/x/crypto v0.15.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
)

replace github.com/huntelaar112/goutils/sched => ./sched

replace github.com/huntelaar112/goutils/timeutils => ./timeutils

replace github.com/huntelaar112/goutils/utils => ./utils
