package logs

import (
	"fmt"
	"github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface

func DisableLog() {
	Logger = seelog.Disabled
}

func loadAppConfig() {
	appConfig := `<seelog minlevel="warn">
    <outputs formatid="common">
        <rollingfile type="size" filename="./runtime/logs/roll.log" maxsize="100000" maxrolls="5"/>
        <filter levels="critical">
            <file path="./runtime/logs/critical.log" formatid="critical"/>
            <smtp formatid="criticalemail" senderaddress="astaxie@gmail.com" sendername="ShortUrl API" hostname="smtp.gmail.com" hostport="587" username="mailusername" password="mailpassword">
                <recipient address="895560756@qq.com"/>
            </smtp>
        </filter>
    </outputs>
    <formats>
        <format id="common" format="%Date/%Time [%LEV] %Msg%n" />
        <format id="critical" format="%File %FullPath %Func %Msg%n" />
        <format id="criticalemail" format="Critical error on our server!\n    %Time %Date %RelFile %Func %Msg \nSent by Seelog"/>
    </formats>
</seelog>`

	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Printf("The LoggerFromConfigAsBytes Error: %v\n", err)
		return
	}
	UseLogger(logger)
}

func UseLogger(logger seelog.LoggerInterface) {
	Logger = logger
}

func init() {
	// 初始化全局变量 Logger 为 seelog 的禁用状态，主要为了防止 Logger 被多次初始化
	DisableLog()
	// 根据配置文件初始化 seelog 的配置信息，这里我们把配置文件通过字符串读取设置好了
	loadAppConfig()
}
