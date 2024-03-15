package sundry

import (
	"github.com/gin-gonic/gin"

	"github.com/opentdp/wechat-rest/wclient/robot"
)

// @Summary 机器人指令集
// @Tags BOT::杂项
// @Param body body RobotHandlersParam true "机器人指令集参数"
// @Success 200 {array} RobotHandler
// @Router /bot/handlers [post]
func handlerList(c *gin.Context) {

	var rq *RobotHandlersParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Reset {
		robot.Reset()
	}

	items := []RobotHandler{}

	for _, v := range robot.GetHandlers() {
		items = append(items, RobotHandler{
			Level:    v.Level,
			Order:    v.Order,
			Roomid:   v.Roomid,
			Command:  v.Command,
			Describe: v.Describe,
		})
	}

	c.Set("Payload", items)

}

type RobotHandler struct {
	Level    int32  `json:"level"`    // 0:不限制 7:群管理 9:创始人
	Order    int32  `json:"order"`    // 排序，越小越靠前
	Roomid   string `json:"roomid"`   // 使用场景 [*:所有,-:私聊,+:群聊,其他:群聊]
	Command  string `json:"command"`  // 指令
	Describe string `json:"describe"` // 指令的描述信息
}

type RobotHandlersParam struct {
	Reset bool `json:"reset"` // 是否重置机器人指令
}