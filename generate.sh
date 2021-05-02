protoc -I src/ --go_out=src/ --go_opt=paths=source_relative src/simple/simple.proto
protoc -I src/ --go_out=src/ --go_opt=paths=source_relative src/enum_example/enum_example.proto
protoc -I src/ --go_out=src/ --go_opt=paths=source_relative src/complex/complex.proto
protoc -I src/ --go_out=src/ --go_opt=paths=source_relative src/addressbook/addressbook.proto