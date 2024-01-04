package render

import (
	"github.com/mlctrez/go-nap/nap"
	"github.com/mlctrez/go-nap/nap/jsa"
	"time"
)

func Run(router nap.Router) {
	jsa.Global().Set("napReleaseId", jsa.FuncOf(func(this jsa.Value, args []jsa.Value) any {
		i := args[0].Int()
		//fmt.Println("release", i)
		nap.ReleaseElem(uint64(i))
		return nil
	}))
	jsa.Call("napInit", true)
	navCurrent := func() { router.Navigate(jsa.CurrentURL()) }
	jsa.OnPopState(navCurrent)
	router.QueueOp(navCurrent)
	ticker := time.NewTicker(time.Second / 120)
	for {
		select {
		case op := <-router.Ops():
			op()
		case <-ticker.C:
			break
		}
	}
}
