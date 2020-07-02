package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"os"
	"strconv"
	"time"
)

//Paciente es la estructura con la que se maneja el algoritmo
type Paciente struct {
	fever               bool
	tiredness           bool
	dry_cough           bool
	difficulty_brithing bool
	sore_throat         bool
	none_symtons        bool
	age0_9              bool
	age10_19            bool
	age20_24            bool
	age25_59            bool
	age60               bool
	gender_female       bool
	gender_male         bool
	severity_mild       bool
	severity_moderate   bool
	severity_none       bool
	severity_severe     bool
	contact_yes         bool
	country             string

	sintoms  int
	others   int
	setClass string
}

//knnPt es la estructura para los puntos del algoritmo KNN
type knnPt struct {
	distance float64
	x        int
	y        int
	contagio string
}

func proccesofChossing(k *knnPt, ATK int, DEF int, p Paciente) {
	absX := math.Abs(float64(ATK - p.sintoms))
	absY := math.Abs(float64(DEF - p.others))
	distance := math.Sqrt(math.Pow(absX, 2) + math.Pow(absY, 2))
	k.distance = distance
	k.x = p.sintoms
	k.y = p.others
	k.contagio = p.setClass
}

func getAbs(p *Paciente) {
	contSim := 0
	contOth := 0

	if p.fever == true {
		contSim += 2
	}
	if p.tiredness == true {
		contSim += 2
	}
	if p.dry_cough == true {
		contSim += 2
	}
	if p.difficulty_brithing == true {
		contSim += 2
	}
	if p.sore_throat == true {
		contSim += 2
	}
	if p.age10_19 == true {
		contOth++
	}
	if p.age20_24 == true {
		contOth++
	}
	if p.age25_59 == true {
		contOth += 2

	}
	if p.age60 == true {
		contOth += 4
	}
	if p.gender_female == true {
		contOth++
	}
	if p.gender_male == true {
		contOth += 2
	}
	if p.severity_mild == true {
		contOth++
	}
	if p.severity_moderate == true {
		contOth += 2
	}
	if p.severity_severe == true {
		contOth += 3
	}
	if p.contact_yes == true {
		contOth += 4
	}

	p.sintoms = contSim
	p.others = contOth

	if p.sintoms+p.others > 10 {
		p.setClass = "contagiado"
	}
	if p.sintoms+p.others <= 10 {
		p.setClass = "no contagiado"
	}
	contSim = 0
	contOth = 0
}

const localAddr = "192.168.0.28:8000"

const (
	cnum = iota
	opContagiado
	opNoContagiado
)

var addrs = []string{
	"192.168.0.27:8000",
	"192.168.0.9:8000"}

var chInfo chan map[string]int

type estadoPaciente struct {
	Code int
	Addr string
	Op   int
}

//KNN no sé
func KNN() int {
	data := "dataset2.csv"
	var i = 0
	var set = [100]Paciente{}
	file, err := os.Open(data)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	i = 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
		}
		fever, _ := strconv.ParseBool(record[0])
		set[i].fever = fever

		tiredness, _ := strconv.ParseBool(record[1])
		set[i].tiredness = tiredness

		dry_cough, _ := strconv.ParseBool(record[2])
		set[i].dry_cough = dry_cough

		difficulty_brithing, _ := strconv.ParseBool(record[3])
		set[i].difficulty_brithing = difficulty_brithing

		sore_throat, _ := strconv.ParseBool(record[4])
		set[i].sore_throat = sore_throat

		none_symtons, _ := strconv.ParseBool(record[5])
		set[i].none_symtons = none_symtons

		age0_9, _ := strconv.ParseBool(record[6])
		set[i].age0_9 = age0_9

		age10_19, _ := strconv.ParseBool(record[7])
		set[i].age10_19 = age10_19

		age20_24, _ := strconv.ParseBool(record[8])
		set[i].age20_24 = age20_24

		age25_59, _ := strconv.ParseBool(record[9])
		set[i].age25_59 = age25_59

		age60, _ := strconv.ParseBool(record[10])
		set[i].age60 = age60

		gender_female, _ := strconv.ParseBool(record[11])
		set[i].gender_female = gender_female

		gender_male, _ := strconv.ParseBool(record[12])
		set[i].gender_male = gender_male

		severity_mild, _ := strconv.ParseBool(record[13])
		set[i].severity_mild = severity_mild

		severity_moderate, _ := strconv.ParseBool(record[14])
		set[i].severity_moderate = severity_moderate

		severity_none, _ := strconv.ParseBool(record[15])
		set[i].severity_none = severity_none

		severity_severe, _ := strconv.ParseBool(record[16])
		set[i].severity_severe = severity_severe

		contact_yes, _ := strconv.ParseBool(record[17])
		set[i].contact_yes = contact_yes

		country := record[18]
		set[i].country = country
		i++
	}

	for i := 0; i < 100; i++ {
		getAbs(&set[i])
	}

	var prueba = Paciente{}
	prueba.fever = true
	prueba.tiredness = true
	prueba.dry_cough = true
	prueba.difficulty_brithing = true
	prueba.sore_throat = false
	prueba.age0_9 = false
	prueba.age10_19 = false
	prueba.age20_24 = false
	prueba.age25_59 = true
	prueba.age60 = false
	prueba.gender_female = false
	prueba.gender_male = true
	prueba.severity_mild = false
	prueba.severity_moderate = false
	prueba.severity_none = true
	prueba.severity_severe = false
	prueba.contact_yes = true
	prueba.country = "Perú"

	getAbs(&prueba)

	var getPoints = [100]knnPt{}
	for i := 0; i < 100; i++ {
		go proccesofChossing(&getPoints[i], prueba.sintoms, prueba.others, set[i])
		time.Sleep(30)
	}
	proccesofChossing(&getPoints[99], prueba.sintoms, prueba.others, set[99])
	for i := 1; i < 100; i++ {
		for j := 0; j < 100-i; j++ {
			if getPoints[j].distance > getPoints[j+1].distance {
				getPoints[j], getPoints[j+1] = getPoints[j+1], getPoints[j]
			}
		}
	}
	count := 0
	var option int

	for i := 0; i < 6; i++ {
		fmt.Printf("(Sintomas:%d, Others:%d, d: %f, estado: %s)\n", getPoints[i].x, getPoints[i].y, getPoints[i].distance, getPoints[i].contagio)
		if getPoints[i].contagio == "contagiado" {
			count++
		}
	}
	fmt.Println("\n")
	if count > 3 /*deberia ser 1.5*/ { //Si el 50% de los casos "cercanos es mas de 50%"
		fmt.Println("Estas posiblemente contagiado")
		option = 0
	} else {
		fmt.Println("Estas sano")
		option = 1
	}
	return option
}

func main() {
	chInfo = make(chan map[string]int)
	go func() { chInfo <- map[string]int{} }()
	go server()
	time.Sleep(time.Millisecond * 100)
	var op int
	for {
		fmt.Print("Generated option: ")
		fmt.Scanf("Choosed %d", KNN())
		msg := estadoPaciente{cnum, localAddr, op}
		for _, addr := range addrs {
			send(addr, msg)
		}
	}

}

func server() {
	if ln, err := net.Listen("tcp", localAddr); err != nil {
		log.Panicln("Can't start listener on", localAddr)
	} else {
		defer ln.Close()
		fmt.Println("Listening on ", localAddr)
		for {
			if conn, err := ln.Accept(); err != nil {
				log.Println("Can't accept", conn.RemoteAddr())
			} else {
				go handle(conn)
			}
		}
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	dec := json.NewDecoder(conn)
	var msg estadoPaciente
	if err := dec.Decode(&msg); err != nil {
		log.Println("Can't decode from", conn.RemoteAddr())
	} else {
		fmt.Println(msg)
		switch msg.Code {
		case cnum:
			concensus(conn, msg)
		}
	}
}
func concensus(conn net.Conn, msg estadoPaciente) {
	info := <-chInfo
	info[msg.Addr] = msg.Op
	if len(info) == len(addrs) {
		cContagio, cNoContagio := 0, 0
		for _, op := range info {
			if op == opContagiado {
				cContagio++
			} else {
				cNoContagio++
			}
		}
		if cContagio > cNoContagio {
			fmt.Println("Esta contagiado!")
		} else {
			fmt.Println("Te salvaste (Esta vez)")
		}
		info = map[string]int{}
	}
	go func() { chInfo <- info }()
}
func send(remoteAddr string, msg estadoPaciente) {
	if conn, err := net.Dial("tcp", remoteAddr); err != nil {
		log.Println("Can't dail", remoteAddr)
	} else {
		defer conn.Close()
		fmt.Println("Sending to ", remoteAddr)
		enc := json.NewEncoder(conn)
		enc.Encode(msg)
	}
}
