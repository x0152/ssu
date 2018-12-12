package commands

type Action int

//type Arguments map[string]string
type Results []Data

const (
	TYPE_RESULT_TEXT = iota
	TYPE_RESULT_IMAGE
	TYPE_RESULT_FILE
	TYPE_RESULT_LINK
)

const (
	CMD_GET_FUNCTIONS = iota
	CMD_IS_LIVE
)

type Data struct {
	Type        int
	Data        []byte
	Description string
}

type CommandServer struct {
	Act Action
	Res Results
	//Args Arguments
}

type ResultCommandClient struct {
	Act Action
	Res Results
}
