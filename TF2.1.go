package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Option struct {
	Value, Id, Text string
	Selected        bool
}

const HTML = `
<!DOCTYPE html>
<html lang="en">
     <head>
        <meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">

        <title>selected attribute</title>
<link rel="stylesheet" type="text/css" href="fondo.css">
<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Lobster">
<link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
    </head>



    <body style=' margin: 0;
    padding: 0;
	background: url("https://www.entornointeligente.com/wp-content/uploads/2020/02/entornointeligente_roberto_pocaterra_pocaterra_buenos_aires_lo_que_tienes_que_saber_sobre_las_nuevas_cifras_del_coronavirus.jpg");    
	background-size: cover;
	background-position: center;
	background-repeat: no-repeat;
    font-family: sans-serif;'>

<div class="w3-container w3-lobster">                                                 
        <h1 style='font-family: "Lobster", serif;  position: absolute; top: 50px; left: 670px;'>Bienvenido al Analizador Covid</h1>
</div>

<div class="w3-container w3-lobster">                                                 
		<h1 style='font-family: "Lobster", serif; color: #05687; position: absolute; top: 110px; left: 100px;'>Por favor</h1>
		<h1 style='font-family: "Lobster", serif; color: #05687; position: absolute; top: 170px; left: 30px;'>tomate tu tiempo y </h1>
		<h1 style='font-family: "Lobster", serif; color: #05687; position: absolute; top: 230px; left: 100px;'>selecciona </h1>
		<h1 style='font-family: "Lobster", serif; color: #05687; position: absolute; top: 300px; left: 30px;'>la opcion correcta </h1>
		<h1 style='font-family: "Lobster", serif; color: #05687; position: absolute; top: 370px; left: 42px;'>para asegurar </h1>
		<h1 style='font-family: "Lobster", serif; color: #05687; position: absolute; top: 440px; left: 35px;'>un buen resultado </h1>
	
</div>


<div class="flip-card style = left: 30px;'">
  <div class="flip-card-inner">
    <div class="flip-card-front">
      <img src="https://i.ytimg.com/vi/mA1qCnk4Lg4/hqdefault.jpg" alt="" width="330" height="600">
    </div>
  </div>
</div>

        <form method="GET">                                                                                                   
            <div  style='color: black; font-weight:  bolder; font-size: 20px; width: 100%; position: absolute; top: 200px; left: 530px; font-family: "Comic Sans MS", cursive, sans-serif;'>
                <label>fever:</label>
				<select id="fever" name="fever">
                    {{range .}}
					<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
                    {{end}}
				</select>
				<label>tiredness:</label>
				<select id="tiredness" name="tiredness">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<label>dryCough:</label>
			<select id="dryCough" name="dryCough">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<label>difficultyBrithing:</label>
			<select id="difficultyBrithing" name="difficultyBrithing">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<br></br>
			<label>soreThroat:</label>
			<select id="soreThroat" name="soreThroat">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>

			

			<label>noneSymtons:</label>
			<select id="noneSymtons" name="noneSymtons">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<label>age0_9:</label>
			<select id="age0_9" name="age0_9">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<label>age10_19:</label>
			<select id="age10_19" name="age10_19">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<br></br>
			<label>age20_24:</label>
			<select id="age20_24" name="age20_24">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<label>age25_59:</label>
			<select id="age25_59" name="age25_59">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<label>age60:</label>
			<select id="age60" name="age60">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<label>genderFemale:</label>
			<select id="genderFemale" name="genderFemale">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<br></br>
			<label>genderMale:</label>
			<select id="genderMale" name="genderMale">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<label>severityMild:</label>
			<select id="severityMild" name="severityMild">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<label>severityModerate:</label>
			<select id="severityModerate" name="severityModerate">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<label>severityNone:</label>
			<select id="severityNone" name="severityNone">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<br></br>
			<label>severitySevere:</label>
			<select id="severitySevere" name="severitySevere">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<label>contactYes:</label>
			<select id="contactYes" name="contactYes">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
			<label>country:</label>
			<select id="country" name="country">
				{{range .}}
				<option value="{{.Value}}" id="{{.Id}}" {{if .Selected}}selected{{end}}>{{.Text}}</option>
				{{end}}
			</select>
            </div>
           <div style="position: absolute;
  top: 500px;
  left: 850px;
  width: 300px;
  height: 200px;">
            <input style='display: inline-block; padding: 15px 25px; font-weight:  bolder;  font-size: 24px; cursor: pointer; text-align: center; text-decoration: none; outline: none; color: black;
                   background-color: #009C8C; border: none; border-radius: 15px; box-shadow: 0 9px #999;' type="submit" value="Analizar" align="center" name = "submit">
           </div>
        </form>

    </body>
</html>
`

var placesPageTmpl *template.Template = template.Must(template.New("	").Parse(HTML))

const localAddr = "192.168.0.9:8000"

const (
	cnum = iota
	opContagiado
	opNoContagiado
)

var chInfo chan map[string]int

//Paciente es la estructura con la que se maneja el algoritmo
type Paciente struct {
	fever              bool
	tiredness          bool
	dryCough           bool
	difficultyBrithing bool
	soreThroat         bool
	noneSymtons        bool
	age0_9             bool
	age10_19           bool
	age20_24           bool
	age25_59           bool
	age60              bool
	genderFemale       bool
	genderMale         bool
	severityMild       bool
	severityModerate   bool
	severityNone       bool
	severitySevere     bool
	contactYes         bool
	country            string

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

type estadoPaciente struct {
	Code int
	Addr string
	Op   int
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

var addrs = []string{
	"192.168.0.27:8000",
	"192.168.0.28:8000"}

func main() {
	fmt.Print(addrs)
	fmt.Println()
	http.HandleFunc("/", name)
	http.ListenAndServe(":8080", nil)
}

//KNN process for a person
func KNN(prueba *Paciente) int {
	data := "covid19.csv"
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
		set[i].dryCough = dry_cough

		difficulty_brithing, _ := strconv.ParseBool(record[3])
		set[i].difficultyBrithing = difficulty_brithing

		sore_throat, _ := strconv.ParseBool(record[4])
		set[i].soreThroat = sore_throat

		none_symtons, _ := strconv.ParseBool(record[5])
		set[i].noneSymtons = none_symtons

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
		set[i].genderFemale = gender_female

		gender_male, _ := strconv.ParseBool(record[12])
		set[i].genderMale = gender_male

		severity_mild, _ := strconv.ParseBool(record[13])
		set[i].severityMild = severity_mild

		severity_moderate, _ := strconv.ParseBool(record[14])
		set[i].severityModerate = severity_moderate

		severity_none, _ := strconv.ParseBool(record[15])
		set[i].severityNone = severity_none

		severity_severe, _ := strconv.ParseBool(record[16])
		set[i].severitySevere = severity_severe

		contact_yes, _ := strconv.ParseBool(record[17])
		set[i].contactYes = contact_yes

		country := record[18]
		set[i].country = country
		i++
	}

	for i := 0; i < 100; i++ {
		getAbs(&set[i])
	}

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
		option = 0
		fmt.Println("Estas posiblemente contagiado")
	} else {
		option = 1
		fmt.Println("No esta contagiado")
	}
	return option
}

func name(w http.ResponseWriter, r *http.Request) {
	var boolean bool
	var paciente = Paciente{}

	if r.FormValue("fever") == "si" {
		boolean = true
		paciente.fever = true
	} else {
		boolean = false
		paciente.fever = false
	}
	if r.FormValue("tiredness") == "si" {
		boolean = true
		paciente.tiredness = true
	} else {
		boolean = false
		paciente.tiredness = false
	}
	if r.FormValue("dryCough") == "si" {
		boolean = true
		paciente.dryCough = true
	} else {
		boolean = false
		paciente.dryCough = false
	}
	if r.FormValue("difficultyBrithing") == "si" {
		boolean = true
		paciente.difficultyBrithing = true
	} else {
		boolean = false
		paciente.difficultyBrithing = false
	}
	if r.FormValue("soreThroat") == "si" {
		boolean = true
		paciente.soreThroat = true
	} else {
		boolean = false
		paciente.soreThroat = false
	}
	if r.FormValue("noneSymtons") == "si" {
		boolean = true
		paciente.noneSymtons = true
	} else {
		boolean = false
		paciente.noneSymtons = false
	}
	if r.FormValue("age0_9") == "si" {
		boolean = true
		paciente.age0_9 = true
	} else {
		boolean = false
		paciente.age0_9 = false
	}
	if r.FormValue("age10_19") == "si" {
		boolean = true
		paciente.age10_19 = true
	} else {
		boolean = false
		paciente.age10_19 = false
	}
	if r.FormValue("age20_24") == "si" {
		boolean = true
		paciente.age20_24 = true
	} else {
		boolean = false
		paciente.age20_24 = false
	}
	if r.FormValue("age25_59") == "si" {
		boolean = true
		paciente.age25_59 = true
	} else {
		boolean = false
		paciente.age25_59 = false
	}
	if r.FormValue("age60") == "si" {
		boolean = true
		paciente.age60 = true
	} else {
		boolean = false
		paciente.age60 = false
	}
	if r.FormValue("genderFemale") == "si" {
		boolean = true
		paciente.genderFemale = true
	} else {
		boolean = false
		paciente.genderFemale = false
	}
	if r.FormValue("genderMale") == "si" {
		boolean = true
		paciente.genderMale = true
	} else {
		boolean = false
		paciente.genderMale = false
	}
	if r.FormValue("severityMild") == "si" {
		boolean = true
		paciente.severityMild = true
	} else {
		boolean = false
		paciente.severityMild = false
	}
	if r.FormValue("severityModerate") == "si" {
		boolean = true
		paciente.severityModerate = true
	} else {
		boolean = false
		paciente.severityModerate = false
	}
	if r.FormValue("severityNone") == "si" {
		boolean = true
		paciente.severityNone = true
	} else {
		boolean = false
		paciente.severityNone = false
	}
	if r.FormValue("severitySevere") == "si" {
		boolean = true
		paciente.severitySevere = true
	} else {
		boolean = false
		paciente.severitySevere = false
	}
	if r.FormValue("contactYes") == "si" {
		boolean = true
		paciente.contactYes = true
	} else {
		boolean = false
		paciente.contactYes = false
	}
	if r.FormValue("submit") == "submit" {
		boolean = true
	}
	getAbs(&paciente)

	if boolean == true {
		fmt.Printf("\nSintomas: %d, Others: %d, Estado: %s\n", paciente.sintoms, paciente.others, paciente.setClass)

		chInfo = make(chan map[string]int)
		go func() { chInfo <- map[string]int{} }()
		go server()
		time.Sleep(time.Millisecond * 100)
		var op int
		for {
			fmt.Print("Generated option: ")
			fmt.Scanf("Choosed %d", KNN(&paciente))
			msg := estadoPaciente{cnum, localAddr, op}
			for _, addr := range addrs {
				send(addr, msg)
			}
		}
	}

	options := []Option{
		Option{"si", "Id1", "si", boolean},
		Option{"no", "Id2", "no", boolean},
	}

	if err := placesPageTmpl.Execute(w, options); err != nil {
		fmt.Println("Failed to build page", err)
	}
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
	if p.dryCough == true {
		contSim += 2
	}
	if p.difficultyBrithing == true {
		contSim += 2
	}
	if p.soreThroat == true {
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
	if p.genderFemale == true {
		contOth++
	}
	if p.genderMale == true {
		contOth += 2
	}
	if p.severityMild == true {
		contOth++
	}
	if p.severityModerate == true {
		contOth += 2
	}
	if p.severitySevere == true {
		contOth += 3
	}
	if p.contactYes == true {
		contOth += 4
	}

	p.sintoms = contSim
	p.others = contOth

	if p.sintoms+p.others > 13 {
		p.setClass = "contagiado"
	}
	if p.sintoms+p.others <= 13 {
		p.setClass = "no contagiado"
	}
	contSim = 0
	contOth = 0
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
