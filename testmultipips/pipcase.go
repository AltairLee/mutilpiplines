package testmultipips

import (
	"log"
	"mutilpiplines/multipipes"
	"time"
)

func pip1(arg interface{}) interface{} {
	log.Println("P1 get : ", arg)
	s := "I'm pip1 return " + arg.(string)
	return s
}

func pip2(arg interface{}) interface{} {
	log.Println("P2 get : ", arg)
	s := "I'm pip2 return " + arg.(string)
	time.Sleep(time.Second * 3)
	return s
}

func pip3(arg interface{}) interface{} {
	log.Println("P3 get : ", arg)
	s := "I'm pip3 return " + arg.(string)
	return s
}

func createPip() (txPip multipipes.Pipeline) {
	pipNodeSlice := make([]*multipipes.Node, 0)
	pipNodeSlice = append(pipNodeSlice, &multipipes.Node{Target: pip1, Name: "pip1", Capacity: 5})
	pipNodeSlice = append(pipNodeSlice, &multipipes.Node{Target: pip2, Name: "pip2", RoutineNum: 2})
	pipNodeSlice = append(pipNodeSlice, &multipipes.Node{Target: pip3, Name: "pip3", Timeout: 10})
	txPip = multipipes.Pipeline{
		Nodes: pipNodeSlice,
	}
	return txPip
}

func startPipCase() {
	txPip := createPip()
	indata := startProduceData()
	outData := startProcessData()
	txPip.Setup(indata, outData)
	//txPip.setup(indata, nil)
	txPip.Start()

	//waitRoutine := sync.WaitGroup{}
	//waitRoutine.Add(1)
	//waitRoutine.Wait()
}