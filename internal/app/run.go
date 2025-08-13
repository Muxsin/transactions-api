package app

func (a *app) Run() error {
	if err := a.router.Run(":" + a.config.HTTPServerPort); err != nil {
		return err
	}

	return nil
}
