package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mgutz/ansi"
)

func main() {
	endresult := []string{}

	Red := ansi.Color("vulnerable", "red")

	logo()

	var api bool
	flag.BoolVar(&api, "a", false, "For the google maps api")
	flag.Parse()

	if api {
		api_key := flag.Arg(0)

		//Falsifying the api key
		url := "https://maps.googleapis.com/maps/api/place/textsearch/json?query=restaurants+in+Sydney&key=" + api_key
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		body_string := string(body)

		if strings.Contains(body_string, "The provided API key is invalid.") {
			fmt.Println("Not a valid api key")

		} else {
			//custom search
			url1 := "https://www.googleapis.com/customsearch/v1?cx=017576662512468239146:omuauf_lfve&q=lectures&key=" + api_key

			resp1, err := http.Get(url1)
			if err != nil {
				fmt.Println("Unable to connect with custom search")
			}
			if resp1 == nil {
				fmt.Println("Checking for the next one")

			} else {
				defer resp1.Body.Close()
				body, err := io.ReadAll(resp1.Body)
				if err != nil {
					panic(err)
				}
				body_string := string(body)

				if strings.Contains(body_string, "errors") {
					fmt.Printf("%v >> Api key is not vulnerable to Custom Search Api\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Custom Search Api\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url1)
					A1 := "Custom Search   >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $5 per 1000 request"
					endresult = append(endresult, A1)
				}
			}

			//static map api

			url2 := "https://maps.googleapis.com/maps/api/staticmap?center=45%2C10&zoom=7&size=400x400&key=" + api_key

			resp2, err := http.Get(url2)
			if err != nil {
				fmt.Println("Unable to connect with Static Map Api")
			}
			if resp2 == nil {
				fmt.Println("Checking for the next one")

			} else {
				defer resp2.Body.Close()
				if resp2.StatusCode == 200 {
					fmt.Printf("%v Api Key is %v to Static Map Api \n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url2)
					A2 := "Staticmap                       >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $2 per 1000 request"
					endresult = append(endresult, A2)
				} else {
					fmt.Printf("%v is not vulnerable to Static Map Api\n\n", api_key)
				}
			}

			//streetview api

			url3 := "https://maps.googleapis.com/maps/api/streetview?size=400x400&location=40.720032,-73.988354&fov=90&heading=235&pitch=10&key=" + api_key

			resp3, err := http.Get(url3)
			if err != nil {
				fmt.Println("Unable to connect with streetview api")
			}
			if resp3 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp3.Body.Close()
				body1, err := io.ReadAll(resp3.Body)
				if err != nil {
					panic(err)
				}
				body_string1 := string(body1)

				if strings.Contains(body_string1, "The Google Maps Platform server rejected your request. This API project is not authorized to use this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to StreetView Api\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to StreetView Api\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url3)
					A3 := "StreetView                      >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $7 per 1000 request"
					endresult = append(endresult, A3)
				}

			}

			//Directions API

			url4 := "https://maps.googleapis.com/maps/api/directions/json?origin=Disneyland&destination=Universal+Studios+Hollywood4&key=" + api_key

			resp4, err := http.Get(url4)
			if err != nil {
				fmt.Println("Unable to connect with Directions API")
			}
			if resp4 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp4.Body.Close()
				body2, err := io.ReadAll(resp4.Body)
				if err != nil {
					panic(err)
				}
				body_string2 := string(body2)

				if strings.Contains(body_string2, "API keys with referer restrictions cannot be used with this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Directions Api\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Directions Api\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url4)
					A4 := "Directions API                  >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $5 per 1000 request"
					A5 := "Directions *Advanced* API       >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $10 per 1000 request"
					endresult = append(endresult, A4)
					endresult = append(endresult, A5)
				}
			}

			//Geocode API

			url5 := "https://maps.googleapis.com/maps/api/geocode/json?latlng=40,30&key=" + api_key

			resp5, err := http.Get(url5)
			if err != nil {
				fmt.Println("Unable to connect with Geocode API")
			}
			if resp5 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp5.Body.Close()
				body3, err := io.ReadAll(resp5.Body)
				if err != nil {
					panic(err)
				}
				body_string3 := string(body3)

				if strings.Contains(body_string3, "API keys with referer restrictions cannot be used with this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Geocode API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Geocode API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url5)
					A6 := "Geocode API                     >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $5 per 1000 request"

					endresult = append(endresult, A6)
				}

			}

			//Distance Matrix API

			url6 := "https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins=40.6655101,-73.89188969999998&destinations=40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.659569%2C-73.933783%7C40.729029%2C-73.851524%7C40.6860072%2C-73.6334271%7C40.598566%2C-73.7527626%7C40.659569%2C-73.933783%7C40.729029%2C-73.851524%7C40.6860072%2C-73.6334271%7C40.598566%2C-73.7527626&key=" + api_key

			resp6, err := http.Get(url6)
			if err != nil {
				fmt.Println("Unable to connect with Distance Matrix API")
			}
			if resp6 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp6.Body.Close()
				body4, err := io.ReadAll(resp6.Body)
				if err != nil {
					panic(err)
				}
				body_string4 := string(body4)

				if strings.Contains(body_string4, "API keys with referer restrictions cannot be used with this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Distance Matrix API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Distance Matrix API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url6)
					A8 := "Distance Matrix API             >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $5 per 1000 request"

					endresult = append(endresult, A8)
				}
			}

			//Find Place From Text API

			url7 := "https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=Museum%20of%20Contemporary%20Art%20Australia&inputtype=textquery&fields=photos,formatted_address,name,rating,opening_hours,geometry&key=" + api_key

			resp7, err := http.Get(url7)
			if err != nil {
				fmt.Println("Unable to connect with Find Place From Text API")
			}
			if resp7 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp7.Body.Close()
				body5, err := io.ReadAll(resp7.Body)
				if err != nil {
					panic(err)
				}
				body_string5 := string(body5)

				if strings.Contains(body_string5, "This API project is not authorized to use this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Find Place From Text API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Find Place From Text API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url7)
					A9 := "Find Place From Text API        >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $17 per 1000 request"

					endresult = append(endresult, A9)
				}
			}

			//Autocomplete API

			url8 := "https://maps.googleapis.com/maps/api/place/autocomplete/json?input=Bingh&types=%28cities%29&key=" + api_key

			resp8, err := http.Get(url8)
			if err != nil {
				fmt.Println("Unable to connect with Autocomplete API")
			}
			if resp8 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp8.Body.Close()
				body6, err := io.ReadAll(resp8.Body)
				if err != nil {
					panic(err)
				}
				body_string6 := string(body6)

				if strings.Contains(body_string6, "This API project is not authorized to use this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Autocomplete API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Autocomplete API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url8)
					A10 := "Autocomplete API                >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $2.83 per 1000 request"
					A11 := "Autocomplete API per session    >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $17 per 1000 request"

					endresult = append(endresult, A10)
					endresult = append(endresult, A11)
				}

			}

			//Elevation API

			url9 := "https://maps.googleapis.com/maps/api/elevation/json?locations=39.7391536,-104.9847034&key=" + api_key

			resp9, err := http.Get(url9)
			if err != nil {
				fmt.Println("Unable to connect with Elevation API")
			}
			if resp9 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp9.Body.Close()
				body7, err := io.ReadAll(resp9.Body)
				if err != nil {
					panic(err)
				}
				body_string7 := string(body7)

				if strings.Contains(body_string7, "API keys with referer restrictions cannot be used with this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Elevation API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Elevation API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url9)
					A12 := "Elevation API                   >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $5 per 1000 request"

					endresult = append(endresult, A12)

				}

			}

			//Timezone API

			url10 := "https://maps.googleapis.com/maps/api/timezone/json?location=39.6034810,-119.6822510&timestamp=1331161200&key=" + api_key

			resp10, err := http.Get(url10)
			if err != nil {
				fmt.Println("Unable to connect with Timezone API")
			}

			if resp10 == nil {
				fmt.Println("checking for the next one")

			} else {
				defer resp10.Body.Close()
				body8, err := io.ReadAll(resp10.Body)
				if err != nil {
					panic(err)
				}
				body_string8 := string(body8)

				if strings.Contains(body_string8, "This API project is not authorized to use this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Timezone API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Timezone API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url10)
					A13 := "Timezone API                    >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $5 per 1000 request"

					endresult = append(endresult, A13)

				}

			}

			//Nearest Roads API

			url11 := "https://roads.googleapis.com/v1/nearestRoads?points=60.170880,24.942795|60.170879,24.942796|60.170877,24.942796&key=" + api_key

			resp11, err := http.Get(url11)
			if err != nil {
				fmt.Println("Unable to connect with Nearest Roads API")
			}
			if resp11 == nil {
				fmt.Println("checking for the next one")
			} else {
				defer resp10.Body.Close()
				body9, err := io.ReadAll(resp11.Body)
				if err != nil {
					panic(err)
				}
				body_string9 := string(body9)

				if strings.Contains(body_string9, "API_KEY_HTTP_REFERRER_BLOCKED") {
					fmt.Printf("%v >> Api key is not vulnerable to Nearest Roads API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Nearest Roads API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url11)
					A14 := "Nearest Roads API               >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $10 per 1000 request"

					endresult = append(endresult, A14)

				}
			}

			//Geolocation API

			url12 := "https://www.googleapis.com/geolocation/v1/geolocate?key=" + api_key
			data := strings.NewReader(`

		{
			"considerIp": "true"
		}
	
	
	`)
			resp12, err := http.Post(url12, "application/x-www-form-urlencoded", data)

			if err != nil {
				fmt.Println("Unable to connect with Geolocation API")
			}
			if resp12 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp12.Body.Close()
				body10, err := io.ReadAll(resp12.Body)
				if err != nil {
					panic(err)
				}
				body_string10 := string(body10)

				if strings.Contains(body_string10, "API_KEY_HTTP_REFERRER_BLOCKED") {
					fmt.Printf("%v >> Api key is not vulnerable to Geolocation API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Geolocation API\n", api_key, Red)
					fmt.Println("POC : This will be a curl Command")
					fmt.Printf("curl -i -s -k  -X $'POST' -H $'Host: www.googleapis.com' -H $'Content-Length: 22' --data-binary $'{\"considerIp\": \"true\"}' $'https://www.googleapis.com/geolocation/v1/geolocate?key=%v'\n\n", api_key)

					A15 := "Geolocation API                 >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $5 per 1000 request"

					endresult = append(endresult, A15)

				}
			}

			//Route to Traveled API

			url13 := "https://roads.googleapis.com/v1/snapToRoads?path=-35.27801,149.12958|-35.28032,149.12907&interpolate=true&key=" + api_key

			resp13, err := http.Get(url13)
			if err != nil {
				fmt.Println("Unable to connect with Route to Traveled API")
			}
			if resp13 == nil {
				fmt.Println("checking for the next one")
			} else {
				defer resp13.Body.Close()
				body13, err := io.ReadAll(resp13.Body)
				if err != nil {
					panic(err)
				}
				body_string13 := string(body13)

				if strings.Contains(body_string13, "API_KEY_HTTP_REFERRER_BLOCKED") {
					fmt.Printf("%v >> Api key is not vulnerable to Route to Traveled API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Route to Traveled API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url13)
					A15 := "Route to Traveled API           >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $10 per 1000 request"

					endresult = append(endresult, A15)

				}

			}

			//Speed Limit-Roads API

			url14 := "https://roads.googleapis.com/v1/speedLimits?path=38.75807927603043,-9.03741754643809&key=" + api_key

			resp14, err := http.Get(url14)
			if err != nil {
				fmt.Println("Unable to connect with Speed Limit-Roads API")
			}
			if resp14 == nil {
				fmt.Println("Checking for the next one")

			} else {
				defer resp14.Body.Close()
				body14, err := io.ReadAll(resp14.Body)
				if err != nil {
					panic(err)
				}
				body_string14 := string(body14)

				if strings.Contains(body_string14, "API_KEY_HTTP_REFERRER_BLOCKED") {
					fmt.Printf("%v >> Api key is not vulnerable to Speed Limit-Roads API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Speed Limit-Roads API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url14)
					A16 := "Speed Limit-Roads API           >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $20 per 1000 request"

					endresult = append(endresult, A16)

				}
			}

			//Place Details API

			url15 := "https://maps.googleapis.com/maps/api/place/details/json?place_id=ChIJN1t_tDeuEmsRUsoyG83frY4&fields=name,rating,formatted_phone_number&key=" + api_key

			resp15, err := http.Get(url15)
			if err != nil {
				fmt.Println("Unable to connect with Place Details API")
			}
			if resp15 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp15.Body.Close()
				body15, err := io.ReadAll(resp15.Body)
				if err != nil {
					panic(err)
				}
				body_string15 := string(body15)

				if strings.Contains(body_string15, "This API project is not authorized to use this API") {
					fmt.Printf("%v >> Api key is not vulnerable to Place Details API\n\n", api_key)

				} else {

					fmt.Printf("%v is %v to Place Details API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url15)
					A17 := "Place Details API               >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $17 per 1000 request"

					endresult = append(endresult, A17)
				}
			}

			//Nearby Search-Places API

			url16 := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=-33.8670522,151.1957362&radius=100&types=food&name=harbour&key=" + api_key

			resp16, err := http.Get(url16)
			if err != nil {
				fmt.Println("Unable to connect with Nearby Search-Places API")
			}
			if resp16 == nil {
				defer resp16.Body.Close()
				body16, err := io.ReadAll(resp16.Body)
				if err != nil {
					panic(err)
				}
				body_string16 := string(body16)

				if strings.Contains(body_string16, "This API project is not authorized to use this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Nearby Search-Places API\n\n", api_key)

				} else {

					fmt.Printf("%v is %v to Nearby Search-Places API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url16)
					A18 := "Nearby Search-Places API        >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $32 per 1000 request"

					endresult = append(endresult, A18)
				}
			}

			//Text Search-Places API

			url17 := "https://maps.googleapis.com/maps/api/place/textsearch/json?query=restaurants+in+Sydney&key=" + api_key

			resp17, err := http.Get(url17)
			if err != nil {
				fmt.Println("Unable to connect with Text Search-Places API")
			}
			if resp17 == nil {
				defer resp17.Body.Close()
				body17, err := io.ReadAll(resp17.Body)
				if err != nil {
					panic(err)
				}
				body_string17 := string(body17)

				if strings.Contains(body_string17, "This API project is not authorized to use this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Text Search-Places API\n\n", api_key)

				} else {

					fmt.Printf("%v is %v to Text Search-Places API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url17)
					A19 := "Text Search-Places API          >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $32 per 1000 request"

					endresult = append(endresult, A19)
				}
			}

			//Places Photo API

			url19 := "https://maps.googleapis.com/maps/api/place/photo?maxwidth=400&photoreference=CnRtAAAATLZNl354RwP_9UKbQ_5Psy40texXePv4oAlgP4qNEkdIrkyse7rPXYGd9D_Uj1rVsQdWT4oRz4QrYAJNpFX7rzqqMlZw2h2E2y5IKMUZ7ouD_SlcHxYq1yL4KbKUv3qtWgTK0A6QbGh87GB3sscrHRIQiG2RrmU_jF4tENr9wGS_YxoUSSDrYjWmrNfeEHSGSc3FyhNLlBU&key=" + api_key

			client := http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				}}

			resp19, err := client.Get(url19)
			if err != nil {
				fmt.Println("Unable to connect with Places Photo API")
			}
			if resp19 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp19.Body.Close()
				if resp19.StatusCode == 302 {
					fmt.Printf("%v Api Key is %v to Places Photo API \n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url2)
					A20 := "Places Photo API                >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $7 per 1000 request"
					endresult = append(endresult, A20)
				} else {
					fmt.Printf("%v is not vulnerable to Places Photo API\n\n", api_key)
				}
			}

			for _, v := range endresult {
				fmt.Println(v)

			}
		}

	} else {
		var api_key string
		fmt.Print("Enter your api key:  ")
		fmt.Scanln(&api_key)

		//Falsifying the api key
		url := "https://maps.googleapis.com/maps/api/place/textsearch/json?query=restaurants+in+Sydney&key=" + api_key
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		body_string := string(body)

		if strings.Contains(body_string, "The provided API key is invalid.") {
			fmt.Println("Not a valid api key")

		} else {
			//custom search
			url1 := "https://www.googleapis.com/customsearch/v1?cx=017576662512468239146:omuauf_lfve&q=lectures&key=" + api_key

			resp1, err := http.Get(url1)
			if err != nil {
				fmt.Println("Unable to connect with custom search")
			}
			if resp1 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp1.Body.Close()
				body, err := io.ReadAll(resp1.Body)
				if err != nil {
					panic(err)
				}
				body_string := string(body)

				if strings.Contains(body_string, "errors") {
					fmt.Printf("%v >> Api key is not vulnerable to Custom Search Api\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Custom Search Api\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url1)
					A1 := "Custom Search   >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $5 per 1000 request"
					endresult = append(endresult, A1)
				}
			}

			//static map api

			url2 := "https://maps.googleapis.com/maps/api/staticmap?center=45%2C10&zoom=7&size=400x400&key=" + api_key

			resp2, err := http.Get(url2)
			if err != nil {
				fmt.Println("Unable to connect with static map api")
			}
			if resp2 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp2.Body.Close()
				if resp2.StatusCode == 200 {
					fmt.Printf("%v Api Key is %v to Static Map Api \n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url2)
					A2 := "Staticmap                       >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $2 per 1000 request"
					endresult = append(endresult, A2)
				} else {
					fmt.Printf("%v is not vulnerable to Static Map Api\n\n", api_key)
				}
			}

			//streetview api

			url3 := "https://maps.googleapis.com/maps/api/streetview?size=400x400&location=40.720032,-73.988354&fov=90&heading=235&pitch=10&key=" + api_key

			resp3, err := http.Get(url3)
			if err != nil {
				fmt.Println("Unable to connect with streetview api")
			}
			if resp3 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp3.Body.Close()
				body1, err := io.ReadAll(resp3.Body)
				if err != nil {
					panic(err)
				}
				body_string1 := string(body1)

				if strings.Contains(body_string1, "The Google Maps Platform server rejected your request. This API project is not authorized to use this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to StreetView Api\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to StreetView Api\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url3)
					A3 := "StreetView                      >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $7 per 1000 request"
					endresult = append(endresult, A3)
				}
			}

			//Directions API

			url4 := "https://maps.googleapis.com/maps/api/directions/json?origin=Disneyland&destination=Universal+Studios+Hollywood4&key=" + api_key

			resp4, err := http.Get(url4)
			if err != nil {
				fmt.Println("Unable to connect with Directions API")
			}
			if resp4 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp4.Body.Close()
				body2, err := io.ReadAll(resp4.Body)
				if err != nil {
					panic(err)
				}
				body_string2 := string(body2)

				if strings.Contains(body_string2, "API keys with referer restrictions cannot be used with this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Directions Api\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Directions Api\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url4)
					A4 := "Directions API                  >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $5 per 1000 request"
					A5 := "Directions *Advanced* API       >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $10 per 1000 request"
					endresult = append(endresult, A4)
					endresult = append(endresult, A5)
				}
			}

			//Geocode API

			url5 := "https://maps.googleapis.com/maps/api/geocode/json?latlng=40,30&key=" + api_key

			resp5, err := http.Get(url5)
			if err != nil {
				fmt.Println("Unable to connect with Geocode API")
			}
			if resp5 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp5.Body.Close()
				body3, err := io.ReadAll(resp5.Body)
				if err != nil {
					panic(err)
				}
				body_string3 := string(body3)

				if strings.Contains(body_string3, "API keys with referer restrictions cannot be used with this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Geocode API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Geocode API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url5)
					A6 := "Geocode API                     >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $5 per 1000 request"

					endresult = append(endresult, A6)
				}
			}

			//Distance Matrix API

			url6 := "https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins=40.6655101,-73.89188969999998&destinations=40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.6905615%2C-73.9976592%7C40.659569%2C-73.933783%7C40.729029%2C-73.851524%7C40.6860072%2C-73.6334271%7C40.598566%2C-73.7527626%7C40.659569%2C-73.933783%7C40.729029%2C-73.851524%7C40.6860072%2C-73.6334271%7C40.598566%2C-73.7527626&key=" + api_key

			resp6, err := http.Get(url6)
			if err != nil {
				fmt.Println("Unable to connect with Distance Matrix API")
			}
			if resp6 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp6.Body.Close()
				body4, err := io.ReadAll(resp6.Body)
				if err != nil {
					panic(err)
				}
				body_string4 := string(body4)

				if strings.Contains(body_string4, "API keys with referer restrictions cannot be used with this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Distance Matrix API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Distance Matrix API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url6)
					A8 := "Distance Matrix API             >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $5 per 1000 request"

					endresult = append(endresult, A8)
				}
			}

			//Find Place From Text API

			url7 := "https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=Museum%20of%20Contemporary%20Art%20Australia&inputtype=textquery&fields=photos,formatted_address,name,rating,opening_hours,geometry&key=" + api_key

			resp7, err := http.Get(url7)
			if err != nil {
				fmt.Println("Unable to connect with Find Place From Text API")
			}
			if resp7 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp7.Body.Close()
				body5, err := io.ReadAll(resp7.Body)
				if err != nil {
					panic(err)
				}
				body_string5 := string(body5)

				if strings.Contains(body_string5, "This API project is not authorized to use this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Find Place From Text API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Find Place From Text API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url7)
					A9 := "Find Place From Text API        >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $17 per 1000 request"

					endresult = append(endresult, A9)
				}
			}

			//Autocomplete API

			url8 := "https://maps.googleapis.com/maps/api/place/autocomplete/json?input=Bingh&types=%28cities%29&key=" + api_key

			resp8, err := http.Get(url8)
			if err != nil {
				fmt.Println("Unable to connect with Autocomplete API")
			}
			if resp8 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp8.Body.Close()
				body6, err := io.ReadAll(resp8.Body)
				if err != nil {
					panic(err)
				}
				body_string6 := string(body6)

				if strings.Contains(body_string6, "This API project is not authorized to use this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Autocomplete API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Autocomplete API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url8)
					A10 := "Autocomplete API                >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $2.83 per 1000 request"
					A11 := "Autocomplete API per session    >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $17 per 1000 request"

					endresult = append(endresult, A10)
					endresult = append(endresult, A11)
				}
			}

			//Elevation API

			url9 := "https://maps.googleapis.com/maps/api/elevation/json?locations=39.7391536,-104.9847034&key=" + api_key

			resp9, err := http.Get(url9)
			if err != nil {
				fmt.Println("Unable to connect with Elevation API")
			}
			if resp9 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp9.Body.Close()
				body7, err := io.ReadAll(resp9.Body)
				if err != nil {
					panic(err)
				}
				body_string7 := string(body7)

				if strings.Contains(body_string7, "API keys with referer restrictions cannot be used with this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Elevation API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Elevation API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url9)
					A12 := "Elevation API                   >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $5 per 1000 request"

					endresult = append(endresult, A12)

				}
			}

			//Timezone API

			url10 := "https://maps.googleapis.com/maps/api/timezone/json?location=39.6034810,-119.6822510&timestamp=1331161200&key=" + api_key

			resp10, err := http.Get(url10)
			if err != nil {
				fmt.Println("Unable to connect with Timezone API")
			}
			if resp10 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp10.Body.Close()
				body8, err := io.ReadAll(resp10.Body)
				if err != nil {
					panic(err)
				}
				body_string8 := string(body8)

				if strings.Contains(body_string8, "This API project is not authorized to use this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Timezone API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Timezone API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url10)
					A13 := "Timezone API                    >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $5 per 1000 request"

					endresult = append(endresult, A13)

				}
			}

			//Nearest Roads API

			url11 := "https://roads.googleapis.com/v1/nearestRoads?points=60.170880,24.942795|60.170879,24.942796|60.170877,24.942796&key=" + api_key

			resp11, err := http.Get(url11)
			if err != nil {
				fmt.Println("Unable to connect with Nearest Roads API")
			}
			if resp11 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp10.Body.Close()
				body9, err := io.ReadAll(resp11.Body)
				if err != nil {
					panic(err)
				}
				body_string9 := string(body9)

				if strings.Contains(body_string9, "API_KEY_HTTP_REFERRER_BLOCKED") {
					fmt.Printf("%v >> Api key is not vulnerable to Nearest Roads API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Nearest Roads API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url11)
					A14 := "Nearest Roads API               >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $10 per 1000 request"

					endresult = append(endresult, A14)

				}
			}

			//Geolocation API

			url12 := "https://www.googleapis.com/geolocation/v1/geolocate?key=" + api_key
			data := strings.NewReader(`

		{
			"considerIp": "true"
		}
	
	
	`)
			resp12, err := http.Post(url12, "application/x-www-form-urlencoded", data)

			if err != nil {
				fmt.Println("Unable to connect with Geolocation API")
			}
			if resp12 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp12.Body.Close()
				body10, err := io.ReadAll(resp12.Body)
				if err != nil {
					panic(err)
				}
				body_string10 := string(body10)

				if strings.Contains(body_string10, "API_KEY_HTTP_REFERRER_BLOCKED") {
					fmt.Printf("%v >> Api key is not vulnerable to Geolocation API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Geolocation API\n", api_key, Red)
					fmt.Println("POC : This will be a curl Command")
					fmt.Printf("curl -i -s -k  -X $'POST' -H $'Host: www.googleapis.com' -H $'Content-Length: 22' --data-binary $'{\"considerIp\": \"true\"}' $'https://www.googleapis.com/geolocation/v1/geolocate?key=%v'\n\n", api_key)

					A15 := "Geolocation API                 >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $5 per 1000 request"

					endresult = append(endresult, A15)

				}
			}

			//Route to Traveled API

			url13 := "https://roads.googleapis.com/v1/snapToRoads?path=-35.27801,149.12958|-35.28032,149.12907&interpolate=true&key=" + api_key

			resp13, err := http.Get(url13)
			if err != nil {
				fmt.Println("Unable to connect with Route to Traveled API")
			}
			if resp13 == nil {
				fmt.Println("Checking for the next one")
			} else {

				defer resp13.Body.Close()
				body13, err := io.ReadAll(resp13.Body)
				if err != nil {
					panic(err)
				}
				body_string13 := string(body13)

				if strings.Contains(body_string13, "API_KEY_HTTP_REFERRER_BLOCKED") {
					fmt.Printf("%v >> Api key is not vulnerable to Route to Traveled API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Route to Traveled API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url13)
					A15 := "Route to Traveled API           >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $10 per 1000 request"

					endresult = append(endresult, A15)

				}
			}

			//Speed Limit-Roads API

			url14 := "https://roads.googleapis.com/v1/speedLimits?path=38.75807927603043,-9.03741754643809&key=" + api_key

			resp14, err := http.Get(url14)
			if err != nil {
				fmt.Println("Unable to connect with Speed Limit-Roads API")
			}
			if resp14 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp14.Body.Close()
				body14, err := io.ReadAll(resp14.Body)
				if err != nil {
					panic(err)
				}
				body_string14 := string(body14)

				if strings.Contains(body_string14, "API_KEY_HTTP_REFERRER_BLOCKED") {
					fmt.Printf("%v >> Api key is not vulnerable to Speed Limit-Roads API\n\n", api_key)
				} else {
					fmt.Printf("%v is %v to Speed Limit-Roads API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url14)
					A16 := "Speed Limit-Roads API           >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $20 per 1000 request"

					endresult = append(endresult, A16)

				}
			}

			//Place Details API

			url15 := "https://maps.googleapis.com/maps/api/place/details/json?place_id=ChIJN1t_tDeuEmsRUsoyG83frY4&fields=name,rating,formatted_phone_number&key=" + api_key

			resp15, err := http.Get(url15)
			if err != nil {
				fmt.Println("Unable to connect with Place Details API")
			}
			if resp15 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp15.Body.Close()
				body15, err := io.ReadAll(resp15.Body)
				if err != nil {
					panic(err)
				}
				body_string15 := string(body15)

				if strings.Contains(body_string15, "This API project is not authorized to use this API") {
					fmt.Printf("%v >> Api key is not vulnerable to Place Details API\n\n", api_key)

				} else {

					fmt.Printf("%v is %v to Place Details API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url15)
					A17 := "Place Details API               >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $17 per 1000 request"

					endresult = append(endresult, A17)
				}
			}

			//Nearby Search-Places API

			url16 := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=-33.8670522,151.1957362&radius=100&types=food&name=harbour&key=" + api_key

			resp16, err := http.Get(url16)
			if err != nil {
				fmt.Println("Unable to connect with Nearby Search-Places API")
			}
			if resp16 == nil {
				fmt.Println("checking for the next one")
			} else {
				defer resp16.Body.Close()
				body16, err := io.ReadAll(resp16.Body)
				if err != nil {
					panic(err)
				}
				body_string16 := string(body16)

				if strings.Contains(body_string16, "This API project is not authorized to use this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Nearby Search-Places API\n\n", api_key)

				} else {

					fmt.Printf("%v is %v to Nearby Search-Places API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url16)
					A18 := "Nearby Search-Places API        >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $32 per 1000 request"

					endresult = append(endresult, A18)
				}
			}

			//Text Search-Places API

			url17 := "https://maps.googleapis.com/maps/api/place/textsearch/json?query=restaurants+in+Sydney&key=" + api_key

			resp17, err := http.Get(url17)
			if err != nil {
				fmt.Println("Unable to connect with Text Search-Places API")
			}
			if resp17 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp17.Body.Close()
				body17, err := io.ReadAll(resp17.Body)
				if err != nil {
					panic(err)
				}
				body_string17 := string(body17)

				if strings.Contains(body_string17, "This API project is not authorized to use this API.") {
					fmt.Printf("%v >> Api key is not vulnerable to Text Search-Places API\n\n", api_key)

				} else {

					fmt.Printf("%v is %v to Text Search-Places API\n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url17)
					A19 := "Text Search-Places API          >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $32 per 1000 request"

					endresult = append(endresult, A19)
				}

			}

			//Places Photo API

			url19 := "https://maps.googleapis.com/maps/api/place/photo?maxwidth=400&photoreference=CnRtAAAATLZNl354RwP_9UKbQ_5Psy40texXePv4oAlgP4qNEkdIrkyse7rPXYGd9D_Uj1rVsQdWT4oRz4QrYAJNpFX7rzqqMlZw2h2E2y5IKMUZ7ouD_SlcHxYq1yL4KbKUv3qtWgTK0A6QbGh87GB3sscrHRIQiG2RrmU_jF4tENr9wGS_YxoUSSDrYjWmrNfeEHSGSc3FyhNLlBU&key=" + api_key

			client := http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				}}

			resp19, err := client.Get(url19)
			if err != nil {
				fmt.Println("Unable to connect with Places Photo API")
			}
			if resp19 == nil {
				fmt.Println("Checking for the next one")
			} else {
				defer resp19.Body.Close()
				if resp19.StatusCode == 302 {
					fmt.Printf("%v Api Key is %v to Places Photo API \n", api_key, Red)
					fmt.Printf("POC :  %v\n\n", url2)
					A20 := "Places Photo API                >>>>>>>>>>   " + api_key + "   >>>>>>>>>>   $7 per 1000 request"
					endresult = append(endresult, A20)
				} else {
					fmt.Printf("%v is not vulnerable to Places Photo API\n\n", api_key)
				}
			}

			for _, v := range endresult {
				fmt.Println(v)

			}
		}

	}
}

func logo() {
	lg1 := ansi.Color("****         ****", "green")
	lg2 := ansi.Color("*****       *****", "green")
	lg3 := ansi.Color("*******    ******", "green")
	lg4 := ansi.Color("*** **** **** ***", "green")
	lg5 := ansi.Color("***  *******  ***", "green")
	lg6 := ansi.Color("***           ***", "green")
	lg7 := ansi.Color("***           ***", "green")
	lg8 := ansi.Color("***           *** >> Replicate by Ractiurd [Mahedi]", "green")
	lg9 := ansi.Color("\nGoogle Maps Api Scanner \n\n", "green")

	fmt.Println(lg1)
	fmt.Println(lg2)
	fmt.Println(lg3)
	fmt.Println(lg4)
	fmt.Println(lg5)
	fmt.Println(lg6)
	fmt.Println(lg7)
	fmt.Println(lg8)
	fmt.Println(lg9)

}
