package actor

import (
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/raspiantoro/go-actor/payload"
	"github.com/raspiantoro/go-actor/random"
)

type message struct {
	id  string
	val uint64
}

type responderActor struct {
	*payload.Counter
}

func (a *responderActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *message:
		random.Sleep()
		a.Count += msg.val
		msg.val = a.Count
		ctx.Send(ctx.Sender(), msg)
		fmt.Printf("reply has been sent to requestor %s with total count: %d\n", msg.id, a.Count)
	}
}

type requestorActor struct{}

func (a *requestorActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *message:
		random.Sleep()
		fmt.Printf("requestor %s receive from %s, current total count: %d\n", ctx.Self().Id, msg.id, msg.val)
	}
}

func ExecuteActor() {
	requestorNums := 10

	system := actor.NewActorSystem()
	rpProps := actor.PropsFromProducer(func() actor.Actor {
		return &responderActor{
			Counter: &payload.Counter{},
		}
	})
	rpActor := system.Root.Spawn(rpProps)

	rqProps := actor.PropsFromProducer(func() actor.Actor { return &requestorActor{} })

	for i := 0; i < requestorNums; i++ {
		rqActor := system.Root.Spawn(rqProps)
		system.Root.RequestFuture(rpActor, &message{id: rqActor.Id, val: 1}, 3*time.Second).PipeTo(rqActor)
		fmt.Printf("requestor %s already sent message\n", rqActor.Id)
	}
}
