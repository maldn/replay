replay
======

parsing dota2 replays/demos in go.

I took this as an opportunity to learn go and protobuf

currently only non-compressed packets from demo.proto are parsed.
nothing fancy, but parses the whole protobuf-stream without errors.

i doubt a single soul will use this, but if you want:

i assume you have a working go environment

`go get code.google.com/p/goprotobuf/proto` if you dont have it already

`go run replay.go <replay file>`