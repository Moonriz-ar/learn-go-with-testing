package concurrency

// an operation that does not block in Go will run in a separate process called a goroutine
// to tell go to start a new goroutine we turn a function call into a go statement by putting the keyword go in front of it: go doSomething()

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// a channel of result
	resultChannel := make(chan result)

	for _, url := range urls {
		// we often use anonymous functions when we want to start a goroutine. Two useful features:
		// 1/ they can be executed at the same time that they are declared. This is what the () at the end of the anonymous func is doing.
		// 2/ they maintain access to the lexical scope in which they are defined, all the variables that are available at the point when you declare the anonymous function are also available in the body of the function

		// the body of the anonymous function is just the same as the loop body was before. The only difference is that each iteration of the loop will start a new goroutine, concurrent with the current process (the WebsiteChecker function). Each goroutine will add its result to the results map.
		go func(u string) {
			// send statement, taking a channel on the left and a vaue on the right
			// by sending results into a channel, we can control the timing of each write into the results map, ensuring that it happens one at a time. Although each of the calls of wc, and each send to the result channel is happening in parallel inside its own process, each of the results is being dealt with one at a time as we take values out of the result channel with the receive expression
			// we have parallelized the part of the code that we wanted to make faster, while making sure that the part that cannot happen in parallel still happens linearly. And we have communicated across the multiple processes involved by using channels
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// receive expression, which assigns a value received from a channel to a variable
		r := <-resultChannel
		// we then use the result received to update the map
		results[r.string] = r.bool
	}

	return results
}
