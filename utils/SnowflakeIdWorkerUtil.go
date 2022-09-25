package utils

import (
	"errors"
	"time"
)

var (
	startTimestamp int64 = 1557489395327                        // 开始时间戳
	sequenceBit    int64 = 12                                   // 序列号占用位数
	machineInt     int64 = 10                                   // 机器标识所占位数
	timestampLeft        = sequenceBit + machineInt             // 时间戳位移位数
	maxSequence    int64 = -1 ^ (-1 << sequenceBit)             // 最大序列号
	maxMachineId   int64 = -1 ^ (-1 << machineInt)              // 最大机器编号
	machineIdPart  int64 = (9123 & maxMachineId) << sequenceBit // 生成id 机器标识部分
	sequence       int64 = 0                                    // 序列号
	lastStamp      int64 = -1
)

func NextId() (int64, error) {
	currentStamp := time.Now().UnixNano() / 1e6 // 当前是时间戳（毫秒）
	// 当前时间小于最后生成的时间
	if currentStamp < lastStamp {
		err := errors.New("时钟已经回拨")
		return 0, err
	}
	// 当前时间等于最后生成时间，阻塞获取下一毫秒
	if currentStamp == lastStamp {
		sequence = (sequence + 1) & maxSequence
		currentStamp = getNextMill()
	} else {
		sequence = 0
	}
	// 修改最后生成时间
	lastStamp = currentStamp

	return (currentStamp-startTimestamp)<<timestampLeft | machineIdPart | sequence, nil
}

func getNextMill() int64 {
	mill := time.Now().UnixNano() / 1e6
	for mill <= lastStamp {
		mill = time.Now().UnixNano() / 1e6
	}
	return mill
}
