package IamAPerson

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Path string
	Description string
	Map map[string]string
}

func NewConfig(path string,description string) (config Config, err error) {
	config.Map = make(map[string]string)
	reader,err := os.Open(path)
	if ! CheckErr(err){
		return Config{},errors.New("路径不正确！")
	}
	defer func(){
		err = reader.Close()
		CheckErr(err)
	}()
	BufferedReader := bufio.NewReader(reader)
	var tempBytes []byte
	lineCount := 0
	for{
		tempBytes,_,err = BufferedReader.ReadLine()
		if err != nil{
			if err == io.EOF{
				break
			}
		}
		lineCount++
		tempString := strings.TrimSpace(string(tempBytes))
		index := strings.Index(tempString,"=")
		if tempString == "" || tempString[0] == '#'{
			continue
		}
		if index == 0{
			return Config{},errors.New("第"+strconv.Itoa(lineCount)+"行没有找到key")
		}
		key := tempString[:index]
		var value string
		if index == len(tempString)-1{
			value = ""
		}else{
			value = tempString[index+1:]
		}
		key = strings.TrimSpace(key)
		key = strings.Trim(key,"\"")
		value = strings.TrimSpace(value)
		value = strings.Trim(value,"\"")
		config.Map[key] = value

	}
	config.Description = description
	config.Path = path
	return
}