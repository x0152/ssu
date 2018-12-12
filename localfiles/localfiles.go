package localfiles

import (
	"../commands"
	"../consts"
	"../untils"
	"fmt"
	"io/ioutil"
)

func SaveLocalFiles(res commands.Results) {
	for _, d := range res {

		if d.Type == commands.TYPE_RESULT_FILE ||
			d.Type == commands.TYPE_RESULT_IMAGE {

			SaveFile(d.Description, d.Data)
		}
	}
}

func SaveFile(name string, data []byte) {
	untils.WriteMsgLog(fmt.Sprintf("save file %s...", name))
	err := ioutil.WriteFile(fmt.Sprintf("%s/%s", consts.DIR_TMP_FILES, name), data, 0644)

	if err != nil {
		untils.WriteMsgLogError(err)
	}
}
