package sample

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/xvbnm48/go-user-management/pb"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}
func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomGPUBrand() string {
	return randomStringFromSet("Nvidia", "AMD")
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	widht := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(widht),
	}

	return resolution
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}

	return pb.Screen_OLED
}

func randomGPUName(brand string) string {
	if brand == "Nvidia" {
		return randomStringFromSet(
			"GeForce GTX 1050",
			"GeForce GTX 1060",
			"GeForce GTX 1070",
			"GeForce GTX 1080",
			"GeForce GTX 1650",
		)
	}
	return randomStringFromSet(
		"Radeon RX 550",
		"Radeon RX 560",
		"Radeon RX 570",
		"Radeon RX 580",
	)
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomCPUbrand(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet("I3", "I5", "I7", "I9")
	}

	return randomStringFromSet(
		"Ryzen 3",
		"Ryzen 5",
		"Ryzen 7",
		"Ryzen 9",
	)
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet(
		"Apple",
		"Lenovo",
		"Dell",
	)
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro", "Macbook")
	case "Dell":
		return randomStringFromSet("XPS 13", "XPS 15")
	default:
		return randomStringFromSet("Thinkpad X1 Carbon", "Thinkpad T480")
	}
}
