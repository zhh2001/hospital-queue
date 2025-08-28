package service

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"hospital-queue/models"
)

var (
	dataFile = filepath.Join("data", "queue.json")
	mutex    sync.Mutex
)

// 初始化数据文件
func init() {
	// 创建data目录（如果不存在）
	if err := os.MkdirAll("data", 0755); err != nil {
		log.Fatalf("创建数据目录失败: %v\n", err)
	}

	// 检查文件是否存在，不存在则创建空文件
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		file, err := os.Create(dataFile)
		if err != nil {
			log.Fatalf("创建数据文件失败: %v\n", err)
		}

		defer func() {
			err = file.Close()
			if err != nil {
				return
			}
		}()

		// 写入空数组
		err = json.NewEncoder(file).Encode([]models.Patient{})
		if err != nil {
			log.Fatalf("创建数据文件失败: %v\n", err)
		}
	}
}

// 读取所有叫号数据
func readAllQueues() ([]models.Patient, error) {
	mutex.Lock()
	defer mutex.Unlock()

	data, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}

	var queues []models.Patient
	if err = json.Unmarshal(data, &queues); err != nil {
		return nil, err
	}

	return queues, nil
}

// 保存叫号数据到文件
func saveQueues(queues []models.Patient) error {
	mutex.Lock()
	defer mutex.Unlock()

	data, err := json.MarshalIndent(queues, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, data, 0644)
}

// CreateNewQueue 新增排队号码
func CreateNewQueue(name string, phone string, department uint) (*models.Patient, error) {
	queues, err := readAllQueues()
	if err != nil {
		return nil, err
	}

	// 计算当前最大ID
	var maxID uint = 0
	for _, q := range queues {
		if q.ID > maxID {
			maxID = q.ID
		}
	}

	// 计算当前最大号码
	var maxNumber uint = 0
	for _, q := range queues {
		if q.Department == department && q.Number > maxNumber {
			maxNumber = q.Number
		}
	}

	// 创建新号码
	newQueue := models.Patient{
		ID:         maxID + 1,
		Number:     maxNumber + 1,
		Name:       name,
		Phone:      phone,
		Department: department,
		Status:     0,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	}

	// 添加到队列并保存
	queues = append(queues, newQueue)
	if err := saveQueues(queues); err != nil {
		return nil, err
	}

	return &newQueue, nil
}

// CallQueue 叫号
func CallQueue(id uint) (*models.Patient, error) {
	queues, err := readAllQueues()
	if err != nil {
		return nil, err
	}

	// 查找被叫号的数据记录
	var patient *models.Patient
	for i, q := range queues {
		if q.ID == id {
			// 更新状态为已叫号
			queues[i].Status = 1
			queues[i].UpdateAt = time.Now()
			patient = &q
		}
	}

	if patient == nil {
		return nil, errors.New("数据不存在")
	}

	if err := CallVoice(patient.Name, patient.Department); err != nil {
		return nil, err
	}

	if err := saveQueues(queues); err != nil {
		return nil, err
	}

	return patient, nil
}
