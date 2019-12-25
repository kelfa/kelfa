package sql

import (
	"time"

	"github.com/jinzhu/gorm"
	"go.kelfa.io/kelfa/pkg/dal/objects"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DP struct {
	conn   *gorm.DB
	prefix string
}

func New(bo objects.BackendOptions) (*DP, error) {
	conn, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}
	conn.AutoMigrate(&objects.DataPoint{})
	d := DP{
		conn: conn,
	}
	return &d, nil
}

func (d *DP) DataBeginTime() (*time.Time, error) {
	dp := new(objects.DataPoint)
	err := d.conn.Order("date-time").First(&dp).Error
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
	err := d.conn.Order("date-time desc").First(&dp).Error
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
	if err := d.conn.Where("date-time >= ?", from).Where("date-time <= ?", to).Find(&items).Error; err != nil {
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
