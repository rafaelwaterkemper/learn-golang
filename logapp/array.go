package main

func main() {
	var servicesVersion *[]string = &[]string{"Rafa", "Loca"}
	var stable *[]string = &[]string{}

	for _, val := range *servicesVersion {
		println(" primeiro, ", val)
	}

	stable = servicesVersion
	
	for _, val := range *stable {
		println(" segundo, ", val)
	}

	temp := append(*servicesVersion, "doente")
	servicesVersion = &temp
	
	for _, val := range *servicesVersion {
		println(" primeiro, ", val)
	}

	stable = servicesVersion

	*servicesVersion = append(*servicesVersion, "careca")

	for _, val := range *stable {
		println(" segundo, ", val)
	}	

}