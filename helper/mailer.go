package helper

import (
	"fmt"
	"test/common"
	"test/config"
	"time"

	"gopkg.in/gomail.v2"
)

func SendEmail(EmailAddress string) ReturnType {
	m := gomail.NewMessage()
	mailConfig := config.GetMailConfig()
	html := "this is a test"
	m.SetAddressHeader("From", mailConfig["username"].(string), mailConfig["from"].(string))
	m.SetHeader("To", EmailAddress)
	m.SetHeader("Subject", "test")
	m.SetBody("text/html", html)
	send := gomail.NewDialer(mailConfig["host"].(string), mailConfig["port"].(int), mailConfig["username"].(string), mailConfig["password"].(string))
	err := send.DialAndSend(m)
	if err != nil {
		return ReturnType{Status: common.CodeError, Msg: "邮件发送失败", Data: err.Error()}
	} else {
		return ReturnType{Status: common.CodeSuccess, Msg: "邮件发送成功", Data: ""}
	}
}

func GetHTMLContent(MsgType int, WishContent string, MessageContent string) string {
	SendTime := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", time.Now().In(common.ChinaTime).Year(), time.Now().In(common.ChinaTime).Month(), time.Now().In(common.ChinaTime).Day(), time.Now().In(common.ChinaTime).Hour(), time.Now().In(common.ChinaTime).Minute(), time.Now().In(common.ChinaTime).Second())
	html := ""
	switch MsgType {
	case common.LightWish:
		html = fmt.Sprintf(`<div>
						<div>
							叮咚ding~
						</div>
						<div>
							<p> 属于你的小幸运已经被签收! </p>
						</div>
						<div>
							<p> 有人点亮你在小幸运活动中的愿望啦~ </p>
							<p> 快去看看Ta是谁吧! </p>
						</div>
						<div>
							<p> 此邮箱为系统邮箱，请勿回复。</p>
							<p> 发送于 %s </p>
						</div>
					<div>`, SendTime)
	case common.CancelLight:
		html = fmt.Sprintf(`<div>
						<div>
							叮咚ding~
						</div>
						<div>
							<p> Ta取消了对你 </p>
						</div>
						<div style="padding: 8px 40px 8px 50px;">
							<p>" %s "</p>
						</div>
						<div> 愿望的点亮，并留言: "%s" </div>
						<div>
							<p>该愿望重新被投入到愿望池</p>
							<p>耐心等待下一个有缘人点亮它吧</p>
							<p>惊喜总在不经意间~</p>
						</div>
						<div>
							<p> 此邮箱为系统邮箱，请勿回复。</p>
							<p> 发送于 %s </p>
						</div>
					<div>`, WishContent, MessageContent, SendTime)
	case common.HaveAchieve:
		html = fmt.Sprintf(`<div>
								<div>
									叮咚ding~
								</div>
								<div>
									<p> 你的愿望已经成功被%s同学实现了 </p>
								</div>
								<div>
									<p> 再去愿望池看看其它愿望吧~ </p>
								</div>
								<div>
									<p> 此邮箱为系统邮箱，请勿回复。</p>
									<p> 发送于 %s </p>
								</div>
							<div>`, WishContent, SendTime)
	case common.DeleteWish:
		html = fmt.Sprintf(`<div>
								<div>
									叮咚ding~
								</div>
								<div>
									<p> Ta删除了愿望 </p>
								</div>
								<div style="padding: 8px 40px 8px 50px;">
									<p> "%s" </p>
								</div>
								<div>
									<p> 再去愿望池看看其它愿望吧~ </p>
								</div>
								<div>
									<p> 此邮箱为系统邮箱，请勿回复。</p>
									<p> 发送于 %s </p>
								</div>
							<div>`, WishContent, SendTime)
	}
	return html
}
