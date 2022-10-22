package service

import (
	"fmt"
	"go-psmp/utils/redis"
	"time"
)

const (
	Agent                  = "agent:agents"
	AgentHeartbeat         = "agent:heartbeat:"
	AgentHeartbeatAlarmKey = "agent:heartbeat_alarm:"
	frequency              = 600
)

var emailService = ServiceGroup.EmailService

func AgentHeartbeatAlarm() {

	agentMap := redis.HGetAll(Agent)

	if agentMap == nil {
		return
	}
	for agent := range agentMap {
		fmt.Println(agent, "agent", agentMap[agent])

		// 心跳缓存
		cacheHeartbeat := redis.Exists(AgentHeartbeat + agent)

		// 告警缓存
		cacheAlarm := redis.Exists(AgentHeartbeatAlarmKey + agent)

		// 存在心跳，agent客户端健康存在，则进入下一循环
		if cacheHeartbeat {
			// 有告警缓存，则表明之前有告警，现已恢复，则发送恢复通知
			if cacheAlarm {

				// 删除缓存
				redis.Delete(AgentHeartbeatAlarmKey + agent)

				content := " 时间:%s \n" + "检测到" + agent + ":已经恢复正常"

				var t = time.Now().Format("2006-01-02 15:04:05") + ""
				content = fmt.Sprintf(content, t)
				// 发送邮件

				emailService.saveEmailRecord("告警平台", "guozhifengvip@163.com", "", "【Agent恢复】", content, "HTML")

			}

			continue
		}

		// frequency分钟内不重复告警
		if !cacheAlarm {

			// 存入缓存，frequency分钟内不重复发告警
			redis.SetEx(AgentHeartbeatAlarmKey+agent, 1, frequency)

			var t = time.Now().Format("2006-01-02 15:04:05") + ""
			// 发送邮件
			content := " 时间:%s\n  检测到:" + agent + ":监控agent服务不可用，请尽快处理"
			content = fmt.Sprintf(content, t)
			emailService.saveEmailRecord("告警平台", "guozhifengvip@163.com", "", "【Agent异常】", content, "HTML")

		}
	}
}

func Heartbeat(agentIp, agentName string) {
	if agentIp == "" || agentName == "" {
		return
	}

	redis.SetEx(AgentHeartbeat+getAgentStr(agentIp, agentName), 1, 185)
	redis.HSet(Agent, getAgentStr(agentIp, agentName), 0)
}

func getAgentStr(agentIp, agentName string) string {
	if agentIp == "" || agentName == "" {
		return ""
	}

	return agentIp + "|" + agentName
}
