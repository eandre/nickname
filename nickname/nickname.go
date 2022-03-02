package nickname

import (
	"context"
	"math/rand"
)


// This is a simple REST API that responds with a nickname.
//
//encore:api public path=/nickname
func Name(ctx context.Context) (*Response, error) {
	
	names := [13]string{"Chief", "Buddy", "Friend", "Ace", "Starlord", "Captain", "Snake",
	 "Senior", "Junior", "Bud", "Ice", "Mario", "Big"}
	poo := rand.Intn(len(names)-0) + 0
	nick := names[poo]
	return &Response{Nickname: nick}, nil
}

type Response struct {
	Nickname string
}



// ==================================================================

// Encore comes with a built-in development dashboard for
// exploring your API, viewing documentation, debugging with
// distributed tracing, and more. Visit your API URL in the browser:
//
//     http://localhost:4000
//

// ==================================================================

// Next steps
//
// 1. Deploy your application to the cloud with a single command:
//
//     git push encore
//
// 2. To continue exploring Encore, check out one of these topics:
//
//    Building a Slack bot:  https://encore.dev/docs/tutorials/slack-bot
//    Building a REST API:   https://encore.dev/docs/tutorials/rest-api
//    Using SQL databases:   https://encore.dev/docs/develop/sql-database
//    Authenticating users:  https://encore.dev/docs/develop/auth
