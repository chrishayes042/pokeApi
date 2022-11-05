package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)
type PokemonResponse struct {
	Text string `json:"flavor_text"`
	PokemonText []Pokemon `json:"flavor_text_entries"`
}
type Response struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNo int `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
	Text string `json:"flavor_text"`
	
}

type PokemonSpecies struct {
	Name string `json:"name"`
	Url string `json:"url"`

}

func main(){
	// call the pokeapi
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	// if error
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// read all the data from the response from the api 
	// put it in the responseData var
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// print all the data
	// fmt.Println(string(responseData))

	// create variable for the Response struct
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	// fmt.Println(responseObject.Name)/* 
	// fmt.Println(len(responseObject.Pokemon)) 

	// for each object in the var above using the Pokemon array
	for _, pokemon := range responseObject.Pokemon {
		// print the entry number and name
		fmt.Println( pokemon.EntryNo, pokemon.Species.Name )
		// if the entry number is 1
		if( pokemon.EntryNo == 1){
			// print
			fmt.Println("Getting data for:", pokemon.Species.Name)
			// call the api url inside the pokemon object
			getPokemon, err := http.Get(pokemon.Species.Url)

			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}
			// get the response
			getPokemonResponseData, err := ioutil.ReadAll(getPokemon.Body)
			if err != nil {
				log.Fatal(err)
			}
			// fmt.Println(string(getPokemonResponseData))
			// create a var with the PokemonResponse struct
			var pokeResponseObj PokemonResponse
			json.Unmarshal(getPokemonResponseData, &pokeResponseObj)
				// fmt.Println(pokeResponseObj.Text)
			
				// print the first text response
				fmt.Println( pokeResponseObj.PokemonText[0] )
				
			
		}
		break
		
	}
	
}











