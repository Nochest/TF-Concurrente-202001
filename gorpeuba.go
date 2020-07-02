package main

import (
	"fmt"
	"html/template"
	"net/http"
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
				<option> Seleccione </option>
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

var addrs = []string{
	"192.168.0.27:8000",
	"192.168.0.28:8000"}

func main() {
	/*var n int
	fmt.Print("Numero de nodos a usar: ")
	fmt.Scanln(&n)
	var ingreso string
	for i := 0; i < n; i++ {
		fmt.Printf("Ingrese la ip del nodo %d :", i+1)
		fmt.Scanln(&ingreso)
	}*/

	fmt.Print(addrs)

	http.HandleFunc("/", name)
	http.ListenAndServe(":8080", nil)
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
		fmt.Printf("Sintomas: %d, Others: %d, Estado: %s\n", paciente.sintoms, paciente.others, paciente.setClass)
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

	if p.sintoms+p.others > 10 {
		p.setClass = "contagiado"
	}
	if p.sintoms+p.others <= 10 {
		p.setClass = "no contagiado"
	}
	contSim = 0
	contOth = 0
}
