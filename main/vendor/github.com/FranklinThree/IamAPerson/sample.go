package IamAPerson

import "errors"

type Sample struct {
	picture *[]byte
	choices [4]Choice
	theTrue int
}

func (sample *Sample) getTheTrue() (err error) {
	for i := 0; i < 4; i++ {
		if sample.choices[i].isRight {
			sample.theTrue = i
			return nil
		}

	}
	return errors.New("找不到正确选项")
}
