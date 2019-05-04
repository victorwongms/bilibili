package ups

import (
	"context"
	"flag"
	"go-common/app/service/main/videoup/conf"
	"os"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

var (
	d *Dao
)

func TestMain(m *testing.M) {
	if os.Getenv("DEPLOY_ENV") != "" {
		flag.Set("app_id", "main.archive.videoup-service")
		flag.Set("conf_token", "4b62721602981eb3635dba3b0d866ac5")
		flag.Set("tree_id", "2308")
		flag.Set("conf_version", "docker-1")
		flag.Set("deploy_env", "uat")
		flag.Set("conf_host", "config.bilibili.co")
		flag.Set("conf_path", "/tmp")
		flag.Set("region", "sh")
		flag.Set("zone", "sh001")
	}
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	d = New(conf.Conf)
	m.Run()
	os.Exit(0)
}

func TestUpPing(t *testing.T) {
	var (
		c = context.TODO()
	)
	convey.Convey("Ping", t, func(ctx convey.C) {
		err := d.Ping(c)
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}

func TestGrayCheckUps(t *testing.T) {
	var (
		c = context.TODO()
	)
	convey.Convey("GrayCheckUps", t, func(ctx convey.C) {
		_, err := d.GrayCheckUps(c, 1)
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}

func TestUpSpecial(t *testing.T) {
	var (
		c = context.TODO()
	)
	convey.Convey("upSpecial", t, func(ctx convey.C) {
		_, err := d.upSpecial(c, "1")
		ctx.Convey("Then err should be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}
