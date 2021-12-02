package util

import (
	"errors"
	"github.com/fire-g/mark-go-util/logger"
	"sync"
	"time"
)

/*
 * 算法解释
 * SnowFlake的结构如下(每部分用-分开):<br>
 * 0 - 0000000000 0000000000 0000000000 0000000000 0 - 00000 - 00000 - 000000000000 <br>
 * 1位标识，由于long基本类型在Java中是带符号的，最高位是符号位，正数是0，负数是1，所以id一般是正数，最高位是0<br>
 * 41位时间截(毫秒级)，注意，41位时间截不是存储当前时间的时间截，而是存储时间截的差值（当前时间截 - 开始时间截)
 * 得到的值），这里的的开始时间截，一般是我们的id生成器开始使用的时间，由我们程序来指定的（如下的epoch属性）。
 * 41位的时间截，可以使用69年，年T = (1L << 41) / (1000L * 60 * 60 * 24 * 365) = 69<br>
 * 10位的数据机器位，可以部署在1024个节点，包括5位dataCenterId和5位workerId<br>
 * 12位序列，毫秒内的计数，12位的计数顺序号支持每个节点每毫秒(同一机器，同一时间截)产生4096个ID序号<br>
 * 加起来刚好64位，为一个Long型。<br>
 * SnowFlake的优点是，整体上按照时间自增排序，并且整个分布式系统内不会产生ID碰撞(由数据中心ID和机器ID作区分)，并且效率较高，经测试，SnowFlake每秒能够产生26万ID左右。
 */
const (
	//开始时间戳 2015-1-1
	epoch int64 = 1420041600000
	//机器id所占的位数
	workerIdBits int64 = 6
	//数据标识id所占的位数
	dataCenterIdBits int64 = 6
	//支持的最大机器id，结果是31(这个移位算法可以很快的计算出几位二进制数所能表示的最大十进制数)
	maxWorkerId     int64 = -1 ^ (-1 << workerIdBits) // 支持的最大数据标识id，结果是31
	maxDataCenterId int64 = -1 ^ (-1 << dataCenterIdBits)
	//序列在id中占的位数
	sequenceBits int64 = 9
	// 机器ID向左移12位
	workerIdShift = sequenceBits
	// 数据标识id向左移17位(12+5)
	dataCenterIdShift = sequenceBits + workerIdBits
	// 时间截向左移22位(5+5+12)
	timestampLeftShift = sequenceBits + workerIdBits + dataCenterIdBits
	// 生成序列的掩码，这里为4095 (0b111111111111=0xfff=4095)
	sequenceMask int64 = -1 ^ (-1 << sequenceBits)
)

var (
	IdWorker *SnowflakeIdWorker
	err      error
)

func init() {
	IdWorker, err = createWorker(0, 0)
	if err != nil {
		logger.Logger.Error.Println(err)
		return
	}
}

// SnowflakeIdWorker 构造
type SnowflakeIdWorker struct {
	mutex         sync.Mutex //添加互斥锁，确保并发安全
	lastTimestamp int64      //上次生成ID的时间戳
	workerId      int64      //工作机器ID(0-31)
	dataCenterId  int64      // 数据中心ID(0-31)
	sequence      int64      //毫秒内序列(0-4095)
}

// createWorker 创建SnowflakeIdWorker
// workerId 工作ID (0~31)
// dataCenterId 数据中心ID (0~31)
func createWorker(wId int64, dId int64) (*SnowflakeIdWorker, error) {
	if wId < 0 || wId > maxWorkerId {
		return nil, errors.New("Worker ID excess of quantity")
	}
	if dId < 0 || dId > maxDataCenterId {
		return nil, errors.New("Datacenter ID excess of quantity")
	}
	// 生成一个新节点
	return &SnowflakeIdWorker{
		lastTimestamp: 0,
		workerId:      wId,
		dataCenterId:  dId,
		sequence:      0,
	}, nil
}

// NextId 获取ID
func (w *SnowflakeIdWorker) NextId() int64 {
	// 保障线程安全 加锁
	w.mutex.Lock()
	// 生成完成后 解锁
	defer w.mutex.Unlock()
	// 获取生成时的时间戳 毫秒
	now := time.Now().UnixNano() / 1e6
	//如果当前时间小于上一次ID生成的时间戳，说明系统时钟回退过这个时候应当抛出异常
	if now < w.lastTimestamp {
		_ = errors.New("Clock moved backwards")
		//根据需要自定义错误码
		return 3001
	}
	if w.lastTimestamp == now {
		w.sequence = (w.sequence + 1) & sequenceMask
		if w.sequence == 0 {
			// 阻塞到下一个毫秒，直到获得新的时间戳
			for now <= w.lastTimestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// 当前时间与工作节点上一次生成ID的时间不一致 则需要重置工作节点生成ID的序号
		w.sequence = 0
	}
	// 将机器上一次生成ID的时间更新为当前时间
	w.lastTimestamp = now
	ID := int64((now-epoch)<<timestampLeftShift | w.dataCenterId<<dataCenterIdShift | (w.workerId << workerIdShift) | w.sequence)
	return ID
}
