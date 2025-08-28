package service

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// TextToSpeech 调用系统原生TTS播放自定义文本
// text: 要转换的文本
// lang: 语言（zh-CN 中文，en-US 英文）
func TextToSpeech(text, lang string) error {
	if text == "" {
		return fmt.Errorf("文本不能为空")
	}

	// 根据操作系统选择不同的TTS命令
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		var voiceName string
		switch lang {
		case "zh-CN", "zh":
			voiceName = "Microsoft Huihui Desktop" // 中文语音
		case "en-US", "en":
			voiceName = "Microsoft David Desktop" // 英文语音
		default:
			voiceName = "Microsoft David Desktop" // 默认语音
		}

		// 使用SelectVoice方法而非SelectVoiceByHints，更简单可靠
		psCmd := fmt.Sprintf(`
			Add-Type -AssemblyName System.Speech;
			$speak = New-Object System.Speech.Synthesis.SpeechSynthesizer;
			$speak.SelectVoice('%s');
			$speak.Speak('%s')
		`, voiceName, escapeWindowsPS(text))
		cmd = exec.Command("powershell", "-NoProfile", "-Command", psCmd)

	case "darwin": // macOS
		// macOS：使用say命令（系统自带，支持多语言）
		voice := getMacVoiceByLang(lang)
		cmd = exec.Command("say", "-v", voice, text)

	case "linux":
		// Linux：使用espeak
		espeakPath, err := exec.LookPath("espeak")
		if err != nil {
			return fmt.Errorf("linux需安装espeak：sudo apt install espeak: %v", err)
		}
		cmd = exec.Command(espeakPath, "-v", lang, "-s", "150", text)

	default:
		return fmt.Errorf("不支持的操作系统：%s", runtime.GOOS)
	}

	// 执行命令（阻塞直到播放完成）
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("播放失败：%w", err)
	}
	return nil
}

// 处理Windows PowerShell文本中的特殊字符
func escapeWindowsPS(text string) string {
	return strings.ReplaceAll(text, "'", "''")
}

// 根据语言获取macOS对应的语音
func getMacVoiceByLang(lang string) string {
	switch lang {
	case "zh-CN", "zh":
		return "Ting-Ting" // 中文语音
	case "en-US", "en":
		return "Alex" // 英文语音
	case "ja-JP", "ja":
		return "Kyoko" // 日语语音
	default:
		return "Alex"
	}
}

// CallVoice 叫号语音
func CallVoice(name string, department uint) error {
	text := fmt.Sprintf("请 %s 到 %d 诊室 就诊。", name, department)

	// 调用TTS播放（中文）
	err := TextToSpeech(text, "zh-CN")
	if err != nil {
		log.Printf("文本转语音失败：%v\n", err)
		return err
	}
	log.Println("语音播放完成！")
	return nil
}
