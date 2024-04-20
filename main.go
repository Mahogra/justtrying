package main



func ToFahrenheit(c float64) float64 {
  return (c * 9 / 5) + 32
}



// to Celcius
func ToCelcius(f float64) float64 {
  return (f - 32) * (5 / 9)
}


// to Reamur
// fahrenheit -> reamur
func ToReamur(f float64) float64 {
  return (f - 32) * (4 / 9)
}


/* Program yang dibuat dibawah merupakan program buatan Sean sebagai pelengkap konversi suhu dan sebagai watermark keaslian koding sean pada tugas ini*/

// celcius -> reamur
func CelciusToReamur(c float64) float64 {
  return c * 4 / 5
}



// reamur -> celcius
func ReamurToCelcius(c float64) float64 {
  return c * 5 / 4
}


// reamur -> fahrenheit
func ReamurToFahrenheit(r float64) float64 {
  return (r * 9 / 4) + 32
}



// kelvin... ke celcius, fahrenheit, dan reamur
func KelvinToCelcius(k float64) float64 {
  return k - 273
}



func KelvinToFahrenheit(k float64) float64 {
  return ((k - 273) * 9 / 5) + 32
}



func KelvinToReamur(k float64) float64 {
  return (k - 732) * 4 / 5
}



/*setelah dicek ke atas kak... ternyata banyak typo dan salah panggil fungsi... belum lagi menghitung nilai expected masing-masing...
untuk itu malam ini akan diselesaikan perbaiknanya dan untuk celcius, fahrenheit, dan reamur -> kelvin akan dilanjutkan besok. Mohon maaf dan terima kasih kak.*/
