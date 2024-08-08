package mail

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/log"
	"github.com/goravel/framework/contracts/mail"
	queuecontract "github.com/goravel/framework/contracts/queue"
	configmock "github.com/goravel/framework/mocks/config"
	logmock "github.com/goravel/framework/mocks/log"
	"github.com/goravel/framework/queue"
	"github.com/goravel/framework/support/carbon"
	"github.com/goravel/framework/support/color"
	testingdocker "github.com/goravel/framework/support/docker"
	"github.com/goravel/framework/support/env"
	"github.com/goravel/framework/support/file"
)

var testBcc, testCc, testTo, testFromAddress, testFromName string

type ApplicationTestSuite struct {
	suite.Suite
	redisPort int
}

func TestApplicationTestSuite(t *testing.T) {
	if env.IsWindows() {
		t.Skip("Skipping tests of using docker")
	}

	if !file.Exists("../.env") && os.Getenv("MAIL_HOST") == "" {
		color.Red().Println("No mail tests run, need create .env based on .env.example, then initialize it")
		return
	}

	redisDocker := testingdocker.NewRedis()
	assert.Nil(t, redisDocker.Build())

	suite.Run(t, &ApplicationTestSuite{
		redisPort: redisDocker.Config().Port,
	})

	assert.Nil(t, redisDocker.Stop())
}

func (s *ApplicationTestSuite) SetupTest() {}

func (s *ApplicationTestSuite) TestSendMailBy465Port() {
	mockConfig := mockConfig(465, s.redisPort)
	app := NewApplication(mockConfig, nil)
	s.Nil(app.To([]string{testTo}).
		Cc([]string{testCc}).
		Bcc([]string{testBcc}).
		Attach([]string{"../logo.png"}).
		Subject("Goravel Test 465").
		Content(Html("<h1>Hello Goravel</h1>")).
		Send())
}

func (s *ApplicationTestSuite) TestSendMailBy587Port() {
	mockConfig := mockConfig(587, s.redisPort)
	app := NewApplication(mockConfig, nil)
	s.Nil(app.To([]string{testTo}).
		Cc([]string{testCc}).
		Bcc([]string{testBcc}).
		Attach([]string{"../logo.png"}).
		Subject("Goravel Test 587").
		Content(Html("<h1>Hello Goravel</h1>")).
		Send())
}

func (s *ApplicationTestSuite) TestSendMailWithFrom() {
	mockConfig := mockConfig(587, s.redisPort)
	app := NewApplication(mockConfig, nil)
	s.Nil(app.From(Address(testFromAddress, testFromName)).
		To([]string{testTo}).
		Cc([]string{testCc}).
		Bcc([]string{testBcc}).
		Attach([]string{"../logo.png"}).
		Subject("Goravel Test 587 With From").
		Content(Html("<h1>Hello Goravel</h1>")).
		Send())
}

func (s *ApplicationTestSuite) TestSendMailWithMailable() {
	mockConfig := mockConfig(587, s.redisPort)
	app := NewApplication(mockConfig, nil)
	s.Nil(app.Send(NewTestMailable()))
}

func (s *ApplicationTestSuite) TestQueueMail() {
	mockConfig := mockConfig(587, s.redisPort)
	mockLog := &logmock.Log{}

	queueFacade := queue.NewApplication(mockConfig, mockLog)
	queueFacade.Register([]queuecontract.Job{
		NewSendMailJob(mockConfig),
	})

	app := NewApplication(mockConfig, queueFacade)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		s.Nil(queueFacade.Worker(nil).Run())

		for range ctx.Done() {
			return
		}
	}(ctx)
	time.Sleep(3 * time.Second)
	s.Nil(app.To([]string{testTo}).
		Cc([]string{testCc}).
		Bcc([]string{testBcc}).
		Attach([]string{"../logo.png"}).
		Subject("Goravel Test Queue").
		Content(Html("<h1>Hello Goravel</h1>")).
		Queue())
	time.Sleep(3 * time.Second)
}

func (s *ApplicationTestSuite) TestQueueMailWithConnection() {
	mockConfig := mockConfig(587, s.redisPort)
	mockLog := NewTestLog()

	queueFacade := queue.NewApplication(mockConfig, mockLog)
	queueFacade.Register([]queuecontract.Job{
		NewSendMailJob(mockConfig),
	})

	app := NewApplication(mockConfig, queueFacade)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		s.Nil(queueFacade.Worker(&queuecontract.Args{
			Connection: "redis",
			Queue:      "test",
		}).Run())

		for range ctx.Done() {
			return
		}
	}(ctx)
	time.Sleep(3 * time.Second)
	s.Nil(app.To([]string{testTo}).
		Cc([]string{testCc}).
		Bcc([]string{testBcc}).
		Attach([]string{"../logo.png"}).
		Subject("Goravel Test Queue with connection").
		Content(Html("<h1>Hello Goravel</h1>")).
		Queue(Query().OnConnection("redis").OnQueue("test")))
	time.Sleep(3 * time.Second)
}

func (s *ApplicationTestSuite) TestQueueMailWithMailable() {
	mockConfig := mockConfig(587, s.redisPort)
	mockLog := &logmock.Log{}

	queueFacade := queue.NewApplication(mockConfig, mockLog)
	queueFacade.Register([]queuecontract.Job{
		NewSendMailJob(mockConfig),
	})

	app := NewApplication(mockConfig, queueFacade)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		s.Nil(queueFacade.Worker(nil).Run())

		for range ctx.Done() {
			return
		}
	}(ctx)
	time.Sleep(3 * time.Second)
	s.Nil(app.Queue(NewTestMailable()))
	time.Sleep(3 * time.Second)
}

func mockConfig(mailPort, redisPort int) *configmock.Config {
	mockConfig := &configmock.Config{}
	mockConfig.On("GetString", "app.name").Return("goravel")
	mockConfig.On("GetBool", "app.debug").Return(true)
	mockConfig.On("GetString", "queue.default").Return("redis")
	mockConfig.On("GetString", "queue.connections.sync.driver").Return("sync")
	mockConfig.On("GetString", "queue.connections.redis.driver").Return("redis")
	mockConfig.On("GetString", "queue.connections.redis.connection").Return("default")
	mockConfig.On("GetString", "queue.connections.redis.queue", "default").Return("default")
	mockConfig.On("GetString", "database.redis.default.host").Return("localhost")
	mockConfig.On("GetString", "database.redis.default.password").Return("")
	mockConfig.On("GetInt", "database.redis.default.port").Return(redisPort)
	mockConfig.On("GetInt", "database.redis.default.database").Return(0)

	if file.Exists("../.env") {
		vip := viper.New()
		vip.SetConfigName("../.env")
		vip.SetConfigType("env")
		vip.AddConfigPath(".")
		_ = vip.ReadInConfig()
		vip.SetEnvPrefix("goravel")
		vip.AutomaticEnv()

		mockConfig.On("GetString", "mail.host").Return(vip.Get("MAIL_HOST"))
		mockConfig.On("GetInt", "mail.port").Return(mailPort)
		mockConfig.On("GetString", "mail.from.address").Return(vip.Get("MAIL_FROM_ADDRESS"))
		mockConfig.On("GetString", "mail.from.name").Return(vip.Get("MAIL_FROM_NAME"))
		mockConfig.On("GetString", "mail.username").Return(vip.Get("MAIL_USERNAME"))
		mockConfig.On("GetString", "mail.password").Return(vip.Get("MAIL_PASSWORD"))
		mockConfig.On("GetString", "mail.to").Return(vip.Get("MAIL_TO"))
		mockConfig.On("GetString", "mail.cc").Return(vip.Get("MAIL_CC"))
		mockConfig.On("GetString", "mail.bcc").Return(vip.Get("MAIL_BCC"))

		testFromAddress = vip.Get("MAIL_FROM_ADDRESS").(string)
		testFromName = vip.Get("MAIL_FROM_NAME").(string)
		testTo = vip.Get("MAIL_TO").(string)
	}
	if os.Getenv("MAIL_HOST") != "" {
		mockConfig.On("GetString", "mail.host").Return(os.Getenv("MAIL_HOST"))
		mockConfig.On("GetInt", "mail.port").Return(mailPort)
		mockConfig.On("GetString", "mail.from.address").Return(os.Getenv("MAIL_FROM_ADDRESS"))
		mockConfig.On("GetString", "mail.from.name").Return(os.Getenv("MAIL_FROM_NAME"))
		mockConfig.On("GetString", "mail.username").Return(os.Getenv("MAIL_USERNAME"))
		mockConfig.On("GetString", "mail.password").Return(os.Getenv("MAIL_PASSWORD"))
		mockConfig.On("GetString", "mail.to").Return(os.Getenv("MAIL_TO"))
		mockConfig.On("GetString", "mail.cc").Return(os.Getenv("MAIL_CC"))
		mockConfig.On("GetString", "mail.bcc").Return(os.Getenv("MAIL_BCC"))

		testFromAddress = os.Getenv("MAIL_FROM_ADDRESS")
		testFromName = os.Getenv("MAIL_FROM_NAME")
		testBcc = os.Getenv("MAIL_BCC")
		testCc = os.Getenv("MAIL_CC")
		testTo = os.Getenv("MAIL_TO")
	}

	return mockConfig
}

type TestMailable struct {
}

func NewTestMailable() *TestMailable {
	return &TestMailable{}
}

func (m *TestMailable) Attachments() []string {
	return []string{"../logo.png"}
}

func (m *TestMailable) Content() *mail.Content {
	html := Html("<h1>Hello Goravel</h1>")

	return &html
}

func (m *TestMailable) Envelope() *mail.Envelope {
	return &mail.Envelope{
		Bcc:     []string{testBcc},
		Cc:      []string{testCc},
		From:    Address(testFromAddress, testFromName),
		Subject: "Goravel Test 587 With Mailable",
		To:      []string{testTo},
	}
}

func (m *TestMailable) Queue() *mail.Queue {
	return &mail.Queue{}
}

type TestLog struct {
	*TestLogWriter
}

func NewTestLog() log.Log {
	return &TestLog{
		TestLogWriter: NewTestLogWriter(),
	}
}

func (r *TestLog) WithContext(ctx context.Context) log.Writer {
	return NewTestLogWriter()
}

func (r *TestLog) Channel(channel string) log.Writer {
	return NewTestLogWriter()
}

func (r *TestLog) Stack(channels []string) log.Writer {
	return NewTestLogWriter()
}

type TestLogWriter struct {
	data map[string]any
}

func NewTestLogWriter() *TestLogWriter {
	return &TestLogWriter{
		data: make(map[string]any),
	}
}

func (r *TestLogWriter) Debug(args ...any) {
	fmt.Print(prefix("debug"))
	fmt.Println(args...)
	r.printData()
}

func (r *TestLogWriter) Debugf(format string, args ...any) {
	fmt.Print(prefix("debug"))
	fmt.Printf(format+"\n", args...)
	r.printData()
}

func (r *TestLogWriter) Info(args ...any) {
	fmt.Print(prefix("info"))
	fmt.Println(args...)
	r.printData()
}

func (r *TestLogWriter) Infof(format string, args ...any) {
	fmt.Print(prefix("info"))
	fmt.Printf(format+"\n", args...)
	r.printData()
}

func (r *TestLogWriter) Warning(args ...any) {
	fmt.Print(prefix("warning"))
	fmt.Println(args...)
	r.printData()
}

func (r *TestLogWriter) Warningf(format string, args ...any) {
	fmt.Print(prefix("warning"))
	fmt.Printf(format+"\n", args...)
	r.printData()
}

func (r *TestLogWriter) Error(args ...any) {
	fmt.Print(prefix("error"))
	fmt.Println(args...)
	r.printData()
}

func (r *TestLogWriter) Errorf(format string, args ...any) {
	fmt.Print(prefix("error"))
	fmt.Printf(format+"\n", args...)
	r.printData()
}

func (r *TestLogWriter) Fatal(args ...any) {
	fmt.Print(prefix("fatal"))
	fmt.Println(args...)
	r.printData()
}

func (r *TestLogWriter) Fatalf(format string, args ...any) {
	fmt.Print(prefix("fatal"))
	fmt.Printf(format+"\n", args...)
	r.printData()
}

func (r *TestLogWriter) Panic(args ...any) {
	fmt.Print(prefix("panic"))
	fmt.Println(args...)
	r.printData()
}

func (r *TestLogWriter) Panicf(format string, args ...any) {
	fmt.Print(prefix("panic"))
	fmt.Printf(format+"\n", args...)
	r.printData()
}

func (r *TestLogWriter) User(user any) log.Writer {
	r.data["user"] = user

	return r
}

func (r *TestLogWriter) Owner(owner any) log.Writer {
	r.data["owner"] = owner

	return r
}

func (r *TestLogWriter) Hint(hint string) log.Writer {
	r.data["hint"] = hint

	return r
}

func (r *TestLogWriter) Code(code string) log.Writer {
	r.data["code"] = code

	return r
}

func (r *TestLogWriter) With(data map[string]any) log.Writer {
	r.data["with"] = data

	return r
}

func (r *TestLogWriter) Tags(tags ...string) log.Writer {
	r.data["tags"] = tags

	return r
}

func (r *TestLogWriter) WithTrace() log.Writer {
	return r
}

func (r *TestLogWriter) Request(req http.ContextRequest) log.Writer {
	r.data["request"] = req

	return r
}

func (r *TestLogWriter) Response(res http.ContextResponse) log.Writer {
	r.data["response"] = res

	return r
}

func (r *TestLogWriter) In(domain string) log.Writer {
	r.data["in"] = domain

	return r
}

func (r *TestLogWriter) printData() {
	if len(r.data) > 0 {
		fmt.Println(r.data)
	}
}

func prefix(model string) string {
	timestamp := carbon.Now().ToDateTimeString()

	return fmt.Sprintf("[%s] %s.%s: ", timestamp, "test", model)
}
