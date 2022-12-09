package nickname

import (
	"context"
	"math/rand"
)



// This is a simple REST API that responds with a nickname.
//
//encore:api public path=/nickname
func Name(ctx context.Context) (*Response, error) {
	
	names := [17]string{"Chief", "Buddy", "Friend", "Ace", "Starlord", "Captain", "Snake",
	 "Senior", "Junior", "Bud", "Ice", "Mario", "Big","Vanilla","Biscuit","Peanut","Bingo","Sting","Goose"}
	poo := rand.Intn(len(names)-0) + 0
	nick := names[poo]
	return &Response{Nickname: nick}, nil
}

type Response struct {
	Nickname string
}
