package boot

import (
	"github.com/yyxing/boot/context"
	"github.com/yyxing/boot/starter"
)

type Application struct {
	context context.ApplicationContext
	isInit  bool
}

// 获取context配置
func (application *Application) Get(key string) interface{} {
	return application.context.Get(key)
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
	application.Run()
	return application
}

// 默认配置启动 config log sql等
func Default() Application {
	application := Application{context: context.ApplicationContext{}}
	application.context.Register(&starter.ConfigStarter{})
	application.context.Register(&starter.DatasourceStarter{})
	application.context.Register(&starter.LogStarter{})
	application.context.Register(&starter.ValidatorStarter{})
	return application
}

// 启动所有starter
func (application *Application) Run() {
	application.context.SortStarter()
	if !application.isInit {
		application.Init()
	}
	for _, starter := range application.context.GetAllStarters() {
		// 调用每个starter的start方法
		starter.Start(application.context)
	}
}

// 初始化starter
func (application *Application) Init() {
	application.context.SortStarter()
	for _, starter := range application.context.GetAllStarters() {
		// 调用每个starter的Init方法
		starter.Init(application.context)
	}
	application.isInit = true
}

func (application *Application) Stop() {
	// 停止所有starter
	for _, starter := range application.context.GetAllStarters() {
		starter.Finalize(application.context)
	}
}
