package jsondeserializer

import "fmt"

type CatHashDeserializer struct {

}

func (chd *CatHashDeserializer) Deserialize(jsonObject map[string]interface{}) map[string]uint32{
	hashList,ok := jsonObject["cat_features_hash"].([]interface{});
	if(!ok){
		return nil;
	}
	hashes := make(map[string]uint32)
	for _,he := range hashList {
		hashElement := he.(map[string]interface{})
		a:= uint32(hashElement["hash"].(float64))
		hashes[hashElement["value"].(string)] = a
		if hashElement["value"] == "fri" {
			fmt.Print(hashElement["hash"])
		}
	}
	return hashes;
}
