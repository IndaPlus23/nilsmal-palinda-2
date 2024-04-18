If written answers are required, you can add them to this file. Just copy the
relevant questions from the root of the repo, preferably in
[Markdown](https://guides.github.com/features/mastering-markdown/) format :)


### bug 1

	The reason for the bug is that there are no gorutines to handle the channel
	By adding a gorutine to handle the channel, the program will work as expected

	Now the ch <- "Hello world!" simply sends the message to the channel and the gorutine will handle the message and print it


### bug 2

 The problem with the code is that the program exits before the goroutine has finished.
 By adding a wait group, the main function will wait for the goroutine to finish before exiting.
 And with that it wont skip the last number and print out the entire series.



### many2many

	If I increase the number of consumers the program would run faster as there are more consumers to handle the data being pushed to the channel.

	If Wait and close change places, the program will throw an err beacuse the channel will be closed before the producers are done.

	If I were to remove the close(ch), the program would still exit fine, but in a larger program having open channels may lead to performance cost, it's always best practise to clean up the channels.

	Because I wait for all the producers to finish but not consumers before closing the channel, there may be consumers who don't get to print all the strings, and so I cant be certain all of the consumers finish printing their strings.

	If I move the close(ch) to Produce, the program will throw an err beacuse the channel will be closed before all the producers are done, only one of the producers will be able to finish.

