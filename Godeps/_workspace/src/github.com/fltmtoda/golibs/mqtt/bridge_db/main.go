package bdb

import (
	"encoding/json"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fltmtoda/golibs/db"
	"github.com/fltmtoda/golibs/log"
	"github.com/fltmtoda/golibs/mqtt"
)

type BridgeDB struct {
	setting    *Setting
	mqttClient *mqtt.MqttClient
	dbm        db.DB
	wg         sync.WaitGroup
	counter    int64
}

func Create(setting *Setting) (*BridgeDB, error) {
	var err error
	dbm, err := db.Create(setting.DB)
	if err != nil {
		return nil, err
	}
	mc := mqtt.Create(setting.MQTT)
	return &BridgeDB{
		setting:    setting,
		mqttClient: mc,
		dbm:        dbm,
		wg:         sync.WaitGroup{},
		counter:    int64(0),
	}, nil
}

type Converter func(topic string) interface{}

func (b *BridgeDB) Start(converter Converter) error {
	if err := b.mqttClient.Connect(); err != nil {
		log.Error(err)
		return err
	}
	defer b.Stop()

	var st time.Time
	var counterLimit int64 = 100000
	sem := make(chan int, b.setting.Concurrency)
	for {
		err := b.mqttClient.Subscribe(
			b.setting.Topic,
			mqtt.QOS_ZERO,
			func(msg mqtt.Message) error {
				atomic.AddInt64(&b.counter, 1)
				if b.counter == 1 {
					log.Infof(
						"Benchmark mqtt-subscriber pg-bridge start...",
					)
					st = time.Now()
				}
				log.Infof("counter=%v", b.counter)
				sem <- 1
				b.wg.Add(1)
				go func(topic string, msgData []byte) {
					log.Info(string(msgData))

					r := converter(topic)
					json.Unmarshal(msgData, r)
					err := b.dbm.Save(r)
					if err != nil {
						log.Error(err)
					}
					b.wg.Done()
					<-sem
				}(msg.Topic(), msg.Payload())
				return nil
			},
		)
		if err != nil {
			log.Error(err)
			break
		}
		if b.counter >= counterLimit {
			break
		}
		time.Sleep(time.Millisecond)
	}
	b.wg.Wait()
	log.Infof("counter=%v", b.counter)
	log.Infof(
		"Benchmark mqtt-subscriber pg-bridge end.[elapsedtime=%f sec]",
		time.Now().Sub(st).Seconds(),
	)
	return nil
}
func (b *BridgeDB) Stop() {
	if b.mqttClient != nil {
		b.mqttClient.Disconnect(100)
	}
	if b.dbm != nil {
		b.dbm.CloseDB()
	}
}
