package space

type Planet string

func Age(seconds float64, planet Planet) (float64) {
	var years = seconds / (60*60*24*365.25)
    switch planet {
		case "Mercury":
		years /= 0.2408467
		case "Venus":
		years /= 0.61519726
		case "Mars":
		years /= 1.8808158
		case "Jupiter":
		years /= 11.862615
		case "Saturn":
		years /= 29.447498
		case "Uranus":
		years /= 84.016846	
		case "Neptune":
		years /= 164.79132
		default:	
	}
	return years
}

