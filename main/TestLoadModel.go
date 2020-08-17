package main

import (
	"encoding/json"
	"fmt"
	"go-catboost/deserializer/jsondeserializer"
	"go-catboost/hash"
	"go-catboost/model"
	"io/ioutil"
	"os"
	"strings"
)

type TestLoadModel struct {

}

func amain() {
	a := int32(-2147483648)
	fmt.Println( hash.CityHash{}.CalcHash(0,uint64(a)))
}

func main() {
	var model model.Model = jsondeserializer.JsonModelDeserializer{}.Deserialize(TestLoadModel{}.read("/home/paras.mal/Downloads/ttm.json"));
	data := TestLoadModel{}.loadInput("/home/paras.mal/Downloads/td.tsv")
	for i,v := range data {
		fmt.Println(i, model.Predict(v))
	}
}


func (tlm TestLoadModel) loadInput(file string) []map[string]string {
	tsv,_ := ioutil.ReadFile(file)
	lines := strings.Split(string(tsv), "\n")
	data := make([]map[string]string,0)
	line := lines[0]
	tokens := strings.Split(line, "\t")
	for i:= 1;i<len(lines);i++ {
		if lines[i] == "" {
			continue
		}
		ma := make(map[string]string)
		t := strings.Split(lines[i],"\t")
		for j:=0;j<len(t);j++ {
			ma[tokens[j]] = t[j]
		}
		data = append(data, ma)
	}
	return data
}
func (tlm TestLoadModel) read(file string) map[string]interface{}{
	jsonFile, _ := os.Open(file)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data map[string]interface{} = make(map[string]interface{})
	json.Unmarshal(byteValue, &data)
	return data;
}