package usanames

import "math/rand"

// NaiveRandomName "naively" generates a random name, based on picking a
// random number based on all of the available names. It is "naive" because
// it does not consider the frequency of the name (which is data that is
// available, and could be used to weight the name generator toward more
// commonly-used names.
//
// This function uses the math/rand package to obtain a random number. The
// random number generator must be seeded for a random name to be chosen,
// otherwise the chosen name will be deterministic.
func NaiveRandomName() string {
	firstNameIndex := rand.Intn(len(FirstNames))
	lastNameIndex := rand.Intn(len(LastNames))

	return FirstNames[firstNameIndex] + " " + LastNames[lastNameIndex]
}
