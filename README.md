# Check the Images in the ACR that image is exist or not

### With the help of this template, you can easily check whether the image is available in the particular ACR or not.

### Follow the below steps to run the terratest code:


You need to define these values in your code as global varibale before running:-

                	acrName = "testacrmk"               // your ACR name
	                test_image = "hello-world:1.1.0"    // Imagw which you are looking 

Step 1:- Run the go initialization command:

            go mod init < name >

Step 2:- Run the tidy command to install the packages:-

            go mod tidy

Step 3:- Run the test command:-

            go test -v
