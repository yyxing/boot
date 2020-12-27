package boot

import (
	"github/yyxing/boot/context"
	"github/yyxing/boot/starter"
)

type Application struct {
	context context.ApplicationContext
}

// TODO 由用户指定starter启动 暂未实现
func New() {

}
func (application *Application) RegisterStarter(s context.Starter) Application {
	application.context.Register(s)
	return *application
}

// 默认配置启动 config log sql等
func TestEnv(configPath string) Application {
	application := Application{context: context.ApplicationContext{}}
	application.context.Register(&starter.ConfigStarter{ConfigPath: configPath})
	application.context.Register(&starter.DatasourceStarter{})
	application.context.Register(&starter.LogStarter{})
	application.context.Register(&starter.ValidatorStarter{})
	application.context.SortStarter()
	application.run()
	return application
}

// 默认配置启动 config log sql等
func Default() Application {
	application := Application{context: context.ApplicationContext{}}
	application.context.Register(&starter.ConfigStarter{})
	application.context.Register(&starter.DatasourceStarter{})
	application.context.Register(&starter.LogStarter{})
	application.context.Register(&starter.ValidatorStarter{})
	application.context.SortStarter()
	application.run()
	return application
}

func (application *Application) run() {
	application.init()
	application.start()
}

// 初始化starter
func (application *Application) init() {
	for _, starter := range application.context.GetAllStarters() {
		// 调用每个starter的Init方法
		starter.Init(application.context)
	}
}

// 启动所有starter
func (application *Application) start() {
	for _, starter := range application.context.GetAllStarters() {
		// 调用每个starter的start方法
		starter.Start(application.context)
	}
}
func (application *Application) Stop() {
	// 停止所有starter
	for _, starter := range application.context.GetAllStarters() {
		starter.Finalize(application.context)
	}
}
