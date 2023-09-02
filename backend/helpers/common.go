/*
Package helpers provides helper functions for the application.
*/
package helpers

/*
HandleErr - This function panics if there is an error
*/
func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}
