package lasagnamaster

func PreparationTime(layers []string, timelayer int) int{
    if timelayer == 0 {
        timelayer = 2
    }
    return len(layers) * timelayer
}

func Quantities(layers []string) (int, float64){
	var noodles int
    var sauce float64
    for _, layer := range layers {
        switch layer {
        case "noodles":
            noodles += 50
        case "sauce":
            sauce += 0.2
        }
    }
    return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64{
    scaled := make([]float64, len(quantities))
    for i, q := range quantities {
        scaled[i] = q * float64(portions) / 2
    } 
    return scaled
}