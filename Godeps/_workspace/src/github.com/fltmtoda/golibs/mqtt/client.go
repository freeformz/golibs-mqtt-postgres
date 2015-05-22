package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"github.com/fltmtoda/golibs/logger"
	"io/ioutil"
	"net/url"
)

const (
	QOS_ZERO QoS = 0 // At most once -- 最高1回。届くかは保証しない
	QOS_ONE      = 1 // At least once -- 少なくとも一回。重複する可能性がある。
	QOS_TWO      = 2 // Exactly once -- 正確に一回
)

var (
	log = logger.GetLogger()
)

type (
	MqttClient struct {
		internalClient *mqtt.Client
		setting        *Setting
	}

	QoS byte

	SubscribeMessageHandler func(Message) error
)

func Create(setting *Setting) *MqttClient {
	uri, err := url.Parse(setting.URL)
	if err != nil {
		log.Errorf("Faild to parse mqtt-url: %v", err)
		return nil
	}
	opts := mqtt.NewClientOptions()

	// tcp://iot.eclipse.org:1883 - connect to iot.eclipse.org on port 1883 using plain TCP
	// ws://iot.eclipse.org:1883  - connect to iot.eclipse.org on port 1883 using WebSockets
	// tls://iot.eclipse.org:8883 - connect to iot.eclipse.org on port 8883 using TLS (ssl:// and tcps:// are synonyms for tls://)
	opts.AddBroker(fmt.Sprintf("tls://%s", uri.Host))
	//opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))

	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(setting.ClientId)
	opts.AutoReconnect = true
	opts.ProtocolVersion = 4
	opts.CleanSession = true

	if false {
		certPool, err := GetCertPool("/vagrant/goapp/src/github.com/fltmtoda/golibs-examples/mosquitto.org.crt")
		if err != nil {
			panic(err)
		}
		opts.SetTLSConfig(
			&tls.Config{
				RootCAs: certPool,
			},
		)
	}

	opts.OnConnectionLost = func(c *mqtt.Client, err error) {
		if err != nil {
			log.Error("Failed to disconnect mqtt-server: %v [%v]", opts.Servers, err)
			return
		} else {
			log.Infof("Disonnect mqtt-server: %v", opts.Servers)
		}
	}

	return &MqttClient{
		internalClient: mqtt.NewClient(opts),
		setting:        setting,
	}
}

func (c *MqttClient) Connect() error {
	if c.internalClient == nil {
		return fmt.Errorf("Already closed mqtt-server")
	}
	if c.internalClient.IsConnected() {
		log.Warn("Already connected mqtt-server")
		return nil
	}
	token := c.internalClient.Connect()
	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("Failed to connect mqtt-server: %v", token.Error())
	}
	log.Infof("Connect mqtt-server: %v", c.setting.URL)
	return nil
}

func (c *MqttClient) Disconnect(quiesce uint) {
	if c.internalClient == nil || !c.internalClient.IsConnected() {
		log.Warn("Already closed mqtt-server")
		return
	}
	c.internalClient.Disconnect(quiesce)
	log.Infof("Disconnect mqtt-server: %v", c.setting.URL)
}
func (c *MqttClient) ForceDisconnect() {
	if c.internalClient == nil || !c.internalClient.IsConnected() {
		log.Warn("Already closed mqtt-server")
		return
	}
	c.internalClient.ForceDisconnect()
	log.Infof("ForceDisconnect mqtt-server: %v", c.setting.URL)
}

func (c *MqttClient) Publish(topic string, qos QoS, retained bool, payload interface{}) error {
	token := c.internalClient.Publish(topic, byte(qos), retained, payload)
	if qos != QOS_ZERO {
		token.Wait()
	}
	if token.Error() != nil {
		return fmt.Errorf("Failed to send message: %v", token.Error())
	}
	return nil
}

func (c *MqttClient) Subscribe(topic string, qos QoS, callback SubscribeMessageHandler) error {
	token := c.internalClient.Subscribe(
		topic,
		byte(qos),
		func(c *mqtt.Client, msg mqtt.Message) {
			if msg == nil {
				log.Warn("Subscribe message is nil")
				return
			}
			if err := callback(msg); err != nil {
				log.Errorf("Failed to subscript mqtt-server: %v", err)
				return
			}
		},
	)
	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("Failed to subscript mqtt-server: %v", token.Error())
	}
	return nil
}
func (c *MqttClient) Unsubscribe(topics ...string) error {
	token := c.internalClient.Unsubscribe(topics...)
	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("Failed to subscript mqtt-server: %v", token.Error())
	}
	return nil
}

func GetCertPool(pemPath string) (*x509.CertPool, error) {
	certs := x509.NewCertPool()

	pemData, err := ioutil.ReadFile(pemPath)
	if err != nil {
		return nil, err
	}
	certs.AppendCertsFromPEM(pemData)
	return certs, nil
}
