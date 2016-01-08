package plugin

import "github.com/shaoshing/train"

func InitAsset() error {
	train.ConfigureHttpHandler(nil)
	train.Config.AssetsPath = "assets"
	train.Config.Verbose = true
	train.Config.SASS.LineNumbers = true
	return nil
}
