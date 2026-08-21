---
updatedAt: 2025-03-26T05:12:17.000Z
---

Fetch the complete documentation index at: https://docs.locationiq.com/llms.txt. Use this file to discover all available pages before exploring further. Append .md to any documentation page URL to get its markdown version.

# Free Form Query

The Search API allows converting addresses, such as a street address, into geographic coordinates (latitude and longitude). These coordinates can serve various use-cases, from placing markers on a map to helping algorithms determine nearby bus stops. This process is also known as Forward Geocoding.

# OpenAPI definition

```json
{
  "openapi": "3.0.0",
  "info": {
    "version": "1.5.1",
    "title": "LocationIQ - API reference",
    "description": "LocationIQ provides flexible enterprise-grade location based solutions. We work with developers, startups and enterprises worldwide serving billions of requests everyday. This page provides an overview of the technical aspects of our API and will help you get started.",
    "contact": {
      "email": "hello@locationiq.com"
    }
  },
  "paths": {
    "/search": {
      "get": {
        "operationId": "search",
        "summary": "Free Form Query",
        "description": "The Search API allows converting addresses, such as a street address, into geographic coordinates (latitude and longitude). These coordinates can serve various use-cases, from placing markers on a map to helping algorithms determine nearby bus stops. This process is also known as Forward Geocoding.",
        "tags": [
          "Search / Forward Geocoding"
        ],
        "parameters": [
          {
            "$ref": "#/components/parameters/query"
          },
          {
            "$ref": "#/components/parameters/format"
          },
          {
            "$ref": "#/components/parameters/addressdetails-default-0"
          },
          {
            "$ref": "#/components/parameters/statecode"
          },
          {
            "$ref": "#/components/parameters/viewbox"
          },
          {
            "$ref": "#/components/parameters/bounded"
          },
          {
            "$ref": "#/components/parameters/limit-max-50"
          },
          {
            "$ref": "#/components/parameters/accept-language"
          },
          {
            "$ref": "#/components/parameters/countrycodes"
          },
          {
            "$ref": "#/components/parameters/normalizeaddress"
          },
          {
            "$ref": "#/components/parameters/normalizecity"
          },
          {
            "$ref": "#/components/parameters/postaladdress"
          },
          {
            "$ref": "#/components/parameters/matchquality"
          },
          {
            "$ref": "#/components/parameters/source"
          },
          {
            "$ref": "#/components/parameters/normalizeimportance"
          },
          {
            "$ref": "#/components/parameters/dedupe-default-1"
          },
          {
            "$ref": "#/components/parameters/namedetails"
          },
          {
            "$ref": "#/components/parameters/extratags"
          },
          {
            "$ref": "#/components/parameters/polygon_geojson"
          },
          {
            "$ref": "#/components/parameters/polygon_kml"
          },
          {
            "$ref": "#/components/parameters/polygon_svg"
          },
          {
            "$ref": "#/components/parameters/polygon_text"
          },
          {
            "$ref": "#/components/parameters/json_callback"
          },
          {
            "$ref": "#/components/parameters/polygon_threshold"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/location-forward"
                },
                "examples": {
                  "response": {
                    "value": [
                      {
                        "place_id": "116136978",
                        "licence": "https://locationiq.com/attribution",
                        "osm_type": "way",
                        "osm_id": "34633854",
                        "boundingbox": [
                          "40.7479255",
                          "40.7489585",
                          "-73.9865012",
                          "-73.9848166"
                        ],
                        "lat": "40.74844205",
                        "lon": "-73.98565890160751",
                        "display_name": "Empire State Building, 350, 5th Avenue, Manhattan Community Board 5, Manhattan, New York County, New York, New York, 10001, USA",
                        "class": "tourism",
                        "type": "attraction",
                        "importance": 0.8515868466874569,
                        "icon": "https://locationiq.org/static/images/mapicons/poi_point_of_interest.p.20.png",
                        "address": {
                          "attraction": "Empire State Building",
                          "house_number": "350",
                          "road": "5th Avenue",
                          "neighbourhood": "Manhattan Community Board 5",
                          "suburb": "Manhattan",
                          "county": "New York County",
                          "city": "New York",
                          "state": "New York",
                          "postcode": "10001",
                          "country": "United States of America",
                          "country_code": "us"
                        }
                      }
                    ]
                  }
                }
              },
              "application/xml": {
                "schema": {
                  "$ref": "#/components/schemas/location-forward"
                },
                "examples": {
                  "Example 1": {
                    "value": "<?xml version=\"1.0\" encoding=\"UTF-8\" ?>\n<searchresults timestamp='Mon, 28 Aug 23 18:14:49 +0530' attribution='https://locationiq.com/attribution' querystring='Empire State Building' polygon='false'>\n    <place place_id='116136978' osm_type='way' osm_id='34633854' boundingbox=\"40.7479255,40.7489585,-73.9865012,-73.9848166\" lat='40.74844205' lon='-73.98565890160751' display_name='Empire State Building, 350, 5th Avenue, Manhattan Community Board 5, Manhattan, New York County, New York, New York, 10001, USA' class='tourism' type='attraction' importance='0.85158684668746' icon='https://locationiq.org/static/images/mapicons/poi_point_of_interest.p.20.png'>\n        <attraction>Empire State Building</attraction>\n        <house_number>350</house_number>\n        <road>5th Avenue</road>\n        <neighbourhood>Manhattan Community Board 5</neighbourhood>\n        <suburb>Manhattan</suburb>\n        <county>New York County</county>\n        <city>New York</city>\n        <state>New York</state>\n        <postcode>10001</postcode>\n        <country>United States of America</country>\n\n        \n        <country_code>us</country_code>\n    </place>\n</searchresults>"
                  }
                }
              }
            }
          },
          "400": {
            "$ref": "#/components/responses/400Error-search"
          },
          "401": {
            "$ref": "#/components/responses/401Error-search"
          },
          "403": {
            "$ref": "#/components/responses/403Error-search"
          },
          "404": {
            "$ref": "#/components/responses/404Error-search"
          },
          "429": {
            "$ref": "#/components/responses/429Error-search"
          },
          "500": {
            "$ref": "#/components/responses/500Error-search"
          }
        },
        "security": [
          {
            "key": []
          }
        ],
        "x-stoplight": {
          "id": "vo6s5vdzq9v1f"
        }
      },
      "parameters": []
    }
  },
  "servers": [
    {
      "url": "https://us1.locationiq.com/v1",
      "description": "US Region Endpoint"
    },
    {
      "url": "https://eu1.locationiq.com/v1",
      "description": "EU Region Endpoint"
    },
    {
      "url": "https://api.locationiq.com/v1",
      "description": "Autocomplete only Endpoint"
    }
  ],
  "components": {
    "securitySchemes": {
      "key": {
        "name": "key",
        "type": "apiKey",
        "in": "query",
        "description": "LocationIQ Access Token"
      }
    },
    "schemas": {
      "namedetails": {
        "type": "object",
        "x-examples": {
          "Example 1": {
            "name": "Empire State Building",
            "name:en": "Empire State Building",
            "name:es": "Edificio Empire State",
            "name:he": "בניין אמפייר סטייט",
            "name:hi": "एम्पायर स्टेट बिल्डिंग",
            "name:ko": "엠파이어 스테이트 빌딩",
            "name:ru": "Эмпайр-Стейт-Билдинг",
            "name:uk": "Емпайр-Стейт-Білдінг",
            "name:zh": "帝国大厦"
          }
        },
        "properties": {
          "name": {
            "type": "string"
          }
        },
        "description": "The dictionary with full list of available names including ref etc. Returned when `namedetails=1` is set in the request.\n"
      },
      "address": {
        "title": "address",
        "type": "object",
        "example": {
          "house_number": "3894",
          "road": "Spring Mill Way",
          "residential": "Hunter’s Point",
          "village": "Landen",
          "city": "Landen",
          "county": "Warren County",
          "state": "Ohio",
          "postcode": "45039",
          "country": "United States of America",
          "country_code": "us",
          "state_code": "oh"
        },
        "description": "Breakdown of the address into elements.\nAll these elements are optional and only those elements that are available for a given location will be returned.",
        "properties": {
          "house_number": {
            "type": "string",
            "description": "House number"
          },
          "road": {
            "type": "string",
            "description": "Road name"
          },
          "neighbourhood": {
            "type": "string",
            "description": "Neighbourhood"
          },
          "hamlet": {
            "type": "string",
            "description": "Hamlet"
          },
          "suburb": {
            "type": "string",
            "description": "Suburb"
          },
          "village": {
            "type": "string",
            "description": "Village name"
          },
          "town": {
            "type": "string",
            "description": "Town name"
          },
          "city_district": {
            "type": "string",
            "description": "Administrative area between city level and town level"
          },
          "city": {
            "type": "string",
            "description": "City name"
          },
          "region": {
            "type": "string",
            "description": "Region name"
          },
          "county": {
            "type": "string",
            "description": "County name"
          },
          "state_district": {
            "type": "string",
            "description": "District name"
          },
          "state": {
            "type": "string",
            "description": "State name"
          },
          "state_code": {
            "type": "string",
            "description": "State code"
          },
          "postcode": {
            "type": "string",
            "description": "Postal code"
          },
          "country": {
            "type": "string",
            "description": "Country name"
          },
          "country_code": {
            "type": "string",
            "description": "Country code"
          },
          "name": {
            "type": "string",
            "description": "Name of the entity/road in the given location"
          },
          "water": {
            "type": "string",
            "x-stoplight": {
              "id": "7nunw3epqpcqo"
            },
            "description": "The name of an ocean or sea, if the location falls within a body of water outside any country's administrative regions."
          }
        },
        "x-examples": {
          "Example 1": {
            "house_number": "3894",
            "road": "Spring Mill Way",
            "residential": "Hunter’s Point",
            "village": "Landen",
            "city": "Landen",
            "county": "Warren County",
            "state": "Ohio",
            "postcode": "45039",
            "country": "United States of America",
            "country_code": "us",
            "state_code": "oh"
          },
          "Result from Ocean": {
            "house_number": "3894",
            "road": "Spring Mill Way",
            "residential": "Hunter’s Point",
            "village": "Landen",
            "city": "Landen",
            "county": "Warren County",
            "state": "Ohio",
            "postcode": "45039",
            "country": "United States of America",
            "country_code": "us",
            "state_code": "oh"
          },
          "Example 2": {
            "name": "South Pacific Ocean",
            "water": "South Pacific Ocean"
          }
        }
      },
      "address-normalized": {
        "title": "address-normalized",
        "x-stoplight": {
          "id": "q798lnglnqkb4"
        },
        "type": "object",
        "example": {
          "house_number": "3894",
          "road": "Spring Mill Way",
          "residential": "Hunter’s Point",
          "village": "Landen",
          "city": "Landen",
          "county": "Warren County",
          "state": "Ohio",
          "postcode": "45039",
          "country": "United States of America",
          "country_code": "us",
          "state_code": "oh"
        },
        "description": "The default address section returns a wide range of elements - from common ones such as `road` and `country` to obscure ones such as `hamlet`, `cycleway`, `park`. This was done to maintain backward compatibility with OpenStreetMap's Nominatim. To make parsing easier for developers, the `normalizeaddress` parameter rolls up elements in the `address` section of the response to the list of elements defined below.\n",
        "properties": {
          "name": {
            "type": "string",
            "description": "House name or Point of Interest (POI)"
          },
          "house_number": {
            "type": "string",
            "description": "House or Building number"
          },
          "road": {
            "type": "string",
            "description": "Roads, Highways, Freeways, Motorways"
          },
          "neighbourhood": {
            "type": "string",
            "description": "Neighbourhoods, Allotments, Quarters, Communities"
          },
          "suburb": {
            "type": "string",
            "description": "Suburbs, Subdivisions"
          },
          "island": {
            "type": "string",
            "description": "Islands, Islets"
          },
          "city": {
            "type": "string",
            "description": "Cities, Towns, Villages, Municipalities, Districts, Boroughs, Hamlets"
          },
          "county": {
            "type": "string",
            "description": "Counties"
          },
          "state": {
            "type": "string",
            "description": "States, Provinces, Regions, State Districts"
          },
          "state_code": {
            "type": "string",
            "description": "State or Province Code"
          },
          "postcode": {
            "type": "string",
            "description": "Postal Codes, Zipcodes"
          },
          "country": {
            "type": "string",
            "description": "Countries, Nation-states"
          },
          "country_code": {
            "type": "string",
            "description": "Country Code - 2 letter (ISO 3166-1 alpha-2)"
          }
        },
        "x-examples": {
          "Example 1": {
            "name": "Empire State Building",
            "house_number": "350",
            "road": "5th Avenue",
            "neighbourhood": "Manhattan Community Board 5",
            "suburb": "Manhattan",
            "city": "New York",
            "county": "New York County",
            "state": "New York",
            "postcode": "10001",
            "country": "United States of America",
            "country_code": "us"
          },
          "Result from Ocean": {
            "name": "South Pacific Ocean"
          },
          "Example 2": {
            "name": "South Pacific Ocean"
          }
        }
      },
      "error": {
        "title": "error",
        "type": "object",
        "properties": {
          "error": {
            "type": "string"
          }
        },
        "example": {
          "error": "Invalid key"
        }
      },
      "location-forward": {
        "title": "location-forward-geocoding",
        "x-stoplight": {
          "id": "360vfacsuwhfx"
        },
        "type": "array",
        "items": {
          "type": "object",
          "properties": {
            "place_id": {
              "type": "string",
              "description": "Unique identifier for the place."
            },
            "licence": {
              "type": "string",
              "description": "License information for the data."
            },
            "osm_type": {
              "type": "string",
              "description": "Type of OpenStreetMap object."
            },
            "osm_id": {
              "type": "string",
              "description": "Unique identifier for the OpenStreetMap object."
            },
            "lat": {
              "type": "string",
              "description": "Latitude of the location."
            },
            "lon": {
              "type": "string",
              "description": "Longitude of the location."
            },
            "display_name": {
              "type": "string",
              "description": "Formatted address for display."
            },
            "class": {
              "type": "string",
              "description": "The category of this result"
            },
            "type": {
              "type": "string",
              "description": "The 'type' of the class/category of this result"
            },
            "importance": {
              "type": "number",
              "description": "Calculated importance of this result compared to the search query the user has provided. Ranges between 0 and 1.",
              "format": "float"
            },
            "address": {
              "anyOf": [
                {
                  "$ref": "#/components/schemas/address"
                },
                {
                  "$ref": "#/components/schemas/address-normalized"
                }
              ]
            },
            "boundingbox": {
              "type": "array",
              "description": "List of bounding box coordinates [min_lat, max_lat, min_lon, max_lon].",
              "items": {
                "type": "string"
              }
            },
            "namedetails": {
              "$ref": "#/components/schemas/namedetails"
            },
            "extratags": {
              "$ref": "#/components/schemas/extratags"
            },
            "geojson": {
              "$ref": "#/components/schemas/geojson"
            },
            "geokml": {
              "$ref": "#/components/schemas/geokml"
            },
            "svg": {
              "$ref": "#/components/schemas/svg"
            },
            "geotext": {
              "$ref": "#/components/schemas/geotext"
            },
            "icon": {
              "type": "string",
              "x-stoplight": {
                "id": "udpck9wsvlgou"
              },
              "description": "The URL of an icon representing this result, if applicable."
            },
            "matchquality": {
              "$ref": "#/components/schemas/matchquality"
            },
            "postaladdress": {
              "$ref": "#/components/schemas/postaladdress"
            }
          },
          "required": [
            "place_id",
            "licence",
            "lat",
            "lon",
            "display_name",
            "boundingbox"
          ]
        },
        "description": "",
        "x-examples": {
          "Example 1": [
            {
              "place_id": "223483692",
              "licence": "© LocationIQ.com CC BY 4.0, Data © OpenStreetMap contributors, ODbL 1.0",
              "osm_type": "way",
              "osm_id": "19301621",
              "boundingbox": [
                "39.307405567782",
                "39.307505567782",
                "-84.292824851595",
                "-84.292724851595"
              ],
              "lat": "39.3074555677816",
              "lon": "-84.2927748515948",
              "display_name": "3894, Spring Mill Way, Hunter’s Point, Landen, Warren County, Ohio, 45039, United States of America",
              "class": "place",
              "type": "house",
              "importance": 0.62025,
              "address": {
                "name": "Empire State Building",
                "house_number": "3894",
                "road": "Spring Mill Way",
                "residential": "Hunter’s Point",
                "village": "Landen",
                "county": "Warren County",
                "state": "Ohio",
                "postcode": "45039",
                "country": "United States of America",
                "country_code": "us",
                "city": "Landen"
              }
            }
          ]
        }
      },
      "extratags": {
        "type": "object",
        "x-examples": {
          "Example 1": {
            "ele": "15",
            "height": "443.2",
            "wikidata": "Q9188",
            "wikipedia": "en:Empire State Building",
            "start_date": "1931",
            "wheelchair": "yes",
            "building:use": "office",
            "opening_hours": "Mo-Su 08:00-02:00",
            "building:levels": "102",
            "construction_date": "1930-1931"
          }
        },
        "description": "The dictionary with additional useful tags like website or maxspeed. Returned when `extratags=1` is set in the request.\n"
      },
      "geojson": {
        "type": "object",
        "properties": {
          "type": {
            "type": "string"
          },
          "coordinates": {
            "type": "array",
            "items": {
              "type": "number"
            }
          }
        },
        "x-examples": {
          "Example 1": {
            "type": "Polygon",
            "coordinates": [
              [
                [
                  -73.9865012,
                  40.748491
                ],
                [
                  -73.9851602,
                  40.7479255
                ],
                [
                  -73.9848166,
                  40.7483931
                ],
                [
                  -73.9861574,
                  40.7489585
                ],
                [
                  -73.9863252,
                  40.7487301
                ],
                [
                  -73.9863554,
                  40.748689
                ],
                [
                  -73.9864839,
                  40.7485145
                ],
                [
                  -73.9865012,
                  40.748491
                ]
              ]
            ]
          }
        },
        "description": "Output geometry of results in geojson format. Returned when `polygon_geojson=1` is set in the request."
      },
      "geokml": {
        "type": "string",
        "x-examples": {
          "Example 1": "<Polygon><outerBoundaryIs><LinearRing><coordinates>-73.986501200000006,40.748491000000001 -73.985160199999996,40.747925500000001 -73.984816600000002,40.748393100000001 -73.986157399999996,40.748958500000001 -73.986325199999996,40.748730100000003 -73.986355399999994,40.748688999999999 -73.986483899999996,40.748514499999999 -73.986501200000006,40.748491000000001</coordinates></LinearRing></outerBoundaryIs></Polygon>"
        },
        "description": "Output geometry of results in kml format. Returned when `polygon_kml=1` is set in the request."
      },
      "svg": {
        "type": "string",
        "x-examples": {
          "Example 1": "M -73.9865012 -40.748491 L -73.9851602 -40.7479255 -73.9848166 -40.7483931 -73.9861574 -40.7489585 -73.9863252 -40.7487301 -73.9863554 -40.748689 -73.9864839 -40.7485145 Z"
        },
        "description": "Output geometry of results in svg format. Returned when `polygon_svg=1` is set in the request."
      },
      "geotext": {
        "type": "string",
        "x-examples": {
          "Example 1": "POLYGON((-73.9865012 40.748491,-73.9851602 40.7479255,-73.9848166 40.7483931,-73.9861574 40.7489585,-73.9863252 40.7487301,-73.9863554 40.748689,-73.9864839 40.7485145,-73.9865012 40.748491))"
        },
        "title": "",
        "description": "Output geometry of results as a WKT. Returned when `polygon_text=1` is set in the request."
      },
      "matchquality": {
        "type": "object",
        "x-examples": {
          "Example 1": {
            "matchcode": "exact",
            "matchtype": "point",
            "matchlevel": "venue"
          }
        },
        "description": "An additional object `matchquality` for every result in the response, containing the following elements: `matchcode`, `matchtype`, `matchlevel`.",
        "properties": {
          "matchcode": {
            "$ref": "#/components/schemas/matchcode"
          },
          "matchtype": {
            "$ref": "#/components/schemas/matchtype"
          },
          "matchlevel": {
            "$ref": "#/components/schemas/matchlevel"
          }
        }
      },
      "matchcode": {
        "title": "matchcode",
        "x-stoplight": {
          "id": "keiun4kwm07z4"
        },
        "type": "string",
        "description": "Specifies the quality of the returned address.\n\n matchcode  | description\n ------------|---------------\n  exact      | The result matches the input query with a high level of probability.\n  fallback   | The result does not exactly match the input but is closely related to it provided there is direct a heierarchial relation.\n  approximate| The result matches the input query with a medium to low level of probability.\n"
      },
      "matchtype": {
        "title": "matchtype",
        "x-stoplight": {
          "id": "20ejxqp1iy2wr"
        },
        "type": "string",
        "description": "Specifies quality of the returned location match\n  \n  matchtype    | description\n --------------|---------------\n  point        | The coordinate returned is a point address, typically with rooftop accuracy.\n  centroid     | The coordinate returned is a centroid of a road or administrative boundary.\n  interpolated | The coordinate returned is a point determined by interpolation."
      },
      "matchlevel": {
        "type": "string",
        "x-examples": {
          "Example 1": "venue"
        },
        "description": "Specifies the most granular address element that matches the geocoding query.\n\n matchlevel       | details\n -----------------|---------------\n  venue           | The returned address is of a Point of Interest (PoI) level.\n  building        | The returned address is of a house level.\n  street          | The returned address is on a street level.\n  neighbourhood   | The returned address is on a neighbourhood level.\n  island          | The returned address is on a island level.\n  borough         | The returned address is on a borough level.\n  city            | The returned address is on a city level.\n  county          | The returned address is on a county level.\n  state           | The returned address is on a state level.\n  country         | The returned address is on a country level.\n  marine          | The returned address is on a marine level.\n  postalcode      | The returned address is on a postalcode level."
      },
      "postaladdress": {
        "title": "postaladdress",
        "x-stoplight": {
          "id": "zq4zgnef9uq9n"
        },
        "type": "string",
        "description": "Returns address specifically formatted for each country. Returned when `postaladdress` is set in the request",
        "x-examples": {
          "Example 1": "5, Avenue Anatole France, 75007, Paris, France"
        }
      }
    },
    "parameters": {
      "query": {
        "name": "q",
        "in": "query",
        "description": "Free-form query string to search for. Commas are optional, but improves performance by reducing the complexity of the search.",
        "schema": {
          "type": "string",
          "example": "Empire State Building"
        },
        "required": true
      },
      "addressdetails-default-0": {
        "name": "addressdetails",
        "in": "query",
        "required": false,
        "schema": {
          "type": "integer",
          "enum": [
            0,
            1
          ],
          "default": 0
        },
        "description": "Include a breakdown of the address of this result into elements. Defaults to `0`."
      },
      "format": {
        "name": "format",
        "in": "query",
        "required": false,
        "schema": {
          "type": "string",
          "default": "xml",
          "enum": [
            "xml",
            "json",
            "xmlv1.1"
          ],
          "example": "json"
        },
        "description": "Output Format. Defaults to xml. \n\n> This version (v1) of our Reverse Geocoding API is compatible with OpenStreetMap's Nominatim Geocoder in both JSON & XML formats. However, all our enhancements such as additional datasets and algorithms are supported only in `json` or `xmlv1.1` format options."
      },
      "accept-language": {
        "name": "accept-language",
        "in": "query",
        "required": false,
        "schema": {
          "type": "string",
          "default": "en",
          "example": "en"
        },
        "description": "Preferred language order for showing search results, overrides the value specified in the `Accept-Language` HTTP header. Defaults to `en`. \n\nTo use native language for the response when available, use `accept-language=native`. \n\nEither uses standard <a href=\"https://tools.ietf.org/html/rfc2616#section-14.4\" target=\"_blank\">rfc2616 accept-language string</a> or a simple comma separated list of language codes."
      },
      "namedetails": {
        "name": "namedetails",
        "in": "query",
        "schema": {
          "type": "integer",
          "enum": [
            0,
            1
          ],
          "default": 0
        },
        "description": "Include a list of alternative names in the results. These may include language variants, references, operator and brand. Defaults to `0`."
      },
      "json_callback": {
        "name": "json_callback",
        "in": "query",
        "required": false,
        "schema": {
          "type": "string"
        },
        "description": "Wrap json output in a callback function (JSONP) i.e. &lt;string&gt;(&lt;json&gt;). Only has an effect for JSON output formats."
      },
      "polygon_geojson": {
        "name": "polygon_geojson",
        "in": "query",
        "required": false,
        "schema": {
          "type": "integer",
          "enum": [
            0,
            1
          ],
          "default": 0
        },
        "description": "Output geometry of results in geojson format. Defaults to `0`."
      },
      "polygon_kml": {
        "name": "polygon_kml",
        "in": "query",
        "required": false,
        "schema": {
          "type": "integer",
          "enum": [
            0,
            1
          ],
          "default": 0
        },
        "description": "Output geometry of results in kml format. Defaults to `0`."
      },
      "polygon_svg": {
        "name": "polygon_svg",
        "in": "query",
        "required": false,
        "schema": {
          "type": "integer",
          "enum": [
            0,
            1
          ],
          "default": 0
        },
        "description": "Output geometry of results in svg format. Defaults to `0`."
      },
      "polygon_text": {
        "name": "polygon_text",
        "in": "query",
        "required": false,
        "schema": {
          "type": "integer",
          "enum": [
            0,
            1
          ],
          "default": 0
        },
        "description": "Output geometry of results as a WKT. Defaults to `0`."
      },
      "extratags": {
        "name": "extratags",
        "in": "query",
        "required": false,
        "schema": {
          "type": "integer",
          "enum": [
            0,
            1
          ],
          "default": 0
        },
        "description": "Include additional information in the result if available, e.g. wikipedia link, opening hours. Defaults to `0`."
      },
      "normalizeaddress": {
        "name": "normalizeaddress",
        "in": "query",
        "required": false,
        "schema": {
          "type": "integer",
          "enum": [
            0,
            1
          ],
          "default": 0
        },
        "description": "Makes parsing of the `address` object easier by returning a predictable and defined list of elements. Defaults to `0` for backward compatibility. We recommend setting this to `1` for new projects.\n\n  Element Name  | Description\n  ------------- | -----------\n  name          | House name or Point of Interest (POI) such as a Cafe or School\n  house_number  | House or Building number\n  road          | Roads, Highways, Freeways, Motorways\n  neighbourhood | Neighbourhoods, Allotments, Quarters, Communities\n  suburb        | Suburbs, Subdivisions\n  island        | Islands, Islets\n  city          | Cities, Towns, Villages, Municipalities, Districts, Boroughs, Hamlets\n  county        | Counties\n  state         | States, Provinces, Regions, State Districts\n  state_code    | State or Province Code\n  postcode      | Postal Codes, Zipcodes\n  country       | Countries, Nation-states\n  country_code  | Country Code - 2 letter (ISO 3166-1 alpha-2)"
      },
      "normalizecity": {
        "name": "normalizecity",
        "in": "query",
        "required": false,
        "schema": {
          "type": "integer",
          "enum": [
            0,
            1
          ],
          "default": 0
        },
        "description": "For responses with no `city` value in the address section, the next available element in this order - `city_district`, `locality`, `town`, `borough`, `municipality`, `village`, `hamlet`, `quarter`, `neighbourhood` - from the address section will be normalized to city. Defaults to `0`."
      },
      "statecode": {
        "name": "statecode",
        "in": "query",
        "required": false,
        "schema": {
          "type": "integer",
          "enum": [
            0,
            1
          ],
          "default": 0
        },
        "description": "Adds state or province code when available to the `state_code` key inside the `address` object when available. Defaults to `0`."
      },
      "postaladdress": {
        "name": "postaladdress",
        "in": "query",
        "required": false,
        "schema": {
          "type": "integer",
          "enum": [
            0,
            1
          ],
          "default": 0
        },
        "description": "Returns address inside the `postaladdress` key, that is specifically formatted for each country. Currently supported for addresses in Belgium, Brazil, France, Germany, Greece, India, Ireland, Italy, Portugal, South Africa, Spain and United Kingdom. Defaults to `0`."
      },
      "source": {
        "name": "source",
        "in": "query",
        "required": false,
        "schema": {
          "type": "string"
        },
        "description": "If this parameter is not specified, LocationIQ uses multiple public and proprietary datasets to return results. If you'd like to restrict results to only OpenStreetMap data, you can set the value of this parameter to `nom`. This will only query our internal cluster of Nominatim servers, and return results. We may still apply some post-processing steps to these results though, so results may vary from the official Nominatim instance."
      },
      "viewbox": {
        "name": "viewbox",
        "in": "query",
        "required": false,
        "schema": {
          "type": "string",
          "example": "-73.9965012,40.7489255,-73.9858166,40.7499585"
        },
        "description": "The preferred area to find search results. Any two corner points of the box - `max_lon,max_lat,min_lon,min_lat` or `min_lon,min_lat,max_lon,max_lat` - are accepted in any order as long as they span a real box. To restrict results to those within the viewbox, use along with the `bounded` option."
      },
      "bounded": {
        "name": "bounded",
        "in": "query",
        "required": false,
        "schema": {
          "type": "integer",
          "enum": [
            0,
            1
          ]
        },
        "description": "Restrict result to items contained within the bounds specified in the `viewbox` parameter. Defaults to `0`."
      },
      "countrycodes": {
        "name": "countrycodes",
        "in": "query",
        "required": false,
        "schema": {
          "type": "string",
          "example": "us,ca,gb"
        },
        "description": "Limit search results to a specific country or a comma-separated list of countries. Should be the ISO 3166-1 alpha-2 code(s)."
      },
      "dedupe-default-1": {
        "name": "dedupe",
        "in": "query",
        "required": false,
        "schema": {
          "type": "integer",
          "enum": [
            0,
            1
          ],
          "default": 1
        },
        "description": "Sometimes you have several objects in OSM identifying the same place or object in reality. The simplest case is a street being split in many different OSM ways due to different characteristics. Our Geocoder will attempt to detect such duplicates and only return one match; this is controlled by the dedupe parameter which defaults to `1`. Since the limit is, for reasons of efficiency, enforced before and not after de-duplicating, it is possible that de-duplicating leaves you with less results than requested."
      },
      "matchquality": {
        "name": "matchquality",
        "in": "query",
        "required": false,
        "schema": {
          "type": "integer",
          "enum": [
            0,
            1
          ],
          "default": 0
        },
        "description": "Returns additional information about quality of the result in a `matchquality` object. Defaults to `0`."
      },
      "limit-max-50": {
        "name": "limit",
        "in": "query",
        "required": false,
        "schema": {
          "type": "integer",
          "minimum": 1,
          "maximum": 50
        },
        "description": "Limit the number of returned results. Accepted value: `1` to `50`. Defaults to `10`."
      },
      "normalizeimportance": {
        "name": "normalizeimportance",
        "in": "query",
        "schema": {
          "type": "integer",
          "default": 1,
          "enum": [
            0,
            1
          ]
        },
        "description": "When this parameter is absent or set to `1`, the `importance` value(s) in the API response is limited to the range of 0 to 1. Values outside this range are adjusted to the nearest boundary (0 or 1). Setting `normalizeimportance` to `0` allows the importance value to be lower or higher than the specified range of 0 to 1. Defaults to `1`"
      },
      "polygon_threshold": {
        "name": "polygon_threshold",
        "in": "query",
        "required": false,
        "schema": {
          "type": "number",
          "example": 0.2,
          "default": 0
        },
        "description": "When one of the polygon_* outputs is chosen, return a simplified version of the output geometry. The parameter describes the tolerance in degrees with which the geometry may differ from the original geometry. Topology is preserved in the geometry."
      }
    },
    "responses": {
      "400Error-search": {
        "description": "Bad Request",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/error"
            },
            "examples": {
              "Invalid Request": {
                "value": {
                  "error": "Invalid Request"
                }
              }
            }
          },
          "application/xml": {
            "schema": {
              "type": "object",
              "properties": {
                "error": {
                  "type": "string"
                }
              },
              "x-examples": {
                "Example 1": {
                  "error": "Invalid Request"
                }
              }
            },
            "examples": {
              "Invalid Request": {
                "value": "<?xml version='1.0' encoding='UTF-8' ?>\n<searchresults timestamp='Wed, 30 Aug 23 07:49:28 +0000' attribution='© LocationIQ.com CC BY 4.0, Data © OpenStreetMap contributors, ODbL 1.0'>\n    <error>Invalid Request</error>\n</searchresults>"
              }
            }
          }
        }
      },
      "401Error-search": {
        "description": "Unauthorized",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/error"
            },
            "examples": {
              "Invalid Key": {
                "value": {
                  "error": "Invalid Key"
                }
              }
            }
          },
          "application/xml": {
            "schema": {
              "type": "object",
              "properties": {
                "error": {
                  "type": "string"
                }
              },
              "x-examples": {
                "Example 1": {
                  "error": "Invalid Key"
                }
              }
            },
            "examples": {
              "Invalid Key": {
                "value": "<?xml version='1.0' encoding='UTF-8' ?>\n<searchresults timestamp='Wed, 30 Aug 23 07:49:28 +0000' attribution='© LocationIQ.com CC BY 4.0, Data © OpenStreetMap contributors, ODbL 1.0'>\n    <error>Invalid Key</error>\n</searchresults>"
              }
            }
          }
        }
      },
      "403Error-search": {
        "description": "The request has been made from an unauthorized domain.",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/error"
            },
            "examples": {
              "Access restricted": {
                "value": {
                  "error": "Access restricted"
                }
              }
            }
          },
          "application/xml": {
            "schema": {
              "type": "object",
              "properties": {
                "error": {
                  "type": "string"
                }
              },
              "x-examples": {
                "Example 1": {
                  "error": "Access restricted"
                }
              }
            },
            "examples": {
              "Access restricted": {
                "value": "<?xml version='1.0' encoding='UTF-8' ?>\n<searchresults timestamp='Wed, 30 Aug 23 07:49:28 +0000' attribution='© LocationIQ.com CC BY 4.0, Data © OpenStreetMap contributors, ODbL 1.0'>\n    <error>Access restricted</error>\n</searchresults>"
              }
            }
          }
        }
      },
      "404Error-search": {
        "description": "No location or places were found for the given input.",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/error"
            },
            "examples": {
              "Unable to geocode": {
                "value": {
                  "error": "Unable to geocode"
                }
              }
            }
          },
          "application/xml": {
            "schema": {
              "type": "object",
              "properties": {
                "error": {
                  "type": "string"
                }
              },
              "x-examples": {
                "Example 1": {
                  "error": "Unable to geocode"
                }
              }
            },
            "examples": {
              "Unable to geocode": {
                "value": "<?xml version='1.0' encoding='UTF-8' ?>\n<searchresults timestamp='Wed, 30 Aug 23 07:49:28 +0000' attribution='© LocationIQ.com CC BY 4.0, Data © OpenStreetMap contributors, ODbL 1.0'>\n    <error>Unable to geocode</error>\n</searchresults>"
              }
            }
          }
        }
      },
      "429Error-search": {
        "description": "Request exceeded the rate-limits set on your account.",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/error"
            },
            "examples": {
              "Rate Limited Day": {
                "value": {
                  "error": "Rate Limited Day"
                }
              },
              "Rate Limited Minute": {
                "value": {
                  "error": "Rate Limited Minute"
                }
              },
              "Rate Limited Second": {
                "value": {
                  "error": "Rate Limited Second"
                }
              }
            }
          },
          "application/xml": {
            "schema": {
              "type": "object",
              "properties": {
                "error": {
                  "type": "string"
                }
              },
              "x-examples": {
                "Example 1": {
                  "error": "Rate Limited Second"
                }
              }
            },
            "examples": {
              "Rate Limited Day": {
                "value": "<?xml version='1.0' encoding='UTF-8' ?>\n<searchresults timestamp='Wed, 30 Aug 23 07:49:28 +0000' attribution='© LocationIQ.com CC BY 4.0, Data © OpenStreetMap contributors, ODbL 1.0'>\n    <error>Rate Limited Day</error>\n</searchresults>"
              },
              "Rate Limited Minute": {
                "value": "<?xml version='1.0' encoding='UTF-8' ?>\n<searchresults timestamp='Wed, 30 Aug 23 07:49:28 +0000' attribution='© LocationIQ.com CC BY 4.0, Data © OpenStreetMap contributors, ODbL 1.0'>\n    <error>Rate Limited Minute</error>\n</searchresults>"
              },
              "Rate Limited Second": {
                "value": "<?xml version='1.0' encoding='UTF-8' ?>\n<searchresults timestamp='Wed, 30 Aug 23 07:49:28 +0000' attribution='© LocationIQ.com CC BY 4.0, Data © OpenStreetMap contributors, ODbL 1.0'>\n    <error>Rate Limited Second</error>\n</searchresults>"
              }
            }
          }
        }
      },
      "500Error-search": {
        "description": "Internal Server Error",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/error"
            },
            "examples": {
              "Internal Server Error": {
                "value": {
                  "error": "Unknown error - Please try again after some time"
                }
              }
            }
          },
          "application/xml": {
            "schema": {
              "type": "object",
              "properties": {
                "error": {
                  "type": "string"
                }
              },
              "x-examples": {
                "Example 1": {
                  "error": "Internal Server Error"
                }
              }
            },
            "examples": {
              "Internal Server Error": {
                "value": "<?xml version='1.0' encoding='UTF-8' ?>\n<searchresults timestamp='Wed, 30 Aug 23 07:49:28 +0000' attribution='© LocationIQ.com CC BY 4.0, Data © OpenStreetMap contributors, ODbL 1.0'>\n    <error>Unknown error - Please try again after some time</error>\n</searchresults>"
              }
            }
          }
        }
      }
    }
  }
}
```
