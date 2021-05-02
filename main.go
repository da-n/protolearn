package main

import (
	"fmt"
	addressbookpb "github.com/da-n/protolearn/src/addressbook"
	complexpb "github.com/da-n/protolearn/src/complex"
	enumpb "github.com/da-n/protolearn/src/enum_example"
	simplepb "github.com/da-n/protolearn/src/simple"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	sm := doSimple()
	readAndWriteDemo(sm)
	jsonDemo(sm)
	doEnum()
	doComplex()
	readAndWritePerson()
}

func readAndWritePerson() {
	pm1 := addressbookpb.Person{
		Name: "Daniel",
		Id: 1,
		Email: "dan@bitmono.com",
		Phones: []*addressbookpb.Person_PhoneNumber{
			&addressbookpb.Person_PhoneNumber{
				Number: "07403123456",
				Type: addressbookpb.Person_MOBILE,
			},
		},
	}
	pm2 := addressbookpb.Person{
		Name: "Michael",
		Id: 2,
		Email: "michael@bitmono.com",
		Phones: []*addressbookpb.Person_PhoneNumber{
			&addressbookpb.Person_PhoneNumber{
				Number: "07403123456",
				Type: addressbookpb.Person_MOBILE,
			},
		},
	}
	am := addressbookpb.AddressBook{
		People: []*addressbookpb.Person{
			&pm1,
			&pm2,
		},
	}
	fmt.Println(am)
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id: 1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id: 2,
				Name: "Second message",
			},
			&complexpb.DummyMessage{
				Id: 3,
				Name: "Third message",
			},
		},
	}
	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id: 42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}
	fmt.Println(em)
}

func jsonDemo(sm  proto.Message) {
	smAsString := toJson(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJson(smAsString, sm2)
	fmt.Println("Successfully created proto struct:", sm2)
}

func toJson(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func fromJson(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON in the pb struct", err)
	}
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}
	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err2)
		return err2
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id: 1,
		IsSimple: false,
		Name: "The Simple Message",
		SampleList: []int32{1,4,7,8,9},
	}
	return &sm
}