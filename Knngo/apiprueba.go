package main

import (
	"encoding/csv"
	"fmt"
	"io"
	//"log"
	"os"
	"strconv"
	"math/rand"
	"time"
	"math"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func loadDataset(test [][]float64, trainingSet *[][]float64, testSet *[][]float64, split int) {
	// Open the file
	csvfile, err := os.Open("wine.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))
	//fmt.Println(len(r))
	
	rand.Seed(time.Now().UnixNano())
    min := 0
	max := 300



	i := 0
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}



		//test[i] = record
		//fmt.Println(record)
		//aa = append(aa,record)
		for j := 0; j < 12; j++ {
			//fmt.Println(v)
			
			if j < 12 {
				f, _ := strconv.ParseFloat(record[j], 64)
				test[i][j] = f
			}

		}



		if (rand.Intn(max - min + 1) + min) < split {
			*trainingSet = append(*trainingSet,test[i])
			//fmt.Println(trainingSet)
		} else {
			//testSet[i] = test[i]
			*testSet = append(*testSet,test[i])
		}

		
		i++

	}	
}


func Distancia(d1 []float64, d2 []float64, longitud int) float64 {
	var distancia float64
	distancia = 0
	for i:=0; i<longitud;i++ {
		distancia = distancia + math.Pow((d1[i]-d2[i]), 2)
		//fmt.Println(distancia)
	} 
	return math.Sqrt(distancia)

}

func Vecinos(trainingSet [][]float64, testSetIns []float64, k int) [][]float64 {
	var distancias []float64
	var total = len(testSetIns)-1
	var sortedarray [][]float64
	sortedarray = trainingSet

	for i:=0;i<len(trainingSet);i++ {
		dist := Distancia(testSetIns,trainingSet[i], total)
		//fmt.Println(dist)
		//fmt.Println(trainingSet[i])
		//trainingSet[i] = append(trainingSet[i],dist)
		//fmt.Println(trainingSet[i])
		distancias = append(distancias, dist)


	}

	//fmt.Println(distancias)

	for i:=0; i<len(distancias); i++ {
        minimum := i
        
        for j := i+1; j < len(distancias); j++ {
            
            if distancias[j] < distancias[minimum] {
				minimum = j
			}
		}

		
		t := distancias[i]
		distancias[i] = distancias[minimum]
		distancias[minimum] = t

		t2 := sortedarray[i]
		sortedarray[i] = sortedarray[minimum]
		sortedarray[minimum] = t2


	}
            
	
	
	//fmt.Println(distancias)
	//fmt.Println(sortedarray)


	var vecinos [][]float64

	for i:=0; i<k;i++ {
		vecinos = append(vecinos, sortedarray[i])
	}
	//fmt.Println(vecinos)
	return vecinos


}

func ClaseResultado(vecinos [][]float64) int {
	votos := [][]int {
		{0,0},
		{0,1},
		{0,2},
		{0,3},
		{0,4},
		{0,5},
		{0,6},
		{0,7},
		{0,8},
		{0,9},
		{0,10},
	}
	for x :=0; x< len(vecinos); x++ {
		resultado := vecinos[x][11]
		fmt.Println(resultado)
		if resultado == 0 {
			votos[0][0]++
		}
		if resultado == 1 {
			votos[1][0]++
		}
		if resultado == 2 {
			votos[2][0]++
		}
		if resultado == 3 {
			votos[3][0]++
		}
		if resultado == 4 {
			votos[4][0]++
		}
		if resultado == 5 {
			votos[5][0]++
		}
		if resultado == 6 {
			votos[6][0]++
		}
		if resultado == 7 {
			votos[7][0]++
		}
		if resultado == 8 {
			votos[8][0]++
		}
		if resultado == 9 {
			votos[9][0]++
		}
		if resultado == 10 {
			votos[10][0]++
		}
	}



	for i:=0; i<len(votos); i++ {
        minimum := i
        
        for j := i+1; j < len(votos); j++ {
            
            if votos[j][0] < votos[minimum][0] {
				minimum = j
			}
		}

		
		t := votos[i]
		votos[i] = votos[minimum]
		votos[minimum] = t


	}

	fmt.Println(votos)
	return votos[10][1]


}

func Probabilidad(testSet [][]float64, predicciones []int)  float64 {
	cont := 0
	var temp []float64
	for i:=0; i<len(testSet); i++ {
		new := float64(predicciones[i])
		temp = append(temp, new)
	} 
	for x := 0; x<len(testSet);x++ {
		if testSet[x][11] == temp[x]{
			cont++
		}
	}
	contflo := float64(cont)
	cantflo := float64(len(testSet))
	fmt.Println(contflo)
	fmt.Println(cantflo)
	pro := (contflo/cantflo) * 100

	return pro
}





type Proba struct {
	Porcent string `json:"porcent,omitempty"`
}

type Submit struct{
	V1 string `json:"v1,omitempty"`
	V2 string `json:"v2,omitempty"`
	V3 string `json:"v3,omitempty"`
	V4 string `json:"v4,omitempty"`
	V5 string `json:"v5,omitempty"`
	V6 string `json:"v6,omitempty"`
	V7 string `json:"v7,omitempty"`
	V8 string `json:"v8,omitempty"`
	V9 string `json:"v9,omitempty"`
	V10 string `json:"v10,omitempty"`
	V11 string `json:"v11,omitempty"`
	Kvariable string `json:"kvariable,omitempty"`
}

type Resultado struct {
	DevoResultado string `json:"devoresultado,omitempty"`
}


var probabi []Proba

func GetProbabilidad(w http.ResponseWriter, req *http.Request) {
	//

	test := make([][]float64, 300)

	for k := range test {
    	test[k] = make([]float64, 12)
	}

	trainingSet := [][]float64{}
	testSet := [][]float64{}




	split:=230



	loadDataset(test, &trainingSet, &testSet,split)
	fmt.Println(len(testSet))
	fmt.Println(len(trainingSet))





	var predicciones []int 
	k := 8
	for i:=0; i<len(testSet); i++ {
		veci := Vecinos(trainingSet, testSet[i], k)
		resultado := ClaseResultado(veci)
		predicciones = append(predicciones,resultado)
		fmt.Println("Resultado: ", resultado, "Actual: ", testSet[i][11])
	}

	proba := Probabilidad(testSet,predicciones)
	
	fmt.Println("probabilidad")

	fmt.Println(proba)

	
	probaString := fmt.Sprintf("%f", proba)

	probabi = append(probabi, Proba{Porcent: probaString})

	json.NewEncoder(w).Encode(probabi)

}


func ProbabilidadDeDatos(w http.ResponseWriter, req *http.Request) {

	//params := mux.Vars(req)
	variables := Submit{V1:"66,3",
	V2:"66,3",
	V3:"66,3",
	V4:"66,3",
	V5:"66,3",
	V6:"66,3",
	V7:"66,3",
	V8:"66,3",
	V9:"66,3",
	V10:"66,3",
	V11:"66,3",
	Kvariable:"66,3",}
	_ = json.NewDecoder(req.Body).Decode(&variables)
	//json.NewEncoder(w).Encode(variables)
	fmt.Println(variables)

	Test1var, _ := strconv.ParseFloat(variables.V1, 64)
	fmt.Println(Test1var)
	Test2var, _ := strconv.ParseFloat(variables.V2, 64)
	Test3var, _ := strconv.ParseFloat(variables.V3, 64)
	Test4var, _ := strconv.ParseFloat(variables.V4, 64)
	Test5var, _ := strconv.ParseFloat(variables.V5, 64)
	Test6var, _ := strconv.ParseFloat(variables.V6, 64)
	Test7var, _ := strconv.ParseFloat(variables.V7, 64)
	Test8var, _ := strconv.ParseFloat(variables.V8, 64)
	Test9var, _ := strconv.ParseFloat(variables.V9, 64)
	Test10var, _ := strconv.ParseFloat(variables.V10, 64)
	Test11var, _ := strconv.ParseFloat(variables.V11, 64)
	KTestvariable, _ := strconv.ParseFloat(variables.Kvariable, 64)

	fmt.Println(KTestvariable)

	test := make([][]float64, 300)

	for k := range test {
    	test[k] = make([]float64, 12)
	}

	trainingSet := [][]float64{}
	testSet := [][]float64{}



	split:=230



	loadDataset(test, &trainingSet, &testSet,split)

	testSet = [][]float64{{Test1var, Test2var, Test3var, Test4var,Test5var,Test6var,Test7var,Test8var,Test9var,Test10var,Test11var,8.0}}
	fmt.Println(len(testSet))
	fmt.Println(testSet)
	

	var predicciones []int 
	var resultado int
	k := int(KTestvariable)
	for i:=0; i<len(testSet); i++ {
		veci := Vecinos(trainingSet, testSet[i], k)
		fmt.Println(veci)
		resultado = ClaseResultado(veci)
		predicciones = append(predicciones,resultado)
		fmt.Println("Resultado: ", resultado, "Actual: ", testSet[i][11])
	}

	resultString := strconv.Itoa(resultado)
	
	var devolverres = Resultado{DevoResultado: resultString}

	json.NewEncoder(w).Encode(devolverres)


}




func main() {


	router := mux.NewRouter()

	

	router.HandleFunc("/probabilidad", GetProbabilidad).Methods("GET")
	router.HandleFunc("/probabilidad", ProbabilidadDeDatos).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))

}
