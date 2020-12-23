package sql

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"go.kelfa.io/pkg/dal/objects"
	"go.kelfa.io/pkg/session"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DP struct {
	conn *gorm.DB
}

func New(bo objects.BackendOptions) (*DP, error) {
	if len(bo.Path) == 0 {
		return nil, errors.New("path option in back-end options is required")
	}
	conn, err := gorm.Open("sqlite3", bo.Path)
	if err != nil {
		return nil, err
	}
	conn.AutoMigrate(&objects.DataPoint{})
	conn.AutoMigrate(&session.Session{})
	d := DP{
		conn: conn,
	}
	return &d, nil
}

func (d *DP) DataBeginTime() (*time.Time, error) {
	dp := new(objects.DataPoint)
	err := d.conn.Order("date_time").First(&dp).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &dp.DateTime, nil
}

func (d *DP) DataEndTime() (*time.Time, error) {
	dp := new(objects.DataPoint)
	err := d.conn.Order("date_time desc").First(&dp).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &dp.DateTime, nil
}

func (d *DP) GetDataPoints(from time.Time, to time.Time) ([]objects.DataPoint, error) {
	var items []objects.DataPoint
	if err := d.conn.Where("date_time >= ?", from).Where("date_time <= ?", to).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (d *DP) AddDataPoint(item objects.DataPoint) error {
	return d.conn.Create(&item).Error
}

func (d *DP) AddDataPoints(items []objects.DataPoint) error {
	for _, item := range items {
		if err := d.AddDataPoint(item); err != nil {
			return err
		}
	}
	return nil
}

func (d *DP) AddSession(item session.Session) error {
	return d.conn.Create(&item).Error
}

func (d *DP) AddSessions(items []session.Session) error {
	for _, item := range items {
		if err := d.AddSession(item); err != nil {
			return err
		}
	}
	return nil
}
