package programscrud

func Mainloopsimple() {
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			//break
			//continue-> é parecido com o break mas é uma interrupcao numa iteraçao do loop for
			continue
		}
		println("Continuing...")
	}
}

func Mainloopsmoresimple() {
	for i := 0; i < 5; i++ {
		println(i)
	}
}

//ugly way
func InfiniteLoop() {
	var i int
	for {
		//Ciclo infinito se tirar esta condicao
		if i == 5 {
			break
		}
		println(i)
		i++
	}
}

func LoopOverCollections() {
	wellKnownPorts := map[string]int{"http": 80, "https": 443}

	for _, v := range wellKnownPorts {
		println(v)
	}
}
