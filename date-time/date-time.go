package main

import (
	"dbus/com/deepin/dde/api/setdatetime"
	"dbus/org/gnome/settingsdaemon/datetimemechanism"
	"dlib"
	"dlib/dbus"
	"dlib/dbus/property"
	"dlib/gio-2.0"
	"fmt"
	"github.com/howeyc/fsnotify"
)

const (
	_DATE_TIME_DEST = "com.deepin.daemon.DateAndTime"
	_DATE_TIME_PATH = "/com/deepin/daemon/DateAndTime"
	_DATA_TIME_IFC  = "com.deepin.daemon.DateAndTime"

	_DATE_TIME_SCHEMA = "com.deepin.dde.datetime"
	_TIME_ZONE_FILE   = "/etc/timezone"
)

var (
	_busConn     *dbus.Conn
	_dtGSettings = gio.NewSettings(_DATE_TIME_SCHEMA)

	_setDT      *setdatetime.SetDateTime
	_gdate      *datetimemechanism.DateTimeMechanism
	zoneWatcher *fsnotify.Watcher
)

type DateTime struct {
	AutoSetTime      bool
	Use24HourDisplay *property.GSettingsBoolProperty `access:"readwrite"`
	CurrentTimeZone  string
}

func (date *DateTime) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{_DATE_TIME_DEST, _DATE_TIME_PATH, _DATA_TIME_IFC}
}

func (date *DateTime) SetDate(d string) bool {
	ret, _ := _setDT.SetCurrentDate(d)
	return ret
}

func (date *DateTime) SetTime(t string) bool {
	ret, _ := _setDT.SetCurrentTime(t)
	return ret
}

func (date *DateTime) SetTimeZone(zone string) {
	_gdate.SetTimezone(zone)
	date.CurrentTimeZone = zone
	dbus.NotifyChange(date, "CurrentTimeZone")
}

func (date *DateTime) SetAutoSetTime(auto bool) bool {
	var ret bool
	if auto {
		date.AutoSetTime = true
		ret, _ = _setDT.SetNtpUsing(true)
	} else {
		date.AutoSetTime = false
		ret, _ = _setDT.SetNtpUsing(false)
	}
	dbus.NotifyChange(date, "AutoSetTime")
	return ret
}

func (date *DateTime) SyncNtpTime() bool {
	ret, _ := _setDT.SyncNtpTime()
	return ret
}

func NewDateAndTime() *DateTime {
	dt := &DateTime{}

	dt.Use24HourDisplay = property.NewGSettingsBoolProperty(dt, "Use24HourDisplay", _dtGSettings, "is-24hour")
	dt.CurrentTimeZone, _ = _gdate.GetTimezone()

	dt.AutoSetTime = _dtGSettings.GetBoolean("is-auto-set")
	_dtGSettings.Connect("changed::is-auto-set", func(s *gio.Settings, name string) {
		fmt.Println("is-auto-set changed:", s.GetBoolean("is-auto-set"))
		dt.SetAutoSetTime(s.GetBoolean("is-auto-set"))
	})

	return dt
}

func ListenZone(dt *DateTime) {
	err := zoneWatcher.Watch(_TIME_ZONE_FILE)
	if err != nil {
		fmt.Printf("Watch '%s' Failed: %s\n", _TIME_ZONE_FILE, err)
		return
	}

	go func() {
		for {
			select {
			case ev := <-zoneWatcher.Event:
				fmt.Println("Watcher Event: ", ev)
				if ev.IsDelete() {
					zoneWatcher.Watch(_TIME_ZONE_FILE)
				} else {
					//if ev.IsModify() {
					dt.CurrentTimeZone, _ = _gdate.GetTimezone()
					dbus.NotifyChange(dt, "CurrentTimeZone")
					fmt.Println(dt.CurrentTimeZone)
					//}
				}
			case err := <-zoneWatcher.Error:
				fmt.Println("Watcher Event: ", err)
			}
		}
	}()
}

func Init() {
	var err error

	_setDT, err = setdatetime.NewSetDateTime("/com/deepin/dde/api/SetDateTime")
	if err != nil {
		fmt.Println("New SetDateTime Failed.")
		panic(err)
	}

	_gdate, err = datetimemechanism.NewDateTimeMechanism("/")
	if err != nil {
		fmt.Println("New DateTimeMechanism Failed.")
		panic(err)
	}

	zoneWatcher, err = fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("New FS Watcher Failed.")
		panic(err)
	}
}

func main() {
	Init()
	defer zoneWatcher.Close()

	date := NewDateAndTime()
	err := dbus.InstallOnSession(date)
	if err != nil {
		panic(err)
	}

	ListenZone(date)
	if date.AutoSetTime {
		go _setDT.SetNtpUsing(true)
	}
	dbus.DealWithUnhandledMessage()
	dlib.StartLoop()
}
