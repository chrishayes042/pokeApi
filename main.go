package main

import (
	pokemonStruct "api/struct"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


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
	var responseObject pokemonStruct.Response
	json.Unmarshal(responseData, &responseObject)

	// for each object in the var above using the Pokemon array
	for _, pokemon := range responseObject.Pokemon {

		// print the entry number and name
		fmt.Println( pokemon.EntryNo, pokemon.Species.Name )

		// if the entry number is 1
		// going to make this when a user clicks on a pokemon
		if( pokemon.EntryNo == 1 ){

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
			
			// create a var with the PokemonResponse struct
			var pokeResponseObj pokemonStruct.PokemonResponse
			var pokeResponseData pokemonStruct.PokemonData
			json.Unmarshal(getPokemonResponseData, &pokeResponseData)
			json.Unmarshal(getPokemonResponseData, &pokeResponseObj)
			
			// print first text data
			// not sure how to print it without the {} around the text
			fmt.Println( pokeResponseObj.PokemonTextData[0] )
			fmt.Println( "The capture rate is:" , pokeResponseData.CapRate )
			fmt.Println( "The happiness rate is:" , pokeResponseData.Happiness )			
		}
		// break the for each loop
		break
		
	}
	
}











