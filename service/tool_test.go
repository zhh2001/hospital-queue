package service

import "testing"

func TestTextToSpeech(t *testing.T) {
	err := TextToSpeech("请张三到2诊室就诊。", "zh")
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Log("success")

	err = TextToSpeech("请张三到2诊室就诊。", "zh-CN")
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Log("success")

	err = TextToSpeech("Please ask Zhang San to go to Clinic 2 for treatment.", "en-US")
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Log("success")

	err = TextToSpeech("Please ask Zhang San to go to Clinic 2 for treatment.", "en")
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Log("success")

	err = TextToSpeech("張三に診療所2に行って治療を受けるよう伝えてください。", "ja-JP")
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Log("success")

	err = TextToSpeech("張三に診療所2に行って治療を受けるよう伝えてください。", "ja")
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Log("success")
}

func TestCallVoice(t *testing.T) {
	err := CallVoice("张三", 2)
	if err != nil {
		t.Errorf("call voice fail: %v", err)
	}
	t.Logf("call voice success")
}
