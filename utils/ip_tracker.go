package utils

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"sync"
	"time"
)

// IPTracker 用于跟踪IP地址的访问次数和黑名单状态
type IPTracker struct {
	mu         sync.Mutex
	ipCountMap map[string]int  // 存储IP地址的访问次数
	ipLastTime map[string]int  // 存储访问的时间
	blacklist  map[string]bool // 存储被加入黑名单的IP地址
}

// NewIPTracker 创建一个新的IPTracker实例
func NewIPTracker() *IPTracker {
	return &IPTracker{
		ipCountMap: make(map[string]int),
		ipLastTime: make(map[string]int),
		blacklist:  make(map[string]bool),
	}
}

// IsBlacklisted 检查给定的IP地址是否在黑名单中
func (tracker *IPTracker) IsBlacklisted(ip string) bool {
	return tracker.blacklist[ip]
}

// AddToBlacklist 将给定的IP地址添加到黑名单
func (tracker *IPTracker) AddToBlacklist(ip string) {
	tracker.blacklist[ip] = true
}

// ResetRecord 重置记录
func (tracker *IPTracker) ResetRecord(ip string) {
	delete(tracker.blacklist, ip)
	delete(tracker.ipLastTime, ip)
	delete(tracker.ipCountMap, ip)
}

// TrackIP 跟踪给定IP地址的访问次数，并在需要时加入黑名单
func (tracker *IPTracker) TrackIP(ip string) bool {
	tracker.mu.Lock()
	defer tracker.mu.Unlock()

	// 如果IP已经被跟踪过，检查访问次数是否超过20次
	ipLastTime := tracker.ipLastTime[ip]
	now := time.Now()
	// 检查IP是否在黑名单中
	if tracker.IsBlacklisted(ip) {
		// 超过了30分钟就把这个IP放出来
		if now.Sub(time.Unix(int64(ipLastTime), 0)) > time.Minute*30 {
			tracker.ResetRecord(ip)
			return true
		}
		return false
	}

	if count, ok := tracker.ipCountMap[ip]; ok {
		fmt.Println(count, "tracker.ResetRecord(ip)")
		// 正常请求的IP就不做记录，避免IP拥挤浪费内存
		if now.Sub(time.Unix(int64(ipLastTime), 0)) > time.Minute {
			tracker.ResetRecord(ip)
			return true
		}
		// 如果在一分钟内访问次数超过60次，加入黑名单
		if now.Sub(time.Unix(int64(ipLastTime), 0)) < time.Minute && count >= 60 {
			logs.Error(fmt.Sprintf("IP %s exceeded 60 requests. Adding to blacklist.", ip))
			tracker.AddToBlacklist(ip)
		}
		// 更新IP的访问时间
		tracker.ipCountMap[ip] = count + 1
		tracker.ipLastTime[ip] = int(now.Unix())
	} else {
		// 如果是第一次访问，添加到计数映射中
		tracker.ipCountMap[ip] = 1
		tracker.ipLastTime[ip] = int(now.Unix())
	}
	return true
}

// RunTracker 防止同一个IP频繁访问
//func RunTracker() {
//	// 模拟IP请求
//	ips := []string{"192.168.0.1", "192.168.0.2", "192.168.0.1", "192.168.0.3", "192.168.0.1"}
//	for _, ip := range ips {
//		if !tracker.TrackIP(ip) {
//			fmt.Printf("IP %s is blacklisted. Access denied.\n", ip)
//		} else {
//			fmt.Printf("Access granted for IP %s.\n", ip)
//		}
//	}
//}
