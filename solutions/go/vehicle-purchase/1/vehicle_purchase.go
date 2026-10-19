package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    license := false
	if kind == "car" { license = true }
    if kind == "truck" { license = true }
    return license
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
    func ChooseVehicle(option1, option2 string) string {
        var choose string
        if option1 < option2 {
            choose = option1 + " is clearly the better choice."
        } else {
            choose = option2 + " is clearly the better choice."
        }
        return choose
    }

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var resellPrice float64
    if age != 0 && age < 3 {
        resellPrice = originalPrice * 0.8
    } else if age < 10 {
        resellPrice = originalPrice * 0.7
    } else if age >= 10 {
        resellPrice = originalPrice * 0.5
    } else {
        resellPrice = originalPrice
    }
    return resellPrice
}
