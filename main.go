
package main

// importing the required packages
import (

	"fmt"
	"os/exec"
	"strings"
	
)


// define the global variables 
var (

	acrName = "testacrmk"
	test_image = "hello-world:1.1.0"

)


// functino to list the images with tags

func listImagesWithTags(acrName string) (bool, error) {
	// Get the list of repositories (images) in the Azure Container Registry
	repoCmd := exec.Command("az", "acr", "repository", "list", "--name", acrName, "--output", "tsv")

	fmt.Println(repoCmd)

	repositories, err := repoCmd.Output()

	fmt.Println(repositories)

	if err != nil {

		return false, fmt.Errorf("error getting list of repositories: %v", err)
	}

	repoNames := strings.Fields(string(repositories))

	// Loop through each repository and get the list of tags
	for _, repo := range repoNames {

		// Get the list of tags for the current repository
		tagCmd := exec.Command("az", "acr", "repository", "show-tags", "--name", acrName, "--repository", repo, "--output", "tsv")

		tags, err := tagCmd.Output()

		if err != nil {

			return false, fmt.Errorf("error getting tags for repository %s: %v", repo, err)
		}

		// Split the tags string into individual tag names
		tagNames := strings.Fields(string(tags))

		// Check if the test image with tag exists in this repository or not
		for _, tag := range tagNames {
			if repo+":"+tag == test_image {
				return true, nil
			}
		}
	}

	return false, nil
}


// main function where we are calling the listImagesWithTags function

func main() {

	exists, err := listImagesWithTags(acrName)

	if err != nil {

		fmt.Println("Error:", err)

	} else if exists {

		fmt.Printf("Image %s exists in the Azure Container Registry.\n", test_image)
	} else {

		fmt.Printf("Image %s does not exist in the Azure Container Registry.\n", test_image)
	}
}

